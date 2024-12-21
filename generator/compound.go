package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleGroup(buf *InstructionBuffer, group ast.Group, pc *pipeContext) {
	var cmdbuf InstructionBuffer

	for _, cmd := range group.Body {
		g.generate(&cmdbuf, cmd, nil)
	}

	*buf = append(*buf, ir.Closure{
		Body: cmdbuf,
	})
}
