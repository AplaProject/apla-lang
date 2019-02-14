package runtime

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
	ASSIGNINT    // vars = int
	ASSIGNADDINT // vars += int
	ASSIGNSUBINT // vars -= int
	ASSIGNMULINT // vars *= int
	ASSIGNDIVINT // vars /= int
	ASSIGNMODINT // vars %= int
	CALLFUNC     // + uint16 call contract function
	GETPARAMS    // + uint16 count of parameters
	RETURN       // return from contract + int16 (type)
	RETFUNC      // return from function
	SIGNINT      // unary minus int
	NOT          // unary logical not
	ADDSTR       // str+str
	PUSH64       // + int64
	DATA         // +uint16 size of data + data
)

// Runtime is a runtime structure
type Runtime struct {
	Vars []int64
}

// NewRuntime creates a new runtime
func NewRuntime() *Runtime {
	return &Runtime{
		Vars: make([]int64, 0, 100),
	}
}
