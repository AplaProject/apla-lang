package parser

// Types of Node
const (
	TContract = iota + 1 // contract
	TData                // Data info
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
	Block *NBlock
}

// Node is a common node structure for yacc
type Node struct {
	Type  int
	Value interface{}
}

func newContract(name string, block *NBlock) *Node {
	return &Node{
		Type: TContract,
		Value: &NContract{
			Name:  name,
			Block: block,
		},
	}
}

func newBlock(vars []NVar) *NBlock {
	return &NBlock{
		Params: vars,
	}
}
