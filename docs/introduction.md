# Introduction

Bunster is a shell compiler that converts shell scripts into secure, portable and static binaries. Unlike other tools (ie. [shc](https://github.com/neurobin/shc)), Bunster does not just wrap your script within a binary. It literally compiles them to standalone shell-independent programs.

Technically speaking, **Bunster** in fact is a `shell-to-Go` [Transpiler](https://en.wikipedia.org/wiki/Source-to-source_compiler) that generates [Go](https://go.dev) source out of your scripts. Then, optionally uses the [Go Toolchain](https://go.dev/dl) to compile the code to an executable.

**Bunster** aims to be compatible with `bash` as a starting move. You should expect your `bash` scripts to just work with bunster. Additional shells will be supported as soon as we release v1.

> [!WARNING]
> This project is in its early stages of development. [Only a subset of features are supported so far](https://bunster.netlify.app/supported-features.html).

## Vision

Bunster has a vision to make shell scripts feel like any modern programming language. With as many features as we could, without any bloating. anything that
makes you feel happy when writing shell scripts. a feeling that shells usually don't provide, a feeling that languages like Go give you, we aim to:

- Improve error handling and messages, we want to help everyone write error-aware scripts. And when they fail, we want to give them clear, concise error messages.
- Introduce a module system that allows you to publish and consume scripts as libraries, with a builtin package manager.
- Add first-class support for a wide collection of builtin commands that just work out of the box. You don't need external programs to use them.
- Add first-class support for `.env` files. Allowing you to load variables from `.env`.
- Support static asset embedding. This feature allows you to embed a file's content to a variable at build time. ([Go has one already](https://pkg.go.dev/embed))
- Support different shells and POSIX.


## Get Started

<img src="./bunster.gif"/>

[Learn more about the usage of bunster.](https://bunster.netlify.app)
