# Test Commands

Test commands are special commands built into the language that allows you to perform conditional expressions. There are three test commands:

- `[[ ... ]]`: inspired from bash and works by wrapping a conditional expression within a pair of brackets like this: `[[ <expression> ]]`.
- `[...]`: works exactly the same as previous command except that it uses the operators `-a` and `-o` instead of `&&` and `||` respectively. (_we will learn more about expressions below_).
- `test`: works by putting a conditional expression in front of the `test` keyword. for example: `test <expression>`. Has same rules as `[...]`.

The exit code of conditional commands is zero if the expression evaluates to true. and one otherwise.

> [!TIP]
> Unlike bash, there is no need to escape the special characters in expressions when using the `test` or `[...]` commands. Enjoy writing code like this: `[ $foo < $bar ]` :tada: .

## Expressions

A conditional expression can be devided into two categories. Unary expressions and binary expressions.

### Unary expressions

The format is :

```sh
<FLAG> word
```

The `<FLAG>` dictates what type of test you want to perform. The following table lists all possible flags with the their functions. The `word` can be any valid expression
the expands to a string, and is the operand over which you want to perform the test.

| Flag | Description                                                       | Example                       |
| ---- | ----------------------------------------------------------------- | ----------------------------- |
| `-a` | True if file exists.                                              | `[ -a /path/to/file ]`        |
| `-b` | True if file exists and is a block special file.                  | `[ -b /dev/sda ]`             |
| `-c` | True if file exists and is a character special file.              | `[ -c /dev/tty ]`             |
| `-d` | True if file exists and is a directory.                           | `[ -d /home/user ]`           |
| `-e` | True if file exists.                                              | `[ -e /path/to/file ]`        |
| `-f` | True if file exists and is a regular file.                        | `[ -f /etc/passwd ]`          |
| `-g` | True if file exists and its set-group-id bit is set.              | `[ -g /path/to/file ]`        |
| `-h` | True if file exists and is a symbolic link.                       | `[ -h /path/to/symlink ]`     |
| `-k` | True if file exists and its "sticky" bit is set.                  | `[ -k /tmp ]`                 |
| `-p` | True if file exists and is a named pipe (FIFO).                   | `[ -p /path/to/pipe ]`        |
| `-r` | True if file exists and is readable.                              | `[ -r /path/to/file ]`        |
| `-s` | True if file exists and has a size greater than zero.             | `[ -s /path/to/file ]`        |
| `-t` | True if file descriptor is open and refers to a terminal.         | `[ -t 1 ]`                    |
| `-u` | True if file exists and its set-user-id bit is set.               | `[ -u /path/to/file ]`        |
| `-w` | True if file exists and is writable.                              | `[ -w /path/to/file ]`        |
| `-x` | True if file exists and is executable.                            | `[ -x /path/to/script.sh ]`   |
| `-G` | True if file exists and is owned by the effective group ID.       | `[ -G /path/to/file ]`        |
| `-L` | True if file exists and is a symbolic link.                       | `[ -L /path/to/symlink ]`     |
| `-N` | True if file exists and has been modified since it was last read. | `[ -N /path/to/file ]`        |
| `-O` | True if file exists and is owned by the effective user ID.        | `[ -O /path/to/file ]`        |
| `-S` | True if file exists and is a socket.                              | `[ -S /var/run/docker.sock ]` |
| `-v` | True if the shell variable is set (has been assigned a value).    | `[ -v MY_VAR ]`               |
| `-z` | True if the length of string is zero.                             | `[ -z $EMPTY_VAR ]`           |
| `-n` | True if the length of string is non-zero.                         | `[ -n $MY_VAR ]`              |

There is an alternative syntax for the `-n` flag. You can omit the flag at all and only keep the operand. for example `[ $var ]` returns true if the length of the value of `$var` is non-zero.

### Binary expressions

The format is:

```sh
word1 <FLAG> word2
```

Binary expressions are used for comparison between two operands. The `<FLAG>` dictates what type of comparison you want to perform. The following table lists all possible flags with the their functions. `word1` and `word2` can be any valid expressions
the expand to string, and are the operands over which you want to perform the comparison.

| Expression           | Description                                                                                                        |
| -------------------- | ------------------------------------------------------------------------------------------------------------------ |
| `file1 -ef file2`    | True if `file1` and `file2` refer to the same device and inode numbers.                                            |
| `file1 -nt file2`    | True if `file1` is newer (according to modification date) than `file2`, or if `file1` exists and `file2` does not. |
| `file1 -ot file2`    | True if `file1` is older than `file2`, or if `file2` exists and `file1` does not.                                  |
| `string1 = string2`  | True if `string1` is equal to `string2` (POSIX-compliant form).                                                    |
| `string1 == string2` | True if `string1` is equal to `string2` (performs pattern matching with `[[...]]` command).                        |
| `string1 != string2` | True if `string1` is not equal to `string2`.                                                                       |
| `string1 < string2`  | True if `string1` sorts before `string2` lexicographically.                                                        |
| `string1 > string2`  | True if `string1` sorts after `string2` lexicographically.                                                         |
| `num1 -eq num2`      | True if `num1` is equal to `num2`.                                                                                 |
| `num1 -ne num2`      | True if `num1` is not equal to `num2`.                                                                             |
| `num1 -lt num2`      | True if `num1` is less than `num2`.                                                                                |
| `num1 -le num2`      | True if `num1` is less than or equal to `num2`.                                                                    |
| `num1 -gt num2`      | True if `num1` is greater than `num2`.                                                                             |
| `num1 -ge num2`      | True if `num1` is greater than or equal to `num2`.                                                                 |

## Combining Expressions

One test command can contain exactly one conditional expression. However, if you want to use more expressions, then you should use expression combination.

Expressions may be combined using the following operators, listed in increasing order of precedence:

**Logical OR**:

The format is:

```sh
expression1 -o expression2
```

Or if used with `[[...]]` command

```sh
expression1 || expression2
```

True if either `expression1` or `expression2` is true. `expression2` is not evaluated if `expression1` evaluates to true.

**Logical AND**:

The format is:

```sh
expression1 -a expression2

```

Or if used with `[[...]]` command

```sh
expression1 && expression2
```

True if both `expression1` and `expression2` are true. `expression2` is not evaluated if `expression1` evaluates to false.

**Expression Inversion**:

The format is:

```sh
! expression
```

True if `expression` is false. And false otherwise.

**Grouping**:

The format is:

```sh
( expression )
```

Returns the value of expression. This may be used to override the normal precedence of operators.

> [!TIP]
> You are not limited to one combination expression at a time. You can combine them as well to create complex conditional expressions. for example: `! (-f file && -n string)`
