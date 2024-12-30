package parser

import (
	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/token"
)

func (p *parser) parseCommandSubstitution() ast.Expression {
	var cmds ast.CommandSubstitution
	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.RIGHT_PAREN && p.curr.Type != token.EOF {
		if p.curr.Type == token.HASH {
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
				p.proceed()
			}
		} else {
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			cmds = append(cmds, cmdList)
			if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if len(cmds) == 0 {
		p.error("expeceted a command list after `$(`")
		return nil
	}

	if p.curr.Type != token.RIGHT_PAREN {
		p.error("unexpected end of file, expeceted `)`")
		return nil
	}

	return cmds
}

func (p *parser) parseProcessSubstitution() ast.Expression {
	tok := p.curr.Literal
	var process ast.ProcessSubstitution

	process.Direction = '>'
	if tok == "<(" {
		process.Direction = '<'
	}

	p.proceed()
	for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
		p.proceed()
	}

	for p.curr.Type != token.RIGHT_PAREN && p.curr.Type != token.EOF {
		if p.curr.Type == token.HASH {
			for p.curr.Type != token.NEWLINE && p.curr.Type != token.EOF {
				p.proceed()
			}
		} else {
			cmdList := p.parseCommandList()
			if cmdList == nil {
				return nil
			}
			process.Body = append(process.Body, cmdList)
			if p.curr.Type == token.SEMICOLON || p.curr.Type == token.AMPERSAND {
				p.proceed()
			}
		}
		for p.curr.Type == token.BLANK || p.curr.Type == token.NEWLINE {
			p.proceed()
		}
	}

	if len(process.Body) == 0 {
		p.error("expeceted a command list after `%s`", tok)
		return nil
	}

	if p.curr.Type != token.RIGHT_PAREN {
		p.error("unexpected end of file, expeceted `)`")
		return nil
	}

	return process
}

func (p *parser) parseParameterExpansion() ast.Expression {
	var exp ast.Expression
	p.proceed()

	if p.curr.Type == token.HASH {
		p.proceed()
		exp = ast.VarCount{Parameter: p.parseParameter()}

		if p.curr.Type != token.RIGHT_BRACE {
			p.error("expected closing brace `}`, found `%s`", p.curr)
			return nil
		}
		return exp
	}

	param := p.parseParameter()

	switch p.curr.Type {
	case token.RIGHT_BRACE:
		if param.Index != nil {
			exp = ast.ParameterExpansion{Name: param.Name, Index: param.Index}
		} else {
			exp = ast.Var(param.Name)
		}
	case token.MINUS, token.COLON_MINUS:
		checkForNull := p.curr.Type == token.COLON_MINUS
		p.proceed()
		exp = ast.VarOrDefault{
			Parameter:    param,
			Default:      p.parseExpansionOperandExpression(0),
			CheckForNull: checkForNull,
		}
	case token.COLON_ASSIGN:
		p.proceed()
		exp = ast.VarOrSet{
			Parameter: param,
			Default:   p.parseExpansionOperandExpression(0),
		}
	case token.COLON_QUESTION:
		p.proceed()
		exp = ast.VarOrFail{
			Parameter: param,
			Error:     p.parseExpansionOperandExpression(0),
		}
	case token.COLON_PLUS:
		p.proceed()
		exp = ast.CheckAndUse{
			Parameter: param,
			Value:     p.parseExpansionOperandExpression(0),
		}
	case token.CIRCUMFLEX, token.DOUBLE_CIRCUMFLEX, token.COMMA, token.DOUBLE_COMMA:
		operator := p.curr.Literal
		p.proceed()
		exp = ast.ChangeCase{
			Parameter: param,
			Operator:  operator,
			Pattern:   p.parseExpansionOperandExpression(0),
		}
	case token.HASH, token.PERCENT, token.DOUBLE_PERCENT:
		operator := p.curr.Literal
		if p.curr.Type == token.HASH && p.next.Type == token.HASH {
			p.proceed()
			operator += p.curr.Literal
		}
		p.proceed()

		exp = ast.MatchAndRemove{
			Parameter: param,
			Operator:  operator,
			Pattern:   p.parseExpansionOperandExpression(0),
		}
	case token.SLASH:
		operator := p.curr.Literal
		p.proceed()
		if p.curr.Type == token.SLASH || p.curr.Type == token.HASH || p.curr.Type == token.PERCENT {
			operator += p.curr.Literal
			p.proceed()
		}

		var pattern ast.Expression
		if p.curr.Type == token.SLASH {
			pattern = ast.Word(p.curr.Literal)
		} else {
			pattern = p.parseExpansionOperandExpression(token.SLASH)
		}

		mar := ast.MatchAndReplace{Parameter: param, Operator: operator, Pattern: pattern}

		if p.curr.Type == token.SLASH {
			p.proceed()
			mar.Value = p.parseExpansionOperandExpression(0)
		}

		exp = mar
	case token.COLON:
		p.proceed()
		slice := ast.Slice{Parameter: param, Offset: p.parseArithmetics()}

		if p.curr.Type == token.COLON {
			p.proceed()
			slice.Length = p.parseArithmetics()
		}

		exp = slice
	case token.AT:
		p.proceed()
		switch p.curr.Literal {
		case "U", "u", "L", "Q", "E", "P", "A", "K", "a", "k":
		default:
			p.error("bad substitution operator `%s`, possible operators are (U, u, L, Q, E, P, A, K, a, k)", p.curr)
			return nil
		}
		exp = ast.Transform{Parameter: param, Operator: p.curr.Literal}
		p.proceed()
	}

	if p.curr.Type != token.RIGHT_BRACE {
		p.error("expected closing brace `}`, found `%s`", p.curr)
		return nil
	}

	return exp
}

func (p *parser) parseExpansionOperandExpression(stopAt token.TokenType) ast.Expression {
	var exprs []ast.Expression
	// TODO: handle special variables in parameter expansion
loop:
	for {
		switch p.curr.Type {
		case token.RIGHT_BRACE, token.EOF:
			break loop
		case token.SIMPLE_EXPANSION:
			exprs = append(exprs, ast.Var(p.curr.Literal))
		case token.SINGLE_QUOTE:
			exprs = append(exprs, p.parseLiteralString())
		case token.DOUBLE_QUOTE:
			exprs = append(exprs, p.parseString())
		case token.DOLLAR_PAREN:
			exprs = append(exprs, p.parseCommandSubstitution())
		case token.GT_PAREN, token.LT_PAREN:
			exprs = append(exprs, p.parseProcessSubstitution())
		case token.DOLLAR_BRACE:
			exprs = append(exprs, p.parseParameterExpansion())
		default:
			if p.curr.Type == stopAt {
				break loop
			}

			exprs = append(exprs, ast.Word(p.curr.Literal))
		}

		p.proceed()
	}

	return concat(exprs, false)
}

func (p *parser) parseParameter() ast.Param {
	if p.curr.Type != token.WORD {
		p.error("couldn't find a valid parameter name, found `%s`", p.curr)
	}

	var param ast.Param
	param.Name = p.curr.Literal

	p.proceed()

	if p.curr.Type == token.LEFT_BRACKET {
		p.proceed()
		param.Index = p.parseArithmetics()
		if p.curr.Type != token.RIGHT_BRACKET {
			p.error("expected a closing bracket `]`, found `%s`", p.curr)
		}
		p.proceed()
	}

	return param
}
