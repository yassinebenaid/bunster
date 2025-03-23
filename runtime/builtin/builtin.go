package builtin

import (
	"github.com/yassinebenaid/bunster/runtime"
)

func Register(shell *runtime.Shell) {
	shell.RegisterBuiltin("true", True)
	shell.RegisterBuiltin("false", False)
	shell.RegisterBuiltin("loadenv", Loadenv)
	shell.RegisterBuiltin("embed", Embed)
	shell.RegisterBuiltin("shift", Shift)
}

func True(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	shell.ExitCode = 0
}

func False(shell *runtime.Shell, stdin, stdout, stderr runtime.Stream) {
	shell.ExitCode = 1
}
