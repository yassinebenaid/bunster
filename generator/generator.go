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

	g.ins(ir.Declare{
		Name: fmt.Sprintf("cmd_%d", id),
		Value: ir.InitCommand{
			Name: fmt.Sprintf("cmd_%d_name", id),
			Args: fmt.Sprintf("cmd_%d_args", id),
		},
	})

	g.ins(ir.Set{
		Name:  fmt.Sprintf("cmd_%d.Stdin", id),
		Value: ir.Literal("os.Stdin"),
	})
	g.ins(ir.Set{
		Name:  fmt.Sprintf("cmd_%d.Stdout", id),
		Value: ir.Literal("os.Stdout"),
	})
	g.ins(ir.Set{
		Name:  fmt.Sprintf("cmd_%d.Stderr", id),
		Value: ir.Literal("os.Stderr"),
	})

	g.ins(ir.RunCommanOrFail{
		Name: fmt.Sprintf("cmd_%d", id),
	})
}

func (g *generator) handleExpression(expression ast.Expression) ir.Instruction {
	switch v := expression.(type) {
	case ast.Word:
		return ir.String(v)
	}

	return nil
}
