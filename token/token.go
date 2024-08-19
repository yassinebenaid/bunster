package token

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	_ TokenType = iota

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

	PLUS                 // '+'
	MINUS                // '-'
	STAR                 // '*'
	SLASH                // '/'
	PERCENT              // '%'
	ASSIGN               // '='
	PLUS_ASSIGN          // '+='
	MINUS_ASSIGN         // '-='
	STAR_ASSIGN          // '*='
	SLASH_ASSIGN         // '/='
	EQ                   // '=='
	NOT_EQ               // '!='
	LT                   // '<'
	LE                   // '<='
	GT                   // '>'
	GE                   // '>='
	AND                  // '&&'
	OR                   // '||'
	PIPE                 // '|'
	AMPERSAND            // '&'
	DOUBLE_GT            // '>>'
	DOUBLE_LT            // '<<'
	DOUBLE_LT_MINUS      // '<<-'
	TRIPLE_LT            // '<<<'
	GT_AMPERSAND         // '>&'
	LT_AMPERSAND         // '<&'
	PIPE_AMPERSAND       // '|&'
	AMPERSAND_GT         // '&>'
	GT_PIPE              // '>|'
	LT_GT                // '<>'
	SEMICOLON            // ';'
	DOUBLE_SEMICOLON     // ';;'
	LEFT_PAREN           // '('
	RIGHT_PAREN          // ')'
	DOUBLE_LEFT_PAREN    // '(('
	DOUBLE_RIGHT_PAREN   // '))'
	LEFT_BRACKET         // '['
	RIGHT_BRACKET        // ']'
	DOUBLE_LEFT_BRACKET  // '[['
	DOUBLE_RIGHT_BRACKET // ']]'
	LEFT_BRACE           // '{'
	RIGHT_BRACE          // '}'
	COMMA                // ','
	COLON                // ':'
	BACKSLASH            // '\'
	DOLLAR               // '$'
	DOUBLE_QUOTE         // '"'
	QUESTION             // '?'
	EXCLAMATION          // '!'
	HASH                 // '#'
	DOLLAR_BRACE         // '${'
	DOLLAR_PAREN         // '$('
	DOLLAR_DOUBLE_PAREN  // '$(('
	LT_PAREN             // '>('
	GT_PAREN             // '<('
	CIRCUMFLEX           // '^'
	DOUBLE_CIRCUMFLEX    // '^^'
	COLON_ASSIGN         // ':='
	COLON_MINUS          // ':-'
	COLON_PLUS           // ':+'
	COLON_QUESTION       // ':?'
	DOUBLE_SLASH         // '//'
	DOUBLE_DOT           // '..'
	INCREMENT            // '++'
	DECREMENT            // '--'
	TILDE                // '~'

	FILE_DESCRIPTOR        // 0-9 in io redirection. eg '>&', '<&', ...
	NAME                   // Variable names, functions and other identifiers.
	FNAME                  // Just like NAME, but contains '-' dash.
	LITERAL_STRING         // Single-quoted string or double-quoted string without any expansion
	PARTIAL_LITERAL_STRING // Double-quoted string before or after expansion, eg. "hello $name mate" : "hello " and " mate" are both partial literal strings
	NUMBER                 // Integer and float numbers
	BLANK                  // Spaces and tabs
	NEWLINE                //  '\n'
	SPECIAL_VAR            // Special variables like $?, $#, $@, $*, $$, $!, $0, $1, $2, ...
	ESCAPE_CHAR            // Escaped characters like \n, \t, \\
	OTHER                  // anything else
	EOF                    // end of file
)

var Keywords = map[string]TokenType{
	"if":       IF,
	"then":     THEN,
	"else":     ELSE,
	"elif":     ELIF,
	"fi":       FI,
	"for":      FOR,
	"in":       IN,
	"do":       DO,
	"done":     DONE,
	"while":    WHILE,
	"until":    UNTIL,
	"case":     CASE,
	"esac":     ESAC,
	"function": FUNCTION,
	"select":   SELECT,
	"trap":     TRAP,
	"return":   RETURN,
	"exit":     EXIT,
	"break":    BREAK,
	"continue": CONTINUE,
	"declare":  DECLARE,
	"local":    LOCAL,
	"export":   EXPORT,
	"readonly": READONLY,
	"unset":    UNSET,
}
