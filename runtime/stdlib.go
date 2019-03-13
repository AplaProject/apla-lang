package runtime

import (
	"fmt"

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
		{LenArr, 1, `Len`, []uint32{parser.VArr}, parser.VInt}, // Len(arr) int
		{LenStr, 1, `Len`, []uint32{parser.VStr}, parser.VInt}, // Len(str) int
		{StrInt, 1, `str`, []uint32{parser.VInt}, parser.VStr}, // str(int) str
	}
)

// LenArr returns the length of the array
func LenArr(rt *Runtime, i int64) int64 {
	return int64(len(rt.Objects[i].([]int64)))
}

// LenStr returns the length of the string
func LenStr(rt *Runtime, i int64) int64 {
	return int64(len(rt.Strings[i]))
	//	return int64(len(*(*string)(unsafe.Pointer(uintptr(ptr)))))
}

// StrInt converts the integer number to the string
func StrInt(rt *Runtime, i int64) int64 {
	rt.Strings = append(rt.Strings, fmt.Sprint(i))
	return int64(len(rt.Strings) - 1)
}
