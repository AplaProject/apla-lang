package runtime

import (
	"fmt"
	"strings"

	"github.com/AplaProject/apla-lang/parser"
)

const (
	NOP          = iota
	PUSH16       // + int16
	PUSH32       // + int32
	PUSHSTR      // + uint64 + uint16  offset + size in Data
	INITVARS     // + uint16 (count) + ... uint16 types
	DELVARS      // + uint16 (new count)
	ADDINT       // int+int
	SUBINT       // int-int
	MULINT       // int*int
	DIVINT       // int/int
	MODINT       // int%int
	EQINT        // int == int
	NEINT        // int != int
	LTINT        // int < int
	LEINT        // int <= int
	GTINT        // int > int
	GEINT        // int >= int
	AND          // bool && bool
	OR           // bool || bool
	DUP          // duplicate top of stack
	GETVAR       // + uint16
	SETVAR       // + uint16
	JMP          // + int16   jump with the offset with clearing stack
	JMPREL       // + int16   jump with the offset
	JZE          // + int16   jump if top equals zero with the offset
	JNZ          // + int16   jump if top does not equal zero with the offset
	ASSIGNINT    // vars = int / vars = str
	ASSIGNADDINT // vars += int
	ASSIGNSUBINT // vars -= int
	ASSIGNMULINT // vars *= int
	ASSIGNDIVINT // vars /= int
	ASSIGNMODINT // vars %= int
	CALLFUNC     // + uint16 call contract function
	EMBEDFUNC    // + uint16 call embedded function
	CALLCONTRACT // + uint16 call contract
	LOADPARS     // load contract parameters
	PARCONTRACT  // + uint16 index of parameter
	GETPARAMS    // + uint16 count of parameters
	RETURN       // return from contract + int16 (type)
	RETFUNC      // return from function
	SIGNINT      // unary minus int
	NOT          // unary logical not
	ADDSTR       // str+str
	ASSIGNADDSTR // vars += str
	APPENDARR    // arr += item
	GETINDEX     // var[]
	SETINDEX     // var[]
	PUSH64       // + int64
	DATA         // +uint16 size of data + data
)

// VarInfo describes a variable
type VarInfo struct {
	Index uint16
	Type  uint16
}

type Var struct {
	Type int64
	Name string
}

// FuncInfo describes a function
type FuncInfo struct {
	Offset int
	Result int64
	Name   string
	Params []Var
}

// Contract contains information about the contract
type Contract struct {
	ID     int64 // External id
	Name   string
	Code   []Bcode
	Vars   map[string]VarInfo
	Funcs  []*FuncInfo
	Params map[string]VarInfo
}

// Runtime is a runtime structure
type Runtime struct {
	//	Vars      []int64
	Contracts *[]*Contract
	Strings   []string
	Objects   []interface{}
}

// NewRuntime creates a new runtime
func NewRuntime(Contracts *[]*Contract) *Runtime {
	return &Runtime{
		//		Vars:      make([]int64, 0, 100),
		Contracts: Contracts,
		Strings:   make([]string, 0, 8),
		Objects:   make([]interface{}, 0, 8),
	}
}

func print(rt *Runtime, val int64, vtype int64) string {
	var result string
	switch vtype & 0xf {
	case parser.VVoid: // skip result
	case parser.VStr:
		result = rt.Strings[val]
	case parser.VInt:
		result = fmt.Sprint(val)
	case parser.VBool:
		if val == 0 {
			result = `false`
		} else {
			result = `true`
		}
	case parser.VArr:
		items := make([]string, len(rt.Objects[val].([]int64)))
		result = `[`
		for i, item := range rt.Objects[val].([]int64) {
			items[i] = print(rt, item, vtype>>4)
		}
		result += strings.Join(items, ` `) + `]`
	default:
		result = fmt.Sprint(val)
	}
	return result
}
