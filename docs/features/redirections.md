# Redirections
When you run a command. there are 3 file descriptors open by default. `0` which referes to the standard input of that command. `1` refers to standard output. And  `2` refers to standard error.

Before the command runs. you can use a special notation to redirect those file descriptors, duplicate them, close them, or make them refer to other files. 

## Input Redirection
The general format is as follows:

```sh
[n]<word
```

This will cause the file whose name results from the expansion of `word` to be opened for reading on file descriptor `n`. Or standard input (file descriptor `0`) if `n` is not specified.

### Example
```sh
cat <file.txt 

# Or
cat 0<file.txt 

# Or
cat 3<file.txt 
```

## Output Redirection
The general format is as follows:

```sh
[n]>word
```

This will cause the file whose name results from the expansion of `word` to be opened for writing on file descriptor `n`. Or standard output (file descriptor `1`) if `n` is not specified.
If the file does not exist it is created. If it does exist it is truncated to zero size.

### Example
```sh
echo "Hello World" >file.txt 

# Or
echo "Hello World" 1>file.txt

# Or
echo "Hello World" 3>file.txt
```

## Appending Redirected Output
The general format for appending output is as follows:

```sh
[n]>>word
```

Redirection of output in this fashion causes the file whose name results from the expansion of `word` to be opened for appending on file descriptor `n`, or the standard output (file descriptor `1`) if n is not specified. If the file does not exist it is created.

### Example
```sh
echo "Hello World" >>file.txt 

# Or
echo "Hello World" 1>>file.txt

# Or
echo "Hello World" 3>>file.txt
```

## Redirecting Standard Output and Standard Error

This construct allows both the standard output (file descriptor `1`) and the standard error output (file descriptor `2`) to be redirected to the file whose name is the expansion of `word`.

The format is:
```sh
&>word
```

This is semantically equivalent to:

```sh
>word 2>&1
```

### Example
```sh
echo "Hello World" &>file.txt 

# Or
echo "Hello World" 1>file.txt 2>&2
```

## Appending Standard Output and Standard Error
This construct allows both the standard output (file descriptor `1`) and the standard error output (file descriptor `2`) to be appended to the file whose name is the expansion of `word`.

The format of appending all output is:
```sh
&>>word
```

This is semantically equivalent to:

```sh
>>word 2>&1
```


### Example
```sh
echo "Hello World" &>>file.txt 

# Or
echo "Hello World" 1>>file.txt 2>&2
```

## Here String
The format is simple as:
```sh
[n]<<<word
```

This notation causes the string resulted from the expansion of `word` to be provided as input at the file descriptor `n`. Or at the standard input (file descriptor `0`) if `n` is not specified.

### Example
```sh
cat <<<"hello World"

# Or
cat 0<<<"Hello World"

# Or
cat 3<<<"Hello World"
```

## Duplicating file descriptors
Duplication means to make a file descriptor refers to the same file that another file descriptor refers to. For example. Duplicating file descriptor `x` on `y` means to make the file descriptor `x` refers to the same file that file descriptor `y` refers to. 

The format is:
```sh
[n]<&word
```

This format causes the file descriptor resulted from the expansion of `word` to be duplicated on file descriptor `n`. Or on standard input (file descriptor `0`) if `n` is not specified.

It is possible to use this other format as well:
```sh
[n]>&word
```

The only difference is that when you omit `n`. The default is standard output (file descriptor `1`).

### Example
In this example, we will duplicate the file descriptor `1` on `2`. 

```sh
echo foobar 1>file.txt 2>&1
```

The file descriptor `1` is made to refer to the file `file.txt`. then we duplicated the file descriptor `1` on file descriptor `2`. This means that when the command `echo` writes to file descriptor `2`. It will write to the file `file.txt`. In other words. the file descriptor `2` is now a copy of the file descriptor `1`.

## Closing file descriptors
The syntax for closing file descriptors is 

```sh
[n]<&-
```

Notice the `-` at the end. This notation causes the file descriptor `n` to be closed. or standard input (file descriptor `0`) if `n` is not specified.

It is possible to use this other format as well:
```sh
[n]>&-
```
The only difference is that when you omit `n`. The default is standard output (file descriptor `1`).

### Example
```sh
echo foobar 1>&-
```
In this example, we closed the file descriptor `1`. when the command `echo` writes to file descriptor `1`. It will fail because the file descriptor is closed.

## Moving file descriptors
This notation is a combination of moving and closing file descriptors. In simple words, moving file descriptor `x` to `y` means to duplicate file descriptor `y` to `x`. then close the file descriptor `y`.

The format is:

```sh
[x]<&y-
```

This means, duplicate the file descriptor `y` to `x`. then close `y`. if `x` is not specified, `0` is the default. An alternative syntax is available:

```sh
[x]>&y-
```
The only difference is that when you omit `x`. The default is standard output (file descriptor `1`).

### Example
In this example, we will move the file descriptor `2` to `1`.

```sh
echo foobar 1>&2-
```
Now, the file descriptor `1` is made a copy of the file descriptor `2`. and the file descriptor `2` was closed. This means. when the command `echo` writes to file descriptor `1`. It will write to the same file to which file descriptor `2` was pointing to. and it will fail to write to file descriptor `2` because it was closed.

## Open file for reading and writing
The format is:
```sh
[n]<>word
```

This notation causes the file whose name is the expansion of `word` to be opened for both reading and writing on file descriptor `n`, or on file descriptor 0 if `n` is not specified. If the file does not exist, it is created.
