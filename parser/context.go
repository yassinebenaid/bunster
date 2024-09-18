package parser

import (
	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/token"
)

func (p *Parser) getCommandContextParser(tt token.TokenType) func(*ast.Command) {
	switch tt {
	case token.GT, token.DOUBLE_GT:
		return p.parseStdoutRedirection
	case token.FILE_DESCRIPTOR:
		return p.parseFileDescriptorRedirection
	default:
		return nil
	}
}

func (p *Parser) parseStdoutRedirection(cmd *ast.Command) {
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

func (p *Parser) parseFileDescriptorRedirection(cmd *ast.Command) {
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
