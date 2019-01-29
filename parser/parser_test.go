package parser

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	syntaxErr := func(v string) error {
		return errors.New(v)
	}

	tests := []struct {
		in  string
		err error
	}{
		{
			`contract {
				}`,
			syntaxErr("file:1:10: syntax error: unexpected LBRACE, expecting IDENT"),
		},
		{
			`contract Test {
				data { 
					int my
				}
						}`,
			nil,
		},
		{
			`contract Тест { data { int My My2
									bool Is
							  }
							}`,
			nil,
		},
		{
			`contract Тест {
										data {
											int my1
											int my2
										}
										int b c
										int a = -1
										b += 20
				}`,
			nil,
		},

		{
			`contract Тест {
							int a = 7 +
							-3 * 4 *(1 + 2) ;}`,
			nil,
		},
		{
			`contract Тест {
				if true { a = 10 
				}
				if false {
					q = 10
				} else {
					q = 20
				}
				if true && true {
					
				} elif false || false {

				} elif true || false {

				}
			}`,
			nil,
		},
	}

	yyErrorVerbose = true

	for _, v := range tests {
		l, err := NewLexer("file", v.in)
		assert.NoError(t, err)
		yyParse(l)
		fmt.Println(`LEXER`, l.result)
		if v.err != nil {
			assert.EqualError(t, l.err, v.err.Error())
			continue
		}

		if l.err != nil {
			t.Error(l.err)
			break
		}
	}
	t.Error(`OK`)
}
