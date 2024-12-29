package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleCommandSubstitution(statements ast.CommandSubstitution) ir.Instruction {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneFDT{})
	cmdbuf.add(ir.CloneShell{})
	cmdbuf.add(ir.Declare{
		Name:  "stdout",
		Value: ir.NewBuffer{Value: ir.String("")},
	})
	cmdbuf.add(ir.AddStream{Fd: "1", StreamName: "stdout"})

	for _, statement := range statements {
		g.generate(&cmdbuf, statement, nil)
	}

	cmdbuf.add(ir.Literal("return stdout.String(true)"))

	return ir.ExpressionClosure(cmdbuf)
}
