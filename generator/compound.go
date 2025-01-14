package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleGroup(buf *InstructionBuffer, group ast.Group, ctx *context) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{DeferDestroy: ctx.pipe == nil})
	g.handleRedirections(&cmdbuf, group.Redirections, ctx)

	if ctx.pipe == nil {
		for _, cmd := range group.Body {
			g.generate(&cmdbuf, cmd, &context{})
		}
	} else {
		cmdbuf.add(ir.Literal("var done = make(chan struct{},1)\n"))
		cmdbuf.add(ir.PushToPipelineWaitgroup{
			Waitgroup: ctx.pipe.waitgroup,
			Value: ir.Literal(`func() error {
				<-done
			 	streamManager.Destroy()
				return nil
			}`),
		})

		var go_routing InstructionBuffer
		for _, cmd := range group.Body {
			g.generate(&go_routing, cmd, &context{})
		}
		go_routing.add(ir.Literal("done<-struct{}{}\n"))
		cmdbuf.add(ir.Gorouting(go_routing))
	}

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleSubshell(buf *InstructionBuffer, subshell ast.SubShell, ctx *context) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneShell{})
	cmdbuf.add(ir.CloneStreamManager{DeferDestroy: ctx.pipe == nil})
	g.handleRedirections(&cmdbuf, subshell.Redirections, ctx)

	if ctx.pipe == nil {
		for _, cmd := range subshell.Body {
			g.generate(&cmdbuf, cmd, &context{})
		}
	} else {
		cmdbuf.add(ir.Literal("var done = make(chan struct{},1)\n"))
		cmdbuf.add(ir.PushToPipelineWaitgroup{
			Waitgroup: ctx.pipe.waitgroup,
			Value: ir.Literal(`func() error {
				<-done
			 	streamManager.Destroy()
				return nil
			}`),
		})

		var go_routing InstructionBuffer
		for _, cmd := range subshell.Body {
			g.generate(&go_routing, cmd, &context{})
		}
		go_routing.add(ir.Literal("done<-struct{}{}\n"))
		cmdbuf.add(ir.Gorouting(go_routing))
	}

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleIf(buf *InstructionBuffer, cond ast.If, ctx *context) {
	var cmdbuf InstructionBuffer
	cmdbuf.add(ir.CloneStreamManager{DeferDestroy: ctx.pipe == nil})

	g.handleRedirections(&cmdbuf, cond.Redirections, ctx)

	var innerBuf InstructionBuffer
	innerBuf.add(ir.Declare{Name: "condition", Value: ir.Literal("false")})
	for _, statement := range cond.Head {
		g.generate(&innerBuf, statement, &context{})
		innerBuf.add(ir.Set{Name: "condition", Value: ir.Literal("shell.ExitCode == 0")})
		innerBuf.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})
	}

	var body InstructionBuffer
	for _, statement := range cond.Body {
		g.generate(&body, statement, &context{})
	}
	innerBuf.add(ir.If{
		Condition: ir.Literal("condition"),
		Body:      body,
		Alternate: g.handleElif(cond.Elifs),
	})

	if cond.Alternate != nil {
		var alt InstructionBuffer
		for _, statement := range cond.Alternate {
			g.generate(&alt, statement, &context{})
		}
		innerBuf.add(ir.If{Condition: ir.Literal("!condition"), Body: alt})
	}

	if ctx.pipe == nil {
		cmdbuf = append(cmdbuf, innerBuf...)
		*buf = append(*buf, ir.Closure(cmdbuf))
		return
	}

	cmdbuf.add(ir.Literal("var done = make(chan struct{},1)\n"))
	cmdbuf.add(ir.PushToPipelineWaitgroup{
		Waitgroup: ctx.pipe.waitgroup,
		Value: ir.Literal(`func() error {
				<-done
			 	streamManager.Destroy()
				return nil
			}`),
	})

	innerBuf.add(ir.Literal("done<-struct{}{}\n"))
	cmdbuf.add(ir.Gorouting(innerBuf))

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleElif(elifs []ast.Elif) []ir.Instruction {
	if len(elifs) == 0 {
		return nil
	}

	var cmdbuf InstructionBuffer

	for _, statement := range elifs[0].Head {
		g.generate(&cmdbuf, statement, &context{})
		cmdbuf.add(ir.Set{Name: "condition", Value: ir.Literal("shell.ExitCode == 0")})
		cmdbuf.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})
	}

	var body InstructionBuffer
	for _, statement := range elifs[0].Body {
		g.generate(&body, statement, &context{})
	}
	cmdbuf.add(ir.If{
		Condition: ir.Literal("condition"),
		Body:      body,
		Alternate: g.handleElif(elifs[1:]),
	})

	return cmdbuf

}

func (g *generator) handleLoop(buf *InstructionBuffer, loop ast.Loop, ctx *context) {
	var cmdbuf InstructionBuffer
	cmdbuf.add(ir.CloneStreamManager{DeferDestroy: ctx.pipe == nil})

	g.handleRedirections(&cmdbuf, loop.Redirections, ctx)

	var innerBuf, body InstructionBuffer

	for _, statement := range loop.Head {
		g.generate(&body, statement, &context{})
		body.add(ir.Declare{Name: "condition", Value: ir.Literal("shell.ExitCode == 0")})
		body.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})

		condition := ir.Literal("! condition")
		if loop.Negate {
			condition = ir.Literal("condition")
		}
		body.add(ir.If{
			Condition: condition,
			Body: []ir.Instruction{
				ir.Literal("break\n"),
			},
		})
	}

	for _, statement := range loop.Body {
		g.generate(&body, statement, &context{})
	}
	innerBuf.add(ir.Loop{
		Condition: ir.Literal(""),
		Body:      body,
	})

	if ctx.pipe == nil {
		cmdbuf = append(cmdbuf, innerBuf...)
	} else {
		cmdbuf.add(ir.Literal("var done = make(chan struct{},1)\n"))
		cmdbuf.add(ir.PushToPipelineWaitgroup{
			Waitgroup: ctx.pipe.waitgroup,
			Value: ir.Literal(`func() error {
				<-done
			 	streamManager.Destroy()
				return nil
			}`),
		})

		innerBuf.add(ir.Literal("done<-struct{}{}\n"))
		cmdbuf.add(ir.Gorouting(innerBuf))
	}

	*buf = append(*buf, ir.Closure(cmdbuf))
}
