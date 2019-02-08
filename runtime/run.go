package runtime

import (
	"fmt"
	"unsafe"

	"github.com/AplaProject/apla-lang/parser"
)

type Bcode uint16

const (
	errDivZero = `integer divide by zero`
	errCommand = `unknown command %d`
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
		case INITVARS:
			top = 0
			count := int64(code[i+1])
			for iVar := int64(0); iVar < count; iVar++ {
				rt.Vars = append(rt.Vars, 0)
			}
			i += count + 1
		case DELVARS:
			i++
			count := int64(code[i])
			rt.Vars = rt.Vars[:count]
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
		case MODINT:
			top--
			if stack[top+1] == 0 {
				return ``, gas, fmt.Errorf(errDivZero)
			}
			stack[top] %= stack[top+1]
		case EQINT:
			var b int64
			top--
			if stack[top] == stack[top+1] {
				b = 1
			}
			stack[top] = b
		case NEINT:
			var b int64
			top--
			if stack[top] != stack[top+1] {
				b = 1
			}
			stack[top] = b
		case LTINT:
			var b int64
			top--
			if stack[top] < stack[top+1] {
				b = 1
			}
			stack[top] = b
		case LEINT:
			var b int64
			top--
			if stack[top] <= stack[top+1] {
				b = 1
			}
			stack[top] = b
		case GTINT:
			var b int64
			top--
			if stack[top] > stack[top+1] {
				b = 1
			}
			stack[top] = b
		case GEINT:
			var b int64
			top--
			if stack[top] >= stack[top+1] {
				b = 1
			}
			stack[top] = b
		case AND:
			var b int64
			top--
			if stack[top] == 1 && stack[top+1] == 1 {
				b = 1
			}
			stack[top] = b
		case OR:
			var b int64
			top--
			if stack[top] == 1 || stack[top+1] == 1 {
				b = 1
			}
			stack[top] = b
		case GETVAR:
			i++
			top++
			stack[top] = rt.Vars[code[i]]
		case SETVAR:
			i++
			top++
			stack[top] = int64(uintptr(unsafe.Pointer(&rt.Vars[code[i]])))
		case JMP:
			i += int64(int16(code[i+1]))
			top = 0
			continue
		case JZE:
			if stack[top] == 0 {
				i += int64(int16(code[i+1]))
				continue
			}
			i++
		case JNZ:
			if stack[top] != 0 {
				i += int64(int16(code[i+1]))
				continue
			}
			i++
		case ASSIGNINT:
			*(*int64)(unsafe.Pointer(uintptr(stack[top-1]))) = stack[top]
			top -= 2
		case ASSIGNADDINT:
			*(*int64)(unsafe.Pointer(uintptr(stack[top-1]))) += stack[top]
			top -= 2
		case ASSIGNSUBINT:
			*(*int64)(unsafe.Pointer(uintptr(stack[top-1]))) -= stack[top]
			top -= 2
		case ASSIGNMULINT:
			*(*int64)(unsafe.Pointer(uintptr(stack[top-1]))) *= stack[top]
			top -= 2
		case ASSIGNDIVINT:
			if stack[top] == 0 {
				return ``, gas, fmt.Errorf(errDivZero)
			}
			*(*int64)(unsafe.Pointer(uintptr(stack[top-1]))) /= stack[top]
			top -= 2
		case ASSIGNMODINT:
			if stack[top] == 0 {
				return ``, gas, fmt.Errorf(errDivZero)
			}
			*(*int64)(unsafe.Pointer(uintptr(stack[top-1]))) %= stack[top]
			top -= 2
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
		default:
			return ``, gas, fmt.Errorf(errCommand, code[i])
		}
		i++
	}
	return result, gas, nil
}
