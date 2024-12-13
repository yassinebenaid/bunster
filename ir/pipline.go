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

type NewPipelineWaitgroup string

func (p NewPipelineWaitgroup) togo() string {
	return fmt.Sprintf("var %s runtime.PiplineWaitgroup\n", p)
}

type PushToPipelineWaitgroup struct {
	Command   string
	Waitgroup string
}

func (p PushToPipelineWaitgroup) togo() string {
	return fmt.Sprintf(
		`%s = append(%s, runtime.PiplineWaitgroupItem{
			Wait: %s.Wait,
		})`,
		p.Waitgroup, p.Waitgroup, p.Command,
	)
}

type WaitPipelineWaitgroup string

func (w WaitPipelineWaitgroup) togo() string {
	return fmt.Sprintf(
		`if err := %s.Wait(); err != nil {
			shell.HandleError("", err)
			return
		}
		//TODO: shell.ExitCode = %%s.ProcessState.ExitCode()
		`, w)
}
