package compiler

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
	rt "github.com/AplaProject/apla-lang/runtime"
)

type compiler struct {
	Contract  *Contract
	NameSpace *map[string]uint32
}

// Contract contains information about the contract
type Contract struct {
	ID   int64 // External id
	Name string
	Code []rt.Bcode
}

func nodeToCode(node *parser.Node, cmpl *compiler) error {
	var err error
	if node == nil {
		return nil
	}
	switch node.Type {
	case parser.TBlock:
		if err = nodeToCode(node.Value.(*parser.NBlock).Statements, cmpl); err != nil {
			return err
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
		cmpl.Contract.Code = append(cmpl.Contract.Code, []rt.Bcode{rt.RETURN, rt.Bcode(vtype)}...)
	case parser.TStatements:
		for _, child := range node.Value.(*parser.NStatements).List {
			if err = nodeToCode(child, cmpl); err != nil {
				return err
			}
		}
	case parser.TUnary:
		nUnary := node.Value.(*parser.NUnary)
		if err = nodeToCode(nUnary.Operand, cmpl); err != nil {
			return err
		}
		code, result := cmpl.findUnary(nUnary)
		if code == rt.NOP {
			return cmpl.ErrorOperator(node)
		}
		cmpl.Contract.Code = append(cmpl.Contract.Code, code)
		node.Result = result
	case parser.TValue:
		switch v := node.Value.(type) {
		case int:
			cmpl.Contract.Code = append(cmpl.Contract.Code, []rt.Bcode{rt.PUSH16, rt.Bcode(v)}...)
			node.Result = parser.VInt
		case bool:
			var bcode rt.Bcode
			if v {
				bcode = 1
			}
			cmpl.Contract.Code = append(cmpl.Contract.Code, []rt.Bcode{rt.PUSH16, bcode}...)
			node.Result = parser.VBool
		default:
			return cmpl.Error(node, fmt.Sprintf(errType, node.Value))
		}
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
