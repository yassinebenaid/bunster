# Environment Files
We have a first-class support for environment files. Known as [.env files](https://dotenvx.com/docs/env-file). This feature allows you
to load environment variables from environment files and make them available as variables in the shell context.

This is possible through the `envload` builtin command.

## Quick Example
You get started, define your variables in a `.env` file.
_.env_:
```sh
KEY=value
```

Then use the `loadenv` builtin command:
_script.sh_:
```sh
loadenv

echo $KEY

# value
```

## Environment file format
We support the standard format of `.env` files. This includes simple keys, quoting, variable references and yaml format.

```sh
# comments are supported
KEY=value # inline comments are supported too

# quotes
KEY="foo bar baz"
KEY='foo bar baz'

# reference other keys
KEY=$KEY
KEY=${KEY}
KEY="$KEY"
KEY="${KEY}"
```

### YAML format
We also support the yaml format:
```yaml
# comments are supported
KEY: value # inline comments are supported too

# quotes
KEY: "foo bar baz"
KEY: 'foo bar baz'

# reference other keys
KEY: $KEY
KEY: ${KEY}
KEY: "$KEY"
KEY: "${KEY}"
```

## Command Usage
The command `loadenv` is very streightforward. When you do not pass any arguments. it will default to load `.env` file in current workin directory.

```sh
loadenv # loads .env in CWD
```

You can pass one or many file paths. In this case, the order of the files matters. Because files that come last override keys in files that come first.

```sh
loadenv file.env file2.yaml file3.env
```

### Exporting variables
By default, the variables are only available in the current shell context. They are not exported to child processes. For example:

_.env_:
```sh
NAME=bunster
```

In your script:

```sh
loadenv

bash -c 'echo env:$NAME'
echo var:$NAME
```

This will output:
```txt
env:
var:bunster
```

To export them, you can use the `-X` flag to export them:
```sh
loadenv -X

bash -c 'echo env:$NAME'
echo var:$NAME
```

This will output:
```txt
env:bunster
var:bunster
```
