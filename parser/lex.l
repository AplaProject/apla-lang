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

\+=				return l.char(ADD_ASSIGN)
-=				return l.char(SUB_ASSIGN)
\*=				return l.char(MUL_ASSIGN)
\/=				return l.char(DIV_ASSIGN)
%=				return l.char(MOD_ASSIGN)
=				return l.char(ASSIGN)

,[ \t]*\n?		return l.char(COMMA)
:				return l.char(COLON)
\([ \t]*\n?     return l.char(LPAREN)
\n[ \t]*\)		return l.char(RPAREN)
\)				return l.char(RPAREN)
\{[ \t]*\n?  	return l.char(LBRACE)
\}				return l.char(RBRACE)
\[[ \t]*\n?		return l.char(LBRAKET)
\]				return l.char(RBRAKET)
\n				return l.char(NEWLINE)
;				return l.char(NEWLINE)

&&				return l.char(AND)
\|\|			return l.char(OR)

\+[ \t]*\n?		return l.char(ADD)
-[ \t]*\n?		return l.char(SUB)
\*[ \t]*\n?		return l.char(MUL)
\/[ \t]*\n?		return l.char(DIV)
%[ \t]*\n?		return l.char(MOD)

==				return l.char(EQ)
!=				return l.char(NOT_EQ)
!				return l.char(NOT)

\<=				return l.char(LTE)
>=				return l.char(GTE)
\<				return l.char(LT)
>				return l.char(GT)

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