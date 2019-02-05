package runtime

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
)

type Bcode int16

// Run executes a bytecode
func (rt *Runtime) Run(code []Bcode) (string, int64, error) {
	var (
		i, top, gas int64
		result      string
	)
	length := int64(len(code))
	stack := make([]int64, 100)
	// top the latest value

main:
	for i < length {
		gas++
		switch code[i] {
		case PUSH16:
			i++
			top++
			stack[top] = int64(code[i])
		case RETURN:
			switch code[i+1] {
			case parser.VVoid: // skip result
			case parser.VInt:
				result = fmt.Sprint(stack[top])
			default:
				fmt.Println(`OOOOPS`)
			}
			break main
		case SIGNINT:
			stack[top] = -stack[top]
		}
		i++
	}
	return result, gas, nil
}
