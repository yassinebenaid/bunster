package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func Generate(script ast.Script) ir.Program {
	g := generator{}
	g.generate(script)

	return g.program
}

type generator struct {
	program  ir.Program
	cmdCount int
}

type InstructionBuffer []ir.Instruction

func (ib *InstructionBuffer) add(ins ir.Instruction) {
	*ib = append(*ib, ins)
}

func (g *generator) generate(script ast.Script) {
	for _, statement := range script {
		switch v := statement.(type) {
		case ast.Command:
			var buf InstructionBuffer
			g.handleSimpleCommand(&buf, v)
			g.program.Instructions = append(g.program.Instructions, ir.Closure{
				Body: buf,
			})
		}
	}
}

func (g *generator) handleSimpleCommand(buf *InstructionBuffer, cmd ast.Command) {
	id := g.cmdCount
	g.cmdCount++

	buf.add(ir.Declare{
		Name:  fmt.Sprintf("cmd_%d_name", id),
		Value: g.handleExpression(cmd.Name),
	})

	buf.add(ir.DeclareSlice(fmt.Sprintf("cmd_%d_args", id)))

	for _, arg := range cmd.Args {
		buf.add(ir.Append{
			Name:  fmt.Sprintf("cmd_%d_args", id),
			Value: g.handleExpression(arg),
		})
	}

	buf.add(ir.Declare{
		Name: fmt.Sprintf("cmd_%d", id),
		Value: ir.InitCommand{
			Name: fmt.Sprintf("cmd_%d_name", id),
			Args: fmt.Sprintf("cmd_%d_args", id),
		},
	})

	g.handleRedirections(buf, fmt.Sprintf("cmd_%d", id), cmd.Redirections)

	buf.add(ir.RunCommanOrFail{
		Command: fmt.Sprintf("cmd_%d", id),
		Name:    fmt.Sprintf("cmd_%d_name", id),
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
	default:
		panic(fmt.Sprintf("unhandled expression type (%T)", expression))
	}
}

func (g *generator) handleRedirections(buf *InstructionBuffer, name string, redirections []ast.Redirection) {
	var fdt = name + "_fdt"
	buf.add(ir.CloneFDT(fdt))

	for i, redirection := range redirections {
		switch redirection.Method {
		case ">", ">|":
			buf.add(ir.OpenStream{
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

	buf.add(ir.GetStream{
		FDT:        fdt,
		Fd:         ir.String("0"),
		StreamName: fmt.Sprintf("%s_stdin", name),
	})
	buf.add(ir.Set{Name: fmt.Sprintf("%s.Stdin", name), Value: ir.Literal(fmt.Sprintf("%s_stdin", name))})
	buf.add(ir.GetStream{
		FDT:        fdt,
		Fd:         ir.String("1"),
		StreamName: fmt.Sprintf("%s_stdout", name),
	})
	buf.add(ir.Set{Name: fmt.Sprintf("%s.Stdout", name), Value: ir.Literal(fmt.Sprintf("%s_stdout", name))})
	buf.add(ir.GetStream{
		FDT:        fdt,
		Fd:         ir.String("2"),
		StreamName: fmt.Sprintf("%s_stderr", name),
	})
	buf.add(ir.Set{Name: fmt.Sprintf("%s.Stderr", name), Value: ir.Literal(fmt.Sprintf("%s_stderr", name))})
}
