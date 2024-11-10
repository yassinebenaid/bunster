package lexer

import (
	"github.com/yassinebenaid/bunny/token"
)

type context byte

const (
	ctx_DEFAULT context = iota
	ctx_LITERAL_STRING
)

type Lexer struct {
	input []byte
	pos   int
	curr  byte
	next  byte
	ctx   context
}

func New(in []byte) Lexer {
	l := Lexer{input: in}

	// read twice so that 'curr' and 'next' get initialized
	l.proceed()
	l.proceed()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

switch_beginning:
	switch {
	case l.ctx == ctx_LITERAL_STRING && l.curr != '\'' && l.curr != 0:
		tok.Type, tok.Literal = token.OTHER, string(l.curr)

		for l.next != 0 && l.next != '\'' {
			l.proceed()
			tok.Literal += string(l.curr)
		}
	case l.curr == ' ' || l.curr == '\t':
		tok.Type, tok.Literal = token.BLANK, string(l.curr)
		for l.next == ' ' || l.next == '\t' {
			l.proceed()
			tok.Literal += string(l.curr)
		}
	case l.curr == '\n':
		tok.Type, tok.Literal = token.NEWLINE, string(l.curr)
	case l.curr == '*':
		if l.next == '=' {
			l.proceed()
			tok.Type, tok.Literal = token.STAR_ASSIGN, "*="
		} else if l.next == '*' {
			l.proceed()
			tok.Type, tok.Literal = token.EXPONENTIATION, "**"
		} else {
			tok.Type, tok.Literal = token.STAR, string(l.curr)
		}
	case l.curr == '^':
		if l.next == '^' {
			l.proceed()
			tok.Type, tok.Literal = token.DOUBLE_CIRCUMFLEX, "^^"
		} else if l.next == '=' {
			l.proceed()
			tok.Type, tok.Literal = token.CIRCUMFLEX_ASSIGN, "^="
		} else {
			tok.Type, tok.Literal = token.CIRCUMFLEX, string(l.curr)
		}
	case l.curr == '%':
		if l.next == '%' {
			l.proceed()
			tok.Type, tok.Literal = token.DOUBLE_PERCENT, "%%"
		} else if l.next == '=' {
			l.proceed()
			tok.Type, tok.Literal = token.PERCENT_ASSIGN, "%="
		} else {
			tok.Type, tok.Literal = token.PERCENT, string(l.curr)
		}
	case l.curr == '[':
		if l.next == '[' {
			l.proceed()
			tok.Type, tok.Literal = token.DOUBLE_LEFT_BRACKET, "[["
		} else {
			tok.Type, tok.Literal = token.LEFT_BRACKET, string(l.curr)
		}
	case l.curr == '<':
		switch l.next {
		case '<':
			l.proceed()
			switch l.next {
			case '-':
				l.proceed()
				tok.Type, tok.Literal = token.DOUBLE_LT_MINUS, "<<-"
			case '<':
				l.proceed()
				tok.Type, tok.Literal = token.TRIPLE_LT, "<<<"
			case '=':
				l.proceed()
				tok.Type, tok.Literal = token.DOUBLE_LT_ASSIGN, "<<="
			default:
				tok.Type, tok.Literal = token.DOUBLE_LT, "<<"
			}
		case '&':
			l.proceed()
			tok.Type, tok.Literal = token.LT_AMPERSAND, "<&"
		case '>':
			l.proceed()
			tok.Type, tok.Literal = token.LT_GT, "<>"
		case '=':
			l.proceed()
			tok.Type, tok.Literal = token.LT_EQ, "<="
		case '(':
			l.proceed()
			tok.Type, tok.Literal = token.LT_PAREN, "<("
		default:
			tok.Type, tok.Literal = token.LT, string(l.curr)
		}
	case l.curr == '>':
		switch l.next {
		case '>':
			l.proceed()
			switch l.next {
			case '=':
				l.proceed()
				tok.Type, tok.Literal = token.DOUBLE_GT_ASSIGN, ">>="
			default:
				tok.Type, tok.Literal = token.DOUBLE_GT, ">>"
			}
		case '=':
			l.proceed()
			tok.Type, tok.Literal = token.GT_EQ, ">="
		case '&':
			l.proceed()
			tok.Type, tok.Literal = token.GT_AMPERSAND, ">&"
		case '|':
			l.proceed()
			tok.Type, tok.Literal = token.GT_PIPE, ">|"
		case '(':
			l.proceed()
			tok.Type, tok.Literal = token.GT_PAREN, ">("
		default:
			tok.Type, tok.Literal = token.GT, ">"
		}
	case l.curr == '&':
		switch l.next {
		case '&':
			l.proceed()
			tok.Type, tok.Literal = token.AND, "&&"
		case '>':
			l.proceed()
			switch l.next {
			case '>':
				l.proceed()
				tok.Type, tok.Literal = token.AMPERSAND_DOUBLE_GT, "&>>"
			default:
				tok.Type, tok.Literal = token.AMPERSAND_GT, "&>"
			}
		default:
			tok.Type, tok.Literal = token.AMPERSAND, string(l.curr)
		}
	case l.curr == '|':
		switch l.next {
		case '|':
			l.proceed()
			tok.Type, tok.Literal = token.OR, "||"
		case '&':
			l.proceed()
			tok.Type, tok.Literal = token.PIPE_AMPERSAND, "|&"
		default:
			tok.Type, tok.Literal = token.PIPE, string(l.curr)
		}
	case l.curr == '+':
		switch l.next {
		case '+':
			l.proceed()
			tok.Type, tok.Literal = token.INCREMENT, "++"
		case '=':
			l.proceed()
			tok.Type, tok.Literal = token.PLUS_ASSIGN, "+="
		default:
			tok.Type, tok.Literal = token.PLUS, string(l.curr)
		}
	case l.curr == '/':
		switch l.next {
		case '=':
			l.proceed()
			tok.Type, tok.Literal = token.SLASH_ASSIGN, "/="
		default:
			tok.Type, tok.Literal = token.SLASH, string(l.curr)
		}
	case l.curr == '-':
		switch l.next {
		case '-':
			l.proceed()
			tok.Type, tok.Literal = token.DECREMENT, "--"
		case '=':
			l.proceed()
			tok.Type, tok.Literal = token.MINUS_ASSIGN, "-="
		default:
			tok.Type, tok.Literal = token.MINUS, string(l.curr)
		}
	case l.curr == ']':
		if l.next == ']' {
			l.proceed()
			tok.Type, tok.Literal = token.DOUBLE_RIGHT_BRACKET, "]]"
		} else {
			tok.Type, tok.Literal = token.RIGHT_BRACKET, string(l.curr)
		}
	case l.curr == ';':
		tok.Type, tok.Literal = token.SEMICOLON, string(l.curr)
	case l.curr == '=':
		switch l.next {
		case '=':
			l.proceed()
			tok.Type, tok.Literal = token.EQ, "=="
		case '~':
			l.proceed()
			tok.Type, tok.Literal = token.EQ_TILDE, "=~"
		default:
			tok.Type, tok.Literal = token.ASSIGN, string(l.curr)
		}
	case l.curr == '(':
		if l.next == '(' {
			l.proceed()
			tok.Type, tok.Literal = token.DOUBLE_LEFT_PAREN, "(("
		} else {
			tok.Type, tok.Literal = token.LEFT_PAREN, string(l.curr)
		}
	case l.curr == ')':
		tok.Type, tok.Literal = token.RIGHT_PAREN, string(l.curr)
	case l.curr == ',':
		if l.next == ',' {
			l.proceed()
			tok.Type, tok.Literal = token.DOUBLE_COMMA, ",,"
		} else {
			tok.Type, tok.Literal = token.COMMA, string(l.curr)
		}
	case l.curr == '{':
		tok.Type, tok.Literal = token.LEFT_BRACE, string(l.curr)
	case l.curr == '}':
		tok.Type, tok.Literal = token.RIGHT_BRACE, string(l.curr)
	case l.curr == ':':
		switch l.next {
		case '=':
			tok.Type, tok.Literal = token.COLON_ASSIGN, ":="
		case '-':
			tok.Type, tok.Literal = token.COLON_MINUS, ":-"
		case '+':
			tok.Type, tok.Literal = token.COLON_PLUS, ":+"
		case '?':
			tok.Type, tok.Literal = token.COLON_QUESTION, ":?"
		default:
			tok.Type, tok.Literal = token.COLON, string(l.curr)
		}

		if tok.Type != token.COLON {
			l.proceed()
		}
	case l.curr == '?':
		tok.Type, tok.Literal = token.QUESTION, string(l.curr)
	case l.curr == '~':
		tok.Type, tok.Literal = token.TILDE, string(l.curr)
	case l.curr == '.' && l.next == '.':
		l.proceed()
		tok.Type, tok.Literal = token.DOUBLE_DOT, ".."
	case l.curr == '!':
		if l.next == '=' {
			l.proceed()
			tok.Type, tok.Literal = token.NOT_EQ, "!="
		} else {
			tok.Type, tok.Literal = token.EXCLAMATION, string(l.curr)
		}
	case l.curr == '#':
		tok.Type, tok.Literal = token.HASH, string(l.curr)
	case l.curr == '\'':
		if l.ctx == ctx_DEFAULT {
			l.ctx = ctx_LITERAL_STRING
		} else {
			l.ctx = ctx_DEFAULT
		}
		tok.Type, tok.Literal = token.SINGLE_QUOTE, string(l.curr)
	case l.curr == '"':
		tok.Type, tok.Literal = token.DOUBLE_QUOTE, string(l.curr)
	case l.curr == '$':
		switch l.next {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', '#', '_', '*', '@', '?', '!':
			l.proceed()
			tok.Type, tok.Literal = token.SPECIAL_VAR, string(l.curr)
		case '{':
			l.proceed()
			tok.Type, tok.Literal = token.DOLLAR_BRACE, "${"
		case '(':
			l.proceed()
			if l.next == '(' {
				l.proceed()
				tok.Type, tok.Literal = token.DOLLAR_DOUBLE_PAREN, "$(("
			} else {
				tok.Type, tok.Literal = token.DOLLAR_PAREN, "$("
			}
		default:
			if isLetter(l.next) {
				tok.Type = token.SIMPLE_EXPANSION
				for isLetter(l.next) || (l.next <= '9' && l.next >= '0') {
					l.proceed()
					tok.Literal += string(l.curr)
				}
			} else {
				tok.Type, tok.Literal = token.OTHER, "$"
			}
		}
	case isLetter(l.curr):
		tok.Literal = string(l.curr)
		for isLetter(l.next) {
			l.proceed()
			tok.Literal += string(l.curr)
		}

		if keyword, ok := token.Keywords[tok.Literal]; ok {
			tok.Type = keyword
		} else {
			tok.Type = token.WORD
		}
	case (l.curr >= '0' && l.curr <= '9') || (l.curr == '.' && (l.next >= '0' && l.next <= '9')):
		tok.Type, tok.Literal = token.INT, string(l.curr)
		isFloat := l.curr == '.'

		for {
			if isFloat && l.next == '.' {
				break
			}

			if !((l.next >= '0' && l.next <= '9') || l.next == '.') {
				break
			}

			if l.next == '.' {
				isFloat = true
			}

			l.proceed()
			tok.Literal += string(l.curr)
		}

		// If numbers appear in file descriptor positions they're treated differently (eg 1>&2)
		if isFloat {
			tok.Type = token.FLOAT
		}
	case l.curr == '\\':
		l.proceed()
		switch l.curr {
		case 0:
			tok.Type, tok.Literal = token.EOF, "end of file"
		case '"', '\\', '$':
			tok.Type, tok.Literal = token.OTHER, string(l.curr)
		case '\n':
			l.proceed()
			goto switch_beginning
		default:
			tok.Type, tok.Literal = token.ESCAPED_CHAR, string(l.curr)
		}
	case l.curr == 0:
		tok.Type, tok.Literal = token.EOF, "end of file"
	default:
		tok.Type, tok.Literal = token.OTHER, string(l.curr)
	}

	l.proceed()

	return tok
}

func isLetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || b == '_'
}

func (l *Lexer) proceed() {
	l.curr = l.next
	if l.pos >= len(l.input) {
		l.next = 0
	} else {
		l.next = l.input[l.pos]
	}
	l.pos++
}
