package runtime

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/AplaProject/apla-lang/parser"
)

type Bcode uint16

const (
	errDivZero = `integer divide by zero`
	errCommand = `unknown command %d`
)

// Run executes a bytecode
func (rt *Runtime) Run(code []Bcode, params []int64) (string, int64, error) {
	var (
		i, top, gas, coff int64
		result            string
		data              []byte
	)
	length := int64(len(code))
	if length == 0 {
		return result, gas, nil
	}
	Vars := make([]int64, 0, 32)
	stack := make([]int64, 100)
	pars := make([]int64, 0, 32)
	calls := make([]int64, 1000)
	// top the latest value
	if code[0] == DATA {
		length := int64(uint64(code[1]))
		data = make([]byte, length<<1)
		length += 2
		var off int
		for i = 2; i < length; i++ {
			data[off] = byte(code[i] >> 8)
			data[off+1] = byte(code[i] & 0xff)
			off += 2
		}
	}
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
		case PUSHSTR:
			dstr := string(data[code[i+1] : code[i+1]+code[i+2]])
			top++
			stack[top] = int64(uintptr(unsafe.Pointer(&dstr)))
			i += 2
		case INITVARS:
			//			top = 0
			count := int64(code[i+1])
			for iVar := int64(0); iVar < count; iVar++ {
				var v int64
				switch code[i+2+iVar] {
				case parser.VStr:
					dstr := ``
					v = int64(uintptr(unsafe.Pointer(&dstr)))
				}
				Vars = append(Vars, v)
			}
			i += count + 1
		case DELVARS:
			i++
			count := int64(code[i])
			Vars = Vars[:count]
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
		case DUP:
			top++
			stack[top] = stack[top-1]
		case GETVAR:
			i++
			top++
			stack[top] = Vars[code[i]]
		case SETVAR:
			i++
			top++
			stack[top] = int64(uintptr(unsafe.Pointer(&Vars[code[i]])))
		case JMP:
			i += int64(int16(code[i+1]))
			top = 0
			continue
		case JMPREL:
			i += int64(int16(code[i+1]))
			continue
		case JZE:
			top--
			if stack[top+1] == 0 {
				i += int64(int16(code[i+1]))
				continue
			}
			i++
		case JNZ:
			top--
			if stack[top+1] != 0 {
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
		case CALLFUNC:
			calls[coff] = i + 2
			calls[coff+1] = int64(len(Vars))
			coff += 2
			i += int64(int16(code[i+1]))
			continue
		case EMBEDFUNC:
			i++
			eFunc := StdLib[code[i]]
			pars := make([]reflect.Value, eFunc.Params)
			/*			if Runtime {
						pars = append(pars, reflect.ValueOf(rt))
					}*/
			top -= eFunc.Params
			for k := int64(0); k < eFunc.Params; k++ {
				pars[k] = reflect.ValueOf(stack[top+k+1])
			}
			var result []reflect.Value
			result = reflect.ValueOf(eFunc.Func).Call(pars)
			if len(result) > 0 {
				last := result[len(result)-1].Interface()
				if last != nil {
					if _, isError := last.(error); isError {
						return ``, gas, result[len(result)-1].Interface().(error)
					}
				}
				top++
				stack[top] = result[0].Interface().(int64)
			}
		case CALLCONTRACT:
			i++
			top++
			result, cgas, cerr := rt.Run((*rt.Contracts)[code[i]].Code, pars)
			pars = pars[:0]
			gas -= cgas
			if cerr != nil {
				return ``, gas, cerr
			}
			stack[top] = int64(uintptr(unsafe.Pointer(&result)))
		case LOADPARS:
			for j := 0; j < (len(params) >> 1); j++ {
				Vars[params[j<<1]] = params[(j<<1)+1]
			}
		case PARCONTRACT:
			i++
			pars = append(pars, int64(code[i]), stack[top])
			top--
		case GETPARAMS:
			i++
			for k := 1; k <= int(code[i]); k++ {
				Vars[len(Vars)-k] = stack[top]
				top--
			}
		case RETURN:
			switch code[i+1] {
			case parser.VVoid: // skip result
			case parser.VStr:
				result = *(*string)(unsafe.Pointer(uintptr(stack[top])))
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
		case RETFUNC:
			Vars = Vars[:calls[coff-1]]
			coff -= 2
			i = calls[coff]
			continue
		case SIGNINT:
			stack[top] = -stack[top]
		case NOT:
			if stack[top] == 0 {
				stack[top] = 1
			} else {
				stack[top] = 0
			}
		case ADDSTR:
			top--
			dstr := *(*string)(unsafe.Pointer(uintptr(stack[top]))) +
				*(*string)(unsafe.Pointer(uintptr(stack[top+1])))
			stack[top] = int64(uintptr(unsafe.Pointer(&dstr)))
		case ASSIGNADDSTR:
			pvar := (*string)(unsafe.Pointer(uintptr(*(*int64)(unsafe.Pointer(uintptr(stack[top-1]))))))
			*pvar += *(*string)(unsafe.Pointer(uintptr(stack[top])))
			top -= 2
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
