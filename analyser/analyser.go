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
	case ast.Pipeline:
		a.analysePipeline(v)
	}
}

type SemanticError struct {
	Line, Position int
	Err            analysisError
}

func (s SemanticError) Error() string {
	return fmt.Sprintf("semantic error: %s. (line: %d, column: %d)", s.Err, s.Line, s.Position)
}

type analysisError string

var (
	ErrorUsingShellParametersWithinPipeline analysisError = "using shell parameters within a pipeline has no effect and is invalid. only statements that perform IO are allowed within pipelines"
)

func (a *analyser) analysePipeline(p ast.Pipeline) {
	for _, cmd := range p {
		switch cmd.Command.(type) {
		case ast.ParameterAssignement:
			a.report(ErrorUsingShellParametersWithinPipeline)
		}
	}
}

func (a *analyser) report(err analysisError) {
	a.errors = append(a.errors, SemanticError{
		Err: err,
	})
}
