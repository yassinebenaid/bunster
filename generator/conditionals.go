package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleTest(buf *InstructionBuffer, test ast.Test, ctx *context) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{DeferDestroy: ctx.pipe == nil})
	g.handleRedirections(&cmdbuf, test.Redirections, ctx)

	g.handleTestExpression(&cmdbuf, test.Expr)

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleTestExpression(buf *InstructionBuffer, test ast.Expression) {
	switch v := test.(type) {
	case ast.Binary:
		g.handleTestBinary(buf, v)
	case ast.Unary:
		g.handleTestUnary(buf, v)
	default:
		buf.add(ir.TestStringIsIsNotZero{String: g.handleExpression(buf, v)})
	}
}

func (g *generator) handleTestBinary(buf *InstructionBuffer, test ast.Binary) {
	left := g.handleExpression(buf, test.Left)
	right := g.handleExpression(buf, test.Right)

	switch test.Operator {
	case "=":
		buf.add(ir.Compare{Left: left, Operator: "==", Right: right})
	case "==":
		buf.add(ir.Compare{Left: left, Operator: "==", Right: right})
	case "!=", "<", ">":
		buf.add(ir.Compare{Left: left, Operator: test.Operator, Right: right})
	case "-eq":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "==", Right: right})
	case "-ne":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "!=", Right: right})
	case "-lt":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "<", Right: right})
	case "-le":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "<=", Right: right})
	case "-gt":
		buf.add(ir.CompareArithmetics{Left: left, Operator: ">", Right: right})
	case "-ge":
		buf.add(ir.CompareArithmetics{Left: left, Operator: ">=", Right: right})
	case "-ef":
		buf.add(ir.TestFilesHaveSameDevAndInoNumbers{File1: left, File2: right})
	case "-ot":
		buf.add(ir.FileIsOlderThan{File1: left, File2: right})
	case "-nt":
		buf.add(ir.FileIsOlderThan{File1: right, File2: left})
	default:
		panic("we do not support the binary operator: " + test.Operator)
	}
}

func (g *generator) handleTestUnary(buf *InstructionBuffer, test ast.Unary) {
	operand := g.handleExpression(buf, test.Operand)

	switch test.Operator {
	case "-n":
		buf.add(ir.TestStringIsIsNotZero{String: operand})
	default:
		panic("we do not support the binary operator: " + test.Operator)
	}
}
