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
	case ast.Loop:
		for _, s := range v.Head {
			if _, ok := s.(ast.Break); ok {
				a.report(fmt.Sprintf("The `break` keyword cannot be used here"))
			}
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
			case ast.Loop:
				withinLoop = true
				break loop
			case ast.If, ast.Break:
			default:
				a.report(fmt.Sprintf("The `break` keyword cannot be used here"))
			}
		}
		if !withinLoop {
			a.report(fmt.Sprintf("The `break` keyword cannot be used here"))
		}
	case ast.Pipeline:
		a.analysePipeline(v)
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
)

func (a *analyser) analysePipeline(p ast.Pipeline) {
	for _, cmd := range p {
		switch cmd.Command.(type) {
		case ast.ParameterAssignement:
			a.report(ErrorUsingShellParametersWithinPipeline)
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
