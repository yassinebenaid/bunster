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
	case l.ch == '^':
		if l.peek == '^' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_CIRCUMFLEX, "^^"
		} else {
			tok.Type, tok.Literal = token.CIRCUMFLEX, string(l.ch)
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
				l.readCh()
				tok.Type, tok.Literal = token.DOUBLE_LT_MINUS, "<<-"
			case '<':
				l.readCh()
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
		case '>':
			l.readCh()
			tok.Type, tok.Literal = token.LT_GT, "<>"
		case '(':
			l.readCh()
			tok.Type, tok.Literal = token.LT_PAREN, "<("
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
		case '|':
			tok.Type, tok.Literal = token.GT_PIPE, ">|"
		case '(':
			tok.Type, tok.Literal = token.GT_PAREN, ">("
		default:
			tok.Type, tok.Literal = token.GT, ">"
		}

		if tok.Type != token.GT {
			l.readCh()
		}
	case l.ch == '&':
		switch l.peek {
		case '&':
			l.readCh()
			tok.Type, tok.Literal = token.AND, "&&"
		case '>':
			l.readCh()
			tok.Type, tok.Literal = token.AMPERSAND_GT, "&>"
		default:
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
	case l.ch == ']':
		if l.peek == ']' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_RIGHT_BRACKET, "]]"
		} else {
			tok.Type, tok.Literal = token.RIGHT_BRACKET, string(l.ch)
		}
	case l.ch == ';':
		if l.peek == ';' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_SEMICOLON, ";;"
		} else {
			tok.Type, tok.Literal = token.SEMICOLON, string(l.ch)
		}
	case l.ch == '=':
		switch l.peek {
		case '=':
			l.readCh()
			tok.Type, tok.Literal = token.EQ, "=="
		case '~':
			l.readCh()
			tok.Type, tok.Literal = token.EQ_TILDE, "=~"
		default:
			tok.Type, tok.Literal = token.ASSIGN, string(l.ch)
		}
	case l.ch == '(':
		if l.peek == '(' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_LEFT_PAREN, "(("
		} else {
			tok.Type, tok.Literal = token.LEFT_PAREN, string(l.ch)
		}
	case l.ch == ')':
		if l.peek == ')' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_RIGHT_PAREN, "))"
		} else {
			tok.Type, tok.Literal = token.RIGHT_PAREN, string(l.ch)
		}
	case l.ch == ',':
		if l.peek == ',' {
			l.readCh()
			tok.Type, tok.Literal = token.DOUBLE_COMMA, ",,"
		} else {
			tok.Type, tok.Literal = token.COMMA, string(l.ch)
		}
	case l.ch == '{':
		tok.Type, tok.Literal = token.LEFT_BRACE, string(l.ch)
	case l.ch == '}':
		tok.Type, tok.Literal = token.RIGHT_BRACE, string(l.ch)
	case l.ch == ':':
		switch l.peek {
		case '=':
			tok.Type, tok.Literal = token.COLON_ASSIGN, ":="
		case '-':
			tok.Type, tok.Literal = token.COLON_MINUS, ":-"
		case '+':
			tok.Type, tok.Literal = token.COLON_PLUS, ":+"
		case '?':
			tok.Type, tok.Literal = token.COLON_QUESTION, ":?"
		default:
			tok.Type, tok.Literal = token.COLON, string(l.ch)
		}

		if tok.Type != token.COLON {
			l.readCh()
		}
	case l.ch == '?':
		tok.Type, tok.Literal = token.QUESTION, string(l.ch)
	case l.ch == '~':
		tok.Type, tok.Literal = token.TILDE, string(l.ch)
	case l.ch == '.' && l.peek == '.':
		l.readCh()
		tok.Type, tok.Literal = token.DOUBLE_DOT, ".."
	case l.ch == '!':
		if l.peek == '=' {
			l.readCh()
			tok.Type, tok.Literal = token.NOT_EQ, "!="
		} else {
			tok.Type, tok.Literal = token.EXCLAMATION, string(l.ch)
		}
	case l.ch == '#':
		tok.Type, tok.Literal = token.HASH, string(l.ch)
	case l.ch == '\'':
		tok.Type = token.LITERAL_STRING
		l.readCh()

		for l.ch != '\'' {
			tok.Literal += string(l.ch)
			l.readCh()
		}
	case l.ch == '$':
		switch l.peek {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', '#', '_', '*', '@', '?', '!':
			l.readCh()
			tok.Type, tok.Literal = token.SPECIAL_VAR, string(l.ch)
		case '{':
			l.readCh()
			tok.Type, tok.Literal = token.DOLLAR_BRACE, "${"
		case '(':
			l.readCh()
			if l.peek == '(' {
				l.readCh()
				tok.Type, tok.Literal = token.DOLLAR_DOUBLE_PAREN, "$(("
			} else {
				tok.Type, tok.Literal = token.DOLLAR_PAREN, "$("
			}
		default:
			if isLetter(l.peek) || l.peek == '_' {
				tok.Type = token.SIMPLE_EXPANSION
				for isLetter(l.peek) || l.peek == '_' {
					l.readCh()
					tok.Literal += string(l.ch)
				}
			}
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
	case l.ch >= '0' && l.ch <= '9':
		tok.Type, tok.Literal = token.NUMBER, string(l.ch)
		for l.peek >= '0' && l.peek <= '9' {
			l.readCh()
			tok.Literal += string(l.ch)
		}
	case l.ch == 0:
		tok.Type = token.EOF
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
