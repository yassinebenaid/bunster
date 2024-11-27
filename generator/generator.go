package generator

import (
	"github.com/yassinebenaid/ryuko/ast"
	"github.com/yassinebenaid/ryuko/ir"
)

func Generate(script ast.Script) ir.Program {
	g := generator{}

	return g.generate(script)
}

type generator struct{}

func (*generator) generate(ast.Script) ir.Program {
	var program ir.Program

	return program
}
