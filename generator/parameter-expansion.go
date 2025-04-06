package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleParameterExpansion(buf *InstructionBuffer, expression ast.Expression) ir.Instruction {
	switch v := expression.(type) {
	case ast.VarLength:
		return ir.VarLength{
			Name: v.Parameter.Name,
		}
	default:
		panic(fmt.Sprintf("Unsupported expansion expression: %T", v))
	}
}
