# Simple Commands

Just like any other shell scripting language. Bunster allows you can run external progams.

The format is as follows:

```sh
command-name [arguments...]
```

The `command-name` can be a full path to a program. or just the name of the binary. In that case, the binary must be accessible in `$PATH`.

## Separator

You can separate the command list by a new line. a semicolon `;`, or ampersand `&`.

```sh
command
command2; command3
command4 & command5
```

We will talk more about the ampersand `&` in [Asynchronous commands](/features/async-commands).

## Quotes

Quotes in bunster work the exact same as in bash.

### Double Quotes

double quotes are used to preserve the literal meaning of tokens except the dollar sign `$`.

```sh
"command name" "foo bar"
```

This will run the comand `command name` and pass the argument `foo bar`.

#### Dollar sign

Dollar sign is a special token that is used for substitution. It does not lose it's meaning in double quotes:

```sh
echo "$HOME"
```

This will run `echo`. and passes an argument which is the value of the `HOME` variable. (we will learn more about variables later).

#### Escaping quotes

You can escape double quotes within double quotes:

```sh
echo "\"foobar\""
```

This will run the command `echo` with the argument `"foobar"`.

### Single Quotes

Single quotes are also used to preserve the literal meaning of tokens. However, unlike double quotes, all tokens within single quotes loose their special meaning.
Including the `$` dollar sign. And you cannot use escaping within single quotes

```sh
echo '$VAR\'
```

Runs the command `echo` with the argument `$VAR\`

## Escaping

Escaping is the same as in bash. you use the backslash `\` to preserve the literal meaning of the next token.
The backslash it self is removed and the escaped token is treated literally.

Newline `\n` is a special case. when you scape a newline. the newline is removed as well:

```sh
command\ name \
    argument argument2
```

This will run the command `command name`. and pass the `argument` and `argument2` arguments.

## Comments

Comments in bunster are (as you may guess) the exact as in any other shell. Lines starting with `#` are ignored.
Parts of lines that start with `#` are treated as comments as well and are ignored.

```sh
# full line comment
command # inline comment
```

## Examples

```sh
echo foo bar

echo "Hello World"

echo 'Hello World'

echo; echo

echo \
    foo bar
```
