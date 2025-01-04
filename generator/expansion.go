package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleCommandSubstitution(buf *InstructionBuffer, statements ast.CommandSubstitution) ir.Instruction {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{DeferDestroy: true})
	cmdbuf.add(ir.CloneShell{})
	cmdbuf.add(ir.Declare{
		Name:  "buffer",
		Value: ir.NewBuffer{Value: ir.String("")},
	})
	cmdbuf.add(ir.AddStream{Fd: "1", StreamName: "buffer"})

	for _, statement := range statements {
		g.generate(&cmdbuf, statement, nil)
	}

	cmdbuf.add(ir.Literal("return buffer.String(true), shell.ExitCode"))

	buf.add(ir.ExpressionClosure{
		Body: cmdbuf,
		Name: "commandSubstitionExpression",
	})

	return ir.Literal("commandSubstitionExpression")
}
