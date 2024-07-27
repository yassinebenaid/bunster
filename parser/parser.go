package parser

import (
	"go/token"

	"github.com/yassinebenaid/nash/lexer"
)

type Pasrer struct {
	l         *lexer.Lexer
	tokenBuff []token.Token
}
