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
