# Quick Start

Have you ever wished your shell scripts could be faster, more portable, and secure ?
**Bunster** brings this to life by transforming your shell scripts into efficient,
standalone binaries that are easy to distribute and deploy across platforms _(only unix is supported at the moment)_.

Unlike other tools, **Bunster** doesn’t just wrap your scripts in a binary—it compiles them down to efficient native machine code,
leveraging the powerful Go toolchain. This ensures performance, portability, and robustness.

Technically speaking, **Bunster** is not a complete compiler, But rather a **Transplier** that generates **GO** code out of your scripts.
Then, opionally uses the **Go Toolchain** to compile the code to an executable program.

**Bunster** targets `bash` scripts in particular. The current syntax and features are all inherited from `bash`.
other shells support will be added soon.
