package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

type InstructionBuffer []ir.Instruction

func (ib *InstructionBuffer) add(ins ir.Instruction) {
	*ib = append(*ib, ins)
}

func Generate(script ast.Script) ir.Program {
	g := generator{}

	var buf InstructionBuffer
	for _, statement := range script {
		g.generate(&buf, statement)
	}

	return ir.Program{
		Instructions: buf,
		Embeds:       g.embeds,
	}
}

type generator struct {
	expressionsCount int
	scopesCount      int
	embeds           []string
}

func (g *generator) generate(buf *InstructionBuffer, statement ast.Statement) {
	switch v := statement.(type) {
	case ast.List:
		g.handleList(buf, v)
	case ast.Pipeline:
		g.handlePipeline(buf, v)
	case ast.Command:
		g.handleSimpleCommand(buf, v)
	case ast.ParameterAssignement:
		g.handleParameterAssignment(buf, v)
	case ast.LocalParameterAssignement:
		g.handleLocalParameterAssignment(buf, v)
	case ast.ExportParameterAssignement:
		g.handleExportParameterAssignment(buf, v)
	case ast.Group:
		g.handleGroup(buf, v)
	case ast.SubShell:
		g.handleSubshell(buf, v)
	case *ast.If:
		g.handleIf(buf, v)
	case *ast.Break:
		g.handleBreak(buf, v)
	case *ast.Continue:
		g.handleContinue(buf, v)
	case *ast.Loop:
		g.handleLoop(buf, v)
	case *ast.For:
		g.handleForLoop(buf, v)
	case *ast.RangeLoop:
		g.handleRangeLoop(buf, v)
	case ast.BackgroundConstruction:
		g.handleBackgroundConstruction(buf, v)
	case ast.InvertExitCode:
		g.generate(buf, v.Statement)
		buf.add(ir.InvertExitCode{})
	case ast.Wait:
		g.handleWait(buf, v)
	case ast.Function:
		g.handleFunction(buf, v)
	case ast.Defer:
		var body InstructionBuffer
		g.generate(&body, v.Command)
		buf.add(ir.Defer{Body: body})
	case ast.Test:
		g.handleTest(buf, v)
	case ast.Embed:
		g.handleEmbed(buf, v)
	case ast.ArithmeticCommand:
		g.handleArithmeticCommand(buf, v)
	case *ast.Case:
		g.handleCase(buf, v)
	case ast.Exit:
		g.handleExit(buf, v)
	default:
		panic(fmt.Sprintf("Unsupported statement: %T", v))
	}
}

func (g *generator) handleList(buf *InstructionBuffer, l ast.List) {
	g.generate(buf, l.Left)

	var bodybuf InstructionBuffer
	g.generate(&bodybuf, l.Right)

	buf.add(ir.IfLastExitCode{
		Zero: l.Operator == "&&",
		Body: bodybuf,
	})
}

func (g *generator) handlePipeline(buf *InstructionBuffer, p ast.Pipeline) {
	var cmdbuf InstructionBuffer
	cmdbuf.add(ir.NewPipelineWaitgroup("pipelineWaitgroup"))

	for i, cmd := range p {
		if i < (len(p) - 1) { //last command doesn't need a pipe
			cmdbuf.add(ir.NewPipe{
				Writer: fmt.Sprintf("pipeWriter%d", i+1),
				Reader: fmt.Sprintf("pipeReader%d", i+1),
			})
		}

		var body, gorouting InstructionBuffer
		body.add(ir.CloneStreamManager{DontDestroy: true})

		if i == 0 {
			body.add(ir.AddStream{Fd: "1", StreamName: fmt.Sprintf("pipeWriter%d", i+1)})

			if cmd.Stderr {
				body.add(ir.AddStream{Fd: "2", StreamName: fmt.Sprintf("pipeWriter%d", i+1)})
			}
		} else if i == (len(p) - 1) {
			body.add(ir.AddStream{Fd: "0", StreamName: fmt.Sprintf("pipeReader%d", i)})
		} else {
			body.add(ir.AddStream{Fd: "0", StreamName: fmt.Sprintf("pipeReader%d", i)})
			body.add(ir.AddStream{Fd: "1", StreamName: fmt.Sprintf("pipeWriter%d", i+1)})
			if cmd.Stderr {
				body.add(ir.AddStream{Fd: "2", StreamName: fmt.Sprintf("pipeWriter%d", i+1)})
			}
		}

		body.add(ir.CloneShell{DontTerminate: true})
		body.add(ir.Literal("var done = make(chan struct{},1)\n"))
		body.add(ir.PushToPipelineWaitgroup{
			Waitgroup: "pipelineWaitgroup",
			Value: ir.Literal(`func() int {
				<-done
				shell.Terminate(streamManager)
			 	streamManager.Destroy()
				return shell.ExitCode
			}`),
		})

		gorouting.add(ir.Literal("defer func(){ done<-struct{}{} }()\n"))
		g.generate(&gorouting, cmd.Command)
		body.add(ir.Gorouting(gorouting))
		cmdbuf.add(ir.Closure(body))
	}

	cmdbuf.add(ir.WaitPipelineWaitgroup("pipelineWaitgroup"))

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleSimpleCommand(buf *InstructionBuffer, cmd ast.Command) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, cmd.Redirections)

	cmdbuf.add(ir.Declare{Name: "commandName", Value: g.handleExpression(&cmdbuf, cmd.Name)})
	cmdbuf.add(ir.DeclareSlice{Name: "arguments"})
	cmdbuf.add(ir.DeclareMap("env"))

	for _, arg := range cmd.Args {
		cmdbuf.add(ir.Append{Name: "arguments", Value: g.handleExpression(&cmdbuf, arg)})
	}
	for _, env := range cmd.Env {
		var value ir.Instruction = ir.String("")
		if env.Value != nil {
			value = g.handleExpression(&cmdbuf, env.Value)
		}
		cmdbuf.add(ir.SetMap{Name: "env", Key: env.Name, Value: value})
	}

	cmdbuf.add(ir.Exec{Name: "commandName", Args: "arguments", Env: "env"})
	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleExpression(buf *InstructionBuffer, expression ast.Expression) ir.Instruction {
	g.expressionsCount++
	switch v := expression.(type) {
	case ast.Word:
		return ir.String(v)
	case ast.Number:
		return ir.String(v)
	case ast.Var:
		return ir.ReadVar(v)
	case ast.SpecialVar:
		return ir.ReadSpecialVar(v)
	case ast.QuotedString:
		var concat ir.Concat
		for _, expr := range v {
			concat = append(concat, g.handleExpression(buf, expr))
		}
		return concat
	case ast.UnquotedString:
		var concat ir.Concat
		for _, expr := range v {
			concat = append(concat, g.handleExpression(buf, expr))
		}
		return concat
	case ast.CommandSubstitution:
		return g.handleCommandSubstitution(buf, v)
	case ast.Arithmetic:
		return g.handleArithmeticSubstitution(buf, v)
	case ast.VarLength:
		return ir.VarLength{Name: v.Parameter.Name}
	case ast.VarOrDefault:
		return g.handleParameterExpansionVarOrDefault(buf, v)
	case ast.VarOrSet:
		return g.handleParameterExpansionVarOrSet(buf, v)
	case ast.CheckAndUse:
		return g.handleParameterExpansionCheckAndUse(buf, v)
	case ast.Slice:
		return g.handleParameterExpansionSlice(buf, v)
	case ast.ChangeCase:
		return g.handleParameterExpansionChangeCase(buf, v)
	case ast.MatchAndRemove:
		return g.handleParameterExpansionMatchAndRemove(buf, v)
	case ast.MatchAndReplace:
		return g.handleParameterExpansionMatchAndReplace(buf, v)
	default:
		panic(fmt.Sprintf("Unsupported expression: %T", v))
	}
}

func (g *generator) handleRedirections(buf *InstructionBuffer, redirections []ast.Redirection) {

	for i, redirection := range redirections {
		switch redirection.Method {
		case ">", ">|":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("stream%d", i),
				Target: g.handleExpression(buf, redirection.Dst),
				Mode:   ir.FLAG_WRITE,
			})
			buf.add(ir.AddStream{Fd: redirection.Src, StreamName: fmt.Sprintf("stream%d", i)})
		case ">>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("stream%d", i),
				Target: g.handleExpression(buf, redirection.Dst),
				Mode:   ir.FLAG_APPEND,
			})
			buf.add(ir.AddStream{Fd: redirection.Src, StreamName: fmt.Sprintf("stream%d", i)})
		case "&>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("stream%d", i),
				Target: g.handleExpression(buf, redirection.Dst),
				Mode:   ir.FLAG_WRITE,
			})
			buf.add(ir.AddStream{Fd: "1", StreamName: fmt.Sprintf("stream%d", i)})
			buf.add(ir.AddStream{Fd: "2", StreamName: fmt.Sprintf("stream%d", i)})
		case "&>>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("stream%d", i),
				Target: g.handleExpression(buf, redirection.Dst),
				Mode:   ir.FLAG_APPEND,
			})
			buf.add(ir.AddStream{Fd: "1", StreamName: fmt.Sprintf("stream%d", i)})
			buf.add(ir.AddStream{Fd: "2", StreamName: fmt.Sprintf("stream%d", i)})
		case ">&", "<&":
			if redirection.Dst == nil && redirection.Close {
				buf.add(ir.CloseStream{
					Fd: ir.String(redirection.Src),
				})
			} else {
				buf.add(ir.DuplicateStream{Old: redirection.Src, New: g.handleExpression(buf, redirection.Dst)})
				if redirection.Close {
					buf.add(ir.CloseStream{Fd: g.handleExpression(buf, redirection.Dst)})
				}
			}
		case "<":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("stream%d", i),
				Target: g.handleExpression(buf, redirection.Dst),
				Mode:   ir.FLAG_READ,
			})
			buf.add(ir.AddStream{Fd: redirection.Src, StreamName: fmt.Sprintf("stream%d", i)})
		case "<<<":
			buf.add(ir.NewPipeBuffer{
				Value: ir.Concat{
					g.handleExpression(buf, redirection.Dst),
					ir.String("\n"),
				},
				Name: fmt.Sprintf("buffer%d", i),
			})
			buf.add(ir.Declare{Name: fmt.Sprintf("stream%d", i), Value: ir.Literal(fmt.Sprintf("buffer%d", i))})
			buf.add(ir.AddStream{Fd: redirection.Src, StreamName: fmt.Sprintf("stream%d", i)})
		case "<>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("stream%d", i),
				Target: g.handleExpression(buf, redirection.Dst),
				Mode:   ir.FLAG_RW,
			})
			buf.add(ir.AddStream{Fd: redirection.Src, StreamName: fmt.Sprintf("stream%d", i)})
		}
	}
}

func (g *generator) handleParameterAssignment(buf *InstructionBuffer, p ast.ParameterAssignement) {
	var scope InstructionBuffer

	for _, assignment := range p {
		ins := ir.SetVar{
			Key:   assignment.Name,
			Value: ir.String(""),
		}
		if assignment.Value != nil {
			ins.Value = g.handleExpression(&scope, assignment.Value)
		}

		scope.add(ins)
	}

	buf.add(ir.Closure(scope))
}

func (g *generator) handleLocalParameterAssignment(buf *InstructionBuffer, p ast.LocalParameterAssignement) {
	var scope InstructionBuffer

	for _, assignment := range p {
		ins := ir.SetLocalVar{
			Key:   assignment.Name,
			Value: ir.String(""),
		}
		if assignment.Value != nil {
			ins.Value = g.handleExpression(&scope, assignment.Value)
		}

		scope.add(ins)
	}

	buf.add(ir.Closure(scope))
}

func (g *generator) handleExportParameterAssignment(buf *InstructionBuffer, p ast.ExportParameterAssignement) {
	var scope InstructionBuffer

	for _, assignment := range p {
		if assignment.Value != nil {
			scope.add(ir.SetExportVar{Key: assignment.Name, Value: g.handleExpression(&scope, assignment.Value)})
		} else {
			scope.add(ir.MarkVarAsExported(assignment.Name))
		}
	}

	buf.add(ir.Closure(scope))
}

func (g *generator) handleBackgroundConstruction(buf *InstructionBuffer, b ast.BackgroundConstruction) {
	var scope InstructionBuffer

	scope.add(ir.CloneStreamManager{DontDestroy: true})
	scope.add(ir.OpenStream{Name: "stdin", Target: ir.String("/dev/null"), Mode: ir.FLAG_READ})
	scope.add(ir.AddStream{Fd: "0", StreamName: "stdin"})

	scope.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})
	scope.add(ir.Literal("shell.WaitGroup.Add(1)\n"))
	scope.add(ir.Declare{Name: "done", Value: ir.Literal("shell.WaitGroup.Done")})
	scope.add(ir.CloneShell{DontTerminate: true})

	var body InstructionBuffer
	body.add(ir.Literal("defer done()\n"))
	body.add(ir.Literal("defer streamManager.Destroy()\n"))
	body.add(ir.DeferTerminateShell{})
	g.generate(&body, b.Statement)
	scope.add(ir.Gorouting(body))

	buf.add(ir.Closure(scope))
}

func (g *generator) handleStatementContext(buf *InstructionBuffer, b ast.BreakPoints) {
	for _, breakpoint := range b {
		switch breakpoint.Type {
		case ast.RETURN:
			buf.add(ir.If{
				Condition: ir.Literal(fmt.Sprintf("breakpoint%d", breakpoint.Id)),
				Body:      []ir.Instruction{ir.Literal("return")},
			})
		case ast.BREAK:
			buf.add(ir.If{
				Condition: ir.Literal(fmt.Sprintf("breakpoint%d", breakpoint.Id)),
				Body:      []ir.Instruction{ir.Literal("break")},
			})
		case ast.CONTINUE:
			buf.add(ir.If{
				Condition: ir.Literal(fmt.Sprintf("breakpoint%d", breakpoint.Id)),
				Body:      []ir.Instruction{ir.Literal("continue")},
			})
		case ast.DECLARE:
			buf.add(ir.Declare{Name: fmt.Sprintf("breakpoint%d", breakpoint.Id), Value: ir.Literal("false")})
			buf.add(ir.Set{Name: "_", Value: ir.Literal(fmt.Sprintf("breakpoint%d", breakpoint.Id))})
		}
	}

}
