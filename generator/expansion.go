package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleCommandSubstitution(buf *InstructionBuffer, statements ast.CommandSubstitution) ir.Instruction {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	cmdbuf.add(ir.CloneShell{})
	cmdbuf.add(ir.Declare{
		Name:  "buffer",
		Value: ir.NewBuffer{Value: ir.String("")},
	})
	cmdbuf.add(ir.AddStream{Fd: "1", StreamName: "buffer"})

	for _, statement := range statements {
		g.generate(&cmdbuf, statement)
	}

	cmdbuf.add(ir.Literal("return buffer.String(true), shell.ExitCode"))

	name := fmt.Sprintf("expr%d", g.expressionsCount)
	buf.add(ir.ExpressionClosure{
		Body: cmdbuf,
		Name: name,
	})

	return ir.Literal(name)
}
