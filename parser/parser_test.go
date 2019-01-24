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
				data { int my
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
			`contract Тест { data { int my1 int my2}
									    1
									}`,
			nil,
		},
		{
			`contract Тест {
								    7
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
