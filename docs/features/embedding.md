# Embedding

How many times you were in a situation where you had a file and you wanted to ship it with your script. Well, you can now.

Bunster offers a simple yet powerful embedding system built on top of [Go's embed package](https://pkg.go.dev/embed). Which allows you to
embed the content of a file or directory within your program at compile time. And use it as if it was on the system at runtime.

## Quick Start

Let's assume you have a js file at `src/main.js`, You want to embed this file within your script and run it with `node` later at runtime.

```sh
@embed src/main.js


embed cat main.js | node
```

The `@embed` special notation tells the compiler to embed the file `src/main.js` at compile time. The `embed` builtin command is used to access embedded files at runtime.

## `@embed` directive

To embed a file or directory, you can use the `@embed` directive followed by one or more relative paths.

```sh
@embed path/to/file
@embed path/to/directory

# or even shorter
@embed path/to/file path/to/directory
```

Compiler directives are special statements processed by the compiler at compile time. And `@embed` is a compiler directive. It's not a command or function that will run at runtime.

The `@embed` directive causes the given files and directories to be embedded within the binary program. The order of paths is not significant.

You are free to arrange the `@embed` directives in any order you want. For example:

```sh
@embed file

command

@embed file2

```

Both `file` and `file2` will be embedded at compile time no matter what the order is. And most importantly, and because `@embed` is a compiler directive. you can **NOT**
nest it, say within a function or condition. For example, You cannot do this:

```sh
if true; then
    @embed file # ❌ compile error
then

function foo() {
    @embed file # ❌ compile error
}

...
```

The path (s) given to the `@embed` directive must be relative and local to the current working directory. And it must not contain the character `\`, `"`, `'`, `<`, `>`, `|`, `?`, `*` or `:`.

> [!TIP]
> If you want to embed the current working directory, you can use the dot `@embed .`

> [!WARNING]
> There are special paths that are ignored and cannot be embedded. This might change in future but for now, the following paths will be ignored from embedding:
> `go.mod`, `.git/*`

## Accessing embedded files

The `embed` builtin command provides and interface to access files embedded in your program.

```sh
├── file
└── src
	├── main.js
	└── utils.js
```

```sh
@embed file src
```

1. Read a file

```sh
embed cat file
embed cat src/main.js
```

2. List content of a directory

```sh
embed ls src
```
