package parser

import (
	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
)

func New(l lexer.Lexer) Parser {
	var p = Parser{
		l: l,
	}
	return p
}

type Parser struct {
	l            lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
}

func (p *Parser) proceed() {
	p.currentToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) ParseScript() ast.Script {
	var script ast.Script

loop:
	for {
		p.proceed()
		switch p.currentToken.Type {
		case token.EOF:
			break loop
		default:
			script.Statements = append(script.Statements, p.parseCommand())
		}
	}

	return script
}

func (p *Parser) parseCommand() ast.Command {
	var cc ast.Command

	cc.Name = p.currentToken.Literal

	for {
		if p.nextToken.Type == token.BLANK {
			p.proceed()
		}
		if p.nextToken.Type != token.Word {
			break
		}

		p.proceed()
		cc.Args = append(cc.Args, p.currentToken.Literal)
	}

	return cc
}
