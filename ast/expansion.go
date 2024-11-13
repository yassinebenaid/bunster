package ast

type Var string

type VarOrDefault struct {
	Name         string
	Default      Expression
	CheckForNull bool
}

type VarOrSet struct {
	Name    string
	Default Expression
}

type VarOrFail struct {
	Name  string
	Error Expression
}

type CheckAndUse struct {
	Name  string
	Value Expression
}

type ChangeCase struct {
	Name     string
	Operator string
	Pattern  Expression
}

type VarCount string

type MatchAndRemove struct {
	Name     string
	Operator string
	Pattern  Expression
}

type MatchAndReplace struct {
	Name     string
	Operator string
	Pattern  Expression
	Value    Expression
}

type Transform struct {
	Name     string
	Operator string
}

func (Var) node()             {}
func (VarOrDefault) node()    {}
func (VarOrSet) node()        {}
func (VarOrFail) node()       {}
func (CheckAndUse) node()     {}
func (ChangeCase) node()      {}
func (VarCount) node()        {}
func (MatchAndRemove) node()  {}
func (MatchAndReplace) node() {}
func (Transform) node()       {}

// Expressions
func (Var) expr()             {}
func (VarOrDefault) expr()    {}
func (VarOrSet) expr()        {}
func (VarOrFail) expr()       {}
func (CheckAndUse) expr()     {}
func (ChangeCase) expr()      {}
func (VarCount) expr()        {}
func (MatchAndRemove) expr()  {}
func (MatchAndReplace) expr() {}
func (Transform) expr()       {}

func (v Var) string() string           { return string(v) }
func (VarOrDefault) string() string    { return "" }
func (VarOrSet) string() string        { return "" }
func (VarOrFail) string() string       { return "" }
func (CheckAndUse) string() string     { return "" }
func (ChangeCase) string() string      { return "" }
func (VarCount) string() string        { return "" }
func (MatchAndRemove) string() string  { return "" }
func (MatchAndReplace) string() string { return "" }
func (Transform) string() string       { return "" }
