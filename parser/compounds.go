package parser

import (
	"github.com/yassinebenaid/bunny/token"
)

func (p *Parser) isCompound() bool {
	switch p.curr.Type {
	case token.WHILE:
		return true
	}

	return false
}
