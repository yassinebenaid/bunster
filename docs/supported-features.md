# Supported Features
This page lists the features and syntax supported by **Bunster**. you must consider any feature that is not mentioned here
as work-in-progress and as not-yet-supported.


## Simple commands
```shell
command argument 'argument' "argument" $VARIABLE
```

## Redirections
### Output redirection
```shell
command >file.txt >|file.txt >>file.txt &>file.txt &>>file.txt
```
### Input redirection
```shell
command <file.txt <<<"foo bar"
```
### File descriptor duplication and closing
```shell
command 2>&3 3>&4- 3>&- 2>&-
```
### Opening files for reading and writing
```shell
command <>file.txt 3<>file.txt
```

::: tip
Unlick bash, you are not limited to only use file descriptors between 0 to 9. Feel free to use any valid 64-bit integer number. (learn why TODO: add this link here)
:::

## Passing shell parameters as environment to commands
```shell
key=value key2="value" command
```

## Pipelines
```shell
command | command2 | command3
```
