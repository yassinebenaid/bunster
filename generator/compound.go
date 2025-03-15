package generator

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleGroup(buf *InstructionBuffer, group ast.Group) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, group.Redirections)

	for _, cmd := range group.Body {
		g.generate(&cmdbuf, cmd)
	}

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleSubshell(buf *InstructionBuffer, subshell ast.SubShell) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, subshell.Redirections)

	cmdbuf.add(ir.Declare{Name: "parentShell", Value: ir.Literal("shell")})
	cmdbuf.add(ir.CloneShell{})
	cmdbuf.add(ir.Literal("defer func() { parentShell.ExitCode = shell.ExitCode }()\n"))

	for _, cmd := range subshell.Body {
		g.generate(&cmdbuf, cmd)
	}

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleIf(buf *InstructionBuffer, cond ast.If) {
	g.scopesCount++

	var cmdbuf, innerBuf InstructionBuffer

	cmdbuf.add(ir.Declare{Name: "condition", Value: ir.Literal("false")})
	cmdbuf.add(ir.CloneStreamManager{})

	g.handleRedirections(&cmdbuf, cond.Redirections)

	for _, statement := range cond.Head {
		g.generate(&innerBuf, statement)
		innerBuf.add(ir.Set{Name: "condition", Value: ir.Literal("shell.ExitCode == 0")})
		innerBuf.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})
	}

	var body InstructionBuffer
	for _, statement := range cond.Body {
		g.generate(&body, statement)
	}
	innerBuf.add(ir.If{
		Condition: ir.Literal("condition"),
		Body:      body,
		Alternate: g.handleElif(cond.Elifs),
	})

	if cond.Alternate != nil {
		var alt InstructionBuffer
		for _, statement := range cond.Alternate {
			g.generate(&alt, statement)
		}
		innerBuf.add(ir.If{Condition: ir.Literal("!condition"), Body: alt})
	}

	cmdbuf = append(cmdbuf, innerBuf...)
	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleElif(elifs []ast.Elif) []ir.Instruction {
	if len(elifs) == 0 {
		return nil
	}

	var cmdbuf InstructionBuffer

	for _, statement := range elifs[0].Head {
		g.generate(&cmdbuf, statement)
		cmdbuf.add(ir.Set{Name: "condition", Value: ir.Literal("shell.ExitCode == 0")})
		cmdbuf.add(ir.Set{Name: "shell.ExitCode", Value: ir.Literal("0")})
	}

	var body InstructionBuffer
	for _, statement := range elifs[0].Body {
		g.generate(&body, statement)
	}
	cmdbuf.add(ir.If{
		Condition: ir.Literal("condition"),
		Body:      body,
		Alternate: g.handleElif(elifs[1:]),
	})

	return cmdbuf

}

func (g *generator) handleLoop(buf *InstructionBuffer, loop ast.Loop) {
	var cmdbuf InstructionBuffer
	cmdbuf.add(ir.CloneStreamManager{})

	g.handleRedirections(&cmdbuf, loop.Redirections)

	var innerBuf, body InstructionBuffer

	for i, statement := range loop.Head {
		g.generate(&body, statement)

		if i < len(loop.Head)-1 {
			continue
		}

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
		g.generate(&body, statement)
	}
	innerBuf.add(ir.Loop{
		Condition: ir.Literal(""),
		Body:      body,
	})

	cmdbuf = append(cmdbuf, innerBuf...)
	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleRangeLoop(buf *InstructionBuffer, loop ast.RangeLoop) {
	var cmdbuf InstructionBuffer
	cmdbuf.add(ir.CloneStreamManager{})

	g.handleRedirections(&cmdbuf, loop.Redirections)

	var innerBuf, body InstructionBuffer

	for _, statement := range loop.Body {
		g.generate(&body, statement)
	}

	var members = ir.Literal("shell.Args")
	if len(loop.Operands) > 0 {
		cmdbuf.add(ir.DeclareSlice{Name: "members"})

		for _, arg := range loop.Operands {
			cmdbuf.add(ir.Append{Name: "members", Value: g.handleExpression(&cmdbuf, arg)})
		}
		members = ir.Literal("members")
	}

	innerBuf.add(ir.RangeLoop{
		Var:     loop.Var,
		Members: members,
		Body:    body,
	})

	cmdbuf = append(cmdbuf, innerBuf...)
	*buf = append(*buf, ir.Closure(cmdbuf))
}
