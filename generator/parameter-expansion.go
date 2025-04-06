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
	case ast.VarOrDefault:
		name := fmt.Sprintf("expr%d", g.expressionsCount)
		buf.add(ir.Declare{Name: name, Value: ir.String("")})

		var def ir.Instruction = ir.String("")
		if v.Default != nil {
			def = g.handleExpression(buf, v.Default)
		}

		_if := ir.If{
			Body:      []ir.Instruction{ir.Set{Name: name, Value: ir.ReadVar(v.Parameter.Name)}},
			Alternate: []ir.Instruction{ir.Set{Name: name, Value: def}},
		}

		if v.UnsetOnly {
			_if.Condition = ir.TestVarIsSet{Name: ir.String(v.Parameter.Name)}
		} else {
			_if.Condition = ir.TestAgainsStringLength{String: ir.ReadVar(v.Parameter.Name)}
		}

		buf.add(_if)
		return ir.Literal(name)
	default:
		panic(fmt.Sprintf("Unsupported expansion expression: %T", v))
	}
}
