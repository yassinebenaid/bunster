package parser

import (
	"github.com/yassinebenaid/bunny/ast"
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) isRedirectionToken() bool {
	switch p.curr.Type {
	case token.GT, token.DOUBLE_GT, token.AMPERSAND_GT, token.GT_AMPERSAND,
		token.LT, token.LT_AMPERSAND, token.TRIPLE_LT, token.GT_PIPE, token.AMPERSAND_DOUBLE_GT:
		return true
	case token.INT:
		// redirections that use file descriptor as source
		switch p.next.Type {
		case token.GT, token.DOUBLE_GT, token.GT_AMPERSAND, token.LT, token.LT_AMPERSAND, token.TRIPLE_LT, token.GT_PIPE:
			return true
		}
	}

	return false
}

func (p *Parser) HandleRedirection(rt *[]ast.Redirection) {
	switch p.curr.Type {
	case token.GT, token.DOUBLE_GT, token.GT_AMPERSAND, token.GT_PIPE:
		p.fromStdout(rt)
	case token.LT_AMPERSAND, token.LT, token.TRIPLE_LT:
		p.toStdin(rt)
	case token.AMPERSAND_GT, token.AMPERSAND_DOUBLE_GT:
		p.allOutputsToFile(rt)
	case token.INT:
		p.fromFileDescriptor(rt)
	}
}

func (p *Parser) fromStdout(rt *[]ast.Redirection) {
	var r ast.Redirection
	r.Src = "1"
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()

	if r.Dst == nil {
		p.error("a redirection operand was not provided after the `%s`", r.Method)
	}
	*rt = append(*rt, r)
}

func (p *Parser) fromFileDescriptor(rt *[]ast.Redirection) {
	var r ast.Redirection
	r.Src = p.curr.Literal

	p.proceed()
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
	if r.Dst == nil {
		p.error("a redirection operand was not provided after the `%s`", r.Method)
	}

	*rt = append(*rt, r)
}

func (p *Parser) allOutputsToFile(rt *[]ast.Redirection) {
	var r ast.Redirection
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
	if r.Dst == nil {
		p.error("a redirection operand was not provided after the `%s`", r.Method)
	}
	*rt = append(*rt, r)
}

func (p *Parser) toStdin(rt *[]ast.Redirection) {
	var r ast.Redirection
	r.Src = "0"
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
	if r.Dst == nil {
		p.error("a redirection operand was not provided after the `%s`", r.Method)
	}

	*rt = append(*rt, r)
}
