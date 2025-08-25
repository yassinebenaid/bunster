<div align="center">
   <img width="200" src="./docs/public/logo.png"/>

# Bunster

</div>

<div align="center">

[![CI](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml/badge.svg)](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml)
[![Read - documentation](https://img.shields.io/badge/Read-documentation-9c2e5c)](https://bunster.netlify.app)

</div>

A shell compiler that converts shell scripts into secure, portable, and static binaries. Unlike other tools (i.e. [shc](https://github.com/neurobin/shc)), Bunster does not just wrap your script within a binary. It literally compiles them to standalone shell-independent programs.

Under the hood, **Bunster** transpiles shell scripts into [Go](https://go.dev) code. Then it uses the [Go Toolchain](https://go.dev/dl) to compile the code to an executable.

**Bunster** aims to be compatible with `bash` as a starting move. Expect that most `bash` scripts will just work with bunster. Additional shells will be supported as soon as we release v1.

> [!WARNING]
> This project is in its early stages of development. [Only a subset of features are supported so far](https://bunster.netlify.app/features/simple-commands).

## Features

In addition to being compatible with bash. bunster offers a lot of additional features that empower its uniqueness:

- **Static binaries**: scripts compiled with bunster are not just wrappers around your script, nor do they rely on any external shells on your system.

- **Modular**: unlike traditional shells scripts that are written in a single file, bunster offers a module system that allows you to distribute code across as many files as needed. [learn more](https://bunster.netlify.app/workspace/modules)

- **Package Manager**: bunster has a builtin package manager that makes it easy to publish and consume modules as libraries. [learn more](https://bunster.netlify.app/workspace/modules)

- **Native `.env` files support**: `.env` files are natively supported in bunster, allowing you to load variables from `.env` files at runtime. [learn more](https://bunster.netlify.app/features/environment-files)

- **Static assets embedding**: bunster allows you to embed files and directories within your compiled program at compile time. Simply use them as if they were normal files in the system at runtime. [learn more](https://bunster.netlify.app/features/embedding)

- **Builtin flags parsing**: You no longer have to bother yourself parsing flags manually. Just declare the flags you expect, and let bunster do the rest. [learn more](https://bunster.netlify.app/features/functions#flags)

- **Static analysis**: bunster statically analyzes your scripts and reports potential bugs at compile time. (_wip_)

## Get Started

<img src="./docs/public/bunster.gif"/>

[Learn more about the usage of bunster.](https://bunster.netlify.app)

## Installation

We have a bash script that installs `bunster` and adds it to your `$PATH`.

```shell
curl -f https://bunster.netlify.app/install.sh | bash
```

The script will install bunster at: 

- `~/.local/bin/bunster` on Linux, and
- `~/bin/bunster` on macOS.

If you want to install the binary system-wide, and make it accessible by all users:

```shell
curl -f https://bunster.netlify.app/install.sh | GLOBAL=1 bash
```

### Homebrew

```sh
brew install bunster
```

Checkout the [documentation](https://bunster.netlify.app/installation) for other ways to install bunster.

## Versioning

Bunster follows [SemVer](https://semver.org/) system for release versioning. On each minor release `v0.x.0`, you can expect adding new features, code optimization, and build improvements. On each patch release `v0.N.x`, you can expect bug fixes and/or other minor enhancements.

Once we reach the stable release `v1.0.0`, you can expect your bash scripts to be fully compatible with bunster (there might be some caveats). All features mentioned above to be implemented unless the community agreed on skipping some of them.

Adding support for additional shells is not planned until our first stable release `v1`. All regarding contributions will remain open until then.

## Developer Guidelines

If you are interested in this project and want to know more about its underlying implementation, or if you want to contribute back but you don't know where to start, [we have a brief article](https://bunster.netlify.app/developers) that explains everything you need to get your hands dirty. Things like:

- the project structure, packages, and their concerns
- how each component works and interacts with other components
- how to add new features
- how to improve existing features
- testing

And anything else in this regard.

## Contributing

Thank you for considering contributing to the Bunster project! The contribution guide can be found in the [documentation](https://bunster.netlify.app/contributing).

This project is developed and maintained by the public community (which includes _you_!). Anything in this repository is subject to criticism. This includes features, the implementation, the code style, the way we manage code reviews, the documentation, and anything else in this regard.

Hence, if you think that we're doing something wrong, or have a suggestion that can make this project better, please consider opening an issue.

## Code Of Conduct

In order to ensure that the Bunster community is welcoming to all, please review and abide by the [Code of Conduct](https://github.com/yassinebenaid/bunster/tree/master/CODE_OF_CONDUCT.md).

## Security

If you discover a security vulnerability within Bunster, please send an e-mail to Yassine Benaid via yassinebenaide3@gmail.com. All security vulnerabilities will be promptly addressed.

Please check out our [Security Policy](https://github.com/yassinebenaid/bunster/tree/master/SECURITY.md) for more details.

## License

The Bunster project is open-sourced software licensed under [The 3-Clause BSD License](https://opensource.org/license/bsd-3-clause).
