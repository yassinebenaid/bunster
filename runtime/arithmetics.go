package runtime

import "strconv"

func VarAdd(shell *Shell, name string, value int) int {
	v := shell.ReadVar(name)
	valueInt, _ := strconv.Atoi(v)
	shell.SetVar(name, strconv.FormatInt(int64(valueInt+value), 10))
	return valueInt + value
}
