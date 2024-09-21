package parser

import (
	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
)

func New(l lexer.Lexer) Parser {
	var p = Parser{l: l}

	// So that both curr and next tokens get initialized.
	p.proceed()
	p.proceed()

	return p
}

type Parser struct {
	l    lexer.Lexer
	curr token.Token
	next token.Token
}

func (p *Parser) proceed() {
	p.curr = p.next
	p.next = p.l.NextToken()
}

func (p *Parser) ParseScript() ast.Script {
	var script ast.Script

	for ; p.curr.Type != token.EOF; p.proceed() {
		switch p.curr.Type {
		default:
			script.Statements = append(script.Statements, p.parseCommand())
		}
	}

	return script
}

func (p *Parser) parseCommand() ast.Command {
	var cmd ast.Command

	cmd.Name = p.parseField()

loop:
	for {
		switch {
		case p.curr.Type.Is(token.BLANK):
			break
		case p.curr.Type.Is(token.EOF):
			break loop
		case p.isRedirectionToken(p.curr):
			p.HandleRedirection(&cmd)
		default:
			cmd.Args = append(cmd.Args, p.parseField())
		}

		if !p.isRedirectionToken(p.curr) {
			p.proceed()
		}
	}

	return cmd
}

func (p *Parser) parseField() ast.Node {
	var nodes []ast.Node

loop:
	for {
		switch p.curr.Type {
		case token.BLANK, token.EOF:
			break loop
		case token.SIMPLE_EXPANSION:
			nodes = append(nodes, ast.SimpleExpansion(p.curr.Literal))
		case token.SINGLE_QUOTE:
			nodes = append(nodes, p.parseLiteralString())
		case token.DOUBLE_QUOTE:
			nodes = append(nodes, p.parseString())
		default:
			if p.isRedirectionToken(p.curr) {
				break loop
			}

			nodes = append(nodes, ast.Word(p.curr.Literal))
			// TODO: handle error
		}

		p.proceed()
	}

	return concat(nodes)
}

func (p *Parser) parseLiteralString() ast.Word {
	p.proceed()

	if p.curr.Type == token.SINGLE_QUOTE {
		return ast.Word("")
	}

	word := p.curr.Literal
	p.proceed()

	if p.curr.Type != token.SINGLE_QUOTE {
		//TODO: handle error here
		panic("TODO: handle error here")
	}

	return ast.Word(word)
}

func (p *Parser) parseString() ast.Node {
	p.proceed()

	if p.curr.Type == token.DOUBLE_QUOTE {
		return ast.Word("")
	}

	var nodes []ast.Node

loop:
	for {
		switch p.curr.Type {
		case token.DOUBLE_QUOTE, token.EOF:
			break loop
		case token.SIMPLE_EXPANSION:
			nodes = append(nodes, ast.SimpleExpansion(p.curr.Literal))
		default:
			nodes = append(nodes, ast.Word(p.curr.Literal))
		}

		p.proceed()
	}

	if p.curr.Type != token.DOUBLE_QUOTE {
		//TODO: handle error here
		panic("TODO: handle error here")
	}

	return concat(nodes)
}

func concat(n []ast.Node) ast.Node {
	var conc ast.Concatination
	var mergedWords ast.Word
	var hasWords bool

	for i, node := range n {

		if w, ok := node.(ast.Word); ok {
			mergedWords += w
			hasWords = true
		} else {
			if hasWords {
				conc.Nodes = append(conc.Nodes, mergedWords)
				mergedWords, hasWords = "", false
			}
			conc.Nodes = append(conc.Nodes, node)
		}

		if i == len(n)-1 && hasWords {
			conc.Nodes = append(conc.Nodes, mergedWords)
		}
	}

	if len(conc.Nodes) == 1 {
		return conc.Nodes[0]
	}

	return conc
}
