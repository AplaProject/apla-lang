package compiler

import (
	"github.com/AplaProject/apla-lang/parser"
)

func newNVar(vtype uint32, name string) parser.NVar {
	return parser.NVar{
		Type: &parser.Node{
			Type: parser.TType,
			Value: &parser.NType{
				Type: int64(vtype),
			},
		},
		Name: name,
	}
}

func forCode(node *parser.Node, cmpl *compiler) error {
	var (
		err      error
		keysName string
	)
	nFor := node.Value.(*parser.NFor)
	isKey := true
	curLen := len(cmpl.Contract.Code)
	if err = nodeToCode(nFor.Expr, cmpl); err != nil {
		return err
	}
	if len(nFor.KeyName) == 0 {
		nFor.KeyName = parser.RandName()
		isKey = false
	}
	cmpl.Contract.Code = cmpl.Contract.Code[:curLen]
	maintype, subtype := parseType(nFor.Expr.Result)
	if maintype != parser.VArr && maintype != parser.VMap {
		return cmpl.ErrorParam(nFor.Expr, errForType, Type2Str(nFor.Expr.Result))
	}
	keyType := parser.VInt
	if maintype == parser.VMap {
		keyType = parser.VStr
		keysName = parser.RandName()
	}
	iKey := parser.RandName()
	objName := parser.RandName()
	vars := []parser.NVar{
		newNVar(nFor.Expr.Result, objName),
		newNVar(subtype, nFor.VarName),
		newNVar(uint32(keyType), nFor.KeyName),
		newNVar(parser.VInt, iKey),
	}
	if maintype == parser.VMap {
		vars = append(vars, newNVar((parser.VStr<<4)|parser.VArr, keysName))
	}

	if nFor.Body == nil {
		nFor.Body = &parser.Node{
			Type: parser.TBlock,
			Value: &parser.NBlock{
				Statements: make([]*parser.Node, 0, 10),
			},
		}
	}
	before := []*parser.Node{&parser.Node{
		Type: parser.TBinary,
		Value: &parser.NBinary{
			Oper: parser.ASSIGN,
			Left: &parser.Node{
				Type: parser.TSetVar,
				Value: &parser.NVarValue{
					Name: nFor.VarName,
				},
			},
			Right: &parser.Node{
				Type: parser.TGetIndex,
				Value: &parser.NGetIndex{
					Name: objName,
					Indexes: []*parser.Node{
						&parser.Node{
							Type: parser.TGetVar,
							Value: &parser.NVarValue{
								Name: iKey,
							},
						},
					},
				},
			},
		},
	}}
	if isKey {
		before = append(before, &parser.Node{
			Type: parser.TBinary,
			Value: &parser.NBinary{
				Oper: parser.ASSIGN,
				Left: &parser.Node{
					Type: parser.TSetVar,
					Value: &parser.NVarValue{
						Name: nFor.KeyName,
					},
				},
				Right: &parser.Node{
					Type: parser.TGetVar,
					Value: &parser.NVarValue{
						Name: iKey,
					},
				},
			},
		})
	}
	nFor.Body.Value.(*parser.NBlock).Statements = append(before, nFor.Body.Value.(*parser.NBlock).Statements...)

	nFor.Body.Value.(*parser.NBlock).Statements = append(nFor.Body.Value.(*parser.NBlock).Statements,
		&parser.Node{
			Type: parser.TBinary,
			Value: &parser.NBinary{
				Oper: parser.ADD_ASSIGN,
				Left: &parser.Node{
					Type: parser.TSetVar,
					Value: &parser.NVarValue{
						Name: iKey,
					},
				},
				Right: &parser.Node{
					Type:   parser.TValue,
					Value:  int64(1),
					Result: parser.VInt,
				},
			},
		})

	code := &parser.Node{
		Type: parser.TBlock,
		Value: &parser.NBlock{
			Statements: []*parser.Node{&parser.Node{
				Line:   node.Line,
				Column: node.Column,
				Type:   parser.TVars,
				Value: &parser.NVars{
					Vars: vars,
				},
			},
				&parser.Node{
					Line:   node.Line,
					Column: node.Column,
					Type:   parser.TBinary,
					Value: &parser.NBinary{
						Oper: parser.ASSIGN,
						Left: &parser.Node{
							Type: parser.TSetVar,
							Value: &parser.NVarValue{
								Name: objName,
							},
						},
						Right: nFor.Expr,
					},
				},
				&parser.Node{
					Line:   node.Line,
					Column: node.Column,
					Type:   parser.TWhile,
					Value: &parser.NWhile{
						Cond: &parser.Node{
							Type: parser.TBinary,
							Value: &parser.NBinary{
								Oper: parser.LT,
								Left: &parser.Node{
									Type: parser.TGetVar,
									Value: &parser.NVarValue{
										Name: iKey,
									},
								},
								Right: &parser.Node{
									Type: parser.TCallFunc,
									Value: &parser.NCallFunc{
										Name: `Len`,
										Params: &parser.Node{
											Type: parser.TParams,
											Value: &parser.NParams{
												Expr: []*parser.Node{
													&parser.Node{
														Type: parser.TGetVar,
														Value: &parser.NVarValue{
															Name: objName,
														},
													},
												},
											},
										},
									},
								},
							},
						},
						Body: nFor.Body,
					},
				},
			},
		}}
	return nodeToCode(code, cmpl)
}
