package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleGroup(buf *InstructionBuffer, group ast.Group, pc *pipeContext) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, group.Redirections, pc)

	if pc == nil {
		for _, cmd := range group.Body {
			g.generate(&cmdbuf, cmd, nil)
		}
	} else {
		cmdbuf.add(ir.Literal("var done = make(chan struct{},1)\n"))
		cmdbuf.add(ir.PushToPipelineWaitgroup{
			Waitgroup: pc.waitgroup,
			Value: ir.Literal(`func() error {
				<-done
			 	streamManager.Destroy()
				return nil
			}`),
		})

		var go_routing InstructionBuffer
		for _, cmd := range group.Body {
			g.generate(&go_routing, cmd, nil)
		}
		go_routing.add(ir.Literal("done<-struct{}{}\n"))
		cmdbuf.add(ir.Closure{
			Async: true,
			Body:  go_routing,
		})
	}

	*buf = append(*buf, ir.Closure{
		Body: cmdbuf,
	})
}

func (g *generator) handleSubshell(buf *InstructionBuffer, subshell ast.SubShell, pc *pipeContext) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneShell{})
	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, subshell.Redirections, pc)

	if pc == nil {
		for _, cmd := range subshell.Body {
			g.generate(&cmdbuf, cmd, nil)
		}
	} else {
		cmdbuf.add(ir.Literal("var done = make(chan struct{},1)\n"))
		cmdbuf.add(ir.PushToPipelineWaitgroup{
			Waitgroup: pc.waitgroup,
			Value: ir.Literal(`func() error {
				<-done
			 	streamManager.Destroy()
				return nil
			}`),
		})

		var go_routing InstructionBuffer
		for _, cmd := range subshell.Body {
			g.generate(&go_routing, cmd, nil)
		}
		go_routing.add(ir.Literal("done<-struct{}{}\n"))
		cmdbuf.add(ir.Closure{
			Async: true,
			Body:  go_routing,
		})
	}

	*buf = append(*buf, ir.Closure{
		Body: cmdbuf,
	})
}

func (g *generator) handleIf(buf *InstructionBuffer, cond ast.If, pc *pipeContext) {
	var cmdbuf InstructionBuffer
	cmdbuf.add(ir.Declare{Name: "condition", Value: ir.Literal("false")})

	for _, statement := range cond.Head {
		g.generate(&cmdbuf, statement, nil)
		cmdbuf.add(ir.Set{Name: "condition", Value: ir.Literal("shell.ExitCode == 0")})
		cmdbuf.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})
	}

	var body InstructionBuffer
	for _, statement := range cond.Body {
		g.generate(&body, statement, nil)
	}
	cmdbuf.add(ir.If{
		Condition: ir.Literal("condition"),
		Body:      body,
	})

	*buf = append(*buf, ir.Closure{
		Body: cmdbuf,
	})
}
