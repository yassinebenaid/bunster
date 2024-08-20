package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		input  string
		tokens []token.Token
	}{
		//Keywords
		{`if`, []token.Token{{Type: token.IF, Literal: `if`}}},
		{`if`, []token.Token{{Type: token.IF, Literal: `if`}}},
		{`then`, []token.Token{{Type: token.THEN, Literal: `then`}}},
		{`else`, []token.Token{{Type: token.ELSE, Literal: `else`}}},
		{`elif`, []token.Token{{Type: token.ELIF, Literal: `elif`}}},
		{`fi`, []token.Token{{Type: token.FI, Literal: `fi`}}},
		{`for`, []token.Token{{Type: token.FOR, Literal: `for`}}},
		{`in`, []token.Token{{Type: token.IN, Literal: `in`}}},
		{`do`, []token.Token{{Type: token.DO, Literal: `do`}}},
		{`done`, []token.Token{{Type: token.DONE, Literal: `done`}}},
		{`while`, []token.Token{{Type: token.WHILE, Literal: `while`}}},
		{`until`, []token.Token{{Type: token.UNTIL, Literal: `until`}}},
		{`case`, []token.Token{{Type: token.CASE, Literal: `case`}}},
		{`esac`, []token.Token{{Type: token.ESAC, Literal: `esac`}}},
		{`function`, []token.Token{{Type: token.FUNCTION, Literal: `function`}}},
		{`select`, []token.Token{{Type: token.SELECT, Literal: `select`}}},
		{`trap`, []token.Token{{Type: token.TRAP, Literal: `trap`}}},
		{`return`, []token.Token{{Type: token.RETURN, Literal: `return`}}},
		{`exit`, []token.Token{{Type: token.EXIT, Literal: `exit`}}},
		{`break`, []token.Token{{Type: token.BREAK, Literal: `break`}}},
		{`continue`, []token.Token{{Type: token.CONTINUE, Literal: `continue`}}},
		{`declare`, []token.Token{{Type: token.DECLARE, Literal: `declare`}}},
		{`local`, []token.Token{{Type: token.LOCAL, Literal: `local`}}},
		{`export`, []token.Token{{Type: token.EXPORT, Literal: `export`}}},
		{`readonly`, []token.Token{{Type: token.READONLY, Literal: `readonly`}}},
		{`unset`, []token.Token{{Type: token.UNSET, Literal: `unset`}}},

		// symbols
		{`+`, []token.Token{{Type: token.PLUS, Literal: `+`}}},
		{`-`, []token.Token{{Type: token.MINUS, Literal: `-`}}},
		{`*`, []token.Token{{Type: token.STAR, Literal: `*`}}},
		{`/`, []token.Token{{Type: token.SLASH, Literal: `/`}}},
		{`%`, []token.Token{{Type: token.PERCENT, Literal: `%`}}},
		{`=`, []token.Token{{Type: token.ASSIGN, Literal: `=`}}},
		{`+=`, []token.Token{{Type: token.PLUS_ASSIGN, Literal: `+=`}}},
		{`-=`, []token.Token{{Type: token.MINUS_ASSIGN, Literal: `-=`}}},
		{`*=`, []token.Token{{Type: token.STAR_ASSIGN, Literal: `*=`}}},
		{`/=`, []token.Token{{Type: token.SLASH_ASSIGN, Literal: `/=`}}},
		{`==`, []token.Token{{Type: token.EQ, Literal: `==`}}},
		{`!=`, []token.Token{{Type: token.NOT_EQ, Literal: `!=`}}},
		{`<`, []token.Token{{Type: token.LT, Literal: `<`}}},
		{`<=`, []token.Token{{Type: token.LE, Literal: `<=`}}},
		{`>`, []token.Token{{Type: token.GT, Literal: `>`}}},
		{`>=`, []token.Token{{Type: token.GE, Literal: `>=`}}},
		{`&&`, []token.Token{{Type: token.AND, Literal: `&&`}}},
		{`||`, []token.Token{{Type: token.OR, Literal: `||`}}},
		{`|`, []token.Token{{Type: token.PIPE, Literal: `|`}}},
		{`&`, []token.Token{{Type: token.AMPERSAND, Literal: `&`}}},
		{`>>`, []token.Token{{Type: token.DOUBLE_GT, Literal: `>>`}}},
		{`<<`, []token.Token{{Type: token.DOUBLE_LT, Literal: `<<`}}},
		{`<<-`, []token.Token{{Type: token.DOUBLE_LT_MINUS, Literal: `<<-`}}},
		{`<<<`, []token.Token{{Type: token.TRIPLE_LT, Literal: `<<<`}}},
		{`>&`, []token.Token{{Type: token.GT_AMPERSAND, Literal: `>&`}}},
		{`<&`, []token.Token{{Type: token.LT_AMPERSAND, Literal: `<&`}}},
		{`|&`, []token.Token{{Type: token.PIPE_AMPERSAND, Literal: `|&`}}},
		{`&>`, []token.Token{{Type: token.AMPERSAND_GT, Literal: `&>`}}},
		{`>|`, []token.Token{{Type: token.GT_PIPE, Literal: `>|`}}},
		{`<>`, []token.Token{{Type: token.LT_GT, Literal: `<>`}}},
		{`;`, []token.Token{{Type: token.SEMICOLON, Literal: `;`}}},
		{`;;`, []token.Token{{Type: token.DOUBLE_SEMICOLON, Literal: `;;`}}},
		{`(`, []token.Token{{Type: token.LEFT_PAREN, Literal: `(`}}},
		{`)`, []token.Token{{Type: token.RIGHT_PAREN, Literal: `)`}}},
		{`((`, []token.Token{{Type: token.DOUBLE_LEFT_PAREN, Literal: `((`}}},
		{`))`, []token.Token{{Type: token.DOUBLE_RIGHT_PAREN, Literal: `))`}}},
		{`[`, []token.Token{{Type: token.LEFT_BRACKET, Literal: `[`}}},
		{`]`, []token.Token{{Type: token.RIGHT_BRACKET, Literal: `]`}}},
		{`[[`, []token.Token{{Type: token.DOUBLE_LEFT_BRACKET, Literal: `[[`}}},
		{`]]`, []token.Token{{Type: token.DOUBLE_RIGHT_BRACKET, Literal: `]]`}}},
		{`{`, []token.Token{{Type: token.LEFT_BRACE, Literal: `{`}}},
		{`}`, []token.Token{{Type: token.RIGHT_BRACE, Literal: `}`}}},
		{`,`, []token.Token{{Type: token.COMMA, Literal: `,`}}},
		{`,,`, []token.Token{{Type: token.DOUBLE_COMMA, Literal: `,,`}}},
		{`:`, []token.Token{{Type: token.COLON, Literal: `:`}}},
		// {`\`, []token.Token{{Type: token.BACKSLASH, Literal: `\`}}}, // TODO: see what to do with this
		// {`$`, []token.Token{{Type: token.DOLLAR, Literal: "$"}}}, // TODO
		// {`"`, []token.Token{{Type: token.DOUBLE_QUOTE, Literal: `"`}}}, // TODO
		{`?`, []token.Token{{Type: token.QUESTION, Literal: `?`}}},
		{`!`, []token.Token{{Type: token.EXCLAMATION, Literal: `!`}}},
		{`#`, []token.Token{{Type: token.HASH, Literal: `#`}}},
		{`${`, []token.Token{{Type: token.DOLLAR_BRACE, Literal: `${`}}},
		{`$(`, []token.Token{{Type: token.DOLLAR_PAREN, Literal: `$(`}}},
		{`$((`, []token.Token{{Type: token.DOLLAR_DOUBLE_PAREN, Literal: `$((`}}},
		{`>(`, []token.Token{{Type: token.LT_PAREN, Literal: `>(`}}},
		{`<(`, []token.Token{{Type: token.GT_PAREN, Literal: `<(`}}},
		{`^`, []token.Token{{Type: token.CIRCUMFLEX, Literal: `^`}}},
		{`^^`, []token.Token{{Type: token.DOUBLE_CIRCUMFLEX, Literal: `^^`}}},
		{`:=`, []token.Token{{Type: token.COLON_ASSIGN, Literal: `:=`}}},
		{`:-`, []token.Token{{Type: token.COLON_MINUS, Literal: `:-`}}},
		{`:+`, []token.Token{{Type: token.COLON_PLUS, Literal: `:+`}}},
		{`:?`, []token.Token{{Type: token.COLON_QUESTION, Literal: `:?`}}},
		{`//`, []token.Token{{Type: token.DOUBLE_SLASH, Literal: `//`}}},
		{`..`, []token.Token{{Type: token.DOUBLE_DOT, Literal: `..`}}},
		{`++`, []token.Token{{Type: token.INCREMENT, Literal: `++`}}},
		{`--`, []token.Token{{Type: token.DECREMENT, Literal: `--`}}},
		{`~`, []token.Token{{Type: token.TILDE, Literal: `~`}}},
	}

	for i, tc := range testCases {
		l := lexer.New([]byte(tc.input))
		for _, tn := range tc.tokens {
			if result := l.NextToken(); tn.Type != result.Type {
				t.Fatalf(`#%d: wrong token type for %q, want=%d got=%d`, i, tn.Literal, tn.Type, result.Type)
			} else if tn.Literal != result.Literal {
				t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.Literal, tn.Literal, i)
			}
		}
	}
}
