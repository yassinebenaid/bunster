package token

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	_ TokenType = iota

	WHITESPACE // Spaces and tabs
	NEWLINE    //  '\n'

	IF       // 'if'
	THEN     // 'then'
	ELSE     // 'else'
	ELIF     // 'elif'
	FI       // 'fi'
	FOR      // 'for'
	IN       // 'in'
	DO       // 'do'
	DONE     // 'done'
	WHILE    // 'while'
	UNTIL    // 'until'
	CASE     // 'case'
	ESAC     // 'esac'
	FUNCTION // 'function'
	SELECT   // 'select'
	TRAP     // 'trap'
	RETURN   // 'return'
	EXIT     // 'exit'
	BREAK    // 'break'
	CONTINUE // 'continue'
	DECLARE  // 'declare'
	LOCAL    // 'local'
	EXPORT   // 'export'
	READONLY // 'readonly'
	UNSET    // 'unset'

	PLUS                // '+'
	MINUS               // '-'
	STAR                // '*'
	SLASH               // '/'
	PERCENT             // '%'
	EQ                  // '='
	PLUS_EQ             // '+='
	MINUS_EQ            // '-='
	STAR_EQ             // '*='
	SLASH_EQ            // '/='
	EQ_EQ               // '=='
	NE                  // '!='
	LT                  // '<'
	LE                  // '<='
	GT                  // '>'
	GE                  // '>='
	AND                 // '&&'
	OR                  // '||'
	NOT                 // '!'
	PIPE                // '|'
	REDIRECT_OUTPUT     // '>'
	APPEND_OUTPUT       // '>>'
	REDIRECT_INPUT      // '<'
	HERE_DOCUMENT       // '<<'
	HERE_DOCUMENT_MINUS // '<<-'
	HERE_STRING         // '<<<'
	FILE_DESCRIPTOR     // 0-9 in io redirection. eg '>&', '<&', ...
	AND_OUTPUT          // '>&'
	AND_INPUT           // '<&'

	NAME  // Variable names, functions and other identifiers.
	FNAME // Just like NAME, but contains '-' dash.

	STRING_SINGLE_QUOTED // Single-quoted string
	NUMBER               // Integer and float numbers

	SEMICOLON                // ';'
	DOUBLE_SEMICOLON         // ';;'
	LEFT_PAREN               // '('
	RIGHT_PAREN              // ')'
	LEFT_BRACKET             // '['
	RIGHT_BRACKET            // ']'
	DOUBLE_LEFT_BRACKET      // '[['
	DOUBLE_RIGHT_BRACKET     // ']]'
	LEFT_BRACE               // '{'
	RIGHT_BRACE              // '}'
	COMMA                    // ','
	COLON                    // ':'
	BACKSLASH                // '\'
	DOLLAR                   // '$'
	DOUBLE_QUOTE             // '"'
	BACKTICK                 // '`'
	QUESTION                 // '?'
	EXCLAMATION              // '!'
	HASH                     // '#'
	DOLLAR_BRACE             // '${'
	DOLLAR_BRACE_HASH        // '${#'
	DOLLAR_BRACE_EXCLAMATION // '${!'
	DOLLAR_PAREN             // '$('
	DOLLAR_DOUBLE_PAREN      // '$(('
	CIRCUMFLEX               // '^'
	DOUBLE_CIRCUMFLEX        // '^^'
	COLON_EQUAL              // ':='
	COLON_MINUS              // ':-'
	COLON_PLUS               // ':+'
	COLON_QUESTION           // ':?'
	DOUBLE_SLASH             // '//'

	SPECIAL_VAR // Special variables like $?, $#, $@, $*, $$, $!, $0, $1, $2, ...

	HEREDOC_END // End marker for here document

	ARRAY_ASSIGN // Array assignment: arr=(...)
	ARRAY_ACCESS // Array access: ${arr[index]}

	COMMAND_SUBSTITUTION_START    // '$(' or `` ` ``
	COMMAND_SUBSTITUTION_END      // ')'
	ARITHMETIC_SUBSTITUTION_START // '$(('
	ARITHMETIC_SUBSTITUTION_END   // '))'

	ESCAPE_CHAR // Escaped characters like \n, \t, \\

	PARAM_EXPANSION // Parameter expansion: ${VAR}, ${VAR:-default}, etc.
	BRACE_EXPANSION // Brace expansion: {a,b,c}, {1..10}
	TILDE_EXPANSION // Tilde expansion: ~, ~user

	ILLIGAL // unknown token or illigal
	EOF     // end of file
)

var literals = map[TokenType]string{
	// NL:      "NEWLINE",
	// IDENT:   "IDENTIFIER",
	// ILLIGAL: "ILLIGAL",
	// EOF:     "EOF",
}

func (t TokenType) String() string {
	return literals[t]
}
