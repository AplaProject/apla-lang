package compiler

import (
	"fmt"
	"math"
	"strings"

	"github.com/AplaProject/apla-lang/parser"
	rt "github.com/AplaProject/apla-lang/runtime"
)

type compiler struct {
	Contract  *rt.Contract
	Blocks    []*parser.Node
	Contracts *[]*rt.Contract
	NameSpace *map[string]uint32
	RetFunc   int
	InFunc    bool
	Data      []byte
}

func (cmpl *compiler) Append(codes ...rt.Bcode) {
	for _, code := range codes {
		cmpl.Contract.Code = append(cmpl.Contract.Code, code)
	}
}

func (cmpl *compiler) JumpOff(node *parser.Node, off int) (rt.Bcode, error) {
	if off < math.MinInt16 || off > math.MaxInt16 {
		return rt.NOP, cmpl.Error(node, errJump)
	}
	return rt.Bcode(off), nil
}

func (cmpl *compiler) ConditionCode(node *parser.Node) (before int, after int, err error) {
	before = len(cmpl.Contract.Code)
	if err = nodeToCode(node, cmpl); err != nil {
		return
	}
	if node.Result != parser.VBool {
		err = cmpl.ErrorParam(node, errCond, Type2Str(node.Result))
		return
	}
	after = len(cmpl.Contract.Code)
	return
}

func (cmpl *compiler) InitVars(node *parser.Node, vars []parser.NVar) error {
	if len(vars) == 0 {
		return nil
	}
	types := make([]rt.Bcode, len(vars))
	for i, v := range vars {
		if _, ok := cmpl.Contract.Vars[v.Name]; ok {
			return cmpl.ErrorParam(node, errVarExists, v.Name)
		}
		types[i] = rt.Bcode(v.Type)
		cmpl.Contract.Vars[v.Name] = rt.VarInfo{
			Index: uint16(len(cmpl.Contract.Vars)),
			Type:  uint16(v.Type),
		}
	}
	cmpl.Append(rt.INITVARS, rt.Bcode(len(types)))
	cmpl.Append(types...)
	return nil
}

func nodeToCode(node *parser.Node, cmpl *compiler) error {
	var (
		err                error
		vinfo              rt.VarInfo
		ok                 bool
		sizeCode, sizeCond int
	)
	if node == nil {
		return nil
	}
	switch node.Type {
	case parser.TBlock:
		varsCount := uint16(len(cmpl.Contract.Vars))
		funcsCount := len(cmpl.Contract.Funcs)
		cmpl.Blocks = append(cmpl.Blocks, node)
		pars := node.Value.(*parser.NBlock).Params
		if len(pars) > 0 {
			if err = cmpl.InitVars(node, pars); err != nil {
				return err
			}
			cmpl.Contract.Params = make(map[string]int)
			for _, ipar := range pars {
				cmpl.Contract.Params[ipar.Name] = ipar.Type
			}
		}
		for _, child := range node.Value.(*parser.NBlock).Statements {
			if err = nodeToCode(child, cmpl); err != nil {
				return err
			}
		}
		cmpl.Blocks = cmpl.Blocks[:len(cmpl.Blocks)-1]
		if uint16(len(cmpl.Contract.Vars)) != varsCount &&
			cmpl.Contract.Code[len(cmpl.Contract.Code)-1] != rt.RETFUNC {
			cmpl.Append(rt.DELVARS, rt.Bcode(varsCount))
		}
		// Remove vars
		for key, vinfo := range cmpl.Contract.Vars {
			if vinfo.Index >= varsCount {
				delete(cmpl.Contract.Vars, key)
			}
		}
		if funcsCount < len(cmpl.Contract.Funcs) {
			// Remove funcs
			for i := funcsCount; i < len(cmpl.Contract.Funcs); i++ {
				delete(*cmpl.NameSpace, getFuncKey(cmpl.Contract.Funcs[i]))
			}
			cmpl.Contract.Funcs = cmpl.Contract.Funcs[:funcsCount]
		}
	case parser.TContract:
		cmpl.Contract.Name = node.Value.(*parser.NContract).Name
		if err = nodeToCode(node.Value.(*parser.NContract).Block, cmpl); err != nil {
			return err
		}
	case parser.TReturn:
		var vtype uint32
		expr := node.Value.(*parser.NReturn).Expr
		if expr != nil {
			if err = nodeToCode(expr, cmpl); err != nil {
				return err
			}
			vtype = expr.Result
		}
		if cmpl.InFunc {
			if vtype != uint32(cmpl.RetFunc) {
				if cmpl.RetFunc == parser.VVoid {
					return cmpl.Error(node, errNotReturn)
				}
				if vtype == parser.VVoid {
					return cmpl.Error(node, errFuncReturn)
				}
				return cmpl.ErrorParam(node, errReturnType, Type2Str(uint32(cmpl.RetFunc)))
			}
			cmpl.Append(rt.RETFUNC)
		} else {
			cmpl.Append(rt.RETURN, rt.Bcode(vtype))
		}
	case parser.TBinary:
		nBinary := node.Value.(*parser.NBinary)
		if err = nodeToCode(nBinary.Left, cmpl); err != nil {
			return err
		}
		if nBinary.Left.Type == parser.TVars { // type varName =
			nBinary.Left = &parser.Node{
				Type: parser.TSetVar,
				Value: &parser.NVarValue{
					Name: nBinary.Left.Value.(*parser.NVars).Vars[0].Name,
				},
			}
			if err = nodeToCode(nBinary.Left, cmpl); err != nil {
				return err
			}
		}
		forJump := len(cmpl.Contract.Code)
		if err = nodeToCode(nBinary.Right, cmpl); err != nil {
			return err
		}
		code, result := cmpl.findBinary(nBinary)
		var jumpCmd rt.Bcode
		switch code {
		case rt.NOP:
			return cmpl.ErrorOperator(node)
		case rt.AND:
			jumpCmd = rt.Bcode(rt.JZE)
		case rt.OR:
			jumpCmd = rt.Bcode(rt.JNZ)
		}
		cmpl.Append(code)
		if jumpCmd != rt.NOP {
			cmpl.Contract.Code = append(cmpl.Contract.Code[:forJump],
				append([]rt.Bcode{rt.DUP, jumpCmd, rt.Bcode(len(cmpl.Contract.Code) - forJump + 2)},
					cmpl.Contract.Code[forJump:]...)...)
		}
		node.Result = result
	case parser.TUnary:
		nUnary := node.Value.(*parser.NUnary)
		if err = nodeToCode(nUnary.Operand, cmpl); err != nil {
			return err
		}
		code, result := cmpl.findUnary(nUnary)
		if code == rt.NOP {
			return cmpl.ErrorOperator(node)
		}
		cmpl.Append(code)
		node.Result = result
	case parser.TQuestion:
		nQuestion := node.Value.(*parser.NQuestion)
		_, sizeCond, err = cmpl.ConditionCode(nQuestion.Cond)
		if err != nil {
			return err
		}
		cmpl.Append(rt.JZE, 0)
		if err = nodeToCode(nQuestion.Left, cmpl); err != nil {
			return err
		}
		sizeCode = len(cmpl.Contract.Code)
		cmpl.Append(rt.JMPREL, 0)
		if err = nodeToCode(nQuestion.Right, cmpl); err != nil {
			return err
		}
		if nQuestion.Left.Result != nQuestion.Right.Result {
			return cmpl.Error(node, errQuestTypes)
		}
		var off rt.Bcode
		if off, err = cmpl.JumpOff(node, sizeCode-sizeCond+2); err != nil {
			return err
		}
		cmpl.Contract.Code[sizeCond+1] = off

		if off, err = cmpl.JumpOff(node, len(cmpl.Contract.Code)-sizeCode); err != nil {
			return err
		}
		cmpl.Contract.Code[sizeCode+1] = off
		node.Result = nQuestion.Left.Result
	case parser.TValue:
		switch v := node.Value.(type) {
		case int:
			if v <= math.MaxInt16 && v >= math.MinInt16 {
				cmpl.Append(rt.PUSH16, rt.Bcode(v))
			} else if v <= math.MaxInt32 && v >= math.MinInt32 {
				u32 := uint32(v)
				cmpl.Append(rt.PUSH32, rt.Bcode(u32>>16), rt.Bcode(u32&0xffff))
			} else {
				u64 := uint64(v)
				cmpl.Append(rt.PUSH64, rt.Bcode(u64>>48), rt.Bcode((u64>>32)&0xffff),
					rt.Bcode((u64>>16)&0xffff), rt.Bcode(u64&0xffff))
			}
			node.Result = parser.VInt
		case bool:
			var bcode rt.Bcode
			if v {
				bcode = 1
			}
			cmpl.Append(rt.PUSH16, bcode)
			node.Result = parser.VBool
		case string:
			if cmpl.Data == nil {
				cmpl.Data = make([]byte, 0, 1024)
			}
			cmpl.Append(rt.PUSHSTR, rt.Bcode(len(cmpl.Data)), rt.Bcode(len(v)))
			cmpl.Data = append(cmpl.Data, []byte(v)...)
		default:
			return cmpl.ErrorParam(node, errType, node.Value)
		}
	case parser.TVars:
		if err = cmpl.InitVars(node, node.Value.(*parser.NVars).Vars); err != nil {
			return err
		}
	case parser.TGetVar:
		name := node.Value.(*parser.NVarValue).Name
		if vinfo, ok = cmpl.Contract.Vars[name]; !ok {
			return cmpl.ErrorParam(node, errVarUnknown, name)
		}
		cmpl.Append(rt.GETVAR, rt.Bcode(vinfo.Index))
		node.Result = uint32(vinfo.Type)
	case parser.TSetVar:
		name := node.Value.(*parser.NVarValue).Name
		if vinfo, ok = cmpl.Contract.Vars[name]; !ok {
			return cmpl.ErrorParam(node, errVarUnknown, name)
		}
		cmpl.Append(rt.SETVAR, rt.Bcode(vinfo.Index))
		node.Result = uint32(vinfo.Type)
	case parser.TWhile:
		nWhile := node.Value.(*parser.NWhile)
		sizeCode, sizeCond, err = cmpl.ConditionCode(nWhile.Cond)
		if err != nil {
			return err
		}
		cmpl.Append(rt.JZE, 0)
		if err = nodeToCode(nWhile.Body, cmpl); err != nil {
			return err
		}
		var off rt.Bcode
		if off, err = cmpl.JumpOff(node, sizeCode-len(cmpl.Contract.Code)); err != nil {
			return err
		}
		cmpl.Append(rt.JMP, off)

		if off, err = cmpl.JumpOff(node, len(cmpl.Contract.Code)-sizeCond); err != nil {
			return err
		}
		cmpl.Contract.Code[sizeCond+1] = off
	case parser.TIf:
		ends := make([]int, 0, 16)
		nIf := node.Value.(*parser.NIf)
		_, sizeCond, err = cmpl.ConditionCode(nIf.Cond)
		if err != nil {
			return err
		}
		cmpl.Append(rt.JZE, 0)
		if err = nodeToCode(nIf.IfBody, cmpl); err != nil {
			return err
		}
		ends = append(ends, len(cmpl.Contract.Code))
		cmpl.Append(rt.JMP, 0)
		var off rt.Bcode
		if off, err = cmpl.JumpOff(node, len(cmpl.Contract.Code)-sizeCond); err != nil {
			return err
		}
		cmpl.Contract.Code[sizeCond+1] = off
		if nIf.ElifBody != nil {
			nElif := nIf.ElifBody.Value.(*parser.NElif)
			for _, child := range nElif.List {
				_, sizeCond, err = cmpl.ConditionCode(child.Cond)
				if err != nil {
					return err
				}
				cmpl.Append(rt.JZE, 0)
				if err = nodeToCode(child.Body, cmpl); err != nil {
					return err
				}
				ends = append(ends, len(cmpl.Contract.Code))
				cmpl.Append(rt.JMP, 0)
				if off, err = cmpl.JumpOff(node, len(cmpl.Contract.Code)-sizeCond); err != nil {
					return err
				}
				cmpl.Contract.Code[sizeCond+1] = off
			}
		}
		if nIf.ElseBody != nil {
			if err = nodeToCode(nIf.ElseBody, cmpl); err != nil {
				return err
			}
		}
		size := len(cmpl.Contract.Code)
		for _, end := range ends {
			if off, err = cmpl.JumpOff(node, size-end); err != nil {
				return err
			}
			cmpl.Contract.Code[end+1] = off
		}
	case parser.TFunc:
		var off rt.Bcode
		nFunc := node.Value.(*parser.NFunc)
		finfo := &rt.FuncInfo{
			Name:   nFunc.Name,
			Result: nFunc.Result,
			Params: nFunc.Params,
		}
		if cmpl.InFunc {
			return cmpl.Error(node, errFuncLevel)
		}
		cmpl.RetFunc = nFunc.Result
		if code, _ := cmpl.findFunc(finfo); code != rt.NOP {
			return cmpl.ErrorParam(node, errFuncExists, nFunc.Name)
		}
		start := len(cmpl.Contract.Code)
		cmpl.Append(rt.JMP, 0)
		finfo.Offset = start + 2
		cmpl.InFunc = true
		if err = cmpl.InitVars(node, nFunc.Params); err != nil {
			return err
		}
		if len(nFunc.Params) > 0 {
			cmpl.Append(rt.GETPARAMS, rt.Bcode(len(nFunc.Params)))
		}
		if err = nodeToCode(nFunc.Body, cmpl); err != nil {
			return err
		}
		cmpl.InFunc = false
		if cmpl.Contract.Code[len(cmpl.Contract.Code)-1] != rt.RETFUNC {
			if cmpl.RetFunc != parser.VVoid {
				return cmpl.Error(node, errFuncReturn)
			}
			cmpl.Append(rt.RETFUNC)
		}
		if off, err = cmpl.JumpOff(nFunc.Body, len(cmpl.Contract.Code)-start); err != nil {
			return err
		}
		cmpl.Contract.Code[start+1] = off
		cmpl.Contract.Funcs = append(cmpl.Contract.Funcs, finfo)
		(*cmpl.NameSpace)[getFuncKey(finfo)] = uint32(len(cmpl.Contract.Funcs) | (finfo.Result << 24))
	case parser.TCallFunc:
		nFunc := node.Value.(*parser.NCallFunc)
		if nFunc.Params != nil {
			for _, expr := range nFunc.Params.Value.(*parser.NParams).Expr {
				if err = nodeToCode(expr, cmpl); err != nil {
					return err
				}
			}
		}
		code, ftype := cmpl.findCallFunc(nFunc)
		if code == rt.NOP {
			pars := make([]string, 0, 10)
			if nFunc.Params != nil {
				for _, par := range nFunc.Params.Value.(*parser.NParams).Expr {
					pars = append(pars, Type2Str(uint32(par.Result)))
				}
			}
			return cmpl.ErrorParam(node, errFuncNotExists, fmt.Sprintf("%s(%s)", nFunc.Name,
				strings.Join(pars, `, `)))
		}
		node.Result = ftype
		if code >= EMBEDDED {
			cmpl.Append(rt.EMBEDFUNC, code-EMBEDDED)
		} else {
			var off rt.Bcode
			if off, err = cmpl.JumpOff(node, cmpl.Contract.Funcs[code-1].Offset-
				len(cmpl.Contract.Code)); err != nil {
				return err
			}
			cmpl.Append(rt.CALLFUNC, off)
		}
	case parser.TCallContract:
		nCallContract := node.Value.(*parser.NCallContract)
		ind, ok := (*cmpl.NameSpace)[nCallContract.Name]
		if !ok {
			return cmpl.ErrorParam(node, errContractNotExists, nCallContract.Name)
		}
		//		cnt := (*cmpl.Contracts)[ind]
		/*		code, ftype := cmpl.findCallFunc(nFunc)
				if code == rt.NOP {
					pars := make([]string, 0, 10)
					if nFunc.Params != nil {
						for _, par := range nFunc.Params.Value.(*parser.NParams).Expr {
							pars = append(pars, Type2Str(uint32(par.Result)))
						}
					}
					return cmpl.ErrorParam(node, errFuncNotExists, fmt.Sprintf("%s(%s)", nFunc.Name,
						strings.Join(pars, `, `)))
				}*/
		cmpl.Append(rt.CALLCONTRACT, rt.Bcode(ind))
		node.Result = parser.VStr
	default:
		fmt.Println(`Ooops`)
		return cmpl.Error(node, errNodeType)
	}
	return nil
}

// Compile compiles contract
func Compile(input string, nameSpace *map[string]uint32, contracts *[]*rt.Contract) (*rt.Contract, error) {
	var root *parser.Node

	if len(*nameSpace) == 0 {
		initNameSpace(nameSpace)
	}
	cmpl := &compiler{
		Contract: &rt.Contract{
			Code: make([]rt.Bcode, 0, 64),
			Vars: make(map[string]rt.VarInfo),
		},
		NameSpace: nameSpace,
		Contracts: contracts,
	}

	root, err := parser.Parser(input)
	if err != nil {
		return nil, err
	}
	defer func() {
		for i := 0; i < len(cmpl.Contract.Funcs); i++ {
			delete(*cmpl.NameSpace, getFuncKey(cmpl.Contract.Funcs[i]))
		}
	}()
	if err = nodeToCode(root, cmpl); err != nil {
		return nil, err
	}
	if len(cmpl.Data) > 0 {
		length := len(cmpl.Data)
		if length > 0xffff {
			return nil, cmpl.Error(root, errData)
		}
		if length&0x1 == 1 {
			cmpl.Data = append(cmpl.Data, 0)
			length++
		}
		length >>= 1
		data := make([]rt.Bcode, length+2)
		data[0] = rt.DATA
		data[1] = rt.Bcode(length)
		var off int
		for i := 0; i < length; i++ {
			data[i+2] = rt.Bcode(uint16(cmpl.Data[off])<<8 | uint16(cmpl.Data[off+1]))
			off += 2
		}
		cmpl.Contract.Code = append(data, cmpl.Contract.Code...)
	}
	return cmpl.Contract, nil
}
