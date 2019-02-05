package compiler

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
)

const (
	errOperator = `Operator %s has not been found`
	errType     = `Unknown type %T`
)

func (cmpl *compiler) Error(node *parser.Node, text string) error {
	return fmt.Errorf("%s %d:%d: %s", cmpl.Contract.Name, node.Line, node.Column, text)
}

func (cmpl *compiler) ErrorOperator(node *parser.Node) error {
	var (
		oper              int
		name, left, right string
	)
	if node.Type == parser.TUnary {
		nUnary := node.Value.(*parser.NUnary)
		oper = nUnary.Oper
		right = Type2Str(nUnary.Operand.Result)
	} else {
		nBinary := node.Value.(*parser.NBinary)
		oper = nBinary.Oper
		left = Type2Str(nBinary.Left.Result)
		right = Type2Str(nBinary.Right.Result)
	}
	switch oper {
	case parser.SUB:
		name = `-`
	}
	return cmpl.Error(node, fmt.Sprintf(errOperator, left+name+right))
}
