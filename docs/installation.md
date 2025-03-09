# Installation

By default, **Bunster** is an independent utility that you can install and start using right away. However,
**Bunster** does assume that you have [the Go toolchain](https://go.dev/dl) installed and accessbile in `PATH`.

We rely on [gofmt](https://pkg.go.dev/cmd/gofmt) provided by the [the Go toolchain](https://go.dev/dl) to format the generated
code. This makes it easy to debug and/or to learn from. Additionally, We use [the Go compiler](https://go.dev/dl) to compile
the code and generate the executable for you.

> [!info]
The absence of [the Go toolchain](https://go.dev/dl) does not affect the working of **Bunster**. It's still going to work just fine.
If you only care about the generated Go code, and don't want **Bunster** to automatically compile the exectuable for you.
That is totally fine and you can go for it.

## Linux/Mac
We have bash script that installs `bunster` and adds it to your `$PATH`.

```shell
curl -f https://bunster.netlify.app/install.sh | bash
```

The script will install bunster at `~/.local/bin/bunster` on linux. And `~/bin/bunster` on mac. If you want to install the binary system wide and make it accessible by all users.

```shell
curl -f https://bunster.netlify.app/install.sh | GLOBAL=1 bash
```

> [!warning]
> Do not trust scripts downloaded from the interne. take a look at the code before running it.

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

> [!warning]
Only `linux` and `macos` binaries are available at the moment. `windows` support is coming soon.

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

> [!info]
If you choose to install using `go install`. make sure that `$HOME/go/bin` is added to your `PATH`. If not yet, Please add
`export PATH=$PATH:$HOME/go/bin` to one of your profile files. eg. *`~/.bashrc` if you're using `bash`, or `~/.zshrc` if you're using `zsh`*.

## Nixpkgs (NixOS & Nix)
If you use [Nix](https://nixos.org), you can install **[Bunster](https://search.nixos.org/packages?channel=unstable&show=bunster&from=0&size=50&sort=relevance&type=packages&query=bunster)** from the `unstable` branch of `nixpkgs`:
### NixOS Config
Add the following Nix code to your NixOS Configuration, usually located in `/etc/nixos/configuration.nix`
```
  environment.systemPackages = [
    pkgs.bunster
  ];
```
### Nix Shell
A nix-shell will temporarily modify your `$PATH` environment variable. This can be used to try a piece of software before deciding to permanently install it.
```shell
nix-shell -p bunster
```
### Nix Env (not recommended)
#### On NixOS
```shell
nix-env -iA nixos.bunster
```
#### On Non-NixOS
```shell
# without flakes:
nix-env -iA nixpkgs.bunster
# with flakes:
nix profile install nixpkgs#bunster
```
> [!warning]
> **Bunster** is currently only available in the `unstable` branch of Nixpkgs but will be coming to the `stable` branch soon.
