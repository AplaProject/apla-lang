package compiler

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
	rt "github.com/AplaProject/apla-lang/runtime"
)

// Contract contains information about the contract
type Contract struct {
	ID   int64 // External id
	Name string
	Code []rt.Bcode
}

func nodeToCode(node *parser.Node, cnt *Contract) error {
	if node == nil {
		return nil
	}
	switch node.Type {
	case parser.TBlock:
		nodeToCode(node.Value.(*parser.NBlock).Statements, cnt)
	case parser.TContract:
		nodeToCode(node.Value.(*parser.NContract).Block, cnt)
		cnt.Name = node.Value.(*parser.NContract).Name
		fmt.Println(`CONT`, cnt.Name)
	case parser.TReturn:
		var vtype int32
		expr := node.Value.(*parser.NReturn).Expr
		if expr != nil {
			nodeToCode(expr, cnt)
			vtype = expr.Result
		}
		cnt.Code = append(cnt.Code, []rt.Bcode{rt.RETURN, rt.Bcode(vtype)}...)
	case parser.TStatements:
		for _, child := range node.Value.(*parser.NStatements).List {
			nodeToCode(child, cnt)
		}
	case parser.TUnary:
		nUnary := node.Value.(*parser.NUnary)
		nodeToCode(nUnary.Operand, cnt)
		switch nUnary.Oper {
		case parser.SUB:
			if nUnary.Operand.Result == parser.VInt {
				cnt.Code = append(cnt.Code, rt.SIGNINT)
				node.Result = parser.VInt
			}
		}
	case parser.TValue:
		switch v := node.Value.(type) {
		case int:
			cnt.Code = append(cnt.Code, []rt.Bcode{rt.PUSH16, rt.Bcode(v)}...)
		default:
			fmt.Printf("%v %T\n", v, node.Value)
		}
	}
	return nil
}

// Compile compiles contract
func Compile(input string) (cnt *Contract, err error) {
	var root *parser.Node

	cnt = &Contract{
		Code: make([]rt.Bcode, 0, 64),
	}
	root, err = parser.Parser(input)
	if err != nil {
		return
	}
	nodeToCode(root, cnt)
	fmt.Println(`ROOT`, cnt)
	return
}
