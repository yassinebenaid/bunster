# Conditional Constructs

Conditional constructs allows you to run certain commands based on a certain condition. This is just the same as `if` and `switch` statements in languages like `C`, `PHP` and others.

Note that in any of the examples below, whenever a semiclon is used. can be replaced by one or more newlines as [mentioned before](/features/simple-commands#separator)

## `if` statement

The most basic syntax of an `if` statement looks like:

```sh
if test-commands; then
	consequent-commands
fi
```

The `test-commands` are invoked first. Then if (and only if) the exit status is zero, `consequent-commands` are then invoked.

For example:

```sh
if true; then
	echo Horrey
fi

if false; then
	echo Booo
fi

```

Outputs:

```txt
Horrey
```

### `elif` branches

The format is so simple:

```sh
if test-commands; then
	consequent-commands
elif alt-test-commands; then
	alt-consequent-commands
fi
```

`elif` act as an alternative branch. if `test-commands` exited with zero. `consequent-commands` are invoked and the statement returns. Otherwise, if `test-commands` exited with non-zero, then `alt-test-commands` are executed. and if the exit status is zero, `alt-consequent-commands` are then executed.

You can have as many `elif` branches as you want. For example:

```sh
if cmd1; then
	echo foo
elif cmd2; then
	echo bar
elif cmd3; then
	echo baz
fi
```

### `else` branch

The `else` branch is the last branch in the `if` statement. looks like:

```sh
if test-commands; then
	consequent-commands
else
	alt-consequent-commands
fi
```

The `else` branch does not have a test commands. because it is the default branch that gets executed when (and only when) none of the other branches get pass the test.

if `test-commands` exited with zero. `consequent-commands` are invoked and the statement returns. Otherwise,`alt-consequent-commands` are then executed.

For example:

```sh
if false; then
	echo foo
elif false; then
	echo bar
else
	echo baz
fi
```
