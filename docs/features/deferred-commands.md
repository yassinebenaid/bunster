# Deferred Commands

You can defer the execution of a command or group of commands until the end of the program or function. Useful for clean up commands.

The `defer` keyword marks the given command, group or sub-shell to be deferred until the end of the program or function.

Format:

```sh
defer command
```

For example:

```sh
defer echo foo
echo bar
```

Will output:

```txt
bar
foo
```

You are not limited to simple commands, you're free to use [groups and sub-shells](/features/groups-and-subshells) as well:

Defer a group:

```sh
defer {
	echo foo
	echo bar
}

echo baz
```

Defer a sub-shell:

```sh
defer (
	echo foo
	echo bar
)

echo baz
```
