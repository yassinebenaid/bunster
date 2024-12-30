<div align="center">
    <img width="200" src="./docs/public/logo.png"/>

# Bunster

</div>

<div align="center">

[![CI](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml/badge.svg)](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml)
[![Documentation](https://img.shields.io/badge/Documentation-e57884?logo=BookStack&logoColor=9c2e5c)](https://bunster.netlify.app)

</div>

Have you ever wished your shell scripts could be faster, more portable, and secure ? **Bunster** brings this to life by transforming your shell scripts into efficient, standalone binaries that are easy to distribute and deploy across platforms _(only unix is supported at the moment)_.

Unlike other tools, **Bunster** doesn’t just wrap your scripts in a binary—it compiles them down to efficient native machine code, leveraging the powerful Go toolchain. This ensures performance, portability, and robustness.

Technically speaking, **Bunster** in fact is a `shell-to-Go` [Transplier](https://en.wikipedia.org/wiki/Source-to-source_compiler) that generates [Go](https://go.dev) source out of your scripts. Then, optionally uses the [Go Toolchain](https://go.dev/dl) to compile the code to an executable program.

**Bunster** targets `bash` scripts in particular. The current syntax and features are all inherited from `bash`. additional shells will be supported as soon as we relase v1.

### Goals

In addition to the shell features, We aim to add several custom features to make shell scripts feel like any modern programming language. These features are either supported or are planned to be implemented in future. (_consider contributing to help us speed up the develpment cycle_)

- **Different Shells support**: Bunster currently aims to be compatible with `bash` as a starting move. Then additional shells in future.
- **Modules**: Something shell scripts lack is a module system, we aim to introduce a module system that allow you to publish and consume scripts as libraries.
- **Static Asset Embedding**: This feature allows you to embed a file's content to a variable at build time. ([Go has one already](https://pkg.go.dev/embed))
- **Password and Expiration Lock**: Surprisingly, some people have asked for this feature. Basically, It allows you to choose an expirity date at build time. the generated program will not work after that date. Also you can choose to lock the script using a password. whenever you try to run it, it prompts for the password.

> [!WARNING]
> This project is in its early stages of development, and is not yet ready for production. Not all features are implemented yet. But, plenty of them are. [see what features are implemented so far](https://bunster.netlify.app/supported-features.html).

### Versionning

Bunster follows [SemVer](https://semver.org/) system for release versionning. On each minor release `v0.x.0`, You should expect adding new features. Code optimization and build improvements. On each patch release `v0.N.x`, you should expect bug fixes and/or other minor enhancements.

Once we reach the stable release `v1.0.0`, you must expect your bash scripts to be fully compatible with Bunster (_there might be some caveates_). All features mentioned above to be implemeted unless the comunity agreed on skipping some of them.

Adding support for additional shells is not planned until our first stable release `v1`. All regarding contributions will remain open until then.

### Installation

Checkout the [documentation](https://bunster.netlify.app/installation) for different ways of installation.

### Contributing

Thank you for considering contributing to the Bunster project! The contribution guide can be found in the [documentation](https://bunster.netlify.app).

This project is developed and maintained by the public community, which includes you. Anything in this repository is subject to criticism. Including features, the implementation, the code style, the way we manage code reviews, The documentation and anything else in this regard.

Hence, if you think that we're doing something wrong, or have a suggestion that can make this project better. Please consider openning an issue.

### Code Of Conduct

In order to ensure that the Bunster community is welcoming to all, please review and abide by the [Code of Conduct](https://github.com/yassinebenaid/bunster/tree/master/CODE_OF_CONDUCT.md).

### Security

If you discover a security vulnerability within Bunster, please send an e-mail to Yassine Benaid via yassinebenaide3@gmail.com. All security vulnerabilities will be promptly addressed.

Plase check out our [Security Policy](https://github.com/yassinebenaid/bunster/tree/master/SECURITY.md) for more details.

### Licence

The Bunster project is open-sourced software licensed under the [GPLv3 license](https://www.gnu.org/licenses/gpl-3.0.en.html).
