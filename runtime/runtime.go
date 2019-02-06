package runtime

const (
	NOP       = iota
	PUSH16    // + int16
	PUSH32    // + int32
	INITVARS  // + uint16 (count) + ... uint16 types
	DELVARS   // + uint16 (new count)
	ADDINT    // int+int
	SUBINT    // int-int
	MULINT    // int*int
	DIVINT    // int/int
	GETVAR    // + uint16
	SETVAR    // + uint16
	ASSIGNINT // vars = int
	RETURN    // return from contract or function + int16 (type)
	SIGNINT   // unary minus int
	NOT       // unary logical not
	PUSH64    // + int64
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
