# Arithmetics

Bunster allows arithmetic expressions to be evaluated, using one of the builtin commands or in areas where arithmetic expressions are expected. we will not address
all possible positions in which arithmetic expressions can be used as this is going to be addressed by the relevant section.

There are 2 builtin constructs that allow you to evaluate arithmetic expressions:

- `(( expr ))`: evaluates `expr` and exits with success if the result is not `0`, or fails otherwise.
- `let expr`: evaluates `expr` and exits with success if the result is not `0`, or fails otherwise.

example:

```sh
(( var = 1 + 2))

echo $var

let var2 = 100

let var2++

echo $var2

```

will output:

```txt
3
101
```

## Expressions

Evaluation is done in fixed-width integers with no check for overflow, The operators and their precedence, associativity, and values are the same as in the `C` language. The following list of operators is grouped into levels of equal-precedence operators. The levels are listed in order of decreasing precedence.

| Expression                                                         | Description                         | Example                        |
| ------------------------------------------------------------------ | ----------------------------------- | ------------------------------ |
| `id++`, `id--`                                                     | Post-increment and post-decrement   | `x++;`                         |
| `++id`, `--id`                                                     | Pre-increment and pre-decrement     | `x = 5; y = ++x;`              |
| `-`, `+`                                                           | Unary minus and plus                | `x = -5; y = +x;`              |
| `!`, `~`                                                           | Logical and bitwise negation        | `x = !true; y = ~5;`           |
| `**`                                                               | Exponentiation                      | `x = 2 ** 3;`                  |
| `*`, `/`, `%`                                                      | Multiplication, division, remainder | `x = 10 * 2; y = 10 % 3;`      |
| `+`, `-`                                                           | Addition, subtraction               | `x = 5 + 3; y = 5 - 2;`        |
| `<<`, `>>`                                                         | Left and right bitwise shifts       | `x = 8 << 2; y = 16 >> 1;`     |
| `<=`, `>=`, `<`, `>`                                               | Comparison operators                | `x = (5 < 10); y = (10 >= 5);` |
| `==`, `!=`                                                         | Equality and inequality             | `x = (5 == 5); y = (5 != 3);`  |
| `&`                                                                | Bitwise AND                         | `x = 5 & 3;`                   |
| `^`                                                                | Bitwise exclusive OR                | `x = 5 ^ 3;`                   |
| `\|`                                                               | Bitwise OR                          | `x = 5 \| 3;`                  |
| `&&`                                                               | Logical AND                         | `x = (true && false);`         |
| `\|\|`                                                             | Logical OR                          | `x = (true \|\| false);`       |
| `expr ? expr : expr`                                               | Conditional (ternary) operator      | `x = (a > b) ? a : b;`         |
| `=`, `*=`, `/=`, `%=`, `+=`, `-=`, `<<=`, `>>=`, `&=`, `^=`, `\|=` | Assignment operators                | `x = 5; x += 3;`               |
| `expr1 , expr2`                                                    | Comma operator                      | `x = (a = 5, a + 2);`          |

Variables are allowed as operands; parameter expansion is performed before the expression is evaluated. Within an expression, shell variables may also be referenced by name without using the parameter expansion syntax. For example, `var + var2` is exacly same as `$var + $var2`.

A shell variable that is null or unset evaluates to `0`. Operators are evaluated in order of precedence. Sub-expressions in parentheses are evaluated first and may override the precedence rules above. for example, `1 + 2 * 3` evaluates to `7`, while `(1 + 2) * 3` evaluates to `9`.

## Expansion

You can the `$(( expr ))` construct to perform arithmetic substitution, which will be substituted by the result of the evaluation of `expr`.

example:

```sh
echo $(( 1 + 2))

var=$(( 10 / 2 ))

echo $var
```

will output:

```txt
3
5
```
