package token

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	_ TokenType = iota
	NL
	IDENT
	ILLIGAL
	EOF
)

var literals = map[TokenType]string{
	NL:      "NEWLINE",
	IDENT:   "IDENTIFIER",
	ILLIGAL: "ILLIGAL",
	EOF:     "EOF",
}

func (t TokenType) String() string {
	return literals[t]
}
