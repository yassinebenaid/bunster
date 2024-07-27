package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/nash/lexer"
	"github.com/yassinebenaid/nash/token"
)

func TestLexer(t *testing.T) {
	input := `
		ls 
		cd
		git branch
	`

	l := lexer.New([]byte(input))

	tokens := []token.Token{
		{Type: token.NL, Literal: "\n"},
		{Type: token.IDENT, Literal: "ls"},
		{Type: token.NL, Literal: "\n"},
		{Type: token.IDENT, Literal: "cd"},
		{Type: token.NL, Literal: "\n"},
		{Type: token.IDENT, Literal: "git"},
		{Type: token.IDENT, Literal: "branch"},
		{Type: token.NL, Literal: "\n"},
		{Type: token.EOF, Literal: "EOF"},
	}

	for i, tn := range tokens {
		if result := l.NextToken(); tn.Type != result.Type {
			t.Fatalf(`wrong token type "%s", expected "%s", case#%d`, result.Type, tn.Type, i)
		} else if tn.Literal != result.Literal {
			t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.Literal, tn.Literal, i)
		}
	}
}
