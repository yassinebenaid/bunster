package parser_test

import (
	"testing"

	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/parser"
)

func TestCanParseCommandCall(t *testing.T) {
	input := `git branch`

	p := parser.New(lexer.New([]byte(input)))
	script := p.ParseScript()

	if len := len(script.Statements); len != 1 {
		t.Fatalf(`expected script to have 1 expression, got "%d".`, len)
	}

}
