%yyc c
%yyn c = l.Next()
%yym l.Mark()

%{
package parser

import (
	"strconv"

	"modernc.org/golex/lex"
)

func (l *lexer) scan(lval *yySymType) lex.Char {
	c := l.Enter()
%}

digit			[0-9]|{unicodeDigit}
identifier		{letter}({letter}|{digit})*
int				[0-9]+
float			{int}\.[0-9]*
string			\"({letter}|{digit})*\"
letter			[_a-zA-Z]|{unicodeLetter}
unicodeDigit	\x81
unicodeLetter	\x80

%%

	c = l.Rule0()

[ \t\r ]+		// ignore all whitespace

,				return l.char(COMMA)
:				return l.char(COLON)
\(				return l.char(LPAREN)
\)				return l.char(RPAREN)
\{				return l.char(LBRACE)
\}				return l.char(RBRACE)
\[				return l.char(LBRAKET)
\]				return l.char(RBRAKET)
\n				return l.char(NEWLINE)
;				return l.char(NEWLINE)

\+\n			return l.char(ADD)
\+				return l.char(ADD)
-				return l.char(SUB)
\*				return l.char(MUL)
\/				return l.char(DIV)
%				return l.char(MOD)


data    		return l.char(DATA)
contract		return l.char(CONTRACT)

bool			return l.char(T_BOOL)
int				return l.char(T_INT)

{int}			{
					lval.i, _ = strconv.Atoi(string(l.TokenBytes(nil)))
					return l.char(INT)
				}
{identifier}	{
					lval.s = string(l.TokenBytes(nil))
					return l.char(IDENT)
				}

%%
	if c, ok := l.Abort(); ok {
		return l.char(c)
	}

	goto yyAction
}