# CLI

Consider this little script.
_script.sh_:

```shell
echo "Hello World"
```

## Build

To build this script, simply run:

```shell
bunster build script.sh -o my-program
```

This will create an executable program named `my-program` in the current working directory.

Now you can run it

```shell
./my-program
# Output: Hello World
```

## Generate

If you want to only generate the Go code, but not compile it:

```shell
bunster generate script.sh -o my-module
```

This will create a directory program named `my-module` in which the Go source is generated.

## Modules mode

> [!INFO]
> You can learn more about bunster modules [here](/workspace/modules)

If no arguments are supplied, the `build` and `generate` commands assume that current directory is a bunster module
and looks for the file `main.sh`.

Then loads all `.sh` files in the curren directory as declaration files. Otherwise, if a file path is supplied as the first argument. only that file is processed.

for example:

```txt
├── main.sh
├── utils.sh
└── functions.sh
```

If you run:

```sh
bunster build -o test
```

bunster assumes this is a module. and starts by loading `main.sh` as an entry point. then processes all of `utils.sh` and `functions.sh` as declaration scripts.

However, if you supplied a path like this:

```sh
bunster build main.sh -o test
```

bunster only processes the `main.sh` file and ignores all other files. and modules mode is turned off. this means no use of `bunster.yml`.

> [!IMPORTANT]
> In first case. the file extension matters. bunster only processes `.sh` files. **BUT**, in the second case. the file extension **does not matter**. This means you can build any file with whatever extension. `bunster build foo.any -o test`.
