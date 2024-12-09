<div align="center">
    <img width="200" src="./docs/public/logo.png"/>
</div>

<div align="center">

# Bunster

</div>

Have you ever wished your shell scripts could be faster, more portable, and secure ? **Bunster** brings this to life by transforming your shell scripts into efficient, standalone binaries that are easy to distribute and deploy across platforms _(only unix is supported at the moment)_.

Unlike other tools, **Bunster** doesn’t just wrap your scripts in a binary—it compiles them down to efficient native machine code, leveraging the powerful Go toolchain. This ensures performance, portability, and robustness, making **Bunster** a unique solution for modern developers.

Technically speaking, **Bunster** is not a complete compiler, But rather a **Transplier** that generates **GO** code out of your scripts. Then, opionally uses the **Go Toolchain** to compile the code to an executable program.

**Bunster** targets `bash` scripts in particular. The current syntax and features are all inherited from `bash`. other shells support will be added soon.  

### Motivation
- **Different Shells support**: Bunster currently aims to be compatible `Bash` as a starting move. then other popular shells in future.
- **Security**: as you may guess, humans cannot read machine code, so why not to compile your scripts to such format.
- **Modules**: something shell scripts lack is a module system, people want to share their code to be used by others, but the infrastructure doesn't allow them. Well, **Bunster** introduces a module system that allow you to publish your scripts as a modules consumed by other users.
- **Performance**: the shell (including bash, zsh ...etc) rely on forking to run your scripts, this means, if you have a script of 3 commands, the shell will have to fork it self 3 times to run each command. This allows the shell to play with file descriptors and other resouces freely. But adds a lot of performance overhead. **Bunster** runs your entire scripts in a single process. and uses [goroutines](https://go.dev/tour/concurrency/1) for background commands. **Bunster** even has its own file descripor system managed by it's runtime. this means less syscalls, thus, better performance. 
