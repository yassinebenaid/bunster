# Asynchronous commands

There is a special notation that when used, allows you to run a command, a pipeline or an entire list asyncronously in the background.
When you run a command in the background. The program will continue executing without waiting the command to exit.

Asynchronous commands run in a `sub-shell`. And read input from the `/dev/null` unless explicitly redirected.

The format is simple:

```sh
command &
```

You simply put an ampersand `&` at the end of the statement.

You can run an entire pipeline in the background:

```sh
command1 | command2 | command3 &
```

You can take this further, and run an entire list in the background:

```sh
command1 | command2 && command3 || command4 &
```

### Waiting for background commands

When you invoke a command in background. the program will continue without waiting for the command to finish. that's why they're called `background commands`.

For example:

```sh
sleep 20 &
echo Hello
```

The command `sleep` will be invoked and put in background. Then the program will continue to run `echo`. In this particular case. the program will exit immediatly after writing `Hello` to `stdout` because the command `sleep` will sleep for `20` seconds. And because the program will not wait for it to finish. The program will move on to `echo` and exit.

If you want to wait for background commands to finish, you can use the `wait` keyword:

```sh
sleep 20 &
echo Hello

wait
```

The `wait` keyword waits for all commands in background to finish. This means. the above example will indeed wait for the `sleep` command to finish before it exits.

### Exit Code

When you run a command in background. the exit code is always set to zero `0`.
