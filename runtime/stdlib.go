package runtime

import (
	"unsafe"

	"github.com/AplaProject/apla-lang/parser"
)

type EmbedFunc struct {
	Func   interface{}
	Params int64 // number of parameters from stack
	Name   string
	PTypes []uint32
	Result uint32
}

var (
	StdLib = []EmbedFunc{
		{LenStr, 1, `Len`, []uint32{parser.VStr}, parser.VInt}, // Len(str) int
	}
)

func LenStr(ptr int64) int64 {
	return int64(len(*(*string)(unsafe.Pointer(uintptr(ptr)))))
}
