package ir

import "fmt"

type VarLength struct {
	Name string
}

func (d VarLength) togo() string {
	return fmt.Sprintf("runtime.FormatInt(len(shell.ReadVar(%q)))", d.Name)
}
