package analyser

import (
	"github.com/yassinebenaid/bunster/ast"
)

func (a *analyser) analyseFor(loop *ast.For) {
	for _, expr := range loop.Head.Init {
		a.analyseArithmeticExpression(expr)
	}
	for _, expr := range loop.Head.Test {
		a.analyseArithmeticExpression(expr)
	}
	for _, expr := range loop.Head.Update {
		a.analyseArithmeticExpression(expr)
	}
	for _, s := range loop.Body {
		a.analyseStatement(s)
	}
	for _, r := range loop.Redirections {
		if r.Dst != nil {
			a.analyseExpression(r.Dst)
		}
	}
}

func (a *analyser) analyseRangeLoop(loop *ast.RangeLoop) {
	for _, expr := range loop.Operands {
		a.analyseExpression(expr)
	}
	for _, s := range loop.Body {
		a.analyseStatement(s)
	}
	for _, r := range loop.Redirections {
		if r.Dst != nil {
			a.analyseExpression(r.Dst)
		}
	}
}

func (a *analyser) analyseLoop(loop *ast.Loop) {
	for _, s := range loop.Head {
		a.analyseStatement(s)
	}
	for _, s := range loop.Body {
		a.analyseStatement(s)
	}
	for _, r := range loop.Redirections {
		if r.Dst != nil {
			a.analyseExpression(r.Dst)
		}
	}
}

func (a *analyser) analyseBreak(b *ast.Break) {
	a.breakpoints++
	b.BreakPoint = a.breakpoints

	var withinLoop bool
	var last int

loop:
	for i := len(a.stack) - 1; i >= 0; i-- {
		switch v := a.stack[i].(type) {
		case *ast.Loop:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinLoop = true
			break loop
		case *ast.RangeLoop:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinLoop = true
			break loop
		case *ast.For:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinLoop = true
			break loop
		case *ast.If:
			v.BreakPoints.Add(a.breakpoints, ast.RETURN)
		case *ast.Break:
			v.Type = ast.RETURN
		case ast.List:
		default:
			break loop
		}
		last = i
	}

	if !withinLoop {
		a.report(Error{Msg: "the `break` keyword cannot be used here"})
		return
	}

	switch v := a.stack[last].(type) {
	case *ast.If:
		v.BreakPoints.Add(a.breakpoints, ast.BREAK)
	case *ast.Break, ast.List:
		b.Type = ast.CONTINUE
	}
}

func (a *analyser) analyseContinue(b *ast.Continue) {
	a.breakpoints++
	b.BreakPoint = a.breakpoints

	var withinLoop bool
	var last int

loop:
	for i := len(a.stack) - 1; i >= 0; i-- {
		switch v := a.stack[i].(type) {
		case *ast.Loop:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinLoop = true
			break loop
		case *ast.RangeLoop:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinLoop = true
			break loop
		case *ast.For:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinLoop = true
			break loop
		case *ast.If:
			v.BreakPoints.Add(a.breakpoints, ast.RETURN)
		case *ast.Continue:
			v.Type = ast.RETURN
		case ast.List:
		default:
			break loop
		}
		last = i
	}

	if !withinLoop {
		a.report(Error{Msg: "the `continue` keyword cannot be used here"})
		return
	}

	switch v := a.stack[last].(type) {
	case *ast.If:
		v.BreakPoints.Add(a.breakpoints, ast.CONTINUE)
	case *ast.Continue, ast.List:
		b.Type = ast.CONTINUE
	}
}
