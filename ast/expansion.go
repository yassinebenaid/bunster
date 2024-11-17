package ast

type Var string

type VarOrDefault struct {
	Parameter    string
	Index        Expression
	Default      Expression
	CheckForNull bool
}

type VarOrSet struct {
	Parameter string
	Index     Expression
	Default   Expression
}

type VarOrFail struct {
	Parameter string
	Index     Expression
	Error     Expression
}

type CheckAndUse struct {
	Parameter string
	Index     Expression
	Value     Expression
}

type ChangeCase struct {
	Parameter string
	Index     Expression
	Operator  string
	Pattern   Expression
}

type VarCount struct {
	Parameter string
	Index     Expression
}

type MatchAndRemove struct {
	Parameter string
	Index     Expression
	Operator  string
	Pattern   Expression
}

type MatchAndReplace struct {
	Parameter string
	Index     Expression
	Operator  string
	Pattern   Expression
	Value     Expression
}

type Transform struct {
	Parameter string
	Index     Expression
	Operator  string
}

type Slice struct {
	Parameter string
	Index     Expression
	Offset    Arithmetic
	Length    Arithmetic
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
