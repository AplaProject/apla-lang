package parser

import ()

// Types of Node
const (
	TContract = iota + 1 // contract
	TData                // Data info
	TBlock
	TStatements
)

// NVar contains type and name of variable or parameter
type NVar struct {
	Type int
	Name string
}

// NBlock contains body of contract
type NBlock struct {
	Params     []NVar
	Statements *Node
}

// NStatements contains statements
type NStatements struct {
	List []*Node
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

func newBlock(vars []NVar, statements *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TBlock,
		Value: &NBlock{
			Params:     vars,
			Statements: statements,
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

func newStatement(statements *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TStatements,
		Value: &NStatements{
			List: make([]*Node, 0, 10),
		},
	}, l)
}

func addStatement(statements *Node, statement *Node, l yyLexer) *Node {
	if statements == nil {
		statements = setPos(&Node{
			Type: TStatements,
			Value: &NStatements{
				List: make([]*Node, 0, 10),
			},
		}, l)
	} else {
		statements.Value.(*NStatements).List = append(statements.Value.(*NStatements).List, statement)
	}
}
