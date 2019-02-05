package runtime

const (
	NOP     = iota
	PUSH16  // + int16
	PUSH32  // + int32
	RETURN  // return from contract or function + int16 (type)
	SIGNINT // unary minus int
)

// Runtime is a runtime structure
type Runtime struct {
}

// NewRuntime creates a new runtime
func NewRuntime() *Runtime {
	return &Runtime{}
}
