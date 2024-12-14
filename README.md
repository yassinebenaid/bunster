<div align="center">
    <img width="200" src="./docs/public/logo.png"/>

# Bunster

</div>

<div align="center">

[![CI](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml/badge.svg)](https://github.com/yassinebenaid/bunster/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/yassinebenaid/bunster/graph/badge.svg?token=56Vp2OyU5t)](https://codecov.io/gh/yassinebenaid/bunster)
[![Documentation](https://img.shields.io/badge/Documentation-e57884?logo=BookStack&logoColor=9c2e5c)](https://bunster.netlify.app)

</div>

Have you ever wished your shell scripts could be faster, more portable, and secure ? **Bunster** brings this to life by transforming your shell scripts into efficient, standalone binaries that are easy to distribute and deploy across platforms _(only unix is supported at the moment)_.

Unlike other tools, **Bunster** doesn’t just wrap your scripts in a binary—it compiles them down to efficient native machine code, leveraging the powerful Go toolchain. This ensures performance, portability, and robustness.

Technically speaking, **Bunster** is not a complete compiler, But rather a **Transplier** that generates **GO** code out of your scripts. Then, opionally uses the **Go Toolchain** to compile the code to an executable program.

**Bunster** targets `bash` scripts in particular. The current syntax and features are all inherited from `bash`. other shells support will be added soon.

### Motivation
- **Different Shells support**: Bunster currently aims to be compatible `Bash` as a starting move. then other popular shells in future.
- **Security**: as you may guess, humans cannot read machine code, so why not to compile your scripts to such format.
- **Modules**: something shell scripts lack is a module system, people want to share their code to be used by others, but the infrastructure doesn't allow them. Well, **Bunster** introduces a module system that allow you to publish your scripts as a modules consumed by other users.
- **Performance**: You should't be concerned about performance that much, But at least scripts compiled by **Bunster** will not suffer to parse and interpret
your scripts everytime you run them.

> [!WARNING]
> This project is in its early phase of development and is not yet ready for production. Not all features are implemented yet. But, plenty of them are. [see what features are features so far](https://bunster.netlify.app/supported-features.html).

### Versionning
**Bunster** follows [SemVer](https://semver.org/) system for release versionning. On each `v0.x.0` release, You would expect adding support for new features (can be new shell feature, improvement in the build process, some custom features ...etc.) . On each `v0.N.x` release, you would expect bug fixes and documentation enhancments.

Once we reach the `v1.0.0`, you must expect a +90% of compatibility with `bash`. A module system, a [static assets embedding capabilities](https://pkg.go.dev/embed) and a plenty of other features to make `shell` scripts feel like any other modern programming language.

Adding support for additional shells is not planned until our first stable release `v1`. All regarding contributions will remain open until then.

### Installation
Checkout the [documentation](https://bunster.netlify.app) for different ways of installation.

### Contributing
Thank you for considering contributing to the **Bunster** project! The contribution guide can be found in the [documentation](https://bunster.netlify.app).

### Code Of Conduct
In order to ensure that the Laravel community is welcoming to all, please review and abide by the [Code of Conduct](https://github.com/yassinebenaid/bunster/tree/master/CODE_OF_CONDUCT.md).

### Security
If you discover a security vulnerability within Bunster, please send an e-mail to Yassine Benaid via yassinebenaide3@gmail.com. All security vulnerabilities will be promptly addressed.

Plase check out our [Security Policy](https://github.com/yassinebenaid/bunster/tree/master/SECURITY.md) for more details.

### Licence
The Bunster project is open-sourced software licensed under the [GPL3.0 license](https://www.gnu.org/licenses/gpl-3.0.en.html).
