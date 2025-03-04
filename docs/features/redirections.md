# Redirections
When you run a command. There are 3 file descriptors open by default. `0` which refers to the standard input of that command. `1` refers to standard output. And  `2` refers to standard error.

You can use a special notation to redirect those file descriptors to other files, duplicate them or close them. 

## Input redirection
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

## Output redirection
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

## Appending redirected output
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

## Redirecting standard output and standard error

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

## Appending standard output and standard error
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

## Here string
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

## Redirections Order
It's worth mentioning that the order of the redirection is significant. For example:
```sh
ls > list.txt 2>&1
```
directs both standard output (file descriptor `1`) and standard error (file descriptor `2`) to the file `list.txt`, while the command:
```sh
ls 2>&1 > list.txt
```
directs only the standard output to file `list.txt`, because the standard error was made a copy of the standard output before the standard output was redirected to `list.txt`.

## Special files
There are some files that are handled specially when used in redirections. These files are not necessarily available on the host system. Bunster will emulate them with the behavior described below.

| File Name | Behavior | Example |
|--|--|--|
| `/dev/stdin` | Refers to the standard input of the command being run | `cmd </dev/stdin` is equivalent to `cmd <&0`|
| `/dev/stdout` | Refers to the standard output of the command being run | `cmd >/dev/stdout` is equivalent to `cmd >&1`|
| `/dev/stderr`  | Refers to the standard error of the command being run | `cmd </dev/stderr` is equivalent to `cmd >&2`|


## Internals of file descriptor management
It is necessary to make something clear about how file descriptors are managed in bunster. `bash` users usually struggle when they see a bunster script using a redirection like this:

```sh
echo foobar 9999999>file.txt >&9999999
```
This code will fail in `bash` because `9999999` is not a valid file descriptor. This is logical because `bash` relies on the kernel to manage file descriptors. Which -- the kernel -- has a strict rules about the format of file descriptors. Also, when you run this code in bash:

```sh
cmd 5>file.txt
```
Bash literally opens the file descriptor `5`. As a result. the command `cmd`  will inherit that file descriptor and can read/write to it.

Bunster is built differently. The following code is going to work totally fine in bunster :
```sh
echo foobar 9999999>file.txt  >&9999999
```
That's because file descriptors are managed by the bunster runtime. And are not real file descriptors. They're just an alias of a file handler. This means that in the above example. The command `echo` will not inherit the file descriptor `9999999` because that file descriptor is not open in reality. 

### Why should I be concerned about file descriptor managment ?
you shouldn't ! there is nothing to worry about regarding the managment of file descriptors. Because the behavior of your script is totally compatible between `bash` and `bunster`. You don't have to change anything in your `bash` scripts to work in bunster. And vice-versa. The only reason why we decided to mention all these information in the documentation is to clarify that while the behavior is similar. The internals are different. and so don't worry when you run `ls /proc/self/fd` and you see a different output in bunster than in bash. 
