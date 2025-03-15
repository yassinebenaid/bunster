# Pipelines

A pipeline is a sequence of one or more commands separated by `|` or `|&`.

The format for a pipeline is:

```sh
[!] command1 [ | command2 ] ...
```

The output of each command in the pipeline is connected via a pipe to the input of the next command.
That is, each command reads the previous command’s output. This connection is performed before any redirections specified by `command1`.

Each command in the pipeline runs asynchronously in a [sub-shell](/features/groups-and-subshells). The pipeline waits for all commands to finish before it exits.

An example of a pipeline can be illustrated as follows:

```sh
cat file.txt | sort | uniq
```

![pipeline](/assets/pipeline.png)

## Piping all output

When using the `|` popeline separator. Only standard output is connected to the pipe. To illustrate this, let's look at this example:

```sh
mkdir foo | tr
```

![pipeline](/assets/pipeline-2.png)

When the separator `|&` is used, `command1`'s standard error, in addition to its standard output, is connected to `command2`'s standard input through the pipe; it is shorthand for `2>&1 |`. This implicit redirection of the standard error to the standard output is performed after any redirections specified by `command1`.

Example:

```sh
mkdir foo |& tr
```

![pipeline](/assets/pipeline-3.png)

## Exit code

The exit status of a pipeline is the exit status of the last command in the pipeline. For example:

```sh
command1 | command2 | command3
```

The exit code of this pipeline is the exit code of the `command3` command. No matter what the exit code of other commands is.

### Exit code negation

If the pipeline is preceeded by an exclamation mark `!`. the exit code of that pipeline is negated. It will be `0` if the last command in the pipeline exits with a non-zero exit code. and `1`otherwise.

For example:

```sh
! command
```

The exit code for this pipeline is zero `0` if the exit code of the command `command` is non zero. (`1`, `2` ...). If the exit code of the command `command` is zero `0`. Then the pipeline's exit code is one `1`.

In other words. the `!` mark negates the exit code of the pipeline.

> [!info]
> Note that a pipeline does not have to contain more than one command. The same rules that apply to pipelines with many commands apply to pipelines with one command.
