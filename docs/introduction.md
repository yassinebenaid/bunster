# Introduction

Bunster is a shell compiler that converts shell scripts into secure, portable, and static binaries. Unlike other tools (e.g., [shc](https://github.com/neurobin/shc)), Bunster does not just wrap your script within a binary—it literally compiles them into standalone, shell-independent programs.

Technically speaking, **Bunster** is, in fact, a `shell-to-Go` [transpiler](https://en.wikipedia.org/wiki/Source-to-source_compiler) that generates [Go](https://go.dev) source code from your scripts. Then, it optionally uses the [Go Toolchain](https://go.dev/dl) to compile the code into an executable.

**Bunster** aims to be compatible with `bash` as a starting point. You should expect your `bash` scripts to just work with Bunster. Additional shells will be supported as soon as we release v1.

> [!WARNING]
> This project is in its early stages of development.

## Vision

Bunster has a vision to make shell scripts feel like any modern programming language, with as many features as possible—without any bloat. Anything that makes you enjoy writing shell scripts, a feeling that shells usually don't provide, a feeling that languages like Go give you. We aim to:

- Improve error handling and messages. We want to help everyone write error-aware scripts, and when they fail, we want to give them clear, concise error messages.
- Introduce a module system that allows you to publish and consume scripts as libraries, with a built-in package manager.
- Add first-class support for a wide collection of built-in commands that just work out of the box. You don’t need external programs to use them.
- Add first-class support for `.env` files, allowing you to load variables from `.env`.
- Support static asset embedding. This feature allows you to embed a file's content into a variable at build time. ([Go has one already](https://pkg.go.dev/embed))
- Support different shells and POSIX.

## Get Started

<img src="/bunster.gif"/>

[Learn more about the usage of Bunster.](/cli)

## Bunster vs. Bash

We are not trying to compete with `bash`, because unlike `bash`, Bunster is not a shell—it is a programming language dedicated to scripting. `bash` has a scripting plan where you can write commands in scripts, but its primary use case is interactive use.

Bunster is trying to take all the good things from `bash`, add more features to them, and make them available in a simple and familiar programming language.

## Why do we have separate documentation?

We use the [bash reference](https://www.gnu.org/software/bash/manual/bash.html) as a source of truth, and so can you.
Because we have a promise to maintain compatibility with `bash`, you are free to refer to that manual for documentation on features supported by Bunster.

However, we decided to create our own documentation so that we can focus on the features we support.
Also, the Bash reference often mentions things that have nothing to do with us—like interactive command-line usage—which can be confusing for our users.
