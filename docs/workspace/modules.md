# Modules

In traditional shells such as bash and zsh, code is written in a single file. which can be called as a command later from other scripts. each script
runs in its own shell which forces encapsulation per-file.

There is nothing wrong with that given the fact that those are interpreted languages. However, once the file exceeds a given count of lines of code. it becomes hard to read, maintain and extend. That's the time where you start thinking of switching to other languages (which usually is python :blush: ).

Bunster is built on the idea of modules. a module consists of one or more files within a directory.

Each bunster module contains a main script (usually named `main.sh`) which is the entry point from which the program will start executing. and declaration files which are all other files in the module that have the `.sh` extension.

The main script looks just like any other shell script. However, declaration files (which are all `.sh` files in the module) are only allowed to declare functions, you cannot run commands in the global scope.

## Creating a module

To create a module, all you need is a single file named `main.sh` in current working directory.

Within a directory of your choice, create a file named `main.sh` with the given content:

_main.sh_

```sh
echo "Hello World"
```

Then run:

```sh
bunster build -o hello
```

this will build an executable program named `hello` in current working directory.

---

### Add more files to this module

within same directory, create another file `math.sh`:

_math.sh_

```sh
function calculatePercentage(){
	total=$1
	x=$1

	echo $(( (x / total) * 100 ))
}

function calculateAge(){
	bith_year=$1
	currentYear=$(date +%Y)

	echo $(( bith_year - currentYear ))
}
```

Then update `main.sh` to use this function:

```sh
calculatePercentage 1024 30

calculateAge 2003
```
