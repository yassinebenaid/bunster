package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleCommandSubstitution(statements ast.CommandSubstitution) ir.Instruction {
	var cmdbuf InstructionBuffer

	// var b bytes.Buffer

	cmdbuf.add(ir.CloneFDT{})

	return ir.ExpressionClosure(cmdbuf)
}
