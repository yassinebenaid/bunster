# Groups & Sub-Shells

There are two ways to group a list of commands to be executed as a unit. When commands are grouped, redirections may be applied to the entire command list. For example, the output of all the commands in the list may be redirected to a single stream. Or you may run the entire command list as a single command in an outer pipeline.

The exit status of grouped commands is the exit code of the last command within the list.

## Subshell

A subshell is an abstract definition that refers to an isolated environment in which commands run without affecting the global scope. Commands that run in a sub-shell are executed in a separate context. This means that all variables mutated or declared within the sub-shell do not affect the global scope.

## Groups

### Ordinary groups

The format is :

```sh
{
	command1
	command2
	...
}

```

Placing a list of commands within curly braces allow you to group those commands into a single unit. Commands can be anything. from simple command to buitlin commands, loops, functions and anything else.

Unless explicitly redirected, redirections applied on the group are applied on all the commands within the list . for example:

```sh
{
	command1
	command2
	command3 < file.txt > file2.txt
} <file3.txt >file4.txt

```

The commands `command1` and `command2` will read from file `file3.txt` and write to `file4.txt`. While `command3` will read from `file.txt` and write to `file2.tx`. This is because the `command3` has explicit redirections. While other commands inherit the redirections from the group.

Because the command group is a single unit, you can treat it just like a simple command. For example you can use it within pipelines:

```sh
{ echo foo; echo bar; } | cat
```

> [!TIP]
> Do not be confused by the semicolon in the above example, semicolons `;` is just one the tokens that you can use to
> separate commands. just like a new line `\n`. [learn more about possible separators in their dedicated section](/features/simple-commands#separator).

> [!WARNING]
> The closing brace must be preceeded by a separator (`;`, `\n`...) to be recognized as end of group. For example this is not valid: `{ command }`.

### Subshell groups

The second way of grouping commands is by wrapping commands within parentheses `(...)`. like this:

```sh
(
	command1
	command2
	...
)

```

Grouping commands this way has the same posibilities as [the ordinary groups](#ordinary-groups). Except that commands run in a [subshell](#subshell).

For example:

```sh
var=foo

(
	var2=bar
	var3=baz
)

echo var:$var var2:$var2 var3:$var3

```

Will output:

```txt
echo var:foo var2: var3:

```

> [!TIP]
> Unlike groups, sub-shells do not require any separator before the closing parenthese.

## Command substitution

Command substitution allows the output of a command to be used as expression, for example as a command name, argument, a variable value etc. Command substitution occurs when a command is enclosed as follows:

```sh
$(command)
```

Bunster performs the substitution by executing `command` in a subshell environment and replacing the command substitution with the standard output of the `command`, with any trailing newlines deleted. Embedded newlines are not deleted.

example:

```sh
echo $( echo foobar )
```

will output:

```txt
foobar
```

You may use as many commands as you like within a command substitution, and all commands, keywords and statments are valid inside command substitution, you can even nest command substitutions.

```sh
echo $( echo $( echo foobar) )
```
