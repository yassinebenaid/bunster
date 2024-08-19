package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		input  string
		tokens []token.Token
	}{
		//Keywords
		{"if", []token.Token{{Type: token.IF, Literal: "if"}}},
		{"if", []token.Token{{Type: token.IF, Literal: "if"}}},
		{"then", []token.Token{{Type: token.THEN, Literal: "then"}}},
		{"else", []token.Token{{Type: token.ELSE, Literal: "else"}}},
		{"elif", []token.Token{{Type: token.ELIF, Literal: "elif"}}},
		{"fi", []token.Token{{Type: token.FI, Literal: "fi"}}},
		{"for", []token.Token{{Type: token.FOR, Literal: "for"}}},
		{"in", []token.Token{{Type: token.IN, Literal: "in"}}},
		{"do", []token.Token{{Type: token.DO, Literal: "do"}}},
		{"done", []token.Token{{Type: token.DONE, Literal: "done"}}},
		{"while", []token.Token{{Type: token.WHILE, Literal: "while"}}},
		{"until", []token.Token{{Type: token.UNTIL, Literal: "until"}}},
		{"case", []token.Token{{Type: token.CASE, Literal: "case"}}},
		{"esac", []token.Token{{Type: token.ESAC, Literal: "esac"}}},
		{"function", []token.Token{{Type: token.FUNCTION, Literal: "function"}}},
		{"select", []token.Token{{Type: token.SELECT, Literal: "select"}}},
		{"trap", []token.Token{{Type: token.TRAP, Literal: "trap"}}},
		{"return", []token.Token{{Type: token.RETURN, Literal: "return"}}},
		{"exit", []token.Token{{Type: token.EXIT, Literal: "exit"}}},
		{"break", []token.Token{{Type: token.BREAK, Literal: "break"}}},
		{"continue", []token.Token{{Type: token.CONTINUE, Literal: "continue"}}},
		{"declare", []token.Token{{Type: token.DECLARE, Literal: "declare"}}},
		{"local", []token.Token{{Type: token.LOCAL, Literal: "local"}}},
		{"export", []token.Token{{Type: token.EXPORT, Literal: "export"}}},
		{"readonly", []token.Token{{Type: token.READONLY, Literal: "readonly"}}},
		{"unset", []token.Token{{Type: token.UNSET, Literal: "unset"}}},
	}

	for i, tc := range testCases {
		l := lexer.New([]byte(tc.input))
		for _, tn := range tc.tokens {
			if result := l.NextToken(); tn.Type != result.Type {
				t.Fatalf(`#%d: wrong token type for %q, want=%d got=%d`, i, tn.Literal, tn.Type, result.Type)
			} else if tn.Literal != result.Literal {
				t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.Literal, tn.Literal, i)
			}
		}
	}
}
