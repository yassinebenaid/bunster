package expr

import (
	"testing"

	"github.com/yassinebenaid/nash/token"
)

func TestLexer(t *testing.T) {
	input := `
		ls 
	`

	l := New([]byte(input))

	tokens := []token.Token{
		{Type: token.NL, Literal: "\n"},
		{Type: token.IDENT, Literal: "ls"},
		{Type: token.NL, Literal: "\n"},
	}

	for i, tn := range tokens {
		if result := l.NextToken(); tn.Type != result.Type {
			t.Fatalf(`wrong token type "%s", expected "%s", case#%d`, result.Type, tn.Type, i)
		} else if tn.Literal != result.Literal {
			t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.Literal, tn.Literal, i)
		}
	}
}
