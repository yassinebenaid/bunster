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
}

func (a *analyser) analyse() {
	for _, statement := range a.script {
		a.analyseStatement(statement)
	}
}

func (a *analyser) analyseStatement(s ast.Statement) {
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
		for _, ass := range v.Env {
			a.analyseExpression(ass.Value)
		}
	case ast.List, ast.If, ast.SubShell, ast.Group, ast.ParameterAssignement:
	case ast.Pipeline:
		a.analysePipeline(v)
	default:
		a.report(fmt.Sprintf("Unsupported statement type: %T", v))
	}
}

func (a *analyser) analyseExpression(s ast.Expression) {
	switch v := s.(type) {
	case ast.Word, ast.Var, ast.CommandSubstitution, ast.QuotedString, ast.UnquotedString, ast.SpecialVar, ast.Number:
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
		}
	}
}

func (a *analyser) report(err string) {
	a.errors = append(a.errors, SemanticError{
		Err: err,
	})
}
