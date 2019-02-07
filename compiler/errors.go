package compiler

import (
	"fmt"

	"github.com/AplaProject/apla-lang/parser"
)

const (
	errOperator   = `Operator %s has not been found`
	errType       = `Unknown type %T`
	errNodeType   = `Unknown node type`
	errVarExists  = `Variable %s has already been defined`
	errVarUnknown = `Variable %s hasn't been defined`
)

func (cmpl *compiler) Error(node *parser.Node, text string) error {
	return fmt.Errorf("%s %d:%d: %s", cmpl.Contract.Name, node.Line, node.Column, text)
}

func (cmpl *compiler) ErrorParam(node *parser.Node, text string, value interface{}) error {
	return cmpl.Error(node, fmt.Sprintf(text, value))
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
	case parser.ADD:
		name = `+`
	case parser.MUL:
		name = `*`
	case parser.DIV:
		name = `/`
	case parser.MOD:
		name = `%`
	case parser.ASSIGN:
		name = `=`
	case parser.ADD_ASSIGN:
		name = `+=`
	case parser.SUB_ASSIGN:
		name = `-=`
	case parser.MUL_ASSIGN:
		name = `*=`
	case parser.DIV_ASSIGN:
		name = `/=`
	case parser.MOD_ASSIGN:
		name = `%=`
	}
	return cmpl.Error(node, fmt.Sprintf(errOperator, left+name+right))
}
