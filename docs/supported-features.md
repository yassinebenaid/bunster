# Supported Features
This page lists the features and syntax supported by **Bunster**. you must consider any feature that is not mentioned here
as work-in-progress and as not-yet-supported.


- [Simple commands](https://www.gnu.org/software/bash/manual/bash.html#Simple-Commands)
- [Redirections](https://www.gnu.org/software/bash/manual/bash.html#Redirections)
- [Passing shell parameters as environment to commands](https://www.gnu.org/software/bash/manual/bash.html#Environment)
- [Pipelines](https://www.gnu.org/software/bash/manual/bash.html#Pipelines)
- [Conditional Execution](https://www.gnu.org/software/bash/manual/bash.html#Lists)
- [Command Grouping](https://www.gnu.org/software/bash/manual/html_node/Command-Grouping.html)
- [Command Substituion](https://www.gnu.org/software/bash/manual/html_node/Command-Substitution.html)
- [`if` command](https://www.gnu.org/software/bash/manual/html_node/Conditional-Constructs.html#index-if)
- [`while` & `until` loops](https://www.gnu.org/software/bash/manual/bash.html#index-until)
- [Shell Parameters](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameters.html)
- Running commands in background
- [Functions](https://www.gnu.org/software/bash/manual/html_node/Shell-Functions.html)

## Code Example
```shell

# simple command
echo "Hello World"

# redirections ('<<' here document is not yet supported)
echo foobar >output.txt >>append.txt &>all.txt
echo foobar 3>file.txt 2>&3 >&3-
cat <input.txt <<<"Hello World" 3<input.txt
cat 3<>io.txt 3>&-

# pipelines
cat file.json | jq '.filename' | grep "*hosts.txt"
! true

# conditional execution
command || command2 && command3

# shell parameters
key=value
key="value"
key='value'

echo "$key"

# passing shell parameters as environment variables
key=value kay2=value command

# subshells
(
  echo foo bar | cat
)

# groups
{
  echo foo bar | cat
}

# command substituion
echo "files: $(ls "$path")"

# `if` command
if true; then
    echo foo bar | cat
elif true; then
    echo baz boo
else
    echo bad
fi


# `while` command
while true; then
  echo foo bar | cat

  if true; then
      break
  else
    continue
  fi
fi


# `until` command
until true; then
  echo foo bar | cat

  if true; then
      break
  else
    continue
  fi
fi

# running commands in backgroud
command &

# functions
function foo() {
  echo foobar
}

```
