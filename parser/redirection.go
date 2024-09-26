package parser

import (
	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/token"
)

func (p *Parser) isRedirectionToken() bool {
	switch p.curr.Type {
	case token.GT, token.DOUBLE_GT, token.AMPERSAND_GT, token.GT_AMPERSAND, token.LT, token.LT_AMPERSAND, token.TRIPLE_LT:
		return true
	case token.INT:
		// redirections that use file descriptor as source
		switch p.next.Type {
		case token.GT, token.DOUBLE_GT, token.GT_AMPERSAND, token.LT, token.LT_AMPERSAND, token.TRIPLE_LT:
			return true
		}
	}

	return false
}

func (p *Parser) HandleRedirection(cmd *ast.Command) {
	switch p.curr.Type {
	case token.GT, token.DOUBLE_GT:
		p.fromStdoutToFile(cmd)
	case token.AMPERSAND_GT:
		p.allOutputsToFile(cmd)
	case token.GT_AMPERSAND:
		p.fromStdoutToFd(cmd)
	case token.LT_AMPERSAND, token.LT, token.TRIPLE_LT:
		p.toStdin(cmd)
	case token.INT:
		p.fromOrToFileDescriptor(cmd)
	}
}

func (p *Parser) fromStdoutToFile(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor("1")
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()

	if r.Dst == nil {
		p.error("a file name was not provided after the `%s`", r.Method)
	}
	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) fromOrToFileDescriptor(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor(p.curr.Literal)

	p.proceed()
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
	if r.Dst == nil {
		p.error("a file name was not provided after the `%s`", r.Method)
	}

	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) allOutputsToFile(cmd *ast.Command) {
	var r ast.Redirection
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) fromStdoutToFd(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor("1")
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
	if r.Dst == nil {
		p.error("a file name was not provided after the `%s`", r.Method)
	}

	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) toStdin(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor("0")
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) toStdinFromFd(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor("0")
	r.Method = p.curr.Literal

	p.proceed()
	r.Dst = ast.FileDescriptor(p.curr.Literal)
	p.proceed()

	cmd.Redirections = append(cmd.Redirections, r)
}
