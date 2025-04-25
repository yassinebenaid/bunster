package analyser

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
)

func (a *analyser) analyseParameterAssignement(pa ast.ParameterAssignement) {
	for _, pa := range pa {
		if pa.Value != nil {
			a.analyseExpression(pa.Value)
		}
	}
}

func (a *analyser) analyseExportParameterAssignement(pa ast.ExportParameterAssignement) {
	for _, pa := range pa {
		if pa.Value != nil {
			a.analyseExpression(pa.Value)
		}
	}
}

func (a *analyser) analyseLocalParameterAssignement(local ast.LocalParameterAssignement) {
	var withinFunction bool
loop:
	for i := len(a.stack) - 1; i >= 0; i-- {
		switch a.stack[i].(type) {
		case *ast.Function:
			withinFunction = true
			break loop
		}
	}
	if !withinFunction {
		a.report(Error{Msg: "the `local` keyword cannot be used outside functions"})
	}

	for _, pa := range local {
		if pa.Value != nil {
			a.analyseExpression(pa.Value)
		}
	}
}

func (a *analyser) analyseParameter(param ast.Parameter) {
	switch v := param.(type) {
	case ast.Var, ast.SpecialVar:
	case ast.ArrayAccess:
		a.analyseExpression(v.Index)
	default:
		a.report(Error{Msg: fmt.Sprintf("unknown parameter kind: %T", param)})
	}
}
