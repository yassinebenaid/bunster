Test: Simple commands
------------

--------input--------
# Nothing but this comment
# we need to make sure empty shell scripts produce empty go programs


--------output--------
package main

import (
    "os/exec"

    "bunster-build/runtime"
)

func Main(shell *runtime.Shell) {
}
