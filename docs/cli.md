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
