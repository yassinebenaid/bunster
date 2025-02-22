# Simple Commands
You may run simple commands just like in a regular shell. There is no difference between bash and bunster when it comes to running external commands.
which are external programs accessible through the `$PATH` environment variable.

**Format**:
```shell
command-name [arguments...]
```

The command name can be either a program name or a full relative or absolute path to the program. if the name is not a path, the program must be in `$PATH`.


## Quotes
### Double Quotes
Quoting in bunster works the exact same as in bash. it preserves the literal value of the characters. for example:

```shell
"command name" "arguments"
```

The dollar sign is an exception.


```shell
echo "$HOME"
```

This will run the command `echo`. the first argument is the value of the variable `HOME`.
