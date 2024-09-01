package lexer_test

import (
	"testing"

	"github.com/yassinebenaid/godump"
	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/token"
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
		{`%%`, []token.Token{{Type: token.DOUBLE_PERCENT, Literal: `%%`}}},
		{`=`, []token.Token{{Type: token.ASSIGN, Literal: `=`}}},
		{`+=`, []token.Token{{Type: token.PLUS_ASSIGN, Literal: `+=`}}},
		{`-=`, []token.Token{{Type: token.MINUS_ASSIGN, Literal: `-=`}}},
		{`*=`, []token.Token{{Type: token.STAR_ASSIGN, Literal: `*=`}}},
		{`/=`, []token.Token{{Type: token.SLASH_ASSIGN, Literal: `/=`}}},
		{`==`, []token.Token{{Type: token.EQ, Literal: `==`}}},
		{`!=`, []token.Token{{Type: token.NOT_EQ, Literal: `!=`}}},
		{`=~`, []token.Token{{Type: token.EQ_TILDE, Literal: `=~`}}},
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
		{`"`, []token.Token{{Type: token.DOUBLE_QUOTE, Literal: `"`}}},
		{`'`, []token.Token{{Type: token.SINGLE_QUOTE, Literal: `'`}}},
		{`?`, []token.Token{{Type: token.QUESTION, Literal: `?`}}},
		{`!`, []token.Token{{Type: token.EXCLAMATION, Literal: `!`}}},
		{`#`, []token.Token{{Type: token.HASH, Literal: `#`}}},
		{`${`, []token.Token{{Type: token.DOLLAR_BRACE, Literal: `${`}}},
		{`$(`, []token.Token{{Type: token.DOLLAR_PAREN, Literal: `$(`}}},
		{`$((`, []token.Token{{Type: token.DOLLAR_DOUBLE_PAREN, Literal: `$((`}}},
		{`>(`, []token.Token{{Type: token.GT_PAREN, Literal: `>(`}}},
		{`<(`, []token.Token{{Type: token.LT_PAREN, Literal: `<(`}}},
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

		// identifiers
		{`foo bar foo-bar`, []token.Token{
			{Type: token.Word, Literal: `foo`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.Word, Literal: `bar`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.Word, Literal: `foo`},
			{Type: token.MINUS, Literal: `-`},
			{Type: token.Word, Literal: `bar`},
		}},

		// Special Variables
		{`$0$1$2 $3$4 $5 $6 $7 $8 $9 $10`, []token.Token{
			{Type: token.SPECIAL_VAR, Literal: "0"},
			{Type: token.SPECIAL_VAR, Literal: "1"},
			{Type: token.SPECIAL_VAR, Literal: "2"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "3"},
			{Type: token.SPECIAL_VAR, Literal: "4"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "5"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "6"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "7"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "8"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "9"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "1"}, // just to emphasize that only first digit is considered
			{Type: token.NUMBER, Literal: "0"},      // just to emphasize that only first digit is considered
		}},
		{`$1something`, []token.Token{
			{Type: token.SPECIAL_VAR, Literal: "1"},
			{Type: token.Word, Literal: "something"},
		}},
		{`$$ $@ $? $# $! $_ $*`, []token.Token{
			{Type: token.SPECIAL_VAR, Literal: "$"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "@"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "?"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "#"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "!"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "_"},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SPECIAL_VAR, Literal: "*"},
		}},
		// Simple expansion
		{`$variable_name $variable-name $concatinated$VAIABLE$VAR_0987654321`, []token.Token{
			{Type: token.SIMPLE_EXPANSION, Literal: `variable_name`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SIMPLE_EXPANSION, Literal: `variable`},
			{Type: token.MINUS, Literal: `-`},
			{Type: token.Word, Literal: `name`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.SIMPLE_EXPANSION, Literal: `concatinated`},
			{Type: token.SIMPLE_EXPANSION, Literal: `VAIABLE`},
			{Type: token.SIMPLE_EXPANSION, Literal: `VAR_0987654321`},
		}},
		// Numbers
		{`0123456789 123.456 .123 123. 1.2.3 .abc 1.c 12.34abc`, []token.Token{
			{Type: token.NUMBER, Literal: `0123456789`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.NUMBER, Literal: `123.456`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.NUMBER, Literal: `.123`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.NUMBER, Literal: `123.`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.NUMBER, Literal: `1.2`},
			{Type: token.NUMBER, Literal: `.3`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.OTHER, Literal: `.`},
			{Type: token.Word, Literal: `abc`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.NUMBER, Literal: `1.`},
			{Type: token.Word, Literal: `c`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.NUMBER, Literal: `12.34`},
			{Type: token.Word, Literal: `abc`},
		}},
		// File descriptors
		{`123< <&45 33<&45 5<< 6<<- 1> 1>&2 7>> 81>| 19<>`, []token.Token{
			{Type: token.FILE_DESCRIPTOR, Literal: `123`},
			{Type: token.LT, Literal: `<`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.LT_AMPERSAND, Literal: `<&`},
			{Type: token.FILE_DESCRIPTOR, Literal: `45`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `33`},
			{Type: token.LT_AMPERSAND, Literal: `<&`},
			{Type: token.FILE_DESCRIPTOR, Literal: `45`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `5`},
			{Type: token.DOUBLE_LT, Literal: `<<`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `6`},
			{Type: token.DOUBLE_LT_MINUS, Literal: `<<-`},

			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `1`},
			{Type: token.GT, Literal: `>`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `1`},
			{Type: token.GT_AMPERSAND, Literal: `>&`},
			{Type: token.FILE_DESCRIPTOR, Literal: `2`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `7`},
			{Type: token.DOUBLE_GT, Literal: `>>`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `81`},
			{Type: token.GT_PIPE, Literal: `>|`},
			{Type: token.BLANK, Literal: ` `},
			{Type: token.FILE_DESCRIPTOR, Literal: `19`},
			{Type: token.LT_GT, Literal: `<>`},
		}},
		// Blank
		{"  	\t", []token.Token{
			{Type: token.BLANK, Literal: "  	\t"},
		}},

		// Escapes
		{`\`, []token.Token{{Type: token.EOF}}},
		{`\\`, []token.Token{{Type: token.OTHER, Literal: `\`}}},
		{`\ `, []token.Token{{Type: token.OTHER, Literal: ` `}}},
		{`\	`, []token.Token{{Type: token.OTHER, Literal: `	`}}}, // this is a tab
		{`\$foo`, []token.Token{{Type: token.OTHER, Literal: `$`}, {Type: token.Word, Literal: `foo`}}},
		{`\|`, []token.Token{{Type: token.OTHER, Literal: `|`}}},
		{`\&`, []token.Token{{Type: token.OTHER, Literal: `&`}}},
		{`\>`, []token.Token{{Type: token.OTHER, Literal: `>`}}},
		{`\<`, []token.Token{{Type: token.OTHER, Literal: `<`}}},
		{`\;`, []token.Token{{Type: token.OTHER, Literal: `;`}}},
		{`\(`, []token.Token{{Type: token.OTHER, Literal: `(`}}},
		{`\)`, []token.Token{{Type: token.OTHER, Literal: `)`}}},
		{`\foo`, []token.Token{{Type: token.OTHER, Literal: `foo`}}},
		{"\\\nfoo", []token.Token{{Type: token.Word, Literal: "foo"}}},

		// escaped keywords
		{`\if`, []token.Token{{Type: token.OTHER, Literal: `if`}}},
		{`\if`, []token.Token{{Type: token.OTHER, Literal: `if`}}},
		{`\then`, []token.Token{{Type: token.OTHER, Literal: `then`}}},
		{`\else`, []token.Token{{Type: token.OTHER, Literal: `else`}}},
		{`\elif`, []token.Token{{Type: token.OTHER, Literal: `elif`}}},
		{`\fi`, []token.Token{{Type: token.OTHER, Literal: `fi`}}},
		{`\for`, []token.Token{{Type: token.OTHER, Literal: `for`}}},
		{`\in`, []token.Token{{Type: token.OTHER, Literal: `in`}}},
		{`\do`, []token.Token{{Type: token.OTHER, Literal: `do`}}},
		{`\done`, []token.Token{{Type: token.OTHER, Literal: `done`}}},
		{`\while`, []token.Token{{Type: token.OTHER, Literal: `while`}}},
		{`\until`, []token.Token{{Type: token.OTHER, Literal: `until`}}},
		{`\case`, []token.Token{{Type: token.OTHER, Literal: `case`}}},
		{`\esac`, []token.Token{{Type: token.OTHER, Literal: `esac`}}},
		{`\function`, []token.Token{{Type: token.OTHER, Literal: `function`}}},
		{`\select`, []token.Token{{Type: token.OTHER, Literal: `select`}}},
		{`\trap`, []token.Token{{Type: token.OTHER, Literal: `trap`}}},
		{`\return`, []token.Token{{Type: token.OTHER, Literal: `return`}}},
		{`\exit`, []token.Token{{Type: token.OTHER, Literal: `exit`}}},
		{`\break`, []token.Token{{Type: token.OTHER, Literal: `break`}}},
		{`\continue`, []token.Token{{Type: token.OTHER, Literal: `continue`}}},
		{`\declare`, []token.Token{{Type: token.OTHER, Literal: `declare`}}},
		{`\local`, []token.Token{{Type: token.OTHER, Literal: `local`}}},
		{`\export`, []token.Token{{Type: token.OTHER, Literal: `export`}}},
		{`\readonly`, []token.Token{{Type: token.OTHER, Literal: `readonly`}}},
		{`\unset`, []token.Token{{Type: token.OTHER, Literal: `unset`}}},

		// Others
		{`$ @`, []token.Token{
			{Type: token.OTHER, Literal: "$"},
			{Type: token.BLANK, Literal: " "},
			{Type: token.OTHER, Literal: "@"},
		}},

		// Others
		{"\n", []token.Token{{Type: token.NEWLINE, Literal: "\n"}}},
		{``, []token.Token{{Type: token.EOF}}},
	}

	for i, tc := range testCases {
		l := lexer.New([]byte(tc.input))

		for _, tn := range tc.tokens {
			result := l.NextToken()
			if tn != result {
				t.Fatalf("\nCase: %d\nWant:\n %s\n Got:\n%s", i, dump(tn), dump(result))
			}
		}

		// EOF
		if result := l.NextToken(); token.EOF != result.Type {
			t.Fatalf("\nCase#%d: expected EOF, got:\n %s ", i, dump(result))
		}
	}
}

func TestLexerLiteralStringContext(t *testing.T) {
	testCases := []struct {
		input  string
		tokens []token.Token
	}{
		{
			`if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset`,
			[]token.Token{{
				Type:    token.OTHER,
				Literal: `if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset`,
			}},
		},
		{
			`+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : \ " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~`,
			[]token.Token{{
				Type:    token.OTHER,
				Literal: `+ - * / % %% = += -= *= /= == != < <= > >= =~ && || | & >> << <<- <<< >& <& |& &> >| <> ; ;; ( ) (( )) [ ] [[ ]] { } , ,, : \ " ? ! # ${ $( $(( >( <( ^ ^^ := :- :+ :? // .. ++ -- ~`,
			}},
		},
		{
			`$$ $@ $? $# $! $_ $* $0$1$2 $3$4 $5 $6 $7 $8 $9 $10 foo bar foo-bar $variable_name $variable-name 
			$concatinated$VAIABLE$VAR_0987654321 0123456789 123.456 .123 123. 1.2.3 .abc 1.c 12.34abc 123< <&45 33<&45 5<< 6<<- 1> 1>&2 7>> 81>| 19<>
			   	\t`,
			[]token.Token{{
				Type: token.OTHER,
				Literal: `$$ $@ $? $# $! $_ $* $0$1$2 $3$4 $5 $6 $7 $8 $9 $10 foo bar foo-bar $variable_name $variable-name 
			$concatinated$VAIABLE$VAR_0987654321 0123456789 123.456 .123 123. 1.2.3 .abc 1.c 12.34abc 123< <&45 33<&45 5<< 6<<- 1> 1>&2 7>> 81>| 19<>
			   	\t`,
			}},
		},
		{`hello world'`, []token.Token{
			{Type: token.OTHER, Literal: `hello world`},
			{Type: token.SINGLE_QUOTE, Literal: `'`},
		}},
		{"\\\n", []token.Token{
			{Type: token.OTHER, Literal: "\\\n"},
		}},
		{`\`, []token.Token{
			{Type: token.OTHER, Literal: `\`},
		}},
		{"", []token.Token{
			{Type: token.OTHER, Literal: ""},
		}},
	}

	for i, tc := range testCases {
		l := lexer.New([]byte(tc.input))
		l.ChangeContext(lexer.CTX_LITERAL_STRING)

		for _, tn := range tc.tokens {
			result := l.NextToken()
			if tn != result {
				t.Fatalf("\nCase: %d\nWant:\n %s\n Got:\n%s", i, dump(tn), dump(result))
			}
		}

		// EOF
		if result := l.NextToken(); token.EOF != result.Type {
			t.Fatalf("\nCase#%d: expected EOF, got:\n %s ", i, dump(result))
		}
	}
}
