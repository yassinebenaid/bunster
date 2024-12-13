package ir

import (
	"fmt"
)

type NewPipe struct {
	Writer, Reader string
}

func (p NewPipe) togo() string {
	return fmt.Sprintf(
		`%s, %s, err := runtime.NewPipe()
		if err != nil {
			shell.HandleError("", err)
			return
		}
		`, p.Reader, p.Writer)
}
