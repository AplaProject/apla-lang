package runtime

import (
	"fmt"
	"sort"
	"strconv"
	"unsafe"

	"github.com/shopspring/decimal"

	"github.com/AplaProject/apla-lang/parser"
)

type EmbedFunc struct {
	Gas    int64
	Func   interface{}
	Params int64 // number of parameters from stack
	Name   string
	PTypes []uint32
	Result uint32
}

var (
	StdLib = []EmbedFunc{
		{20, KeysMap, 1, `Keys`, []uint32{parser.VMap},
			(parser.VStr << 4) | parser.VArr}, // Keys(map) arr
		{5, LenArr, 1, `Len`, []uint32{parser.VArr}, parser.VInt},           // Len(arr) int
		{5, LenMap, 1, `Len`, []uint32{parser.VMap}, parser.VInt},           // Len(map) int
		{5, LenStr, 1, `Len`, []uint32{parser.VStr}, parser.VInt},           // Len(str) int
		{5, StrInt, 1, `str`, []uint32{parser.VInt}, parser.VStr},           // str(int) str
		{5, IntStr, 1, `int`, []uint32{parser.VStr}, parser.VInt},           // int(str) int
		{5, FloatInt, 1, `float`, []uint32{parser.VInt}, parser.VFloat},     // float(int) float
		{5, IntFloat, 1, `int`, []uint32{parser.VFloat}, parser.VInt},       // int(float) int
		{7, MoneyInt, 1, `money`, []uint32{parser.VInt}, parser.VMoney},     // money(int) money
		{7, MoneyFloat, 1, `money`, []uint32{parser.VFloat}, parser.VMoney}, // money(float) money
		{7, MoneyStr, 1, `money`, []uint32{parser.VStr}, parser.VMoney},     // money(str) money
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

// FloatInt converts an integer number to float
func FloatInt(rt *Runtime, i int64) int64 {
	f := float64(i)
	return *(*int64)(unsafe.Pointer(&f))
}

// IntFloat converts a float to the integer number
func IntFloat(rt *Runtime, i int64) int64 {
	return int64(*(*float64)(unsafe.Pointer(&i)))
}

// MoneyInt converts an integer number to money
func MoneyInt(rt *Runtime, i int64) int64 {
	rt.Objects = append(rt.Objects, decimal.New(i, 0))
	return int64(len(rt.Objects) - 1)
}

// MoneyFloat converts a float number to money
func MoneyFloat(rt *Runtime, i int64) int64 {
	rt.Objects = append(rt.Objects, decimal.NewFromFloat(*(*float64)(unsafe.Pointer(&i))).Floor())
	return int64(len(rt.Objects) - 1)
}

// MoneyStr converts a string to money
func MoneyStr(rt *Runtime, i int64) (int64, error) {
	d, err := decimal.NewFromString(rt.Strings[i])
	if err != nil {
		return 0, err
	}
	rt.Objects = append(rt.Objects, d.Floor())
	return int64(len(rt.Objects) - 1), nil
}
