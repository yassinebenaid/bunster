package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleCommandSubstitution(statements ast.CommandSubstitution) ir.Instruction {
	var cmdbuf InstructionBuffer

	// var b bytes.Buffer

	cmdbuf.add(ir.CloneFDT{})
	cmdbuf.add(ir.Declare{
		Name:  "stdout",
		Value: ir.NewBuffer{Value: ir.String("")},
	})

	cmdbuf.add(ir.Literal("return stdout.String()"))

	return ir.ExpressionClosure(cmdbuf)
}
