# Builtin Commands

There are a bunch of commands built into the language. meaning that they are not external programs in the system. Thry are implemented internally in bunster.

[`true`](#true) [`false`](#false) [`shift`](#shift) [`loadenv`](#loadenv) [`embed`](#embed)

### `true`

Exits with zero.

### `false`

Exits with non-zero exit code. (`1`).

### `shift`

Used to shift the positional arguments to left by `N`. if no arguments passed, then `N` is 1. if an argument is passed. it should be an integer. and in that case `N` is the first argument.

For example, calling `shift` with no arguments causes `$1` to become `$2`, `$2` to become `$3`, `$3` to become `$4` and so on. If an argument is passed. for example `3`. The arguments are shifted by `3` steps. `$1` becomes `$4`, `$2` becomes `$5`, `$3` becomes `$6` and so on

### `loadenv`

Used to deal with `.env` files. more information about it in [Environment Files section](/features/environment-files)

### `embed`

Used to access embedded files. more information about it in [Embedding section](/features/embedding)
