package runtime

import (
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
	PUSH64       // + int64
	DATA         // +uint16 size of data + data
)

// VarInfo describes a variable
type VarInfo struct {
	Index uint16
	Type  uint16
}

// FuncInfo describes a function
type FuncInfo struct {
	Offset int
	Result int
	Name   string
	Params []parser.NVar
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
}

// NewRuntime creates a new runtime
func NewRuntime(Contracts *[]*Contract) *Runtime {
	return &Runtime{
		//		Vars:      make([]int64, 0, 100),
		Contracts: Contracts,
	}
}
