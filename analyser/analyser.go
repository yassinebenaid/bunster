package analyser

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
)

func Analyse(s ast.Script) error {
	a := analyser{script: s}
	a.analyse()
	if len(a.errors) == 0 {
		return nil
	}
	return a.errors[0]
}

type analyser struct {
	script ast.Script
	errors []error
	stack  []ast.Statement
}

func (a *analyser) analyse() {
	for _, statement := range a.script {
		a.analyseStatement(statement)
	}
}

func (a *analyser) analyseStatement(s ast.Statement) {
	a.stack = append(a.stack, s)

	switch v := s.(type) {
	case ast.Command:
		a.analyseExpression(v.Name)
		for _, arg := range v.Args {
			a.analyseExpression(arg)
		}
		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
		for _, env := range v.Env {
			if env.Value != nil {
				a.analyseExpression(env.Value)
			}
		}
	case ast.List:
		a.analyseStatement(v.Left)
		a.analyseStatement(v.Right)
	case ast.If:
		for _, s := range v.Head {
			a.analyseStatement(s)
		}
		for _, s := range v.Body {
			a.analyseStatement(s)
		}
		for _, elif := range v.Elifs {
			for _, s := range elif.Head {
				a.analyseStatement(s)
			}
			for _, s := range elif.Body {
				a.analyseStatement(s)
			}
		}
		for _, s := range v.Alternate {
			a.analyseStatement(s)
		}
		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case ast.SubShell:
		for _, s := range v.Body {
			a.analyseStatement(s)
		}
		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case ast.Group:
		for _, s := range v.Body {
			a.analyseStatement(s)
		}
		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case ast.ParameterAssignement:
		for _, pa := range v {
			if pa.Value != nil {
				a.analyseExpression(pa.Value)
			}
		}
	case ast.ExportParameterAssignement:
		for _, pa := range v {
			if pa.Value != nil {
				a.analyseExpression(pa.Value)
			}
		}
	case ast.LocalParameterAssignement:
		var withinFunction bool
	funcLoop:
		for i := len(a.stack) - 1; i >= 0; i-- {
			switch a.stack[i].(type) {
			case ast.Function:
				withinFunction = true
				break funcLoop
			}
		}
		if !withinFunction {
			a.report("The `local` keyword cannot be used outside functions")
		}

		for _, pa := range v {
			if pa.Value != nil {
				a.analyseExpression(pa.Value)
			}
		}
	case ast.Loop:
		for _, s := range v.Head {
			a.analyseStatement(s)
		}
		for _, s := range v.Body {
			a.analyseStatement(s)
		}
		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case ast.Break:
		var withinLoop bool
	loop:
		for i := len(a.stack) - 1; i >= 0; i-- {
			switch a.stack[i].(type) {
			case ast.Loop, ast.RangeLoop:
				withinLoop = true
				break loop
			case ast.List, ast.Break:
			default:
				a.report("The `break` keyword cannot be used here")
			}
		}
		if !withinLoop {
			a.report("The `break` keyword cannot be used here")
		}
	case ast.Continue:
		var withinLoop bool
	loop2:
		for i := len(a.stack) - 1; i >= 0; i-- {
			switch a.stack[i].(type) {
			case ast.Loop, ast.RangeLoop:
				withinLoop = true
				break loop2
			case ast.List, ast.Continue:
			default:
				a.report("The `continue` keyword cannot be used here")
			}
		}
		if !withinLoop {
			a.report("The `continue` keyword cannot be used here")
		}
	case ast.Pipeline:
		a.analysePipeline(v)
	case ast.BackgroundConstruction:
		a.analyseStatement(v.Statement)
	case ast.Wait:
		//TODO: ensure 'wait' is not invokes when no commands are put in background.
	case ast.InvertExitCode:
		a.analyseStatement(v.Statement)
	case ast.Function:
		a.analyseStatement(v.Command)
	case ast.RangeLoop:
		for _, expr := range v.Operands {
			a.analyseExpression(expr)
		}
		for _, s := range v.Body {
			a.analyseStatement(s)
		}
		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case ast.Test:
		a.analyseExpression(v.Expr)

		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	default:
		a.report(fmt.Sprintf("Unsupported statement type: %T", v))
	}

	a.stack = a.stack[:len(a.stack)-1]
}

func (a *analyser) analyseExpression(s ast.Expression) {
	switch v := s.(type) {
	case ast.Word, ast.Var, ast.SpecialVar, ast.Number:
	case ast.CommandSubstitution:
		for _, s := range v {
			a.analyseStatement(s)
		}
	case ast.UnquotedString:
		for _, exp := range v {
			a.analyseExpression(exp)
		}
	case ast.QuotedString:
		for _, exp := range v {
			a.analyseExpression(exp)
		}
	case ast.Binary:
		if v.Operator == "=~" {
			a.report(fmt.Sprintf("Unsupported test operator: %s", v.Operator))
		}
		a.analyseExpression(v.Left)
		a.analyseExpression(v.Right)
	default:
		a.report(fmt.Sprintf("Unsupported statement type: %T", v))
	}
}

type SemanticError struct {
	Line, Position int
	Err            string
}

func (s SemanticError) Error() string {
	return fmt.Sprintf("semantic error: %s. (line: %d, column: %d)", s.Err, s.Line, s.Position)
}

var (
	ErrorUsingShellParametersWithinPipeline = "using shell parameters within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines"
	ErrorUsingWaitWithinPipeline            = "using 'wait' command within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines"
	ErrorUsingLocalWithinPipeline           = "using 'local' command within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines"
	ErrorUsingExportWithinPipeline          = "using 'export' command within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines"
)

func (a *analyser) analysePipeline(p ast.Pipeline) {
	for _, cmd := range p {
		switch cmd.Command.(type) {
		case ast.ParameterAssignement:
			a.report(ErrorUsingShellParametersWithinPipeline)
		case ast.Wait:
			a.report(ErrorUsingWaitWithinPipeline)
		case ast.LocalParameterAssignement:
			a.report(ErrorUsingLocalWithinPipeline)
		case ast.ExportParameterAssignement:
			a.report(ErrorUsingExportWithinPipeline)
		default:
			a.analyseStatement(cmd.Command)
		}
	}
}

func (a *analyser) report(err string) {
	a.errors = append(a.errors, SemanticError{
		Err: err,
	})
}
