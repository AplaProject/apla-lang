package parser

import ()

// Types of Node
const (
	TContract = iota + 1 // contract
	TData                // Data info
	TBlock
	TStatements
	TValue
	TVars
	TBinary
	TUnary
	TVarValue
)

// NVar contains type and name of variable or parameter
type NVar struct {
	Type int
	Name string
}

// NVarValue contains the name of variable
type NVarValue struct {
	Name string
}

// NBlock contains body of contract
type NBlock struct {
	Params     []NVar
	Statements *Node
}

// NBinary contains binary operator
type NBinary struct {
	Oper  int
	Left  *Node
	Right *Node
}

// NUnary contains an unary operator
type NUnary struct {
	Oper    int
	Operand *Node
}

// NVars contains type and name of variable or parameter
type NVars struct {
	Vars []NVar
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
	return statements
}

func newValue(value interface{}, l yyLexer) *Node {
	return setPos(&Node{
		Type:  TValue,
		Value: value,
	}, l)
}

func newVarDecl(itype int, vars []string, l yyLexer) *Node {
	list := make([]NVar, len(vars))
	for i, v := range vars {
		list[i] = NVar{
			Type: itype,
			Name: v,
		}
	}
	return setPos(&Node{
		Type: TVars,
		Value: &NVars{
			Vars: list,
		},
	}, l)
}

func newBinary(left *Node, right *Node, oper int, l yyLexer) *Node {
	return setPos(&Node{
		Type: TBinary,
		Value: &NBinary{
			Oper:  oper,
			Left:  left,
			Right: right,
		},
	}, l)
}

func newVarValue(name string, l yyLexer) *Node {
	return setPos(&Node{
		Type: TVarValue,
		Value: &NVarValue{
			Name: name,
		},
	}, l)
}

func newUnary(operand *Node, oper int, l yyLexer) *Node {
	return setPos(&Node{
		Type: TUnary,
		Value: &NUnary{
			Oper:    oper,
			Operand: operand,
		},
	}, l)
}
