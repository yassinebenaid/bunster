package analyser

import (
	"github.com/yassinebenaid/bunster/ast"
)

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
	*b = ast.Break{
		BreakPoint: a.breakpoints,
	}

	var withinLoop bool
	var last int

loop:
	for i := len(a.stack) - 1; i >= 0; i-- {
		switch v := a.stack[i].(type) {
		case *ast.Loop:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinLoop = true
			break loop
		case ast.RangeLoop, ast.For:
			withinLoop = true
			break loop
		case *ast.If:
			v.BreakPoints.Add(a.breakpoints, ast.RETURN)
		case *ast.Break:
			v.Type = ast.RETURN
		case ast.List:
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
	case *ast.Break:
		v.Type = ast.BREAK
	}
}
