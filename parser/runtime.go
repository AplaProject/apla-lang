package parser

import (
	"fmt"
)

type Runtime struct {
	Stack []int
}

const (
	SETVAR = iota + 1 // ind
	GETVAR            // ind
	ADDI32
	PUSH16 // value
	PUSH32 // value
	GTI32
	JNZ
	JMP
)

func Run(Code []uint16) (int, int) {
	var i, top, gas int

	stack := make([]int, 100)

	vars := make([]int, 32)
	length := len(Code)
	// top the latest value
	for i < length {
		//		fmt.Println(i, top, Code[i], vars[:5], stack[:top+1])
		gas++
		switch Code[i] {
		case SETVAR:
			i++
			vars[Code[i]] = stack[top]
			top--
		case GETVAR:
			i++
			top++
			stack[top] = vars[Code[i]]
		case ADDI32:
			top--
			stack[top] = stack[top] + stack[top+1]
		case PUSH16:
			i++
			top++
			stack[top] = int(Code[i])
		case PUSH32:
			i += 2
			top++
			stack[top] = int((int(Code[i-1]) << 16) | int(Code[i]))
		case GTI32:
			top--
			if stack[top] > stack[top+1] {
				stack[top] = 1
			} else {
				stack[top] = 0
			}
		case JNZ:
			i++
			top--
			if stack[top+1] != 0 {
				i = int(Code[i])
				continue
			}
		case JMP:
			i = int(Code[i+1])
			continue
		}
		i++
	}
	if vars[3] < 10000000 {
		fmt.Println(`OOOPS`, vars[3])
	}
	if len(stack) > 0 {
		return stack[top], gas
	}
	return 0, gas
}
