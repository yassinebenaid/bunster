package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func Generate(script ast.Script) ir.Program {
	g := generator{}

	var buf InstructionBuffer
	for _, statement := range script {
		g.generate(&buf, statement, nil)
	}

	return ir.Program{
		Instructions: buf,
	}
}

type generator struct{}

type InstructionBuffer []ir.Instruction

func (ib *InstructionBuffer) add(ins ir.Instruction) {
	*ib = append(*ib, ins)
}

func (g *generator) generate(buf *InstructionBuffer, statement ast.Statement, pc *pipeContext) {
	switch v := statement.(type) {
	case ast.List:
		g.handleList(buf, v)
	case ast.Pipeline:
		g.handlePipeline(buf, v)
	case ast.Command:
		g.handleSimpleCommand(buf, v, pc)
	case ast.ParameterAssignement:
		g.handleParameterAssignment(buf, v)
	case ast.Group:
		g.handleGroup(buf, v, pc)
	default:
		panic(fmt.Sprintf("unhandled statement type (%T)", statement))
	}
}

func (g *generator) handleList(buf *InstructionBuffer, l ast.List) {
	g.generate(buf, l.Left, nil)

	var bodybuf InstructionBuffer
	g.generate(&bodybuf, l.Right, nil)

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

		var pc pipeContext
		if i == 0 {
			pc = pipeContext{
				writer: fmt.Sprintf("pipeWriter%d", i+1),
				stderr: cmd.Stderr,
			}
		} else if i == (len(p) - 1) {
			pc = pipeContext{
				reader: fmt.Sprintf("pipeReader%d", i),
			}
		} else {
			pc = pipeContext{
				writer: fmt.Sprintf("pipeWriter%d", i+1),
				reader: fmt.Sprintf("pipeReader%d", i),
				stderr: cmd.Stderr,
			}
		}

		pc.waitgroup = "pipelineWaitgroup"
		g.generate(&cmdbuf, cmd.Command, &pc)
	}

	cmdbuf.add(ir.WaitPipelineWaitgroup("pipelineWaitgroup"))

	*buf = append(*buf, ir.Closure{
		Body: cmdbuf,
	})
}

func (g *generator) handleSimpleCommand(buf *InstructionBuffer, cmd ast.Command, pc *pipeContext) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.Declare{Name: "commandName", Value: g.handleExpression(cmd.Name)})
	cmdbuf.add(ir.DeclareSlice{Name: "arguments"})

	for _, arg := range cmd.Args {
		cmdbuf.add(ir.Append{Name: "arguments", Value: g.handleExpression(arg)})
	}

	cmdbuf.add(ir.Declare{
		Name:  "command",
		Value: ir.InitCommand{Name: "commandName", Args: "arguments"},
	})

	for _, env := range cmd.Env {
		cmdbuf.add(ir.SetCmdEnv{
			Command: "command",
			Key:     env.Name,
			Value:   g.handleExpression(env.Value),
		})
	}

	g.handleRedirections(&cmdbuf, "command", cmd.Redirections, pc, false)
	cmdbuf.add(ir.SetStream{
		Name: "command.Stdin",
		Fd:   ir.String("0"),
	})
	cmdbuf.add(ir.SetStream{
		Name: "command.Stdout",
		Fd:   ir.String("1"),
	})
	cmdbuf.add(ir.SetStream{
		Name: "command.Stderr",
		Fd:   ir.String("2"),
	})

	if pc != nil {
		cmdbuf.add(ir.StartCommand("command"))
		cmdbuf.add(ir.PushToPipelineWaitgroup{
			Waitgroup: pc.waitgroup,
			Command:   "command",
		})
	} else {
		cmdbuf.add(ir.RunCommand("command"))
	}

	*buf = append(*buf, ir.Closure{
		Body: cmdbuf,
	})
}

func (g *generator) handleExpression(expression ast.Expression) ir.Instruction {
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
			concat = append(concat, g.handleExpression(expr))
		}
		return concat
	case ast.UnquotedString:
		var concat ir.Concat
		for _, expr := range v {
			concat = append(concat, g.handleExpression(expr))
		}
		return concat
	case ast.CommandSubstitution:
		return g.handleCommandSubstitution(v)
	default:
		panic(fmt.Sprintf("unhandled expression type (%T)", expression))
	}
}

func (g *generator) handleRedirections(buf *InstructionBuffer, name string, redirections []ast.Redirection, pc *pipeContext, nd bool) {
	buf.add(ir.CloneFDT{ND: nd})

	// if we're inside a pipline, we need to connect the pipe to the command.(before any other redirection)
	if pc != nil {
		if pc.writer != "" {
			buf.add(ir.AddStream{Fd: "1", StreamName: pc.writer})

			if pc.stderr {
				buf.add(ir.AddStream{Fd: "2", StreamName: pc.writer})
			}
		}

		if pc.reader != "" {
			buf.add(ir.AddStream{Fd: "0", StreamName: pc.reader})
		}
	}

	for i, redirection := range redirections {
		switch redirection.Method {
		case ">", ">|":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_WRITE,
			})
			buf.add(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case ">>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_APPEND,
			})
			buf.add(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "&>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_WRITE,
			})
			buf.add(ir.AddStream{
				Fd:         "1",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
			buf.add(ir.AddStream{
				Fd:         "2",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "&>>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_APPEND,
			})
			buf.add(ir.AddStream{
				Fd:         "1",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
			buf.add(ir.AddStream{
				Fd:         "2",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case ">&", "<&":
			if redirection.Dst == nil && redirection.Close {
				buf.add(ir.CloseStream{
					Fd: ir.String(redirection.Src),
				})
			} else {
				buf.add(ir.DuplicateStream{
					Old: redirection.Src,
					New: g.handleExpression(redirection.Dst),
				})
				if redirection.Close {
					buf.add(ir.CloseStream{
						Fd: g.handleExpression(redirection.Dst),
					})
				}
			}
		case "<":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_READ,
			})
			buf.add(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "<<<":
			buf.add(ir.Declare{
				Name: fmt.Sprintf("%s_file_%d", name, i),
				Value: ir.NewBuffer{
					Readonly: true,
					Value:    g.handleExpression(redirection.Dst),
				},
			})
			buf.add(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "<>":
			buf.add(ir.OpenStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_RW,
			})
			buf.add(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		}
	}
}

type pipeContext struct {
	writer    string
	reader    string
	waitgroup string
	stderr    bool
}

func (g *generator) handleParameterAssignment(buf *InstructionBuffer, p ast.ParameterAssignement) {
	buf.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})
	for _, assignment := range p {
		ins := ir.SetVar{
			Key:   assignment.Name,
			Value: ir.String(""),
		}
		if assignment.Value != nil {
			ins.Value = g.handleExpression(assignment.Value)
		}

		buf.add(ins)
	}
}
