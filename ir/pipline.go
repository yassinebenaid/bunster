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
			shell.HandleError(err)
			return
		}
		`, p.Reader, p.Writer)
}

type NewPipelineWaitgroup string

func (p NewPipelineWaitgroup) togo() string {
	return fmt.Sprintf("var %s []func() error\n", p)
}

type PushToPipelineWaitgroup struct {
	Command   string
	Waitgroup string
}

func (p PushToPipelineWaitgroup) togo() string {
	return fmt.Sprintf(
		`%s = append(%s, %s.Wait)`,
		p.Waitgroup, p.Waitgroup, p.Command,
	)
}

type WaitPipelineWaitgroup string

func (w WaitPipelineWaitgroup) togo() string {
	return fmt.Sprintf(
		`for i, wait := range %s {
			if err := wait(); err != nil {
				shell.HandleError(err)
			}
			if i < (len(%s) - 1){
				shell.ExitCode = 0
			}
		}
		`, w, w)
}
