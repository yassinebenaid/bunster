package parser_test

import (
	"github.com/yassinebenaid/bunster/ast"
)

var embeddingTests = []testCase{
	{`@embed file`, ast.Script{ast.Embed{"file"}}},
}
