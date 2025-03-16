package runtime

import (
	"math"
	"strconv"
)

func ParseInt(value string) int {
	valueFloat, _ := strconv.ParseFloat(value, 64)
	return int(valueFloat) // if the value is 5.5, we should return 5 not 0.
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

func NegateInt(value int) int {
	if value == 0 {
		return 1
	}
	return 0
}

func IntPower(operand, pow int) int {
	return int(math.Pow(float64(operand), float64(pow)))
}

func CompareInt(x int, op string, y int) int {
	var result bool
	switch op {
	case "==":
		result = x == y
	case "!=":
		result = x != y
	case "<":
		result = x < y
	case ">":
		result = x > y
	case "<=":
		result = x <= y
	case ">=":
		result = x >= y
	case "&&":
		result = x != 0 && y != 0
	case "||":
		result = x != 0 || y != 0
	}

	if result {
		return 1
	}
	return 0
}

func ConditionalInt(a int, b int, c int) int {
	if a != 0 {
		return b
	}
	return c
}
