package lexer

import (
	"github.com/yassinebenaid/nash/token"
)

type Lexer struct {
	input []byte
	pos   int
	ch    byte
	peek  byte
}

func New(in []byte) *Lexer {
	l := &Lexer{input: in}
	l.readCh()
	l.readCh()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	for l.ch == ' ' || l.ch == '\t' {
		l.readCh()
	}

	switch {
	case l.ch == '\n':
		tok.Type, tok.Literal = token.NL, "\n"
	case isIdentifier(l.ch):
		tok.Literal += string(l.ch)

		for isIdentifier(l.peek) {
			l.readCh()
			tok.Literal += string(l.ch)
		}

		tok.Type = token.IDENT
	case l.ch == 0:
		tok.Type, tok.Literal = token.EOF, "EOF"
	default:
		tok.Type, tok.Literal = token.ILLIGAL, string(l.ch)
	}

	l.readCh()

	return tok
}

func isIdentifier(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

func (l *Lexer) readCh() {
	l.ch = l.peek
	if l.pos >= len(l.input) {
		l.peek = 0
	} else {
		l.peek = l.input[l.pos]
	}
	l.pos++
}
