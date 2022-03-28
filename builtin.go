package pointable

import (
	"strings"
)

// Builtin represents the built-in types in Go.
type Builtin struct {
	Type string
}

// Symbol returns the type name as it is.
func (b Builtin) Symbol() string {
	return b.Type
}

// FuncName converts the type name into upper-case and return it.
func (b Builtin) FuncName() string {
	return strings.Title(b.Type)
}

// BuiltinData provides information of built-in types in Go.
// For more, see: https://go.dev/ref/spec#Types
func BuiltinData() []Type {
	res := make([]Type, 0, len(builtinTypes))

	for _, t := range builtinTypes {
		res = append(res, &Builtin{Type: t})
	}

	return res
}

var builtinTypes = []string{
	// Boolean types
	"bool",

	// Numeric types
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"int8",
	"int16",
	"int32",
	"int64",
	"float32",
	"float64",
	"complex64",
	"complex128",
	"byte",
	"rune",
	"uint",
	"int",

	// String types
	"string",
}
