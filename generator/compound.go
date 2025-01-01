package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleGroup(buf *InstructionBuffer, group ast.Group, pc *pipeContext) {
	var cmdbuf InstructionBuffer

	g.handleRedirections(&cmdbuf, "group", group.Redirections, pc, true)

	if pc == nil {
		for _, cmd := range group.Body {
			g.generate(&cmdbuf, cmd, nil)
		}
	} else {
		cmdbuf.add(ir.Literal("var done = make(chan struct{},1)"))
		cmdbuf.add(ir.Literal(fmt.Sprintf(`
			pipelineWaitgroup = append(pipelineWaitgroup,  func() error {
				<-done
				%s.Close()
				return nil
			})
		`, pc.writer)))

		var go_routing InstructionBuffer
		go_routing.add(ir.Literal("defer streamManager.Destroy()\n"))
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

	g.handleRedirections(&cmdbuf, "subshell", subshell.Redirections, pc, true)

	if pc == nil {
		for _, cmd := range subshell.Body {
			g.generate(&cmdbuf, cmd, nil)
		}
	} else {
		cmdbuf.add(ir.Literal("var done = make(chan struct{},1)"))
		cmdbuf.add(ir.Literal(fmt.Sprintf(`
			pipelineWaitgroup = append(pipelineWaitgroup,  func() error {
				<-done
				%s.Close()
				return nil
			})
		`, pc.writer)))

		var go_routing InstructionBuffer
		go_routing.add(ir.Literal("defer streamManager.Destroy()\n"))
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
