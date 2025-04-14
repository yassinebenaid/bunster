package analyser

import "github.com/yassinebenaid/bunster/ast"

func (a *analyser) analyseFunction(fn *ast.Function) {
	for _, s := range fn.Body {
		a.analyseStatement(s)
	}
	for _, r := range fn.Redirections {
		if r.Dst != nil {
			a.analyseExpression(r.Dst)
		}
	}
}

func (a *analyser) analyseReturn(b *ast.Return) {
	a.breakpoints++
	b.BreakPoint = a.breakpoints

	var withinFunction bool

loop:
	for i := len(a.stack) - 1; i >= 0; i-- {
		switch v := a.stack[i].(type) {
		case *ast.Function:
			v.BreakPoints.Add(a.breakpoints, ast.DECLARE)
			withinFunction = true
			break loop
		case *ast.If:
			v.BreakPoints.Add(a.breakpoints, ast.RETURN)
		case *ast.Case:
			v.BreakPoints.Add(a.breakpoints, ast.RETURN)
		case *ast.Return:
		default:
			break loop
		}
	}

	if !withinFunction {
		a.report(Error{Msg: "the `return` keyword cannot be used here"})
	}
}
