package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleParameterExpansionVarLength(buf *InstructionBuffer, expression ast.VarLength) ir.Instruction {
	switch v := expression.Parameter.(type) {
	case ast.Var:
		return ir.VarLength{Name: string(v)}
	case ast.ArrayAccess:
		return ir.VarLength{Name: string(v.Name), Index: ir.ParseInt{Value: g.handleExpression(buf, v.Index)}}
	}
	panic("unknown parameter kind")
}

func (g *generator) handleParameterExpansionVarOrDefault(buf *InstructionBuffer, expression ast.VarOrDefault) ir.Instruction {
	name := fmt.Sprintf("expr%d", g.expressionsCount)
	buf.add(ir.Declare{Name: name, Value: ir.String("")})

	var def ir.Instruction = ir.String("")
	if expression.Default != nil {
		def = g.handleExpression(buf, expression.Default)
	}

	_if := ir.If{
		Body:      []ir.Instruction{ir.Set{Name: name, Value: g.handleExpression(buf, expression.Parameter)}},
		Alternate: []ir.Instruction{ir.Set{Name: name, Value: def}},
	}

	if expression.UnsetOnly {
		_if.Condition = ir.TestVarIsSet{Name: ir.String(string(expression.Parameter.(ast.Var)))}
	} else {
		_if.Condition = ir.TestAgainsStringLength{String: g.handleExpression(buf, expression.Parameter)}
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
			ir.SetVar{Key: string(expression.Parameter.(ast.Var)), Value: def},
		},
	}

	if expression.UnsetOnly {
		_if.Condition = ir.TestVarIsSet{Name: ir.String(string(expression.Parameter.(ast.Var)))}
	} else {
		_if.Condition = ir.TestAgainsStringLength{String: ir.ReadVar(string(expression.Parameter.(ast.Var)))}
	}

	buf.add(_if)
	return ir.ReadVar(string(expression.Parameter.(ast.Var)))
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
		_if.Condition = ir.TestVarIsSet{Name: ir.String(string(expression.Parameter.(ast.Var)))}
	} else {
		_if.Condition = ir.TestAgainsStringLength{String: ir.ReadVar(string(expression.Parameter.(ast.Var)))}
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
		String: ir.ReadVar(string(expression.Parameter.(ast.Var))),
		Offset: ir.Literal(offset),
		Length: ir.Literal(length),
	}
}

func (g *generator) handleParameterExpansionChangeCase(buf *InstructionBuffer, expression ast.ChangeCase) ir.Instruction {
	var pattern ir.Instruction = ir.String("?")
	if expression.Pattern != nil {
		pattern = g.handleExpression(buf, expression.Pattern)
	}

	if expression.Operator == "^" || expression.Operator == "^^" {
		return ir.StringToUpperCase{
			String:  ir.ReadVar(string(expression.Parameter.(ast.Var))),
			Pattern: pattern,
			All:     expression.Operator == "^^",
		}
	}

	return ir.StringToLowerCase{
		String:  ir.ReadVar(string(expression.Parameter.(ast.Var))),
		Pattern: pattern,
		All:     expression.Operator == ",,",
	}
}

func (g *generator) handleParameterExpansionMatchAndRemove(buf *InstructionBuffer, expression ast.MatchAndRemove) ir.Instruction {
	var pattern ir.Instruction = ir.String("")
	if expression.Pattern != nil {
		pattern = g.handleExpression(buf, expression.Pattern)
	}

	if expression.Operator == "#" || expression.Operator == "##" {
		return ir.RemoveMatchingPrefix{
			String:  ir.ReadVar(string(expression.Parameter.(ast.Var))),
			Pattern: pattern,
			Longest: expression.Operator == "##",
		}
	}

	return ir.RemoveMatchingSuffix{
		String:  ir.ReadVar(string(expression.Parameter.(ast.Var))),
		Pattern: pattern,
		Longest: expression.Operator == "%%",
	}
}

func (g *generator) handleParameterExpansionMatchAndReplace(buf *InstructionBuffer, expression ast.MatchAndReplace) ir.Instruction {
	var pattern ir.Instruction = ir.String("")
	if expression.Pattern != nil {
		pattern = g.handleExpression(buf, expression.Pattern)
	}

	var repl ir.Instruction = ir.String("")
	if expression.Value != nil {
		repl = g.handleExpression(buf, expression.Value)
	}

	switch expression.Operator {
	case "/", "//":
		return ir.ReplaceMatching{
			String:  ir.ReadVar(string(expression.Parameter.(ast.Var))),
			Pattern: pattern,
			Value:   repl,
			All:     expression.Operator == "//",
		}
	case "/#":
		return ir.ReplaceMatchingPrefix{
			String:  ir.ReadVar(string(expression.Parameter.(ast.Var))),
			Pattern: pattern,
			Value:   repl,
		}
	case "/%":
		return ir.ReplaceMatchingSuffix{
			String:  ir.ReadVar(string(expression.Parameter.(ast.Var))),
			Pattern: pattern,
			Value:   repl,
		}
	}

	return nil
}
