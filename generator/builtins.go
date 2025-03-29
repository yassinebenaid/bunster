package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleFunction(buf *InstructionBuffer, function ast.Function) {
	var body InstructionBuffer

	g.handleRedirections(&body, function.Redirections)

	for _, cmd := range function.Body {
		g.generate(&body, cmd)
	}

	buf.add(ir.Function{Name: function.Name, Body: body, Subshell: function.SubShell})
}

func (g *generator) handleExit(buf *InstructionBuffer, exit ast.Exit) {
	buf.add(ir.Exit(exit))
}

func (g *generator) handleBreak(buf *InstructionBuffer, _ ast.Break) {
	buf.add(ir.Literal("break\n"))
}

func (g *generator) handleContinue(buf *InstructionBuffer, _ ast.Continue) {
	buf.add(ir.Literal("continue\n"))
}

func (g *generator) handleWait(buf *InstructionBuffer, _ ast.Wait) {
	buf.add(ir.Literal("shell.WaitGroup.Wait()\n"))
}

func (g *generator) handleEmbed(_ *InstructionBuffer, embed ast.Embed) {
	g.embeds = append(g.embeds, embed...)
}
