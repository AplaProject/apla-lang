package runtime

const (
	NOP     = iota
	PUSH16  // + int16
	PUSH32  // + int32
	ADDINT  // int+int
	SUBINT  // int-int
	MULINT  // int*int
	DIVINT  // int/int
	RETURN  // return from contract or function + int16 (type)
	SIGNINT // unary minus int
	NOT     // unary logical not
	PUSH64  // + int64
)

// Runtime is a runtime structure
type Runtime struct {
}

// NewRuntime creates a new runtime
func NewRuntime() *Runtime {
	return &Runtime{}
}
