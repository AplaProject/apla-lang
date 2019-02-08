package compiler

import (
	"fmt"
	"math"

	"github.com/AplaProject/apla-lang/parser"
	rt "github.com/AplaProject/apla-lang/runtime"
)

type compiler struct {
	Contract  *Contract
	Blocks    []*parser.Node
	NameSpace *map[string]uint32
}

// VarInfo describes a variable
type VarInfo struct {
	Index uint16
	Type  uint16
}

// Contract contains information about the contract
type Contract struct {
	ID   int64 // External id
	Name string
	Code []rt.Bcode
	Vars map[string]VarInfo
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

func nodeToCode(node *parser.Node, cmpl *compiler) error {
	var err error
	if node == nil {
		return nil
	}
	switch node.Type {
	case parser.TBlock:
		//		startSize := len(cmpl.Contract.Code)
		varsCount := uint16(len(cmpl.Contract.Vars))
		cmpl.Blocks = append(cmpl.Blocks, node)
		for _, child := range node.Value.(*parser.NBlock).Statements {
			if err = nodeToCode(child, cmpl); err != nil {
				return err
			}
		}
		cmpl.Blocks = cmpl.Blocks[:len(cmpl.Blocks)-1]
		if uint16(len(cmpl.Contract.Vars)) != varsCount {
			cmpl.Append(rt.DELVARS, rt.Bcode(varsCount))
		}
		// Remove vars
		for key, vinfo := range cmpl.Contract.Vars {
			if vinfo.Index >= varsCount {
				delete(cmpl.Contract.Vars, key)
			}
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
		cmpl.Append(rt.RETURN, rt.Bcode(vtype))
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
		if err = nodeToCode(nBinary.Right, cmpl); err != nil {
			return err
		}
		code, result := cmpl.findBinary(nBinary)
		if code == rt.NOP {
			return cmpl.ErrorOperator(node)
		}
		cmpl.Append(code)
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
		default:
			return cmpl.ErrorParam(node, errType, node.Value)
		}
	case parser.TVars:
		types := make([]rt.Bcode, len(node.Value.(*parser.NVars).Vars))
		for i, v := range node.Value.(*parser.NVars).Vars {
			if _, ok := cmpl.Contract.Vars[v.Name]; ok {
				return cmpl.ErrorParam(node, errVarExists, v.Name)
			}
			types[i] = rt.Bcode(v.Type)
			cmpl.Contract.Vars[v.Name] = VarInfo{
				Index: uint16(len(cmpl.Contract.Vars)),
				Type:  uint16(v.Type),
			}
		}
		cmpl.Append(rt.INITVARS, rt.Bcode(len(types)))
		cmpl.Append(types...)
	case parser.TGetVar:
		name := node.Value.(*parser.NVarValue).Name
		if vinfo, ok := cmpl.Contract.Vars[name]; !ok {
			return cmpl.ErrorParam(node, errVarUnknown, name)
		} else {
			cmpl.Append(rt.GETVAR, rt.Bcode(vinfo.Index))
			node.Result = uint32(vinfo.Type)
		}
	case parser.TSetVar:
		name := node.Value.(*parser.NVarValue).Name
		if vinfo, ok := cmpl.Contract.Vars[name]; !ok {
			return cmpl.ErrorParam(node, errVarUnknown, name)
		} else {
			cmpl.Append(rt.SETVAR, rt.Bcode(vinfo.Index))
			node.Result = uint32(vinfo.Type)
		}
	case parser.TWhile:
		nWhile := node.Value.(*parser.NWhile)
		sizeCode := len(cmpl.Contract.Code)
		if err = nodeToCode(nWhile.Cond, cmpl); err != nil {
			return err
		}
		if nWhile.Cond.Result != parser.VBool {
			return cmpl.ErrorParam(nWhile.Cond, errCond, Type2Str(nWhile.Cond.Result))
		}
		sizeCond := len(cmpl.Contract.Code)
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
		if err = nodeToCode(nIf.Cond, cmpl); err != nil {
			return err
		}
		if nIf.Cond.Result != parser.VBool {
			return cmpl.ErrorParam(nIf.Cond, errCond, Type2Str(nIf.Cond.Result))
		}
		sizeCond := len(cmpl.Contract.Code)
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
				if err = nodeToCode(child.Cond, cmpl); err != nil {
					return err
				}
				if child.Cond.Result != parser.VBool {
					return cmpl.ErrorParam(child.Cond, errCond, Type2Str(child.Cond.Result))
				}
				sizeCond = len(cmpl.Contract.Code)
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
	default:
		fmt.Println(`Ooops`)
		return cmpl.Error(node, errNodeType)
	}
	return nil
}

// Compile compiles contract
func Compile(input string, nameSpace *map[string]uint32) (*Contract, error) {
	var root *parser.Node

	if len(*nameSpace) == 0 {
		initNameSpace(nameSpace)
	}
	cmpl := &compiler{
		Contract: &Contract{
			Code: make([]rt.Bcode, 0, 64),
			Vars: make(map[string]VarInfo),
		},
		NameSpace: nameSpace,
	}

	root, err := parser.Parser(input)
	if err != nil {
		return nil, err
	}
	if err = nodeToCode(root, cmpl); err != nil {
		return nil, err
	}
	return cmpl.Contract, nil
}
