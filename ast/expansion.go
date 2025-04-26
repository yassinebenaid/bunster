package ast

type Parameter interface {
	Expression
	param()
}

type ArrayAccess struct {
	Name  string
	Index Expression
}

type VarOrDefault struct {
	Parameter Parameter
	Default   Expression
	UnsetOnly bool
}

type VarOrSet struct {
	Parameter Parameter
	Default   Expression
	UnsetOnly bool
}

type VarOrFail struct {
	Parameter Parameter
	Index     Expression
	Error     Expression
}

type CheckAndUse struct {
	Parameter Parameter
	Value     Expression
	UnsetOnly bool
}

type ChangeCase struct {
	Parameter Parameter
	Operator  string
	Pattern   Expression
}

type VarLength struct {
	Parameter Parameter
}

type MatchAndRemove struct {
	Parameter Parameter
	Operator  string
	Pattern   Expression
}

type MatchAndReplace struct {
	Parameter Parameter
	Operator  string
	Pattern   Expression
	Value     Expression
}

type Transform struct {
	Parameter Parameter
	Index     Expression
	Operator  string
}

type Slice struct {
	Parameter Parameter
	Offset    Arithmetic
	Length    Arithmetic
}

func (ArrayAccess) node()     {}
func (VarOrDefault) node()    {}
func (VarOrSet) node()        {}
func (VarOrFail) node()       {}
func (CheckAndUse) node()     {}
func (ChangeCase) node()      {}
func (VarLength) node()       {}
func (MatchAndRemove) node()  {}
func (MatchAndReplace) node() {}
func (Transform) node()       {}
func (Slice) node()           {}

// Expressions
func (ArrayAccess) expr()     {}
func (VarOrDefault) expr()    {}
func (VarOrSet) expr()        {}
func (VarOrFail) expr()       {}
func (CheckAndUse) expr()     {}
func (ChangeCase) expr()      {}
func (VarLength) expr()       {}
func (MatchAndRemove) expr()  {}
func (MatchAndReplace) expr() {}
func (Transform) expr()       {}
func (Slice) expr()           {}

func (v ArrayAccess) string() string   { return string(v.Name) }
func (VarOrDefault) string() string    { return "" }
func (VarOrSet) string() string        { return "" }
func (VarOrFail) string() string       { return "" }
func (CheckAndUse) string() string     { return "" }
func (ChangeCase) string() string      { return "" }
func (VarLength) string() string       { return "" }
func (MatchAndRemove) string() string  { return "" }
func (MatchAndReplace) string() string { return "" }
func (Transform) string() string       { return "" }
func (Slice) string() string           { return "" }

func (Var) param()              {}
func (ArrayAccess) param()      {}
func (SpecialVar) param()       {}
func (PositionalSpread) param() {}
