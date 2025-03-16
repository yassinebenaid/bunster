package runtime

import "strconv"

func ParseInt(value string) int {
	valueInt, _ := strconv.Atoi(value)
	return valueInt
}

func FormatInt(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func VarIncrement(shell *Shell, name string, value int, post bool) int {
	v := shell.ReadVar(name)
	valueInt, _ := strconv.Atoi(v)
	shell.SetVar(name, strconv.FormatInt(int64(valueInt+value), 10))
	if post {
		return valueInt
	}

	return valueInt + value
}
