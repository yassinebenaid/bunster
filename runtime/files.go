package runtime

import "os"

func OpenFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
}
