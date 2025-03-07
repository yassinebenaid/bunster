# Variables & Environment

A shell script is less useful if you cannot store data and operate on it. Variables allow you to store data in a named variable to be used later. If you have
a decent experience in programming. Then you should be familiar with what variables are and how they're used.

## Quick overview

You can define a variable like this:

```sh
varname=value
```

Reading a variable looks like this:

```sh
echo $varname
```

To read the value of a variable, we use the dollar sign `$` followed by the name of the variable. Reading an undefined variable will result in a null string.

## Variable declaration

You can declare a variable using the form:

```sh
name=[value]
```

The `name` is the name of the variable to be used later to read the value. It should only contain letters `A-Z`, `a-z`, underscore `_` or numbers `0-9`.

The value is optional. When omited, the variable is initialized with a null string. You can use any expression in the value.

Examples:

```sh
var=
var2=foo
var3="foo"
var4='foo'

#...
```

## Environment variables

When a process runs in the system. It inherits a vector (array) of key-value pairs called `Environment Variables` or `ENVs`.
These variables are usually meant to pass additional data to the program. But it's up to the program to decide what to do with them.
And the program is also free to edit those variables before it passes them to its child processes.

On unix systems, there are many environment variables that are set by default, like `$HOME` and `$PATH` etc.

When you run a bunster script, all environment variables passed to the process are set as variables. This means you can read them just normal:

```sh
echo $HOME
```

> [!info]
> Unless altered. All environment variables passed to the script are passed to it's child processes (commands you run).

### Passing environment variables

When you run a command in your script. you may want to pass additional environment variables to it. You can do that by proceeding the command invokation by a
list of variable declarations. like this:

```sh
key=value key2=value command arguments ...
```

The variables `key` and `key2` will be passed to the command `command` as environment variables. You can list as many variables as you want.

Note that when you pass environment variales to a command, they will not be set as variables in the script.

For example:

```sh
key=value command

echo $key # $key is unset
```

The variable `key` is second line will result in a null string because the variable is not set.

### Exporting variables

If you have a variable defined and you want it to be passed to child commands as environment variable.
This is possible using the `export` keyword.

The format is :

```sh
export key[=value] ...
```

The `export` keyword is followed by one or more variable declaration. This will set the variables in the script. Meanwhile, those variables will be passed to all commands as environment variables.

For example:

```sh
export var=foobar

echo $var # works
sh -c 'echo $var' # works as well
```

If you have a variable set already. And want to only export it. you can pass the name of the variable to the `export` keyword:

```sh
var=foo
var2=bar

export var var2

```

Or, you can mix them if you like:

```sh
var=foo

export var var2=bar

```

## Positional variables

Positional variables are the arguments passed to the script during invokation. You can access them using the dollar sign followed by a single digit.

For example, if the script was run as `script arg1 arg2 arg3` :

```sh
echo $1 $2 $3
```

the variables `$1`, `$2` and `$3` will be expanded to `arg1`, `arg2` and `arg3` respectivily.

## Special variables

There are some special variables that are treated specially. Assignment to them is not allowed.

| Variable  | Description                                                                                                                                             |
| --------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `$0`      | Expands to the name of the program as it was invoked. for example, if the program was invoked as `program-name args...`, `$0` expands to `program-name` |
| `$*`      | Expands to the positional arguments concatinated by a space                                                                                             |
| `$@`      | Same as `$*`                                                                                                                                            |
| `$#`      | Expands to the number of positional parameters in decimal.                                                                                              |
| `$?`      | Expands to the exit status of the most recently executed command.                                                                                       |
| `$$`      | Expands to the process ID of the program                                                                                                                |
