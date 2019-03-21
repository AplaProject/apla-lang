%yyc c
%yyn c = l.Next()
%yym l.Mark()

%{
package parser

import (
	"strconv"
	"strings"

	"modernc.org/golex/lex"
)

func (l *lexer) scan(lval *yySymType) lex.Char {
	c := l.Enter()
%}

hexint	    	0[xX][0-9a-fA-F]+
digit			[0-9]|{unicodeDigit}
int				[0-9]+
float			{int}\.[0-9]*
identifier		{letter}({letter}|{digit})*
callcontract    @{identifier}\(
call    		{identifier}\(
index           {identifier}\[
string 			\"([^\\"]|\\.)*\"
qstring 		`([^`])*`
letter			[_a-zA-Z]|{unicodeLetter}
unicodeDigit	\x81
unicodeLetter	\x80

%%

	c = l.Rule0()

[ \t\r ]+		 // ignore all whitespace
\/\*(.|\n)*\*\/  // ignore comments
\/\/(.)*\n      return l.char(NEWLINE)

\+=				return l.char(ADD_ASSIGN)
-=				return l.char(SUB_ASSIGN)
\*=				return l.char(MUL_ASSIGN)
\/=				return l.char(DIV_ASSIGN)
%=				return l.char(MOD_ASSIGN)
=				return l.char(ASSIGN)

,[ \t]*\n?		return l.char(COMMA)
:				return l.char(COLON)
\.				return l.char(DOT)
\?         	    return l.char(QUESTION)
\([ \t]*\n?     return l.char(LPAREN)
\n[ \t]*\)		return l.char(RPAREN)
\)				return l.char(RPAREN)
\{[ \t]*\n?  	return l.char(LBRACE)
\}		return l.char(RBRACE)
\[[ \t]*\n?		return l.char(LBRACKET)
\]				return l.char(RBRACKET)
\n				return l.char(NEWLINE)
;				return l.char(NEWLINE)

&&				return l.char(AND)
\|\|			return l.char(OR)

\+[ \t]*\n?		return l.char(ADD)
-[ \t]*\n?		return l.char(SUB)
\*[ \t]*\n?		return l.char(MUL)
\/[ \t]*\n?		return l.char(DIV)
%[ \t]*\n?		return l.char(MOD)

==[ \t]*\n?		return l.char(EQ)
!=[ \t]*\n?		return l.char(NOT_EQ)
!				return l.char(NOT)

\<=[ \t]*\n?	return l.char(LTE)
>=[ \t]*\n?		return l.char(GTE)
\<[ \t]*\n?		return l.char(LT)
>[ \t]*\n?  	return l.char(GT)

data    		return l.char(DATA)
contract		return l.char(CONTRACT)
while           return l.char(WHILE)
if				return l.char(IF)
elif			return l.char(ELIF)
else			return l.char(ELSE)
return			return l.char(RETURN)
true			return l.char(TRUE)
false			return l.char(FALSE)
func			return l.char(FUNC)

bool			return l.char(T_BOOL)
int				return l.char(T_INT)
hexint			return l.char(T_INT)
str			    return l.char(T_STR)
arr			    return l.char(T_ARR)
map			    return l.char(T_MAP)

{hexint}			{
					val, _ := strconv.ParseInt(string(l.TokenBytes(nil)), 0, 64)
					lval.i = int64(val)
					return l.char(INT)
				}

{int}			{
					ai, _ := strconv.Atoi(string(l.TokenBytes(nil)))
					lval.i = int64(ai)
					return l.char(INT)
				}
{identifier}	{
					lval.s = string(l.TokenBytes(nil))
					return l.char(IDENT)
				}
{string}		{
					var err error
					s := strings.Replace( string(l.TokenBytes(nil)), "\n", `\n`, -1 )
					lval.s, err = strconv.Unquote(`"` + s[1:len(s)-1] + `"`)
					if err != nil  {
						l.Error(err.Error())
					}
					return l.char(STRING)
				}				
{qstring}		{
					s := string(l.TokenBytes(nil))
					lval.s = s[1:len(s)-1]
					return l.char(STRING)
				}				
{call}	{
					lval.s = string(l.TokenBytes(nil))
					lval.s = lval.s[:len(lval.s)-1]
					return l.char(CALL)
				}
{callcontract}	{
					lval.s = string(l.TokenBytes(nil))
					lval.s = lval.s[:len(lval.s)-1]
					return l.char(CALLCONTRACT)
				}
{index}	{
					lval.s = string(l.TokenBytes(nil))
					lval.s = lval.s[:len(lval.s)-1]
					return l.char(INDEX)
				}

%%
	if c, ok := l.Abort(); ok {
		return l.char(c)
	}

	goto yyAction
}