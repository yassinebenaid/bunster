package runtime

import "strconv"

func VarIncrement(shell *Shell, name string, value int, post bool) int {
	v := shell.ReadVar(name)
	valueInt, _ := strconv.Atoi(v)
	shell.SetVar(name, strconv.FormatInt(int64(valueInt+value), 10))
	if post {
		return valueInt
	}

	return valueInt + value
}
