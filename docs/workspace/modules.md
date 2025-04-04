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

Publishing a bunster module is as easy as hosting it in a git repository on your preferred git registry such as `github`, `gitlab`, `bitbucket` or even your own server.

We call bunster modules published for others to use as `libraries`. as a library, the module may only contain declaration files. in other words. only functions declaraion are allowed in global scope.

### Example library

Create a file named `hello.sh` with given content:

```sh
function hello {
   echo "Hello World ✨"
}
```

that's it, your module is ready. go ahead an put it on a git repository of your choice. in our case we will put it at: `https://github.com/yassinebenaid/hello-bunster`

## Using external libraries

As described above, bunster libraries are hosted on git repositories. for example, we have previously published a library at `https://github.com/yassinebenaid/hello-bunster`.

Now, in your project directory, run the following command:

```sh
bunster get github.com/yassinebenaid/hello-bunster@684da09acea05d9351c4c61d4296bc696f729533
```

this command fetches the library from the repository `github.com/yassinebenaid/hello-bunster` at commit `684da09acea05d9351c4c61d4296bc696f729533`. And will update the file `bunster.yml` with the content:

```yaml
require:
  github.com/yassinebenaid/hello-bunster: 684da09acea05d9351c4c61d4296bc696f729533
```

That's it, you can use the functions form the library in your own module :

_main.sh_

```sh
hello
```

Output:

```txt
Hello World ✨
```

### Editing `bunster.yml` manually

You can list your dependencies manually in `bunster.yml` file. for example:

```yaml
require:
  github.com/foo/bar: 684da09acea05d9351c4c61d4296bc696f729533
  gitlab.com/baz/boo: 0ef176a380bb9c2410298c7444c93d62fd915357
  bitbucket.com/baz/boo: 086d715616f6ec157fb8f2544aa025883acba649
```

You can download these modules using the command:

```sh
bunster get --missing
```

### Why commit hash as version

Security, yes, commit hash is the most secure way to trust the content of a library. it's not the most readable, beutiful or friendly. but that's not as important as security.

If we were to use semantic versionning, library authors can still edit the release content. and no one can trust what the authors may change.

However, commit hash is unique enough that can never be altered. worst thing that can happen is the commit to be deleted.

At least for now, we will continue to use commit hash as version. I don't have the budget to host a checksum database and allow use of semantic versionning as well.
