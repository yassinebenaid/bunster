# Developers Guideline

This document is your go-to if you want to know more about the internals of bunster. Things like how the compiler boots up. How the code goes from the shell file to
being a static binary. What each package in the repository is responsible for and how it interacts with other packages.

This document is dedicated to educate new contributors who struggle to get their hands dirty with bunster. But will also be very benificial for anyone who is intersted
in writing compilers (or transpilers in particular).

## project Tree

```txt
.
├── token
├── lexer
├── ast
├── parser
├── analyser
├── ir
├── generator
├── runtime
├── pkg
├── stubs
├── cmd
├── tests
├── Dockerfile
├── Makefile
├── bunster_test.go
├── go.mod
└── embed.go

14 directories, 10 files

```

This project tree is very simple. It is flatten with encapsulated functionalities. Each directory is dedicated for a specific purpose. And usually only export one function or type.

### Packages

#### `token`

This package only defines a list of constants that represent tokens. Things like keywords, symbols and so on. It doesn't export any functionality. But serves as an asset for the `parser` and `lexer`.


#### `lexer`

This package is responsible on the generation of tokens. It reads the input file (as text) and converts it into a feed of tokens. For example:
```shell
echo foobar >file | cat

```

will yield a list of tokens similar to the following:
```
WORD: "echo"
BLANK
WORD: "foobar"
BLANK
GT: ">"
WORD: "file"
BLANK
PIPE: "|"
BLANK
WORD: "cat"
NEWLINE
EOF
```

This package gives each character in the input a meaning. for example, `>` means `GT`,  `|` means `PIPE`, but `>|` is considered one token `GTPIPE`. and `>>` is also considered one token `APPEND`. In other words, the `lexer` applies lexical analysis over the input and decides what each combination of characters means.  

> [!TIP]
> This package does not generate all tokens at once. It exports a function (`NextToken`) that returns one token at a time. 

**Depends on:** 

`token`



