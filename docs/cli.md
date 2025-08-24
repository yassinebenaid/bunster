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

Now you can run it:

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
> You can [learn more about bunster modules here](/workspace/modules)

If no arguments are supplied, the `build` and `generate` commands assume that current directory 
is a bunster module and looks a file named `main.sh`.

Then bunster loads all `.sh` files in the current directory as declaration files.

Otherwise, if a file path is supplied as the first argument, _only_ that file is processed.

For example, given a directory with the following:

```txt
├── main.sh
├── utils.sh
└── functions.sh
```

If you run:

```sh
bunster build -o test
```

bunster assumes this is a module, and starts by loading `main.sh` as an entry point. 
It then processes all of `utils.sh` and `functions.sh` as declaration scripts.

However, if you supply a path like this:

```sh
bunster build main.sh -o test
```

bunster only processes `main.sh` and ignores all other files. 
In fact, bunster’s modules mode is turned off. This means no use of `bunster.yml`.

> [!IMPORTANT]
> In first case, the file extension matters. bunster only processes `.sh` files.
> **BUT**, in the second case, the file extension **does not matter**.
> 
> This means you can build any file with whatever extension (e.g., `bunster build foo.any -o test`).
