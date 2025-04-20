package analyser

import (
	"fmt"
	"path/filepath"

	"github.com/yassinebenaid/bunster/ast"
)

type Error struct {
	File           string
	Line, Position int
	Msg            string
}

func (s Error) Error() string {
	s.File = "main.sh"

	return fmt.Sprintf("%s(%d:%d): semantic error: %s.", s.File, s.Line, s.Position, s.Msg)
}

func (a *analyser) report(err Error) {
	a.errors = append(a.errors, err)
}

func Analyse(s ast.Script, main bool) error {
	a := analyser{script: s}
	a.analyse(main)
	if len(a.errors) == 0 {
		return nil
	}
	return a.errors[0]
}

type analyser struct {
	script      ast.Script
	errors      []error
	stack       []ast.Statement
	breakpoints int
}

func (a *analyser) analyse(main bool) {
	for _, statement := range a.script {
		if !main {
			_, ok := statement.(*ast.Function)
			if !ok {
				a.report(Error{Msg: "only functions can exist in global scope"})
				return
			}
		}
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
	case *ast.If:
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
		a.analyseLocalParameterAssignement(v)
	case *ast.For:
		a.analyseFor(v)
	case *ast.RangeLoop:
		a.analyseRangeLoop(v)
	case *ast.Loop:
		a.analyseLoop(v)
	case *ast.Break:
		a.analyseBreak(v)
	case *ast.Continue:
		a.analyseContinue(v)
	case ast.Pipeline:
		for _, cmd := range v {
			a.analyseStatement(cmd.Command)
		}
	case ast.BackgroundConstruction:
		a.analyseStatement(v.Statement)
	case ast.Wait, ast.Exit:
		//TODO: ensure 'wait' is not invokes when no commands are put in background.
	case ast.InvertExitCode:
		a.analyseStatement(v.Statement)
	case *ast.Function:
		a.analyseFunction(v)
	case *ast.Return:
		a.analyseReturn(v)
	case ast.Test:
		a.analyseExpression(v.Expr)

		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case ast.Embed:
		if len(a.stack) != 1 {
			a.report(Error{Msg: "using '@embed' directive is only valid in global scope"})
		}

		for _, path := range v {
			if !filepath.IsLocal(path) {
				a.report(Error{Msg: fmt.Sprintf("the path %q cannot be embeded because it is not local to the module", path)})
			}
		}
	case ast.Defer:
		a.analyseStatement(v.Command)
	case ast.ArithmeticCommand:
		for _, expr := range v.Arithmetic {
			a.analyseArithmeticExpression(expr)
		}
		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case *ast.Case:
		a.analyseExpression(v.Word)

		for _, _case := range v.Cases {
			for _, pattern := range _case.Patterns {
				a.analyseExpression(pattern)
			}
			for _, s := range _case.Body {
				a.analyseStatement(s)
			}
		}

		for _, r := range v.Redirections {
			if r.Dst != nil {
				a.analyseExpression(r.Dst)
			}
		}
	case ast.Unset:
		for _, name := range v.Names {
			a.analyseExpression(name)
		}
	default:
		a.report(Error{Msg: fmt.Sprintf("Unsupported statement type: %T", v)})
	}

	a.stack = a.stack[:len(a.stack)-1]
}

func (a *analyser) analyseExpression(s ast.Expression) {
	switch v := s.(type) {
	case ast.Word, ast.Var, ast.SpecialVar, ast.Number:
	case ast.VarLength:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
	case ast.VarOrDefault:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
		if v.Default != nil {
			a.analyseExpression(v.Default)
		}
	case ast.VarOrSet:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
		if v.Default != nil {
			a.analyseExpression(v.Default)
		}
	case ast.ChangeCase:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
		if v.Pattern != nil {
			a.analyseExpression(v.Pattern)
		}
	case ast.MatchAndRemove:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
		if v.Pattern != nil {
			a.analyseExpression(v.Pattern)
		}
	case ast.MatchAndReplace:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
		if v.Pattern != nil {
			a.analyseExpression(v.Pattern)
		}
	case ast.CheckAndUse:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
		if v.Value != nil {
			a.analyseExpression(v.Value)
		}
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
			a.report(Error{Msg: fmt.Sprintf("Unsupported test operator: %s", v.Operator)})
		}
		a.analyseExpression(v.Left)
		a.analyseExpression(v.Right)
	case ast.Unary:
		a.analyseExpression(v.Operand)
	case ast.Negation:
		a.analyseExpression(v.Operand)
	case ast.Arithmetic:
		for _, expr := range v {
			a.analyseArithmeticExpression(expr)
		}
	case ast.Slice:
		if v.Parameter.Index != nil {
			a.analyseExpression(v.Parameter.Index)
		}
		a.analyseExpression(v.Offset)
		if v.Length != nil {
			a.analyseExpression(v.Length)
		}
	case ast.ArrayLiteral:
		for _, exp := range v {
			a.analyseExpression(exp)
		}
	case ast.ParameterExpansion:
		a.analyseExpression(v.Index)
	default:
		a.report(Error{Msg: fmt.Sprintf("Unsupported expression type: %T", v)})
	}
}

func (a *analyser) analyseArithmeticExpression(s ast.Expression) {
	switch v := s.(type) {
	case ast.PostIncDecArithmetic, ast.PreIncDecArithmetic, ast.Number, ast.Var:
	case ast.Unary:
		a.analyseArithmeticExpression(v.Operand)
	case ast.Negation:
		a.analyseArithmeticExpression(v.Operand)
	case ast.BitFlip:
		a.analyseArithmeticExpression(v.Operand)
	case ast.Binary:
		a.analyseArithmeticExpression(v.Left)
		a.analyseArithmeticExpression(v.Right)
	case ast.Conditional:
		a.analyseArithmeticExpression(v.Test)
		a.analyseArithmeticExpression(v.Body)
		a.analyseArithmeticExpression(v.Alternate)
	default:
		a.report(Error{Msg: fmt.Sprintf("Unsupported arithmetic expression type: %T", v)})
	}
}
