package builtin

import (
	"github.com/yassinebenaid/bunster/runtime"
)

func Register(shell *runtime.Shell) {
	shell.RegisterFunction("true", True)
	shell.RegisterFunction("false", False)
	shell.RegisterFunction("loadenv", Loadenv)
}

func True(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	shell.ExitCode = 0
}

func False(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	shell.ExitCode = 1
}
