# Conditional Statements

Conditional statements allows you to run certain commands based on a certain condition. This is just the same as `if` and `switch` statements in languages like `C`, `PHP` and others.

The input and output of conditional statements are connected to the input and output of all commands within the statement. This means that any redirections applied on conditional statements are automatically applied on all commands within the statement. The same rule applies when you use conditional statement within pipelines.

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

## `case` statement

The `case` statement is similar to the `switch` statement in languages like `C` and `Go`.

The most basic format is:

```sh
case WORD in
	(PATTERN) consequent-commands ;;
esac
```

- `WORD` is the value to match against, can be any string.
- `PATTERN` is a string that represents a [glob pattern](https://www.gnu.org/software/bash/manual/html_node/Pattern-Matching.html).
- `consequent-commands` is a list of one or commands.
- `(PATTERN) consequent-commands ;;` is known as a **CLAUSE**.
- the left parenthese `(` in the **clause** is optional

The case statement can have one or more **clauses**, and each **clause** can have or or more patterns separated by `|`.

The case statement will go through each clause, and match `WORD` against its patterns. if one of the patterns matches, then the `consequent-commands` corresponding to that clause will be executed.

For example:

```sh
CWD=$(pwd)

case $CWD in
	/root ) echo "we are in the 'root' home directory";;
	/home/* | /Users/*)  echo "we are in home directory";;
	/var/www/html/*)  echo "we are in nginx default directory";;
esac
```

### Default clause

It's common to use the `*` as the final pattern to define the default **clause**, since that pattern will always match.

for example:

```sh
OS=$(uname)

case $OS in
	Linux | Darwin ) echo "Unix system, cool";;
	*)  echo "the operating system '$OS' is unknown";;
esac
```

### Control the match flow

Each clause must be terminated by `;;`, `;&` or `;;&`.

- If the `;;` operator is used, no subsequent matches are attempted after the first pattern match.

  ```sh
  case foo in
  f*) echo "it's foo";;
  *oo) echo "it's foo as well";;
  esac
  ```

  Outputs:

  ```txt
  it's foo
  ```

- The `;&` causes execution to continue with the `consequent-commands` associated with the next clause, if any. You can think of as the `fallthrough` keyword in `Go`.

  ```sh
  case foo in
  	f*) echo "it's foo";&
  	bar) echo "it's bar";;
  esac
  ```

  Outputs:

  ```txt
  it's foo
  it's bar
  ```

- The `;;&` causes the case statement to test the patterns in the next clause, if any, and execute any associated `consequent-commands` on a successful match, continuing the execution as if the pattern list had not matched.

  ```sh
  case foo in
  	f*) echo "it's foo";;&
  	*oo) echo "it's foo as well";;
  esac
  ```

  Outputs:

  ```txt
  it's foo
  it's foo as well
  ```
