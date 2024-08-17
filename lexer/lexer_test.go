package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
)

func TestLexer(t *testing.T) {
	input := `
		echo $HOME
		git branch --show-current
	`

	l := lexer.New([]byte(input))

	tokens := []token.Token{
		{Type: token.NL, Literal: "\n"},
		{Type: token.IDENT, Literal: "echo"},
		{Type: token.DOLLAR, Literal: "$"},
		{Type: token.NAME, Literal: "HOME"},
		{Type: token.NL, Literal: "\n"},
		{Type: token.IDENT, Literal: "$"},
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
