package compiler

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
	rt "github.com/AplaProject/apla-lang/runtime"
)

const (
	EMBEDDED = 0x1000
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
		{rt.ADDSTR, parser.VStr, parser.ADD, parser.VStr, parser.VStr},               // str+str
		{rt.ASSIGNINT, parser.VVoid, parser.ASSIGN, parser.VStr, parser.VStr},        // str = str
		{rt.ASSIGNADDSTR, parser.VVoid, parser.ADD_ASSIGN, parser.VStr, parser.VStr}, // str += str
	}
)

func parseType(intype uint32) (outtype, subtype uint32) {
	if intype > 0xf {
		subtype = intype >> 4
		outtype = intype & 0xf
	}
	return
}

func (cmpl *compiler) findBinary(binary *parser.NBinary) (rt.Bcode, uint32) {
	key := fmt.Sprintf("#%d#%d#%d", binary.Oper, binary.Left.Result, binary.Right.Result)
	if v, ok := (*cmpl.NameSpace)[key]; ok {
		return rt.Bcode(v & 0xffff), v >> 24
	}
	if binary.Oper == parser.ADD_ASSIGN {
		outtype, subtype := parseType(binary.Left.Result)
		if outtype&0xf == parser.VArr && subtype == binary.Right.Result {
			return rt.APPENDARR, parser.VVoid
		}
	}
	if binary.Oper == parser.ASSIGN && (binary.Left.Result&0xf == parser.VArr) {
		if binary.Left.Result == binary.Right.Result {
			return rt.ASSIGNINT, parser.VVoid
		}
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

func getFuncKey(nfunc *rt.FuncInfo) string {
	ret := fmt.Sprintf("$%s", nfunc.Name)
	for _, par := range nfunc.Params {
		ret += fmt.Sprintf(`$%d`, par.Type)
	}
	return ret
}

func (cmpl *compiler) findFunc(nfunc *rt.FuncInfo) (rt.Bcode, uint32) {
	key := getFuncKey(nfunc)
	if v, ok := (*cmpl.NameSpace)[key]; ok {
		return rt.Bcode(v & 0xffff), v >> 24
	}
	return rt.NOP, 0
}

func (cmpl *compiler) findCallFunc(nfunc *parser.NCallFunc) (rt.Bcode, uint32) {
	key := fmt.Sprintf("$%s", nfunc.Name)
	softkey := key
	if nfunc.Params != nil {
		for _, par := range nfunc.Params.Value.(*parser.NParams).Expr {
			parkey := fmt.Sprintf(`$%d`, par.Result)
			if par.Result > 0xf {
				outtype, _ := parseType(par.Result)
				softkey += fmt.Sprintf(`$%d`, outtype)
			} else {
				softkey += parkey
			}
			key += parkey
		}
	}
	if v, ok := (*cmpl.NameSpace)[key]; ok {
		return rt.Bcode(v & 0xffff), v >> 24
	} else if len(softkey) > 0 {
		if v, ok := (*cmpl.NameSpace)[softkey]; ok {
			return rt.Bcode(v & 0xffff), v >> 24
		}
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

	for i, eFunc := range rt.StdLib {
		key := fmt.Sprintf(`$%s`, eFunc.Name)
		for _, par := range eFunc.PTypes {
			key += fmt.Sprintf(`$%d`, par)
		}
		(*nameSpace)[key] = uint32(i+EMBEDDED) | (eFunc.Result << 24)
	}
}

// Type2Str return a name of the type
func Type2Str(vtype uint32) (ret string) {
main:
	for i := 0; i < 4; i++ {
		itype := vtype & 0xf
		if i > 0 && itype != parser.VVoid {
			ret += `.`
		}
		switch itype {
		case parser.VInt:
			ret += `int`
		case parser.VBool:
			ret += `bool`
		case parser.VStr:
			ret += `str`
		case parser.VArr:
			ret += `arr`
		default:
			break main
		}
		vtype >>= 4
	}
	if len(ret) == 0 {
		ret = `unknown`
	}
	return
}
