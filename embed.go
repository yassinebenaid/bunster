package bunster

import "embed"

//go:embed runtime
var RuntimeFS embed.FS

//go:embed stubs/main.go.stub
var MainGoStub []byte

//go:embed stubs/go.mod.stub
var GoModStub []byte
