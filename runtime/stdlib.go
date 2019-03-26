package runtime

import (
	"fmt"
	"sort"
	"strconv"

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
		{KeysMap, 1, `Keys`, []uint32{parser.VMap}, (parser.VStr << 4) | parser.VArr}, // Keys(map) arr
		{LenArr, 1, `Len`, []uint32{parser.VArr}, parser.VInt},                        // Len(arr) int
		{LenMap, 1, `Len`, []uint32{parser.VMap}, parser.VInt},                        // Len(map) int
		{LenStr, 1, `Len`, []uint32{parser.VStr}, parser.VInt},                        // Len(str) int
		{StrInt, 1, `str`, []uint32{parser.VInt}, parser.VStr},                        // str(int) str
		{IntStr, 1, `int`, []uint32{parser.VStr}, parser.VInt},                        // int(str) int
	}
)

// KeysMap returns the array of map keys
func KeysMap(rt *Runtime, i int64) int64 {

	out := make([]string, len(rt.Objects[i].(map[string]int64)))
	var j int64
	for key := range rt.Objects[i].(map[string]int64) {
		out[j] = key
		j++
	}
	sort.Strings(out)
	ret := make([]int64, len(out))
	for i, val := range out {
		rt.Strings = append(rt.Strings, val)
		ret[i] = int64(len(rt.Strings) - 1)
	}
	rt.Objects = append(rt.Objects, ret)
	return int64(len(rt.Objects) - 1)
}

// LenArr returns the length of the array
func LenArr(rt *Runtime, i int64) int64 {
	return int64(len(rt.Objects[i].([]int64)))
}

// LenMap returns the length of the map
func LenMap(rt *Runtime, i int64) int64 {
	return int64(len(rt.Objects[i].(map[string]int64)))
}

// LenStr returns the length of the string
func LenStr(rt *Runtime, i int64) int64 {
	return int64(len(rt.Strings[i]))
}

// IntStr converts a string to the integer number
func IntStr(rt *Runtime, i int64) (int64, error) {
	val, err := strconv.ParseInt(rt.Strings[i], 0, 64)
	if err != nil {
		err = fmt.Errorf(errStr2Int, rt.Strings[i])
	}
	return int64(val), err
}

// StrInt converts the integer number to the string
func StrInt(rt *Runtime, i int64) int64 {
	rt.Strings = append(rt.Strings, fmt.Sprint(i))
	return int64(len(rt.Strings) - 1)
}
