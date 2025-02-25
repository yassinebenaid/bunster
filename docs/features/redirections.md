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
