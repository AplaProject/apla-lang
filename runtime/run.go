package runtime

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
)

type Bcode uint16

const (
	errDivZero = `integer divide by zero`
)

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
		case PUSH32:
			i += 2
			top++
			stack[top] = int64((uint64(code[i-1]) << 16) | uint64(code[i]&0xffff))
		case ADDINT:
			top--
			stack[top] += stack[top+1]
		case SUBINT:
			top--
			stack[top] -= stack[top+1]
		case MULINT:
			top--
			stack[top] *= stack[top+1]
		case DIVINT:
			top--
			if stack[top+1] == 0 {
				return ``, gas, fmt.Errorf(errDivZero)
			}
			stack[top] /= stack[top+1]
		case RETURN:
			switch code[i+1] {
			case parser.VVoid: // skip result
			case parser.VInt:
				result = fmt.Sprint(stack[top])
			case parser.VBool:
				if stack[top] == 0 {
					result = `false`
				} else {
					result = `true`
				}
			default:
				result = fmt.Sprint(stack[top])
			}
			break main
		case SIGNINT:
			stack[top] = -stack[top]
		case NOT:
			if stack[top] == 0 {
				stack[top] = 1
			} else {
				stack[top] = 0
			}
		case PUSH64:
			i += 4
			top++
			stack[top] = int64((uint64(code[i-3]) << 48) | (uint64(code[i-2]) << 32) |
				(uint64(code[i-1]) << 16) | (uint64(code[i]) & 0xffff))
		}
		i++
	}
	return result, gas, nil
}
