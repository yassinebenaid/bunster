package parser

import (
	"fmt"

	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/token"
)

func (p *Parser) isRedirectionToken() bool {
	switch p.curr.Type {
	case token.GT, token.DOUBLE_GT, token.AMPERSAND_GT:
		return true
	case token.INT:
		switch p.next.Type {
		case token.GT, token.DOUBLE_GT:
			return true
		}
		fallthrough
	default:
		return false
	}
}

func (p *Parser) HandleRedirection(cmd *ast.Command) {
	switch p.curr.Type {
	case token.GT, token.DOUBLE_GT:
		p.fromStdoutToFile(cmd)
	case token.AMPERSAND_GT:
		p.allOutputsToFile(cmd)
	case token.INT:
		p.fromFdToFile(cmd)
	default:
		panic(fmt.Sprintf("unhandled redirection token: %q", p.curr.Literal))
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
	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) fromFdToFile(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor(p.curr.Literal)

	p.proceed()
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	r.Dst = p.parseField()
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

func (p *Parser) parseStdoutToFileDescriptorRedirection(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor("1")
	r.Method = p.curr.Literal

	p.proceed()
	r.Dst = ast.FileDescriptor(p.curr.Literal)
	p.proceed()

	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) parseFileDescriptorRedirection(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor(p.curr.Literal)

	p.proceed()
	r.Method = p.curr.Literal
	p.proceed()

	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.FILE_DESCRIPTOR {
		r.Dst = ast.FileDescriptor(p.curr.Literal)
		p.proceed()
	} else {
		r.Dst = p.parseField()
	}

	cmd.Redirections = append(cmd.Redirections, r)
}

func (p *Parser) parseStdinFromFileRedirection(cmd *ast.Command) {
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
