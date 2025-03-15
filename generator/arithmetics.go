package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleArithmeticCommand(buf *InstructionBuffer, cmd ast.ArithmeticCommand) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, cmd.Redirections)

	cmdbuf.add(ir.Declare{Name: "arithmeticResult", Value: ir.Literal("0")})
	g.handleArithmeticExpression(&cmdbuf, cmd.Arithmetic)
	cmdbuf.add(ir.Literal("if arithmeticResult == 0 { shell.ExitCode = 1 } else { shell.ExitCode = 0  }\n"))

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleArithmeticExpression(buf *InstructionBuffer, arithmetics ast.Arithmetic) {
	for _, arithmetic := range arithmetics {

		switch v := arithmetic.(type) {
		case ast.PostIncDecArithmetic:
			buf.add(ir.VarIncDec{Operand: v.Operand, Operator: v.Operator})
		default:
			panic(fmt.Sprintf("what the f**k is this shit: %T", v))
			// buf.add(ir.ToInt{String: g.handleExpression(buf, v)})
		}

	}
}
