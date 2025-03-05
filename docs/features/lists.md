# Lists

A list in bunster is a sequence of commands separated by `&&` or `||`. These operators control the execution of the commands based on success or failure.

## Logical AND (`&&`)

The format looks like this:

```sh
command1 && command2
```

The command `command2` is only executed when the last exit status is `0`.

### Example

In this example, the command `echo` will run because the command `true` exits with success (status code `0`).

```sh
true && echo yes

```

In this example, the command `echo` **will NOT run** because the command `false` exits with failure (status code `1`).

```sh
false && echo yes

```

## Logical OR (`||`)

The logical OR `||` is the opposite of the `&&` operator. The format looks similar:

```sh
command1 || command2
```

The command `command2` is only executed when the last exit status is **NOT** `0`.

### Example

In this example, the command `echo` will run because the command `false` exits with failure (status code `1`).

```sh
false || echo yes

```

In this example, the command `echo` **will NOT run** because the command `true` exits with success (status code `0`).

```sh
false || echo yes

```

## Mixing operators

A list can contain as many commands as you like. And you can mix the `&&` and `||` operators however you like. for example, this is a valid list:

```sh
command1 && command2 || command3 && command4 || command5 || comand6
```

Also, you are not limited to only simple commands within lists. You can use pipelines, functions, builtins or even compounds if you want. we will talk more about
these terms later in docs.

### Example

```sh
command1 | command2 && command3 || command4 | command5
```

## Exit Code

The exit code of a list is the exit code of the last command executed within the list.

For example:

```sh
false && true
```

will exit with `1` because the last command that was executed was `false`. However:

```sh
false || true
```

will exit with `0` because the last command that was executed was `true`.
