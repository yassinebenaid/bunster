package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleTest(buf *InstructionBuffer, test ast.Test) {
	var cmdbuf, body InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, test.Redirections)

	body.add(ir.Declare{Name: "testResult", Value: ir.Literal("false")})
	g.handleTestExpression(&body, test.Expr)
	body.add(ir.Literal("if testResult { shell.ExitCode = 0 } else { shell.ExitCode = 1  }\n"))

	cmdbuf = append(cmdbuf, body...)
	*buf = append(*buf, ir.Closure(cmdbuf))

}

func (g *generator) handleTestExpression(buf *InstructionBuffer, test ast.Expression) {
	switch v := test.(type) {
	case ast.Binary:
		g.handleTestBinary(buf, v)
	case ast.Unary:
		g.handleTestUnary(buf, v)
	case ast.Negation:
		g.handleTestExpression(buf, v.Operand)
		buf.add(ir.Literal("testResult = !testResult \n"))
	default:
		buf.add(ir.TestAgainsStringLength{String: g.handleExpression(buf, v)})
	}
}

func (g *generator) handleTestBinary(buf *InstructionBuffer, test ast.Binary) {
	switch test.Operator {
	case "&&":
		g.handleTestExpression(buf, test.Left)
		var right InstructionBuffer
		g.handleTestExpression(&right, test.Right)
		buf.add(ir.If{Condition: ir.Literal("testResult"), Body: right})
		return

	case "||":
		g.handleTestExpression(buf, test.Left)
		var right InstructionBuffer
		g.handleTestExpression(&right, test.Right)
		buf.add(ir.If{Condition: ir.Literal("! testResult"), Body: right})
		return
	}

	left := g.handleExpression(buf, test.Left)
	right := g.handleExpression(buf, test.Right)

	switch test.Operator {
	case "=":
		buf.add(ir.Compare{Left: left, Operator: "==", Right: right})
	case "==":
		buf.add(ir.Compare{Left: left, Operator: "==", Right: right})
	case "!=", "<", ">":
		buf.add(ir.Compare{Left: left, Operator: test.Operator, Right: right})
	case "-eq":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "==", Right: right})
	case "-ne":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "!=", Right: right})
	case "-lt":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "<", Right: right})
	case "-le":
		buf.add(ir.CompareArithmetics{Left: left, Operator: "<=", Right: right})
	case "-gt":
		buf.add(ir.CompareArithmetics{Left: left, Operator: ">", Right: right})
	case "-ge":
		buf.add(ir.CompareArithmetics{Left: left, Operator: ">=", Right: right})
	case "-ef":
		buf.add(ir.TestFilesHaveSameDevAndInoNumbers{File1: left, File2: right})
	case "-ot":
		buf.add(ir.FileIsOlderThan{File1: left, File2: right})
	case "-nt":
		buf.add(ir.FileIsOlderThan{File1: right, File2: left})
	default:
		panic("we do not support the binary operator: " + test.Operator)
	}
}

func (g *generator) handleTestUnary(buf *InstructionBuffer, test ast.Unary) {
	operand := g.handleExpression(buf, test.Operand)

	switch test.Operator {
	case "-n":
		buf.add(ir.TestAgainsStringLength{String: operand})
	case "-z":
		buf.add(ir.TestAgainsStringLength{String: operand, Zero: true})
	case "-e", "-a":
		buf.add(ir.TestFileExists{File: operand})
	case "-d":
		buf.add(ir.TestDirectoryExists{File: operand})
	case "-b":
		buf.add(ir.TestBlockSpecialFileExists{File: operand})
	case "-c":
		buf.add(ir.TestCharacterSpecialFileExists{File: operand})
	case "-f":
		buf.add(ir.TestRegularFileExists{File: operand})
	case "-g":
		buf.add(ir.TestFileSGIDIsSet{File: operand})
	case "-G":
		buf.add(ir.TestFileIsOwnedByEffectiveGroup{File: operand})
	case "-O":
		buf.add(ir.TestFileIsOwnedByEffectiveUser{File: operand})
	case "-u":
		buf.add(ir.TestFileSUIDIsSet{File: operand})
	case "-h", "-L":
		buf.add(ir.TestFileIsSymbolic{File: operand})
	case "-k":
		buf.add(ir.TestFileIsSticky{File: operand})
	case "-p":
		buf.add(ir.TestFileIsFIFO{File: operand})
	case "-r":
		buf.add(ir.TestFileIsReadable{File: operand})
	case "-x":
		buf.add(ir.TestFileIsExecutable{File: operand})
	case "-w":
		buf.add(ir.TestFileIsWritable{File: operand})
	case "-s":
		buf.add(ir.TestFileHasAPositiveSize{File: operand})
	case "-t":
		buf.add(ir.TestFileDescriptorIsTerminal{File: operand})
	case "-N":
		buf.add(ir.TestFileHasBeenModifiedSinceLastRead{File: operand})
	case "-S":
		buf.add(ir.TestFileIsSocket{File: operand})
	case "-v":
		buf.add(ir.TestVarIsSet{Name: operand})
	default:
		panic("we do not support the unary operator: " + test.Operator)
	}
}
