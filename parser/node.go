package parser

// Types of Node
const (
	TContract = iota + 1 // contract
	TData                // Data info
	TBlock
	TValue
	TVars
	TBinary
	TUnary
	TSetVar
	TIf
	TElif
	TReturn
	TGetVar
	TWhile
	TQuestion
	TFunc
	TCallFunc
	TParams
	TCallContract
	TContractParams
	TType
	TGetIndex
	TSetIndex
)

const (
	VVoid = iota
	VInt  // int
	VBool // bool
	VStr  // str
	VArr  // arr
	VMap  // map
)

// NVar contains type and name of variable or parameter
type NVar struct {
	Type *Node
	Name string
}

// NType contains the type
type NType struct {
	Type int64
	Def  bool
}

// NVarValue contains the name of variable
type NVarValue struct {
	Name string
}

// NWhile - while statement
type NWhile struct {
	Cond *Node
	Body *Node
}

// NGetIndex - getting index
type NGetIndex struct {
	Name    string
	Indexes []*Node
}

// NFunc - function
type NFunc struct {
	Name   string
	Result *Node
	Params []NVar
	Body   *Node
}

// NCallFunc - call function
type NCallFunc struct {
	Name   string
	Params *Node
}

// NCallContract - call contract
type NCallContract struct {
	Name   string
	Params []ContractParam
}

// NIf - if statement
type NIf struct {
	Cond     *Node
	IfBody   *Node
	ElifBody *Node
	ElseBody *Node
}

// NElifBody - elif statement
type NElifBody struct {
	Cond *Node
	Body *Node
}

// NElif - elif statement
type NElif struct {
	List []*NElifBody
}

// NBlock contains body of contract
type NBlock struct {
	Params     []NVar
	Statements []*Node
}

// NParams contains param expressions
type NParams struct {
	Expr []*Node
}

type ContractParam struct {
	Name string
	Expr *Node
}

// NContractParams contains contract param expressions
type NContractParams struct {
	Params []ContractParam
}

// NBinary contains binary operator
type NBinary struct {
	Oper  int
	Left  *Node
	Right *Node
}

// NQuestion contains ? operator
type NQuestion struct {
	Cond  *Node
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

// NReturn is a return statement
type NReturn struct {
	Expr *Node
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
	Column uint32
	Result uint32
	Value  interface{}
}

func setPos(node *Node, l yyLexer) *Node {
	pos := l.(*lexer).FilePosition()
	node.Line = pos.Line
	node.Column = uint32(pos.Column)
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

func newBlock(vars []NVar, block *Node, l yyLexer) *Node {
	if block == nil {
		block = setPos(&Node{
			Type:  TBlock,
			Value: &NBlock{},
		}, l)
	}
	if len(vars) > 0 {
		block.Value.(*NBlock).Params = vars
	}
	return block
}

func newParam(expr *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TParams,
		Value: &NParams{
			Expr: []*Node{expr},
		},
	}, l)
}

func newType(itype int64, l yyLexer) *Node {
	var def bool
	if itype == VArr || itype == VMap {
		def = true
		itype |= VStr << 4
	}
	return setPos(&Node{
		Type: TType,
		Value: &NType{
			Type: itype,
			Def:  def,
		},
	}, l)
}

func addSubtype(tNode *Node, ichild int64, l yyLexer) *Node {
	itype := tNode.Value.(*NType).Type
	var (
		i uint64
	)
	if tNode.Value.(*NType).Def {
		for i = 12; i > 0; i -= 4 {
			if (itype >> i) != 0 {
				itype &= ^(0xf << i)
				break
			}
		}
		tNode.Value.(*NType).Def = false
	}
	if (itype >> 12) == 0 {
		for i = 4; i < 16; i += 4 {
			if itype&(0xf<<i) == 0 {
				if (itype>>(i-4)) != VArr && (itype>>(i-4)) != VMap {
					itype = VVoid
				} else {
					itype |= ichild << i
				}
				break
			}
		}
	} else {
		itype = VVoid
	}
	if itype != VVoid && (ichild == VArr || ichild == VMap) {
		if itype&0xf000 != 0 {
			itype = VVoid
		} else {
			tNode.Value.(*NType).Def = true
			itype |= VStr << (i + 4)
		}
	}
	tNode.Value.(*NType).Type = itype
	return tNode
}

func addParam(node *Node, expr *Node) *Node {
	node.Value.(*NParams).Expr = append(node.Value.(*NParams).Expr, expr)
	return node
}

func newContractParam(name string, expr *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TContractParams,
		Value: &NContractParams{
			Params: []ContractParam{
				{
					Name: name,
					Expr: expr,
				},
			},
		},
	}, l)
}

func addContractParam(params *Node, name string, expr *Node) *Node {
	params.Value.(*NContractParams).Params = append(params.Value.(*NContractParams).Params,
		ContractParam{
			Name: name,
			Expr: expr,
		})
	return params
}

func newReturn(expr *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TReturn,
		Value: &NReturn{
			Expr: expr,
		},
	}, l)
}

func newVars(vtype *Node, vars []string) []NVar {
	va := make([]NVar, len(vars))
	for i, name := range vars {
		va[i] = NVar{
			Type: vtype,
			Name: name,
		}
	}
	return va
}

func addStatement(statements *Node, statement *Node, l yyLexer) *Node {
	if statements == nil {
		list := &NBlock{
			Statements: make([]*Node, 1, 10),
		}
		list.Statements[0] = statement
		statements = setPos(&Node{
			Type:  TBlock,
			Value: list,
		}, l)
	} else {
		statements.Value.(*NBlock).Statements = append(statements.Value.(*NBlock).Statements, statement)
	}
	return statements
}

func newValue(value interface{}, l yyLexer) *Node {
	var vtype uint32
	switch value.(type) {
	case int64:
		vtype = VInt
	case string:
		vtype = VStr
	case bool:
		vtype = VBool
	}
	return setPos(&Node{
		Type:   TValue,
		Value:  value,
		Result: vtype,
	}, l)
}

func newVarDecl(tNode *Node, vars []string, l yyLexer) *Node {
	list := make([]NVar, len(vars))
	for i, v := range vars {
		list[i] = NVar{
			Type: tNode,
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
		Type: TSetVar,
		Value: &NVarValue{
			Name: name,
		},
	}, l)
}

func newGetVar(name string, l yyLexer) *Node {
	return setPos(&Node{
		Type: TGetVar,
		Value: &NVarValue{
			Name: name,
		},
	}, l)
}

func newIndex(name string, index *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TGetIndex,
		Value: &NGetIndex{
			Name:    name,
			Indexes: []*Node{index},
		},
	}, l)
}

func addIndex(left *Node, index *Node, l yyLexer) *Node {
	left.Value.(*NGetIndex).Indexes = append(left.Value.(*NGetIndex).Indexes, index)
	return left
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

func newWhile(cond *Node, body *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TWhile,
		Value: &NWhile{
			Cond: cond,
			Body: body,
		},
	}, l)
}

func newIf(cond *Node, ifbody *Node, elif *Node, elsebody *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TIf,
		Value: &NIf{
			Cond:     cond,
			IfBody:   ifbody,
			ElifBody: elif,
			ElseBody: elsebody,
		},
	}, l)
}

func newElif(statements *Node, cond *Node, statement *Node, l yyLexer) *Node {
	if statements == nil {
		list := &NElif{
			List: make([]*NElifBody, 1, 10),
		}
		list.List[0] = &NElifBody{
			Cond: cond,
			Body: statement,
		}
		statements = setPos(&Node{
			Type:  TElif,
			Value: list,
		}, l)
	} else {
		statements.Value.(*NElif).List = append(statements.Value.(*NElif).List, &NElifBody{
			Cond: cond,
			Body: statement,
		})
	}
	return statements
}

func newQuestion(cond *Node, left *Node, right *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TQuestion,
		Value: &NQuestion{
			Cond:  cond,
			Left:  left,
			Right: right,
		},
	}, l)
}

func newFunc(name string, pars []NVar, retType *Node, body *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TFunc,
		Value: &NFunc{
			Name:   name,
			Result: retType,
			Body:   body,
			Params: pars,
		},
	}, l)
}

func newCallFunc(name string, params *Node, l yyLexer) *Node {
	return setPos(&Node{
		Type: TCallFunc,
		Value: &NCallFunc{
			Name:   name,
			Params: params,
		},
	}, l)
}

func newCallContract(name string, params *Node, l yyLexer) *Node {
	var list []ContractParam
	if params != nil {
		list = params.Value.(*NContractParams).Params
	}
	return setPos(&Node{
		Type: TCallContract,
		Value: &NCallContract{
			Name:   name[1:],
			Params: list,
		},
	}, l)
}

// Parser creates AST
func Parser(input string) (*Node, error) {
	yyErrorVerbose = true

	l, err := NewLexer(``, input)
	if err != nil {
		return nil, err
	}
	yyParse(l)
	if l.err != nil {
		return nil, l.err
	}
	return l.result.(*Node), nil
}
