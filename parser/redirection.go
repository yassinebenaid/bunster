package parser

import (
	"fmt"

	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/token"
)

func (p *Parser) isRedirectionToken() bool {
	switch p.curr.Type {
	case token.GT, token.DOUBLE_GT, token.AMPERSAND_GT, token.GT_AMPERSAND, token.LT:
		return true
	case token.INT:
		// redirections that use file descriptor as source
		switch p.next.Type {
		case token.GT, token.DOUBLE_GT, token.GT_AMPERSAND, token.LT:
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
	case token.GT_AMPERSAND:
		p.fromStdoutToFd(cmd)
	case token.LT:
		p.toStdin(cmd)
	case token.INT:
		p.fromFileDescriptor(cmd)
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

func (p *Parser) fromFileDescriptor(cmd *ast.Command) {
	var r ast.Redirection
	r.Src = ast.FileDescriptor(p.curr.Literal)

	p.proceed()
	r.Method = p.curr.Literal

	p.proceed()
	if p.curr.Type == token.BLANK {
		p.proceed()
	}

	if p.curr.Type == token.INT {
		r.Dst = ast.FileDescriptor(p.curr.Literal)
	} else {
		r.Dst = p.parseField()
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
	r.Dst = ast.FileDescriptor(p.curr.Literal)
	p.proceed()

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
