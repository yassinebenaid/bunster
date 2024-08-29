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

	cmd.Name = p.parseSentence()

loop:
	for ; ; p.proceed() {
		switch p.curr.Type {
		case token.BLANK:
			continue
		case token.EOF:
			break loop
		default:
			cmd.Args = append(cmd.Args, p.parseSentence())
		}

	}

	return cmd
}

func (p *Parser) parseSentence() ast.Node {
	var nodes []ast.Node

loop:
	for {
		switch p.curr.Type {
		case token.BLANK, token.EOF:
			break loop
		case token.SIMPLE_EXPANSION:
			nodes = append(nodes, ast.SimpleExpansion(p.curr.Literal))
		default:
			nodes = append(nodes, ast.Word(p.curr.Literal))
			// TODO: handle error
		}

		p.proceed()
	}

	var conc ast.Concatination
	var word ast.Word

	for _, node := range nodes {
		w, ok := node.(ast.Word)
		if ok {
			word += w
		} else {
			if word != "" {
				conc.Nodes = append(conc.Nodes, word)
				word = ""
			}
			conc.Nodes = append(conc.Nodes, node)
		}
	}

	if word != "" {
		conc.Nodes = append(conc.Nodes, word)
	}

	if len(conc.Nodes) == 1 {
		return conc.Nodes[0]
	}

	return conc
}
