package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleParameterExpansion(buf *InstructionBuffer, expression ast.Expression) ir.Instruction {
	switch v := expression.(type) {
	case ast.VarLength:
		return ir.VarLength{Name: v.Parameter.Name}
	case ast.VarOrDefault:
		return g.handleParameterExpansionVarOrDefault(buf, v)
	case ast.VarOrSet:
		return g.handleParameterExpansionVarOrSet(buf, v)
	default:
		panic(fmt.Sprintf("Unsupported expansion expression: %T", v))
	}
}

func (g *generator) handleParameterExpansionVarOrDefault(buf *InstructionBuffer, expression ast.VarOrDefault) ir.Instruction {
	name := fmt.Sprintf("expr%d", g.expressionsCount)
	buf.add(ir.Declare{Name: name, Value: ir.String("")})

	var def ir.Instruction = ir.String("")
	if expression.Default != nil {
		def = g.handleExpression(buf, expression.Default)
	}

	_if := ir.If{
		Body:      []ir.Instruction{ir.Set{Name: name, Value: ir.ReadVar(expression.Parameter.Name)}},
		Alternate: []ir.Instruction{ir.Set{Name: name, Value: def}},
	}

	if expression.UnsetOnly {
		_if.Condition = ir.TestVarIsSet{Name: ir.String(expression.Parameter.Name)}
	} else {
		_if.Condition = ir.TestAgainsStringLength{String: ir.ReadVar(expression.Parameter.Name)}
	}

	buf.add(_if)
	return ir.Literal(name)
}

func (g *generator) handleParameterExpansionVarOrSet(buf *InstructionBuffer, expression ast.VarOrSet) ir.Instruction {
	var def ir.Instruction = ir.String("")
	if expression.Default != nil {
		def = g.handleExpression(buf, expression.Default)
	}

	_if := ir.If{
		Not: true,
		Body: []ir.Instruction{
			ir.SetVar{Key: expression.Parameter.Name, Value: def},
		},
	}

	if expression.UnsetOnly {
		_if.Condition = ir.TestVarIsSet{Name: ir.String(expression.Parameter.Name)}
	} else {
		_if.Condition = ir.TestAgainsStringLength{String: ir.ReadVar(expression.Parameter.Name)}
	}

	buf.add(_if)
	return ir.ReadVar(expression.Parameter.Name)
}
