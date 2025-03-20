# Loops

There are three statements built into the language that allow you to run commands in a loop. All redirections applied on each loop statment is applied on all commands within the loop.

Whenever a semicolon `;` appears in the examples below, it may be replaced with one or more newlines.

## `until`

The syntax of the `until` loop is:

```sh
until test-commands; do consequent-commands; done
```

The until loop executes `consequent-commands` as long as `test-commands` has an exit status which is not zero. The return status is the exit status of the last command executed in `consequent-commands`.

For example:

```sh
until false; do
	echo foobar
done

```

## `while`

The syntax of the `while` loop is:

```sh
while test-commands; do consequent-commands; done
```

The while is the opposite of `until` loop. it executes `consequent-commands` as long as `test-commands` has an exit **status of zero**. The return status is the exit status of the last command executed in `consequent-commands`.

For example:

```sh
while true; do
	echo foobar
done

```

## `for`

The for loop can be constructed in 3 different formats.

### Loop over positionals

The format is as follows:

```sh
for NAME do
	consequent-commands
done
```

This format will execute `consequent-commands` once for each [positional argument](/features/variables-and-environment#positional-variables) that is set. the `NAME` is a variable name that will hold the value of the positional
argument under examination.

For example:

```sh
for arg do
	echo $arg
done
```

If the positional arguments are `foo bar baz`, the output would be:

```txt
foo
bar
baz
```

### Loop over arguments

The format looks similar to above:

```sh
for NAME in arguments; do
	consequent-commands
done
```

This format will execute `consequent-commands` once for each field of the list resulted from the expansion of `arguments`.

For example:

```sh
for user in bob yassine phank; do
	echo $user
done
```

the output would be:

```txt
bob
yassine
phank
```
