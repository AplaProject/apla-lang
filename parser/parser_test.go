package parser

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRT(t *testing.T) {
	code := []uint16{ // a b c i n
		PUSH16, 1,
		SETVAR, 0, // a = 1
		PUSH16, 1,
		SETVAR, 1, // b = 1
		PUSH16, 3, // i = 3
		SETVAR, 3,
		PUSH32, 0x98, 0x9680, // n = 10000000
		SETVAR, 4,
		// 17
		GETVAR, 3,
		GETVAR, 4,
		GTI32,   // i > n
		JNZ, 48, // break if i > n
		GETVAR, 0,
		GETVAR, 1,
		ADDI32,
		SETVAR, 2, // c = a + b
		GETVAR, 1,
		SETVAR, 0, // a = b
		GETVAR, 2,
		SETVAR, 1, // b = c
		PUSH16, 1,
		GETVAR, 3,
		ADDI32,
		SETVAR, 3, // i = i + 1
		JMP, 17, // loop again
		// 48
		GETVAR, 1, // return b
	}
	start := time.Now()
	i, gas := Run(code)
	fmt.Println(i, gas, time.Since(start))
	t.Error(`OK`)
}

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
