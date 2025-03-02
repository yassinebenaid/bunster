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
	}
}

func (g *generator) handleTestBinary(buf *InstructionBuffer, test ast.Binary) {
	switch test.Operator {
	case "=":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.Compare{Left: l, Operator: "==", Right: r})
	case "==":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.Compare{Left: l, Operator: "==", Right: r})
	case "!=", "<", ">":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.Compare{Left: l, Operator: test.Operator, Right: r})
	case "-eq":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.CompareArithmetics{Left: l, Operator: "==", Right: r})
	case "-ne":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.CompareArithmetics{Left: l, Operator: "!=", Right: r})
	case "-lt":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.CompareArithmetics{Left: l, Operator: "<", Right: r})
	case "-le":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.CompareArithmetics{Left: l, Operator: "<=", Right: r})
	case "-gt":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.CompareArithmetics{Left: l, Operator: ">", Right: r})
	case "-ge":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.CompareArithmetics{Left: l, Operator: ">=", Right: r})
	case "-ef":
		l := g.handleExpression(buf, test.Left)
		r := g.handleExpression(buf, test.Right)

		buf.add(ir.TestFilesHaveSameDevAndInoNumbers{File1: l, File2: r})
	default:
		panic("we do not support the binary operator: " + test.Operator)
	}
}
