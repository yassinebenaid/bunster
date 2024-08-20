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
	case l.ch == '*':
		if l.peek == '=' {
			l.readCh()
			tok.Type, tok.Literal = token.STAR_ASSIGN, "*="
		} else {
			tok.Type, tok.Literal = token.STAR, string(l.ch)
		}
	case l.ch == '%':
		tok.Type, tok.Literal = token.PERCENT, string(l.ch)
	case l.ch == '[':
		if l.peek == '[' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_LEFT_BRACKET, "[["
		} else {
			tok.Type, tok.Literal = token.LEFT_BRACKET, string(l.ch)
		}
	case l.ch == '<':
		switch l.peek {
		case '<':
			l.readCh()
			switch l.peek {
			case '-':
				tok.Type, tok.Literal = token.DOUBLE_LT_MINUS, "<<-"
			case '<':
				tok.Type, tok.Literal = token.TRIPLE_LT, "<<<"
			default:
				tok.Type, tok.Literal = token.DOUBLE_LT, "<<"
			}
		case '=':
			l.readCh()
			tok.Type, tok.Literal = token.LE, "<="
		case '&':
			l.readCh()
			tok.Type, tok.Literal = token.LT_AMPERSAND, "<&"
		default:
			tok.Type, tok.Literal = token.LT, string(l.ch)
		}
	case l.ch == '>':
		switch l.peek {
		case '>':
			tok.Type, tok.Literal = token.DOUBLE_GT, ">>"
		case '=':
			tok.Type, tok.Literal = token.GE, ">="
		case '&':
			tok.Type, tok.Literal = token.GT_AMPERSAND, ">&"
		default:
			tok.Type, tok.Literal = token.GT, ">"
		}
	case l.ch == '&':
		if l.peek == '&' {
			l.readCh()
			tok.Type, tok.Literal = token.AND, "&&"
		} else {
			tok.Type, tok.Literal = token.AMPERSAND, string(l.ch)
		}
	case l.ch == '|':
		switch l.peek {
		case '|':
			l.readCh()
			tok.Type, tok.Literal = token.OR, "||"
		case '&':
			l.readCh()
			tok.Type, tok.Literal = token.PIPE_AMPERSAND, "|&"
		default:
			tok.Type, tok.Literal = token.PIPE, string(l.ch)
		}
	case l.ch == '+':
		if l.peek == '+' {
			l.readCh()
			tok.Type, tok.Literal = token.INCREMENT, "++"
		} else if l.peek == '=' {
			l.readCh()
			tok.Type, tok.Literal = token.PLUS_ASSIGN, "+="
		} else {
			tok.Type, tok.Literal = token.PLUS, string(l.ch)
		}
	case l.ch == '/':
		if l.peek == '/' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_SLASH, "//"
		} else if l.peek == '=' {
			l.readCh()
			tok.Type, tok.Literal = token.SLASH_ASSIGN, "/="
		} else {
			tok.Type, tok.Literal = token.SLASH, string(l.ch)
		}
	case l.ch == '-':
		if l.peek == '-' {
			l.readCh()
			tok.Type, tok.Literal = token.DECREMENT, "--"
		} else if l.peek == '=' {
			l.readCh()
			tok.Type, tok.Literal = token.MINUS_ASSIGN, "-="
		} else {
			tok.Type, tok.Literal = token.MINUS, string(l.ch)
		}
	case l.ch == '!' && l.peek == '=':
		tok.Type, tok.Literal = token.NOT_EQ, "!="
	case l.ch == ']':
		if l.peek == ']' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_RIGHT_BRACKET, "]]"
		} else {
			tok.Type, tok.Literal = token.RIGHT_BRACKET, string(l.ch)
		}
	case l.ch == ';':
		tok.Type, tok.Literal = token.SEMICOLON, string(l.ch)
	case l.ch == '=':
		if l.peek == '=' {
			l.readCh()
			tok.Type, tok.Literal = token.EQ, "=="
		} else {
			tok.Type, tok.Literal = token.ASSIGN, string(l.ch)
		}
	case l.ch == '\'':
		tok.Type = token.LITERAL_STRING
		l.readCh()

		for l.ch != '\'' {
			tok.Literal += string(l.ch)
			l.readCh()
		}
	case l.ch == '$':
		switch {
		case l.peek >= '0' && l.peek <= '9':
			l.readCh()
			tok.Type, tok.Literal = token.SPECIAL_VAR, string(l.ch)
		}
	case isLetter(l.ch):
		tok.Literal = string(l.ch)

		for isLetter(l.peek) {
			l.readCh()
			tok.Literal += string(l.ch)
		}

		if keyword, ok := token.Keywords[tok.Literal]; ok {
			tok.Type = keyword
		} else {
			tok.Type = token.NAME
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
