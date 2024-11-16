package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/bunny/lexer"
	"github.com/yassinebenaid/bunny/token"
	"github.com/yassinebenaid/godump"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

func TestLexer(t *testing.T) {
	testCases := []struct {
		input  string
		tokens []token.Token
	}{
		//Keywords
		{`if`, []token.Token{{Type: token.IF, Literal: `if`, Line: 1, Position: 1}}},
		{`if`, []token.Token{{Type: token.IF, Literal: `if`, Line: 1, Position: 1}}},
		{`then`, []token.Token{{Type: token.THEN, Literal: `then`, Line: 1, Position: 1}}},
		{`else`, []token.Token{{Type: token.ELSE, Literal: `else`, Line: 1, Position: 1}}},
		{`elif`, []token.Token{{Type: token.ELIF, Literal: `elif`, Line: 1, Position: 1}}},
		{`fi`, []token.Token{{Type: token.FI, Literal: `fi`, Line: 1, Position: 1}}},
		{`for`, []token.Token{{Type: token.FOR, Literal: `for`, Line: 1, Position: 1}}},
		{`in`, []token.Token{{Type: token.IN, Literal: `in`, Line: 1, Position: 1}}},
		{`do`, []token.Token{{Type: token.DO, Literal: `do`, Line: 1, Position: 1}}},
		{`done`, []token.Token{{Type: token.DONE, Literal: `done`, Line: 1, Position: 1}}},
		{`while`, []token.Token{{Type: token.WHILE, Literal: `while`, Line: 1, Position: 1}}},
		{`until`, []token.Token{{Type: token.UNTIL, Literal: `until`, Line: 1, Position: 1}}},
		{`case`, []token.Token{{Type: token.CASE, Literal: `case`, Line: 1, Position: 1}}},
		{`esac`, []token.Token{{Type: token.ESAC, Literal: `esac`, Line: 1, Position: 1}}},
		{`function`, []token.Token{{Type: token.FUNCTION, Literal: `function`, Line: 1, Position: 1}}},
		{`select`, []token.Token{{Type: token.SELECT, Literal: `select`, Line: 1, Position: 1}}},
		{`trap`, []token.Token{{Type: token.TRAP, Literal: `trap`, Line: 1, Position: 1}}},
		{`return`, []token.Token{{Type: token.RETURN, Literal: `return`, Line: 1, Position: 1}}},
		{`exit`, []token.Token{{Type: token.EXIT, Literal: `exit`, Line: 1, Position: 1}}},
		{`break`, []token.Token{{Type: token.BREAK, Literal: `break`, Line: 1, Position: 1}}},
		{`continue`, []token.Token{{Type: token.CONTINUE, Literal: `continue`, Line: 1, Position: 1}}},
		{`declare`, []token.Token{{Type: token.DECLARE, Literal: `declare`, Line: 1, Position: 1}}},
		{`local`, []token.Token{{Type: token.LOCAL, Literal: `local`, Line: 1, Position: 1}}},
		{`export`, []token.Token{{Type: token.EXPORT, Literal: `export`, Line: 1, Position: 1}}},
		{`readonly`, []token.Token{{Type: token.READONLY, Literal: `readonly`, Line: 1, Position: 1}}},
		{`unset`, []token.Token{{Type: token.UNSET, Literal: `unset`, Line: 1, Position: 1}}},

		// symbols
		{`+`, []token.Token{{Type: token.PLUS, Literal: `+`, Line: 1, Position: 1}}},
		{`-`, []token.Token{{Type: token.MINUS, Literal: `-`, Line: 1, Position: 1}}},
		{`*`, []token.Token{{Type: token.STAR, Literal: `*`, Line: 1, Position: 1}}},
		{`**`, []token.Token{{Type: token.EXPONENTIATION, Literal: `**`, Line: 1, Position: 1}}},
		{`/`, []token.Token{{Type: token.SLASH, Literal: `/`, Line: 1, Position: 1}}},
		{`%`, []token.Token{{Type: token.PERCENT, Literal: `%`, Line: 1, Position: 1}}},
		{`%=`, []token.Token{{Type: token.PERCENT_ASSIGN, Literal: `%=`, Line: 1, Position: 1}}},
		{`%%`, []token.Token{{Type: token.DOUBLE_PERCENT, Literal: `%%`, Line: 1, Position: 1}}},
		{`=`, []token.Token{{Type: token.ASSIGN, Literal: `=`, Line: 1, Position: 1}}},
		{`+=`, []token.Token{{Type: token.PLUS_ASSIGN, Literal: `+=`, Line: 1, Position: 1}}},
		{`-=`, []token.Token{{Type: token.MINUS_ASSIGN, Literal: `-=`, Line: 1, Position: 1}}},
		{`*=`, []token.Token{{Type: token.STAR_ASSIGN, Literal: `*=`, Line: 1, Position: 1}}},
		{`/=`, []token.Token{{Type: token.SLASH_ASSIGN, Literal: `/=`, Line: 1, Position: 1}}},
		{`==`, []token.Token{{Type: token.EQ, Literal: `==`, Line: 1, Position: 1}}},
		{`!=`, []token.Token{{Type: token.NOT_EQ, Literal: `!=`, Line: 1, Position: 1}}},
		{`=~`, []token.Token{{Type: token.EQ_TILDE, Literal: `=~`, Line: 1, Position: 1}}},
		{`<`, []token.Token{{Type: token.LT, Literal: `<`, Line: 1, Position: 1}}},
		{`<=`, []token.Token{{Type: token.LT_EQ, Literal: `<=`, Line: 1, Position: 1}}},
		{`>`, []token.Token{{Type: token.GT, Literal: `>`, Line: 1, Position: 1}}},
		{`>=`, []token.Token{{Type: token.GT_EQ, Literal: `>=`, Line: 1, Position: 1}}},
		{`&&`, []token.Token{{Type: token.AND, Literal: `&&`, Line: 1, Position: 1}}},
		{`||`, []token.Token{{Type: token.OR, Literal: `||`, Line: 1, Position: 1}}},
		{`|`, []token.Token{{Type: token.PIPE, Literal: `|`, Line: 1, Position: 1}}},
		{`|=`, []token.Token{{Type: token.PIPE_ASSIGN, Literal: `|=`, Line: 1, Position: 1}}},
		{`&`, []token.Token{{Type: token.AMPERSAND, Literal: `&`, Line: 1, Position: 1}}},
		{`&=`, []token.Token{{Type: token.AMPERSAND_ASSIGN, Literal: `&=`, Line: 1, Position: 1}}},
		{`>>`, []token.Token{{Type: token.DOUBLE_GT, Literal: `>>`, Line: 1, Position: 1}}},
		{`>>=`, []token.Token{{Type: token.DOUBLE_GT_ASSIGN, Literal: `>>=`, Line: 1, Position: 1}}},
		{`<<`, []token.Token{{Type: token.DOUBLE_LT, Literal: `<<`, Line: 1, Position: 1}}},
		{`<<=`, []token.Token{{Type: token.DOUBLE_LT_ASSIGN, Literal: `<<=`, Line: 1, Position: 1}}},
		{`<<-`, []token.Token{{Type: token.DOUBLE_LT_MINUS, Literal: `<<-`, Line: 1, Position: 1}}},
		{`<<<`, []token.Token{{Type: token.TRIPLE_LT, Literal: `<<<`, Line: 1, Position: 1}}},
		{`>&`, []token.Token{{Type: token.GT_AMPERSAND, Literal: `>&`, Line: 1, Position: 1}}},
		{`<&`, []token.Token{{Type: token.LT_AMPERSAND, Literal: `<&`, Line: 1, Position: 1}}},
		{`|&`, []token.Token{{Type: token.PIPE_AMPERSAND, Literal: `|&`, Line: 1, Position: 1}}},
		{`&>`, []token.Token{{Type: token.AMPERSAND_GT, Literal: `&>`, Line: 1, Position: 1}}},
		{`&>>`, []token.Token{{Type: token.AMPERSAND_DOUBLE_GT, Literal: `&>>`, Line: 1, Position: 1}}},
		{`>|`, []token.Token{{Type: token.GT_PIPE, Literal: `>|`, Line: 1, Position: 1}}},
		{`<>`, []token.Token{{Type: token.LT_GT, Literal: `<>`, Line: 1, Position: 1}}},
		{`;`, []token.Token{{Type: token.SEMICOLON, Literal: `;`, Line: 1, Position: 1}}},
		{`(`, []token.Token{{Type: token.LEFT_PAREN, Literal: `(`, Line: 1, Position: 1}}},
		{`)`, []token.Token{{Type: token.RIGHT_PAREN, Literal: `)`, Line: 1, Position: 1}}},
		{`((`, []token.Token{{Type: token.DOUBLE_LEFT_PAREN, Literal: `((`, Line: 1, Position: 1}}},
		{`[`, []token.Token{{Type: token.LEFT_BRACKET, Literal: `[`, Line: 1, Position: 1}}},
		{`]`, []token.Token{{Type: token.RIGHT_BRACKET, Literal: `]`, Line: 1, Position: 1}}},
		{`[[`, []token.Token{{Type: token.DOUBLE_LEFT_BRACKET, Literal: `[[`, Line: 1, Position: 1}}},
		{`]]`, []token.Token{{Type: token.DOUBLE_RIGHT_BRACKET, Literal: `]]`, Line: 1, Position: 1}}},
		{`{`, []token.Token{{Type: token.LEFT_BRACE, Literal: `{`, Line: 1, Position: 1}}},
		{`}`, []token.Token{{Type: token.RIGHT_BRACE, Literal: `}`, Line: 1, Position: 1}}},
		{`,`, []token.Token{{Type: token.COMMA, Literal: `,`, Line: 1, Position: 1}}},
		{`,,`, []token.Token{{Type: token.DOUBLE_COMMA, Literal: `,,`, Line: 1, Position: 1}}},
		{`:`, []token.Token{{Type: token.COLON, Literal: `:`, Line: 1, Position: 1}}},
		{`"`, []token.Token{{Type: token.DOUBLE_QUOTE, Literal: `"`, Line: 1, Position: 1}}},
		{`'`, []token.Token{{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1}}},
		{`?`, []token.Token{{Type: token.QUESTION, Literal: `?`, Line: 1, Position: 1}}},
		{`!`, []token.Token{{Type: token.EXCLAMATION, Literal: `!`, Line: 1, Position: 1}}},
		{`#`, []token.Token{{Type: token.HASH, Literal: `#`, Line: 1, Position: 1}}},
		{`${`, []token.Token{{Type: token.DOLLAR_BRACE, Literal: `${`, Line: 1, Position: 1}}},
		{`$(`, []token.Token{{Type: token.DOLLAR_PAREN, Literal: `$(`, Line: 1, Position: 1}}},
		{`$((`, []token.Token{{Type: token.DOLLAR_DOUBLE_PAREN, Literal: `$((`, Line: 1, Position: 1}}},
		{`>(`, []token.Token{{Type: token.GT_PAREN, Literal: `>(`, Line: 1, Position: 1}}},
		{`<(`, []token.Token{{Type: token.LT_PAREN, Literal: `<(`, Line: 1, Position: 1}}},
		{`^`, []token.Token{{Type: token.CIRCUMFLEX, Literal: `^`, Line: 1, Position: 1}}},
		{`^^`, []token.Token{{Type: token.DOUBLE_CIRCUMFLEX, Literal: `^^`, Line: 1, Position: 1}}},
		{`^=`, []token.Token{{Type: token.CIRCUMFLEX_ASSIGN, Literal: `^=`, Line: 1, Position: 1}}},
		{`:=`, []token.Token{{Type: token.COLON_ASSIGN, Literal: `:=`, Line: 1, Position: 1}}},
		{`:-`, []token.Token{{Type: token.COLON_MINUS, Literal: `:-`, Line: 1, Position: 1}}},
		{`:+`, []token.Token{{Type: token.COLON_PLUS, Literal: `:+`, Line: 1, Position: 1}}},
		{`:?`, []token.Token{{Type: token.COLON_QUESTION, Literal: `:?`, Line: 1, Position: 1}}},
		{`..`, []token.Token{{Type: token.DOUBLE_DOT, Literal: `..`, Line: 1, Position: 1}}},
		{`++`, []token.Token{{Type: token.INCREMENT, Literal: `++`, Line: 1, Position: 1}}},
		{`--`, []token.Token{{Type: token.DECREMENT, Literal: `--`, Line: 1, Position: 1}}},
		{`~`, []token.Token{{Type: token.TILDE, Literal: `~`, Line: 1, Position: 1}}},
		{`@`, []token.Token{{Type: token.AT, Literal: `@`, Line: 1, Position: 1}}},

		// identifiers
		{`foo bar foo-bar foo_bar`, []token.Token{
			{Type: token.WORD, Literal: `foo`, Line: 1, Position: 1},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 4},
			{Type: token.WORD, Literal: `bar`, Line: 1, Position: 5},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 8},
			{Type: token.WORD, Literal: `foo`, Line: 1, Position: 9},
			{Type: token.MINUS, Literal: `-`, Line: 1, Position: 12},
			{Type: token.WORD, Literal: `bar`, Line: 1, Position: 13},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 16},
			{Type: token.WORD, Literal: `foo_bar`, Line: 1, Position: 17},
		}},

		// Special Variables
		{`$0$1$2 $3$4 $5 $6 $7 $8 $9 $10`, []token.Token{
			{Type: token.SPECIAL_VAR, Literal: "0", Line: 1, Position: 1},
			{Type: token.SPECIAL_VAR, Literal: "1", Line: 1, Position: 3},
			{Type: token.SPECIAL_VAR, Literal: "2", Line: 1, Position: 5},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 7},
			{Type: token.SPECIAL_VAR, Literal: "3", Line: 1, Position: 8},
			{Type: token.SPECIAL_VAR, Literal: "4", Line: 1, Position: 10},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 12},
			{Type: token.SPECIAL_VAR, Literal: "5", Line: 1, Position: 13},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 15},
			{Type: token.SPECIAL_VAR, Literal: "6", Line: 1, Position: 16},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 18},
			{Type: token.SPECIAL_VAR, Literal: "7", Line: 1, Position: 19},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 21},
			{Type: token.SPECIAL_VAR, Literal: "8", Line: 1, Position: 22},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 24},
			{Type: token.SPECIAL_VAR, Literal: "9", Line: 1, Position: 25},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 27},
			{Type: token.SPECIAL_VAR, Literal: "1", Line: 1, Position: 28}, // just to emphasize that only first digit is considered
			{Type: token.INT, Literal: "0", Line: 1, Position: 30},         // just to emphasize that only first digit is considered
		}},
		{`$1something`, []token.Token{
			{Type: token.SPECIAL_VAR, Literal: "1", Line: 1, Position: 1},
			{Type: token.WORD, Literal: "something", Line: 1, Position: 3},
		}},
		{`$$ $@ $? $# $! $_ $*`, []token.Token{
			{Type: token.SPECIAL_VAR, Literal: "$", Line: 1, Position: 1},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 3},
			{Type: token.SPECIAL_VAR, Literal: "@", Line: 1, Position: 4},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 6},
			{Type: token.SPECIAL_VAR, Literal: "?", Line: 1, Position: 7},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 9},
			{Type: token.SPECIAL_VAR, Literal: "#", Line: 1, Position: 10},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 12},
			{Type: token.SPECIAL_VAR, Literal: "!", Line: 1, Position: 13},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 15},
			{Type: token.SPECIAL_VAR, Literal: "_", Line: 1, Position: 16},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 18},
			{Type: token.SPECIAL_VAR, Literal: "*", Line: 1, Position: 19},
		}},
		// Simple expansion
		{`$variable_name $variable-name $concatinated$VAIABLE$VAR_0987654321`, []token.Token{
			{Type: token.SIMPLE_EXPANSION, Literal: `variable_name`, Line: 1, Position: 1},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 15},
			{Type: token.SIMPLE_EXPANSION, Literal: `variable`, Line: 1, Position: 16},
			{Type: token.MINUS, Literal: `-`, Line: 1, Position: 25},
			{Type: token.WORD, Literal: `name`, Line: 1, Position: 26},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 30},
			{Type: token.SIMPLE_EXPANSION, Literal: `concatinated`, Line: 1, Position: 31},
			{Type: token.SIMPLE_EXPANSION, Literal: `VAIABLE`, Line: 1, Position: 44},
			{Type: token.SIMPLE_EXPANSION, Literal: `VAR_0987654321`, Line: 1, Position: 52},
		}},
		// Numbers
		{`0123456789 abc1234 123.456 .123 123. 1.2.3 .abc 1.c 12.34abc`, []token.Token{
			{Type: token.INT, Literal: `0123456789`, Line: 1, Position: 1},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 11},
			{Type: token.WORD, Literal: `abc`, Line: 1, Position: 12},
			{Type: token.INT, Literal: `1234`, Line: 1, Position: 15},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 19},
			{Type: token.FLOAT, Literal: `123.456`, Line: 1, Position: 20},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 27},
			{Type: token.FLOAT, Literal: `.123`, Line: 1, Position: 28},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 32},
			{Type: token.FLOAT, Literal: `123.`, Line: 1, Position: 33},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 37},
			{Type: token.FLOAT, Literal: `1.2`, Line: 1, Position: 38},
			{Type: token.FLOAT, Literal: `.3`, Line: 1, Position: 41},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 43},
			{Type: token.OTHER, Literal: `.`, Line: 1, Position: 44},
			{Type: token.WORD, Literal: `abc`, Line: 1, Position: 45},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 48},
			{Type: token.FLOAT, Literal: `1.`, Line: 1, Position: 49},
			{Type: token.WORD, Literal: `c`, Line: 1, Position: 51},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 52},
			{Type: token.FLOAT, Literal: `12.34`, Line: 1, Position: 53},
			{Type: token.WORD, Literal: `abc`, Line: 1, Position: 58},
		}},
		// Blank
		{"  	\t", []token.Token{
			{Type: token.BLANK, Literal: "  	\t", Line: 1, Position: 1},
		}},

		// Escapes
		{`\`, []token.Token{{Type: token.EOF, Literal: "end of file", Line: 1, Position: 1}}},
		{`\\`, []token.Token{{Type: token.OTHER, Literal: `\`, Line: 1, Position: 1}}},
		{`\"`, []token.Token{{Type: token.OTHER, Literal: `"`, Line: 1, Position: 1}}},
		{`\$foo`, []token.Token{
			{Type: token.OTHER, Literal: `$`, Line: 1, Position: 1},
			{Type: token.WORD, Literal: `foo`, Line: 1, Position: 3},
		}},
		{`\ `, []token.Token{{Type: token.ESCAPED_CHAR, Literal: ` `, Line: 1, Position: 1}}},
		{`\	`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `	`, Line: 1, Position: 1}}}, // this is a tab
		{`\|`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `|`, Line: 1, Position: 1}}},
		{`\&`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `&`, Line: 1, Position: 1}}},
		{`\>`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `>`, Line: 1, Position: 1}}},
		{`\<`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `<`, Line: 1, Position: 1}}},
		{`\;`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `;`, Line: 1, Position: 1}}},
		{`\(`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `(`, Line: 1, Position: 1}}},
		{`\)`, []token.Token{{Type: token.ESCAPED_CHAR, Literal: `)`, Line: 1, Position: 1}}},
		{`\foo`, []token.Token{
			{Type: token.ESCAPED_CHAR, Literal: `f`, Line: 1, Position: 1},
			{Type: token.WORD, Literal: `oo`, Line: 1, Position: 3},
		}},
		{"\\\nfoo", []token.Token{{Type: token.WORD, Literal: "foo", Line: 2, Position: 1}}},

		// Literal strings
		{`'hello world'`, []token.Token{
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.OTHER, Literal: `hello world`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
		}},
		{`''`, []token.Token{
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
		}},
		{`'\'`, []token.Token{
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.OTHER, Literal: `\`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
		}},
		{`'''''x' '  '`, []token.Token{
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.OTHER, Literal: `x`, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.BLANK, Literal: ` `, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.OTHER, Literal: `  `, Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
		}},
		{
			`'if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset'`,
			[]token.Token{
				{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
				{
					Type:     token.OTHER,
					Literal:  `if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset`,
					Line:     1,
					Position: 1,
				},
				{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			},
		},
		{
			`'+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : \ " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~'`,
			[]token.Token{
				{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
				{
					Type:     token.OTHER,
					Literal:  `+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : \ " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~`,
					Line:     1,
					Position: 1,
				},
				{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			},
		},
		{
			`'$$ $@ $? $# $! $_ $* $0$1$2 $3$4 $5 $6 $7 $8 $9 $10 foo bar foo-bar $variable_name $variable-name
					$concatinated$VAIABLE$VAR_0987654321 0123456789 123.456 .123 123. 1.2.3 .abc 1.c 12.34abc 123< <&45 33<&45 5<< 6<<- 1> 1>&2 7>> 81>| 19<>
					   	\t'`,
			[]token.Token{
				{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
				{
					Type: token.OTHER,
					Literal: `$$ $@ $? $# $! $_ $* $0$1$2 $3$4 $5 $6 $7 $8 $9 $10 foo bar foo-bar $variable_name $variable-name
					$concatinated$VAIABLE$VAR_0987654321 0123456789 123.456 .123 123. 1.2.3 .abc 1.c 12.34abc 123< <&45 33<&45 5<< 6<<- 1> 1>&2 7>> 81>| 19<>
					   	\t`,
					Line:     1,
					Position: 1,
				},
				{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 3, Position: 1},
			},
		},
		{"'\\\n'", []token.Token{
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 1, Position: 1},
			{Type: token.OTHER, Literal: "\\\n", Line: 1, Position: 1},
			{Type: token.SINGLE_QUOTE, Literal: `'`, Line: 2, Position: 1},
		}},
		// Others
		{`$ @`, []token.Token{
			{Type: token.OTHER, Literal: "$", Line: 1, Position: 1},
			{Type: token.BLANK, Literal: " ", Line: 1, Position: 1},
			{Type: token.AT, Literal: "@", Line: 1, Position: 1},
		}},
		{"\n", []token.Token{{Type: token.NEWLINE, Literal: "\n", Line: 2, Position: 1}}},
		{``, []token.Token{{Type: token.EOF, Literal: "end of file", Line: 1, Position: 1}}},
	}

	for i, tc := range testCases {
		l := lexer.New([]byte(tc.input))

		for j, tn := range tc.tokens {
			result := l.NextToken()
			if tn != result {
				t.Fatalf("\nCase: %d:%d\nInput: %s\nWant:\n%s\nGot:\n%s", i, j, dump(tc.input), dump(tn), dump(result))
			}
		}

		// EOF
		if result := l.NextToken(); token.EOF != result.Type {
			t.Fatalf("\nCase:%d, expected EOF, got:\n %s ", i, dump(result))
		}
	}
}
