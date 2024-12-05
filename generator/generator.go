package generator

import (
	"fmt"

	"github.com/yassinebenaid/ryuko/ast"
	"github.com/yassinebenaid/ryuko/ir"
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

func (g *generator) ins(ins ir.Instruction) {
	g.program.Instructions = append(g.program.Instructions, ins)
}

func (g *generator) generate(script ast.Script) {
	for _, statement := range script {
		switch v := statement.(type) {
		case ast.Command:
			g.handleSimpleCommand(v)
		}
	}
}

func (g *generator) handleSimpleCommand(cmd ast.Command) {
	id := g.cmdCount
	g.cmdCount++

	g.ins(ir.Declare{
		Name:  fmt.Sprintf("cmd_%d_name", id),
		Value: g.handleExpression(cmd.Name),
	})

	g.ins(ir.DeclareSlice(fmt.Sprintf("cmd_%d_args", id)))

	for _, arg := range cmd.Args {
		g.ins(ir.Append{
			Name:  fmt.Sprintf("cmd_%d_args", id),
			Value: g.handleExpression(arg),
		})
	}

	g.ins(ir.Declare{
		Name: fmt.Sprintf("cmd_%d", id),
		Value: ir.InitCommand{
			Name: fmt.Sprintf("cmd_%d_name", id),
			Args: fmt.Sprintf("cmd_%d_args", id),
		},
	})

	g.handleRedirections(fmt.Sprintf("cmd_%d", id), cmd.Redirections)

	g.ins(ir.RunCommanOrFail{
		Command: fmt.Sprintf("cmd_%d", id),
		Name:    fmt.Sprintf("cmd_%d_name", id),
	})
}

func (g *generator) handleExpression(expression ast.Expression) ir.Instruction {
	switch v := expression.(type) {
	case ast.Word:
		return ir.String(v)
	case ast.Var:
		return ir.ReadVar(v)
	default:
		panic(fmt.Sprintf("unhandled expression type (%T)", expression))
	}
}

func (g *generator) handleRedirections(name string, redirections []ast.Redirection) {
	g.ins(ir.AddStream{Fd: "0", StreamName: "shell.Stdin"})
	g.ins(ir.AddStream{Fd: "1", StreamName: "shell.Stdout"})
	g.ins(ir.AddStream{Fd: "2", StreamName: "shell.Stderr"})

	for i, redirection := range redirections {
		switch redirection.Method {
		case ">", ">|":
			g.ins(ir.OpenWritableStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
			})
			g.ins(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case ">>":
			g.ins(ir.OpenAppendableStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
			})
			g.ins(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "&>":
			g.ins(ir.OpenWritableStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
			})
			g.ins(ir.AddStream{
				Fd:         "1",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
			g.ins(ir.AddStream{
				Fd:         "2",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "&>>":
			g.ins(ir.OpenAppendableStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
			})
			g.ins(ir.AddStream{
				Fd:         "1",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
			g.ins(ir.AddStream{
				Fd:         "2",
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case ">&":
			g.ins(ir.DuplicateStream{
				Old: redirection.Src,
				New: g.handleExpression(redirection.Dst),
			})
		case "<&":
			g.ins(ir.Set{
				Name: fmt.Sprintf("%s.Stdin", name),
				Value: ir.NewStreamFromFD{
					Fd: g.handleExpression(redirection.Dst),
				},
			})
		case "<":
			g.ins(ir.OpenReadableStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
			})
			g.ins(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "<<<":
			g.ins(ir.Declare{
				Name: fmt.Sprintf("%s_file_%d", name, i),
				Value: ir.NewStringStream{
					Target: g.handleExpression(redirection.Dst),
				},
			})
			g.ins(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		case "<>":
			g.ins(ir.OpenReadWritableStream{
				Name:   fmt.Sprintf("%s_file_%d", name, i),
				Target: g.handleExpression(redirection.Dst),
			})
			g.ins(ir.AddStream{
				Fd:         redirection.Src,
				StreamName: fmt.Sprintf("%s_file_%d", name, i),
			})
		}
	}

	g.ins(ir.GetStream{Fd: ir.String("0"), StreamName: fmt.Sprintf("%s_stdin", name)})
	g.ins(ir.Set{Name: fmt.Sprintf("%s.Stdin", name), Value: ir.Literal(fmt.Sprintf("%s_stdin", name))})
	g.ins(ir.GetStream{Fd: ir.String("1"), StreamName: fmt.Sprintf("%s_stdout", name)})
	g.ins(ir.Set{Name: fmt.Sprintf("%s.Stdout", name), Value: ir.Literal(fmt.Sprintf("%s_stdout", name))})
	g.ins(ir.GetStream{Fd: ir.String("2"), StreamName: fmt.Sprintf("%s_stderr", name)})
	g.ins(ir.Set{Name: fmt.Sprintf("%s.Stderr", name), Value: ir.Literal(fmt.Sprintf("%s_stderr", name))})
}
