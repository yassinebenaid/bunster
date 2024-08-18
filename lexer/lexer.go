package lexer

import (
	"github.com/yassinebenaid/nbs/token"
)

type Lexer struct {
	input []byte
	pos   int
	ch    byte
	peek  byte
}

func New(in []byte) Lexer {
	l := Lexer{input: in}
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
		tok.Type, tok.Literal = token.NEWLINE, string(l.ch)
	case isLetter(l.ch):
		tok.Literal = string(l.ch)

		for isLetter(l.peek) {
			l.readCh()
			tok.Literal += string(l.ch)
		}

		if keyword, ok := token.Keywords[tok.Literal]; ok {
			tok.Type = keyword
		}
	}

	l.readCh()

	return tok
}

func isLetter(b byte) bool {
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
