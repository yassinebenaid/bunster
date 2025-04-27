# Functions

Functions are a way to group commands for later execution using a single name for the group. They are executed just like a "regular" command. When the name of a function is used as a simple command name, the list of commands associated with that function name is executed. Functions are executed in the current shell context.

## Declaring functions

You can declare a function like this:

```sh
function functionName() {
	# commands
}

```

The parentheses are optional, the above example can be refactored to:

```sh
function functionName {
	# commands
}

```

Even the `function` keyword is optional, and can be omited, in this case, the parentheses `()` are required, the above examples can be refactored to:

```sh
functionName() {
	# commands
}

```

### Declaring redirections

When you declare a function in any of the possible forms, you may apply redirections on the function declaration. when you do, those redirections will be applied on the commands within the function each time the function in invoked.

for example:

```sh
function functionName() {
	# commands
} >output.txt <input.txt

```

## Invoking functions

Functions can invoked just like any other simple command, you just call a command using the function name, this means that all rules that apply to simple commands like arguments, environment variables and redirections will be applied to the function.

```sh
function functionName() {
	echo $env_var $1 $2
}

env_var=foo functionName bar baz

# Output: fo bar baz
```

## Isolated functions

Because functions are executed in current context. variables declared or mutated within functions affect the global scope.

for example:

```sh
var=foo

function functionName() {
	var=bar
}

echo $var

functionName

echo $var

```

Outputs:

```txt
foo
bar
```

You can declare a function to run in a [subshell](/features/groups-and-subshells#subshell) by replacing the braces `{}` by parentheses `()`:

```sh
var=foo

function functionName() (
	var=bar
)


echo $var

functionName

echo $var

```

Outputs:

```txt
foo
foo
```

## Return from function

the `return` keyword is used to return from a function.

for example:

```sh
function foo(){
	echo foo

	return

	echo bar
}

foo
```

Outputs:

```txt
foo
```

an optional argument specifies the return status.

```sh
function foo(){
	# ...
	return 3
	# ...
}

foo

```

This function exits with `3` exit code.

## Flags

Bunster has builtin flags parsing capability. it's part of the function declaration.

for example:

```sh
function myFunc( -a -b -c ){
	echo -a = $fflags_a
	echo -b = $fflags_b
	echo -c = $fflags_c
}

myFunc -a -c
```

Outputs:

```txt
-a = 1
-b = 0
-c = 1
```

You declare the flags within the parentheses in function declaration, and when the function is called, bunster will parse them and make them available as local variables
prefixed with `fflags_`. for example, in above example, we declared three boolean flags `-a`, `-b` and `-c`, and they became available as `$fflags_a`, `$fflags_b` and `$fflags_c` variables.

### Short boolean flags

short boolean flags are one character long flags that don't accept a value, they're assigned `1` when they're supplied. and `0` otherwise.

example:

```sh
function myFunc( -a -b -c ){
	echo -a = $fflags_a, -b = $fflags_b, -c = $fflags_c
}

myFunc -a -c
myFunc -c -b
myFunc -a
```

Outputs:

```txt
-a = 1, -b = 0, -c = 1
-a = 0, -b = 1, -c = 1
-a = 1, -b = 0, -c = 0
```

### Short string flags

short string flags are one character long flags that accept a value. example:

```sh
function myFunc( -a= -b= -c= ){
	echo -a = $fflags_a, -b = $fflags_b, -c = $fflags_c
}

myFunc -a avalue -b bvalue -c cvalue
```

Outputs:

```txt
-a = avalue, -b = bvalue, -c = cvalue
```

### Short flags grouping

short flags can be grouped together when supplied. for example:

```sh
function myFunc( -a -b= -c= ){
	echo -a = $fflags_a, -b = $fflags_b, -c = $fflags_c
}

myFunc -acb cvalue bvalue
```

Outputs:

```txt
-a = 1, -b = bvalue, -c = cvalue
```

### Long boolean flags

long boolean flags are sometimes called options, they're usually a word or more long, that don't accept a value, they're assigned `1` when they're supplied. and `0` otherwise.

example:

```sh
function myFunc( --foo --bar --baz ){
	echo --foo = $fflags_foo, --bar = $fflags_bar, --baz = $fflags_baz
}

myFunc --foo --bar
myFunc --bar --baz
```

Outputs:

```txt
--foo = 1, --bar = 1, --baz = 0
--foo = 0, --bar = 1, --baz = 1
```

### Long string flags

long string flags are just long boolean flags,but they accept a value. example:

```sh
function myFunc( --foo= --bar= --baz= ){
	echo --foo = $fflags_foo, --bar = $fflags_bar, --baz = $fflags_baz
}

myFunc --foo fooValue --bar barValue --baz bazValue
```

Outputs:

```txt
--foo = fooValue, --bar = barValue, --baz = bazValue
```

### Optional flags

All boolean flags are optional by default, their value is `0` unless they're supplied. However, all string flags short or long are required by default.
an error occurs if they're not supplied.

You can mark a flag as optional using the `[=]` instead of `=` in declaration. optional flags are not set if not supplied.

```sh
function myFunc( -a[=] --foo[=] ){
	echo -a = $fflags_a, --foo = $fflags_foo
}

myFunc --foo fooValue -a avalue
myFunc --foo fooValue
myFunc -a avalue
myFunc
```

Outputs:

```txt
-a = avalue, --foo = fooValue
-a = , --foo = fooValue
-a = avalue, --foo =
-a = , --foo =
```

### Arguments after flag parsing

The arguments that are not flags or arguments to flags are kept available as positional variables to that function.

for example:

```sh
function myFunc( -a[=] --foo[=] ){
	echo -a = $fflags_a, --foo = $fflags_foo
	echo arguments = [$@]
}

myFunc arg --foo fooValue arg -a avalue arg
```

Outputs:

```txt
-a = avalue, --foo = fooValue
arguments = [arg arg arg]
```

### Terminating flags parsing (`--`)

by default, an error occurs when passing a flag that is not declared. for example:

```sh
function myFunc( -a --boo ){
	return;
}

myFunc -x

# myFunc: unknown short flag: x
```

Sometimes, you want to pass an argument that starts with a dash `-`, but you don't want bunster to treat it as a flag. you can use the `--` to terminate flags parsing. all arguments that come after the `--` are treated literally as arguments.

for example:

```sh
function myFunc( -a --boo ){
	echo arguments = [$@]
}

myFunc arg -- -x --foo -abc
```

Outputs:

```txt
arguments = [arg -x --foo -abc]
```
