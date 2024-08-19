package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
)

func TestLexer(t *testing.T) {
	input := `
	if [ $1 = 'foo' ]; then
    	echo 'Foo Bar'
	fi
	`

	l := lexer.New([]byte(input))

	tokens := []token.Token{
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.IF, Literal: "if"},
		{Type: token.LEFT_BRACKET, Literal: "["},
		{Type: token.SPECIAL_VAR, Literal: "1"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.LITERAL_STRING, Literal: "foo"},
		{Type: token.RIGHT_BRACKET, Literal: "]"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.THEN, Literal: "then"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.NAME, Literal: "echo"},
		{Type: token.LITERAL_STRING, Literal: "Foo Bar"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.FI, Literal: "fi"},
		{Type: token.NEWLINE, Literal: "\n"},
	}

	for i, tn := range tokens {
		if result := l.NextToken(); tn.Type != result.Type {
			t.Fatalf(`#%d: wrong token type for %q, want=%d got=%d`, i, tn.Literal, tn.Type, result.Type)
		} else if tn.Literal != result.Literal {
			t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.Literal, tn.Literal, i)
		}
	}
}
