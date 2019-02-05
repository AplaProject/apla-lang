package compiler

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
	rt "github.com/AplaProject/apla-lang/runtime"
)

var (
	operators = [][]uint32{
		// BCode, Result, Operator, Type of parameters...
		{rt.SIGNINT, parser.VInt, parser.SUB, parser.VInt}, // -int
		{rt.NOT, parser.VBool, parser.NOT, parser.VBool},   // !bool
	}
)

func (cmpl *compiler) findUnary(unary *parser.NUnary) (rt.Bcode, uint32) {
	key := fmt.Sprintf("#%d#%d", unary.Oper, unary.Operand.Result)
	if v, ok := (*cmpl.NameSpace)[key]; ok {
		return rt.Bcode(v & 0xffff), v >> 24
	}
	return rt.NOP, 0
}

func initNameSpace(nameSpace *map[string]uint32) {
	for _, oper := range operators {
		var key string
		for i := 2; i < len(oper); i++ {
			key += fmt.Sprintf(`#%d`, oper[i])
		}
		(*nameSpace)[key] = oper[0] | (oper[1] << 24)
	}
}

// Type2Str return a name of the type
func Type2Str(vtype uint32) (ret string) {
	switch vtype {
	case parser.VInt:
		ret = `int`
	case parser.VBool:
		ret = `bool`
	default:
		ret = `unknown`
	}
	return
}
