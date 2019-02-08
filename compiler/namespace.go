package compiler

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
	rt "github.com/AplaProject/apla-lang/runtime"
)

var (
	operators = [][]uint32{
		// BCode, Result, Operator, Type of parameters...
		{rt.SIGNINT, parser.VInt, parser.SUB, parser.VInt},                           // -int
		{rt.NOT, parser.VBool, parser.NOT, parser.VBool},                             // !bool
		{rt.ADDINT, parser.VInt, parser.ADD, parser.VInt, parser.VInt},               // int+int
		{rt.SUBINT, parser.VInt, parser.SUB, parser.VInt, parser.VInt},               // int-int
		{rt.MULINT, parser.VInt, parser.MUL, parser.VInt, parser.VInt},               // int*int
		{rt.DIVINT, parser.VInt, parser.DIV, parser.VInt, parser.VInt},               // int/int
		{rt.MODINT, parser.VInt, parser.MOD, parser.VInt, parser.VInt},               // int%int
		{rt.ASSIGNINT, parser.VVoid, parser.ASSIGN, parser.VInt, parser.VInt},        // int = int
		{rt.ASSIGNADDINT, parser.VVoid, parser.ADD_ASSIGN, parser.VInt, parser.VInt}, // int += int
		{rt.ASSIGNSUBINT, parser.VVoid, parser.SUB_ASSIGN, parser.VInt, parser.VInt}, // int -= int
		{rt.ASSIGNMULINT, parser.VVoid, parser.MUL_ASSIGN, parser.VInt, parser.VInt}, // int *= int
		{rt.ASSIGNDIVINT, parser.VVoid, parser.DIV_ASSIGN, parser.VInt, parser.VInt}, // int /= int
		{rt.ASSIGNMODINT, parser.VVoid, parser.MOD_ASSIGN, parser.VInt, parser.VInt}, // int %= int
		{rt.EQINT, parser.VBool, parser.EQ, parser.VInt, parser.VInt},                // int == int
		{rt.NEINT, parser.VBool, parser.NOT_EQ, parser.VInt, parser.VInt},            // int != int
		{rt.LTINT, parser.VBool, parser.LT, parser.VInt, parser.VInt},                // int < int
		{rt.LEINT, parser.VBool, parser.LTE, parser.VInt, parser.VInt},               // int <= int
		{rt.GTINT, parser.VBool, parser.GT, parser.VInt, parser.VInt},                // int > int
		{rt.GEINT, parser.VBool, parser.GTE, parser.VInt, parser.VInt},               // int >= int
		{rt.ASSIGNINT, parser.VVoid, parser.ASSIGN, parser.VBool, parser.VBool},      // bool = bool
		{rt.AND, parser.VBool, parser.AND, parser.VBool, parser.VBool},               // bool && bool
		{rt.OR, parser.VBool, parser.OR, parser.VBool, parser.VBool},                 // bool || bool
	}
)

func (cmpl *compiler) findBinary(binary *parser.NBinary) (rt.Bcode, uint32) {
	key := fmt.Sprintf("#%d#%d#%d", binary.Oper, binary.Left.Result, binary.Right.Result)
	if v, ok := (*cmpl.NameSpace)[key]; ok {
		return rt.Bcode(v & 0xffff), v >> 24
	}
	return rt.NOP, 0
}

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
