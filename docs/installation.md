# Installation

By default, **Bunster** is an independent utility that you can install and start using right away. However,
**Bunster** does assume that you have [the Go toolchain](https://go.dev/dl) installed and accessbile in `PATH`.

We rely on [gofmt](https://pkg.go.dev/cmd/gofmt) provided by the [the Go toolchain](https://go.dev/dl) to format the generated
code. This makes it easy to debug and/or to learn from. Additionally, We use [the Go compiler](https://go.dev/dl) to compile
the code and generate the executable for you.

::: info
The absence of [the Go toolchain](https://go.dev/dl) does not affect the working of **Bunster**. It's still going to work just fine.
If you only care about the generated Go code, and don't want **Bunster** to automatically compile the exectuable for you.
That is totally fine and you can go for it.
:::


## Docker Image
The easiest way to get **Bunster** is through our official [Docker Image](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-an-image/).
It comes with everything needed. Including the **Bunster** compiler and [the Go toolchain](https://go.dev/dl).

```shell
docker pull ghcr.io/yassinebenaid/bunster:latest
```

Or, if you want a specific version (v.0.3.0 for example):

```shell
docker pull ghcr.io/yassinebenaid/bunster:v0.3.0
```

## Github Release
You can get the latest version of `bunster` from [github releases](https://github.com/yassinebenaid/bunster/releases).

::: warning
Only `linux` and `macos` binaries are available at the moment. `windows` support is coming soon.
:::

## Using Go
If you already have [the Go toolchain](https://go.dev/dl) installed. You can use the `go install` command to get **Bunster** on your machine.

```shell
go install github.com/yassinebenaid/bunster/cmd/bunster@latest
```

Or, if you want a specific version (v.0.3.0 for example):

```shell
go install github.com/yassinebenaid/bunster/cmd/bunster@v0.3.0
```

This will build the binary at `$HOME/go/bin/bunster`, if you want to make it accessible by all users, you can move it to `/usr/local/bin`
```shell
mv $HOME/go/bin/bunster /usr/local/bin # you may need to use `sudo`.
```

::: info
If you choose to install using `go install`. make sure that `$HOME/go/bin` is added to your `PATH`. If not yet, Please add
`export PATH=$PATH:$HOME/go/bin` to one of your profile files. eg. *`~/.bashrc` if you're using `bash`, or `~/.zshrc` if you're using `zsh`*.
:::
