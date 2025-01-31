<div align="center">
    <img width="200" src="./docs/public/logo.png"/>

# Bunster

</div>

<div align="center">

[![CI](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml/badge.svg)](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml)
[![Documentation](https://img.shields.io/badge/Documentation-e57884?logo=BookStack&logoColor=9c2e5c)](https://bunster.netlify.app)

</div>

A shell compiler that converts shell scripts into secure, portable static binaries. Unlike other tools (ie. [shc](https://github.com/neurobin/shc)), Bunster does not just wrap your script within a binary. It literally compiles them to standalone shell-independent programs.

Technically speaking, **Bunster** in fact is a `shell-to-Go` [Transpiler](https://en.wikipedia.org/wiki/Source-to-source_compiler) that generates [Go](https://go.dev) source out of your scripts. Then, optionally uses the [Go Toolchain](https://go.dev/dl) to compile the code to an executable.

**Bunster** aims to be compatible with `bash` as a staring move. You should expect your `bash` scripts to just work with bunster. additional shells will be supported as soon as we release v1.

> [!WARNING]
> This project is in its early stages of development. [Only a subset of features are supported so far](https://bunster.netlify.app/supported-features.html).

## Goals

Bunster's primary goal is to make shell scripts feel like any modern programming language. Additional goals include (_but not limited to_) :

- **Modules**: We aim to introduce a module system that allow you to publish and consume scripts as libraries. with a builtin package manager.
- **Standard library**: we aim to add first-class support for a wide collection of builtin commands that just work out of the box. you don't need external programs to use them.
- **Builtin support for .env files**: Bunster aims to allow you to load variables from `.env` files. 
- **Static Asset Embedding**: This feature allows you to embed a file's content to a variable at build time. ([Go has one already](https://pkg.go.dev/embed))
- **Different Shells support**: Bunster aims to add support for a variety of shells in the future. But as a starting move.

## Get Started

 <img src="./docs/public/bunster.gif"/>

[Learn more about the usage of bunster.](https://bunster.netlify.app)

## Installation

Checkout the [documentation](https://bunster.netlify.app/installation) for different ways of installation.

## FAQ
### How does bunster make your scripts more secure?
It does not, Bunster does not make your script more secure, Instead, it makes your environment more secure. Bunster programs run in environments where the shell doesn't exist. for example in a cloud server. By eliminating the shell, You remain safe from a variety of security risks (examples include: [RCE](https://www.invicti.com/learn/remote-code-execution-rce/), [Command Injection](https://www.imperva.com/learn/application-security/command-injection/#:~:text=Code%20injection%20is%20a%20generic,proper%20input%2Foutput%20data%20validation.), [Reverse Shell](https://www.wiz.io/academy/reverse-shell-attacks), ...)

_more about this topic on this discussion:_ https://github.com/yassinebenaid/bunster/discussions/126.

### Does bunster replace the programs in my script?
No, bunster replaces the shell, but you still need to have the programs available in `$PATH` to use them. Later we might add first-class support for a subset of commands builtin.  

### Is bunster a drop-in replacement for bash ?
Yes and No. Yes bacause bunster aims to be compatible with bash, with a lot of additional features, Bunster worth to be your primary tool for shell scripting.
No because bunster programs are binary, they're cool if they're yours. But very bad if downloaded from the internet. Also, if you only have a short script that runs few commands. It doesn't worth it to install an entire toolchain, You can use the shell.

- use bunster for long, complex scripts.
- don't use bunster for short simple scripts.
  
## Versioning

Bunster follows [SemVer](https://semver.org/) system for release versioning. On each minor release `v0.x.0`, You should expect adding new features. Code optimization and build improvements. On each patch release `v0.N.x`, you should expect bug fixes and/or other minor enhancements.

Once we reach the stable release `v1.0.0`, you must expect your bash scripts to be fully compatible with Bunster (_there might be some caveats_). All features mentioned above to be implemented unless the community agreed on skipping some of them.

Adding support for additional shells is not planned until our first stable release `v1`. All regarding contributions will remain open until then.


## Contributing

Thank you for considering contributing to the Bunster project! The contribution guide can be found in the [documentation](https://bunster.netlify.app/contributing).

This project is developed and maintained by the public community, which includes you. Anything in this repository is subject to criticism. Including features, the implementation, the code style, the way we manage code reviews, The documentation and anything else in this regard.

Hence, if you think that we're doing something wrong, or have a suggestion that can make this project better. Please consider opening an issue.

## Code Of Conduct

In order to ensure that the Bunster community is welcoming to all, please review and abide by the [Code of Conduct](https://github.com/yassinebenaid/bunster/tree/master/CODE_OF_CONDUCT.md).

## Security

If you discover a security vulnerability within Bunster, please send an e-mail to Yassine Benaid via yassinebenaide3@gmail.com. All security vulnerabilities will be promptly addressed.

Please check out our [Security Policy](https://github.com/yassinebenaid/bunster/tree/master/SECURITY.md) for more details.

## License

The Bunster project is open-sourced software licensed under [The 3-Clause BSD License](https://opensource.org/license/bsd-3-clause).
