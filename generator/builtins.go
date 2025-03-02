package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleBreak(buf *InstructionBuffer, _ ast.Break) {
	buf.add(ir.Literal("break\n"))
}

func (g *generator) handleContinue(buf *InstructionBuffer, _ ast.Continue) {
	buf.add(ir.Literal("continue\n"))
}

func (g *generator) handleWait(buf *InstructionBuffer, _ ast.Wait) {
	buf.add(ir.Literal("shell.WaitGroup.Wait()\n"))
}
