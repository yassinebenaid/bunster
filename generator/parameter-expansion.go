package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

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

func (g *generator) handleParameterExpansionCheckAndUse(buf *InstructionBuffer, expression ast.CheckAndUse) ir.Instruction {
	name := fmt.Sprintf("expr%d", g.expressionsCount)
	buf.add(ir.Declare{Name: name, Value: ir.String("")})

	var value ir.Instruction = ir.String("")
	if expression.Value != nil {
		value = g.handleExpression(buf, expression.Value)
	}

	_if := ir.If{
		Body: []ir.Instruction{
			ir.Set{Name: name, Value: value},
		},
	}

	if expression.UnsetOnly {
		_if.Condition = ir.TestVarIsSet{Name: ir.String(expression.Parameter.Name)}
	} else {
		_if.Condition = ir.TestAgainsStringLength{String: ir.ReadVar(expression.Parameter.Name)}
	}

	buf.add(_if)
	return ir.Literal(name)
}

func (g *generator) handleParameterExpansionSlice(buf *InstructionBuffer, expression ast.Slice) ir.Instruction {
	offset := fmt.Sprintf("offset%d", g.expressionsCount)
	length := fmt.Sprintf("length%d", g.expressionsCount)

	buf.add(ir.Declare{Name: offset, Value: ir.Literal("0")})
	buf.add(ir.Declare{Name: length, Value: ir.Literal("int(^uint32(0))")})

	for _, arithmetic := range expression.Offset {
		buf.add(ir.Set{Name: offset, Value: g.handleArithmeticExpression(buf, arithmetic)})
	}

	for _, arithmetic := range expression.Length {
		buf.add(ir.Set{Name: length, Value: g.handleArithmeticExpression(buf, arithmetic)})
	}

	return ir.Substring{
		String: ir.ReadVar(expression.Parameter.Name),
		Offset: ir.Literal(offset),
		Length: ir.Literal(length),
	}
}
