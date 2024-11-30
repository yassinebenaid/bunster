package runtime

import (
	"fmt"
	"os"
	"os/exec"
)

func HandleCommandRunError(err error) {
	if e, ok := err.(*exec.Error); ok {
		fmt.Fprintf(os.Stderr, "failed to recognize command %q, %v", e.Name, e.Err)
	}

	os.Exit(1)
}
