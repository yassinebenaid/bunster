package runtime

import "strconv"

func NumberCompare(x, op, y string) bool {
	xv, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		return false
	}
	yv, err := strconv.ParseInt(y, 10, 64)
	if err != nil {
		return false
	}

	switch op {
	case "==":
		return xv == yv
	case "!=":
		return xv != yv
	case "<":
		return xv < yv
	case ">":
		return xv > yv
	case "<=":
		return xv <= yv
	case ">=":
		return xv >= yv
	default:
		panic("unsupported arithmetic operator: " + op)
	}
}
