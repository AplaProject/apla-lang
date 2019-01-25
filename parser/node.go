package parser

import ()

// Types of Node
const (
	TContract = iota + 1 // contract
	TData                // Data info
	TBlock
)

// NVar contains type and name of variable or parameter
type NVar struct {
	Type int
	Name string
}

// NBlock contains statements
type NBlock struct {
	Params []NVar
	List   []*Node
}

// NContract is a root node
type NContract struct {
	Name  string // the name of the contract
	Block *Node
}

// Node is a common node structure for yacc
type Node struct {
	Type   int
	Line   int
	Column int
	Value  interface{}
}

func setPos(node *Node, l yyLexer) *Node {
	pos := l.(*lexer).FilePosition()
	node.Line = pos.Line
	node.Column = pos.Column
	return node
}

func newContract(name string, block *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TContract,
		Value: &NContract{
			Name:  name,
			Block: block,
		},
	}, l)
}

func newBlock(vars []NVar, l yyLexer) *Node {
	return setPos(&Node{
		Type: TBlock,
		Value: &NBlock{
			Params: vars,
		},
	}, l)
}

func newVars(vtype int, vars []string) []NVar {
	va := make([]NVar, len(vars))
	for i, name := range vars {
		va[i] = NVar{
			Type: vtype,
			Name: name,
		}
	}
	return va
}
