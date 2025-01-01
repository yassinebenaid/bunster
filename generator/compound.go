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
