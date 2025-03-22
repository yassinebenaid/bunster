package bunster

import "embed"

//go:embed VERSION
var Version string

//go:embed runtime
var RuntimeFS embed.FS

//go:embed stubs/go.mod.stub
var Gomod []byte

//go:embed stubs/main.go.stub
var MainGo []byte
