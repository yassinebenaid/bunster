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

type generator struct {
	program ir.Program
}

type InstructionBuffer []ir.Instruction

func (ib *InstructionBuffer) add(ins ir.Instruction) {
	*ib = append(*ib, ins)
}

func (g *generator) generate(buf *InstructionBuffer, statement ast.Statement, pc *pipeContext) {
	switch v := statement.(type) {
	case ast.Command:
		var cmdbuf InstructionBuffer
		g.handleSimpleCommand(&cmdbuf, v, pc)
		*buf = append(*buf, ir.Closure{
			Body: cmdbuf,
		})
	case ast.Pipeline:
		var cmdbuf InstructionBuffer
		g.handlePipeline(&cmdbuf, v)
		*buf = append(*buf, ir.Closure{
			Body: cmdbuf,
		})
	case ast.List:
		g.handleList(buf, v)
	case ast.ParameterAssignement:
		g.handleParameterAssignment(buf, v)

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
	buf.add(ir.NewPipelineWaitgroup("pipelineWaitgroup"))

	for i, cmd := range p {
		if i < (len(p) - 1) { //last command doesn't need a pipe
			buf.add(ir.NewPipe{
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
		g.generate(buf, cmd.Command, &pc)
	}

	buf.add(ir.WaitPipelineWaitgroup("pipelineWaitgroup"))

}

func (g *generator) handleSimpleCommand(buf *InstructionBuffer, cmd ast.Command, pc *pipeContext) {
	buf.add(ir.Declare{Name: "commandName", Value: g.handleExpression(cmd.Name)})
	buf.add(ir.DeclareSlice{Name: "arguments"})

	for _, arg := range cmd.Args {
		buf.add(ir.Append{Name: "arguments", Value: g.handleExpression(arg)})
	}

	buf.add(ir.Declare{
		Name:  "command",
		Value: ir.InitCommand{Name: "commandName", Args: "arguments"},
	})

	for _, env := range cmd.Env {
		buf.add(ir.SetCmdEnv{
			Command: "command",
			Key:     env.Name,
			Value:   g.handleExpression(env.Value),
		})
	}

	g.handleRedirections(buf, "command", cmd.Redirections, pc)

	if pc != nil {
		buf.add(ir.StartCommand("command"))
		buf.add(ir.PushToPipelineWaitgroup{
			Waitgroup: pc.waitgroup,
			Command:   "command",
		})
	} else {
		buf.add(ir.RunCommand("command"))
	}
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
	default:
		panic(fmt.Sprintf("unhandled expression type (%T)", expression))
	}
}

func (g *generator) handleRedirections(buf *InstructionBuffer, name string, redirections []ast.Redirection, pc *pipeContext) {
	var fdt = name + "FDT"
	buf.add(ir.CloneFDT(fdt))

	// if we're inside a pipline, we need to connect the pipe to the command.(before any other redirection)
	if pc != nil {
		if pc.writer != "" {
			buf.add(ir.AddStream{FDT: fdt, Fd: "1", StreamName: pc.writer})

			if pc.stderr {
				buf.add(ir.AddStream{FDT: fdt, Fd: "2", StreamName: pc.writer})
			}
		}

		if pc.reader != "" {
			buf.add(ir.AddStream{FDT: fdt, Fd: "0", StreamName: pc.reader})
		}
	}

	for i, redirection := range redirections {
		switch redirection.Method {
		case ">", ">|":
			buf.add(ir.OpenStream{
				FDT:    fdt,
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_WRITE,
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case ">>":
			buf.add(ir.OpenStream{
				FDT:    fdt,
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_APPEND,
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "&>":
			buf.add(ir.OpenStream{
				FDT:    fdt,
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_WRITE,
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         "1",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         "2",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "&>>":
			buf.add(ir.OpenStream{
				FDT:    fdt,
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_APPEND,
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         "1",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         "2",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case ">&", "<&":
			if redirection.Dst == nil && redirection.Close {
				buf.add(ir.CloseStream{
					FDT: fdt,
					Fd:  ir.String(redirection.Src),
				})
			} else {
				buf.add(ir.DuplicateStream{
					FDT: fdt,
					Old: redirection.Src,
					New: g.handleExpression(redirection.Dst),
				})

				if redirection.Close {
					buf.add(ir.CloseStream{
						FDT: fdt,
						Fd:  g.handleExpression(redirection.Dst),
					})
				}
			}
		case "<":
			buf.add(ir.OpenStream{
				FDT:    fdt,
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_READ,
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "<<<":
			buf.add(ir.Declare{
				Name: fmt.Sprintf("%s_file_%d", name, i),
				Value: ir.NewStringStream{
					Target: g.handleExpression(redirection.Dst),
				},
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "<>":
			buf.add(ir.OpenStream{
				FDT:    fdt,
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
				Mode:   ir.FLAG_RW,
			})
			buf.add(ir.AddStream{
				FDT:        fdt,
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		}
	}

	buf.add(ir.Set{
		Name:  fmt.Sprintf("%s.Stdin", name),
		Value: ir.GetStream{FDT: fdt, Fd: ir.String("0")},
	})
	buf.add(ir.Set{
		Name:  fmt.Sprintf("%s.Stdout", name),
		Value: ir.GetStream{FDT: fdt, Fd: ir.String("1")},
	})
	buf.add(ir.Set{
		Name:  fmt.Sprintf("%s.Stderr", name),
		Value: ir.GetStream{FDT: fdt, Fd: ir.String("2")},
	})
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
