package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
)

func TestLexer(t *testing.T) {
	input := `
	    if []
	`

	l := lexer.New([]byte(input))

	tokens := []token.Token{
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.IF, Literal: "if"},
	}

	for i, tn := range tokens {
		if result := l.NextToken(); tn.Type != result.Type {
			t.Fatalf(`#%d: wrong token type for %q, want=%d got=%d`, i, tn.Literal, tn.Type, result.Type)
		} else if tn.Literal != result.Literal {
			t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.Literal, tn.Literal, i)
		}
	}
}
