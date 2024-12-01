package runtime

import "os"

func OpenFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR, 0644)
}

func OpenOrCreateFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
}

func OpenOrCreateFileForAppending(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}
