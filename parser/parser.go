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

func (p *Parser) ParseProgram() ast.Program {
	var program ast.Program

loop:
	for {
		p.proceed()
		switch p.currentToken.Type {
		case token.IDENT:
			program.Nodes = append(program.Nodes, p.parseCommandCall())
		case token.EOF:
			break loop
		}
	}

	return program
}

func (p *Parser) parseCommandCall() ast.CommandCall {
	var cc ast.CommandCall

	cc.Command = p.currentToken.Literal

	for {
		if p.nextToken.Type != token.IDENT {
			break
		}

		p.proceed()
		cc.Args = append(cc.Args, p.currentToken.Literal)
	}

	return cc
}
