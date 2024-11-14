package ast

type Parameter interface {
	Expression
	param()
}

type Var string

type VarOrDefault struct {
	Parameter
	Default      Expression
	CheckForNull bool
}

type VarOrSet struct {
	Parameter
	Default Expression
}

type VarOrFail struct {
	Parameter
	Error Expression
}

type CheckAndUse struct {
	Parameter
	Value Expression
}

type ChangeCase struct {
	Parameter
	Operator string
	Pattern  Expression
}

type VarCount struct {
	Parameter
}

type MatchAndRemove struct {
	Parameter
	Operator string
	Pattern  Expression
}

type MatchAndReplace struct {
	Parameter
	Operator string
	Pattern  Expression
	Value    Expression
}

type Transform struct {
	Parameter
	Operator string
}

type Slice struct {
	Parameter
	Offset Arithmetic
	Length Arithmetic
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
func (Slice) node()           {}

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
func (Slice) expr()           {}

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
func (Slice) string() string           { return "" }

func (v Var) param() {}
