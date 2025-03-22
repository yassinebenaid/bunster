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

Compile this module:

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
	x=$2


	echo $(( (x * 100) / total ))
}

function calculateAge(){
	bith_year=$1
	currentYear=$(date +%Y)

	echo $(( currentYear - bith_year ))
}
```

Then update `main.sh` to use these functions:

```sh
calculatePercentage 1024 30

calculateAge 2003
```

Compile this module:

```sh
bunster build -o hello
```

Now run the program

```sh
./hello
```

This will output:

```txt
2
22
```

> [!TIP]
> Each module can have as many declaration scripts as needed.

## Custom main script

Sometimes the main script is not named `main.sh`. When you compile the program, you can pass the name of the main script as the first argument:

```sh
├── foo.sh
├── bar.sh
└── baz.sh
```

Assuming your main script is `foo.sh`, you can compile this module like this:

```sh
bunster build foo.sh -o hello
```

## Publishing module as library

You may publish your module publically in any git registry (such as github, gitlab etc). This allows others to use your module in their projects by requiring it as dependencies.

> [!WARNING]
> This is a work-in-progrss feature, will be release soon.

## Using external libraries

If you want to use an external library, it must first be available in a git registry such `githab`, `gitlab` etc.

> [!WARNING]
> This is a work-in-progrss feature, will be release soon.
