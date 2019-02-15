// Code generated by goyacc -o parser.go -v y.output parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

func setResult(l yyLexer, v *Node) {
	l.(*lexer).result = v
}

//line parser.y:10
type yySymType struct {
	yys int
	n   *Node
	b   bool
	i   int
	f   float64
	s   string
	sa  []string
	va  []NVar
}

const IDENT = 57346
const CALL = 57347
const CALLCONTRACT = 57348
const INT = 57349
const STRING = 57350
const QSTRING = 57351
const TRUE = 57352
const FALSE = 57353
const NEWLINE = 57354
const COMMA = 57355
const COLON = 57356
const LPAREN = 57357
const RPAREN = 57358
const LBRACE = 57359
const RBRACE = 57360
const LBRAKET = 57361
const RBRAKET = 57362
const QUESTION = 57363
const ADD = 57364
const SUB = 57365
const MUL = 57366
const DIV = 57367
const MOD = 57368
const ADD_ASSIGN = 57369
const SUB_ASSIGN = 57370
const MUL_ASSIGN = 57371
const DIV_ASSIGN = 57372
const MOD_ASSIGN = 57373
const ASSIGN = 57374
const AND = 57375
const OR = 57376
const EQ = 57377
const NOT_EQ = 57378
const NOT = 57379
const LT = 57380
const GT = 57381
const LTE = 57382
const GTE = 57383
const DATA = 57384
const CONTRACT = 57385
const IF = 57386
const ELIF = 57387
const ELSE = 57388
const RETURN = 57389
const WHILE = 57390
const FUNC = 57391
const T_INT = 57392
const T_BOOL = 57393
const T_STR = 57394
const UNARYMINUS = 57395
const UNARYNOT = 57396

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"CALL",
	"CALLCONTRACT",
	"INT",
	"STRING",
	"QSTRING",
	"TRUE",
	"FALSE",
	"NEWLINE",
	"COMMA",
	"COLON",
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"LBRAKET",
	"RBRAKET",
	"QUESTION",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"MOD_ASSIGN",
	"ASSIGN",
	"AND",
	"OR",
	"EQ",
	"NOT_EQ",
	"NOT",
	"LT",
	"GT",
	"LTE",
	"GTE",
	"DATA",
	"CONTRACT",
	"IF",
	"ELIF",
	"ELSE",
	"RETURN",
	"WHILE",
	"FUNC",
	"T_INT",
	"T_BOOL",
	"T_STR",
	"UNARYMINUS",
	"UNARYNOT",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 455

var yyAct = [...]int{

	55, 15, 90, 27, 24, 23, 25, 54, 37, 95,
	129, 131, 2, 28, 10, 8, 66, 38, 51, 52,
	31, 32, 33, 34, 35, 30, 71, 72, 69, 70,
	73, 60, 61, 62, 63, 64, 65, 58, 9, 135,
	82, 24, 23, 25, 22, 20, 21, 69, 70, 73,
	86, 87, 12, 83, 128, 91, 11, 116, 143, 28,
	115, 94, 92, 92, 5, 112, 93, 96, 84, 56,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 97, 16, 85, 113, 17, 18, 19,
	24, 23, 25, 118, 22, 20, 21, 119, 57, 29,
	117, 3, 12, 114, 53, 67, 59, 36, 142, 4,
	1, 6, 22, 20, 21, 13, 130, 124, 91, 125,
	12, 126, 127, 14, 7, 89, 137, 26, 123, 0,
	134, 0, 0, 136, 16, 0, 0, 17, 18, 19,
	24, 23, 25, 133, 0, 0, 22, 20, 21, 0,
	139, 0, 16, 141, 12, 17, 18, 19, 24, 23,
	25, 22, 20, 21, 0, 0, 0, 0, 0, 12,
	22, 20, 21, 0, 0, 122, 0, 0, 12, 0,
	0, 0, 0, 0, 120, 0, 16, 0, 0, 17,
	18, 19, 24, 23, 25, 0, 0, 0, 0, 0,
	0, 16, 0, 0, 17, 18, 19, 24, 23, 25,
	16, 0, 0, 17, 18, 19, 24, 23, 25, 140,
	0, 0, 0, 0, 0, 71, 72, 69, 70, 73,
	0, 0, 0, 0, 0, 0, 74, 75, 76, 77,
	138, 80, 81, 78, 79, 71, 72, 69, 70, 73,
	0, 0, 0, 0, 0, 0, 74, 75, 76, 77,
	132, 80, 81, 78, 79, 0, 0, 0, 0, 71,
	72, 69, 70, 73, 0, 0, 0, 0, 0, 0,
	74, 75, 76, 77, 121, 80, 81, 78, 79, 0,
	0, 0, 0, 71, 72, 69, 70, 73, 0, 0,
	0, 0, 0, 0, 74, 75, 76, 77, 111, 80,
	81, 78, 79, 0, 71, 72, 69, 70, 73, 0,
	0, 0, 0, 0, 0, 74, 75, 76, 77, 88,
	80, 81, 78, 79, 71, 72, 69, 70, 73, 0,
	0, 0, 0, 0, 0, 74, 75, 76, 77, 68,
	80, 81, 78, 79, 71, 72, 69, 70, 73, 0,
	0, 0, 0, 0, 0, 74, 75, 76, 77, 0,
	80, 81, 78, 79, 71, 72, 69, 70, 73, 0,
	0, 0, 0, 0, 0, 74, 75, 76, 77, 0,
	80, 81, 78, 79, 71, 72, 69, 70, 73, 0,
	0, 71, 72, 69, 70, 73, 75, 76, 77, 0,
	80, 81, 78, 79, 76, 77, 0, 80, 81, 78,
	79, 47, 45, 46, 40, 41, 42, 43, 44, 0,
	0, 0, 39, 0, 0, 0, 0, 0, 48, 0,
	49, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 50,
}
var yyPact = [...]int{

	-31, 89, 105, -1000, 47, -27, 20, -1000, 39, -1000,
	142, -46, -1000, 87, -7, 103, 417, 417, 417, 99,
	417, 53, -1000, -1000, -1000, -1000, 86, -1000, 102, -1000,
	417, 417, 417, 417, 417, 417, -16, 101, 332, 417,
	-1000, -1000, -1000, -1000, -1000, 417, 52, -1000, 70, 417,
	417, 352, 312, -46, 50, 352, -1000, -9, 101, -1000,
	352, 352, 352, 352, 352, 352, 417, -1000, -1000, 417,
	417, 417, 417, 417, 417, 417, 417, 417, 417, 417,
	417, 417, 292, 49, -1000, 417, -1000, -1000, -1000, 44,
	-1000, 102, 417, -1000, -1000, 85, 352, 166, -1000, -1000,
	23, 23, -1000, 372, 379, 4, 4, 4, 4, 4,
	4, -1000, -1000, 271, 157, -46, -46, 101, 352, -1000,
	-1000, 417, -1000, 37, -1000, -1000, -35, 247, -1000, 417,
	-1000, 22, 417, 108, 223, -1000, 203, -1000, -1000, 90,
	-1000, 40, -1000, -1000,
}
var yyPgo = [...]int{

	0, 1, 128, 8, 3, 127, 2, 125, 124, 7,
	123, 0, 121, 116, 115, 14, 111, 110,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 2, 2, 15, 15, 15, 9,
	9, 9, 10, 13, 13, 12, 12, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 3, 3, 6,
	7, 7, 7, 4, 5, 5, 8, 8, 16, 17,
	17,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 0, 1, 0, 2, 3, 0,
	1, 3, 1, 0, 4, 0, 6, 3, 3, 3,
	3, 3, 3, 4, 2, 7, 1, 2, 5, 8,
	3, 2, 3, 1, 1, 1, 1, 1, 3, 2,
	1, 8, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 1, 2, 2,
	0, 1, 3, 2, 1, 3, 0, 6, 2, 5,
	2,
}
var yyChk = [...]int{

	-1000, -17, 43, 12, 4, 17, -16, -8, 42, 18,
	-15, 17, 12, -14, -10, -1, 44, 47, 48, 49,
	5, 6, 4, 51, 50, 52, -5, -4, -1, 12,
	32, 27, 28, 29, 30, 31, 4, -3, -11, 15,
	7, 8, 9, 10, 11, 5, 6, 4, 21, 23,
	37, -11, -11, 5, -9, -11, 16, 12, -3, 4,
	-11, -11, -11, -11, -11, -11, 32, 4, 17, 24,
	25, 22, 23, 26, 33, 34, 35, 36, 40, 41,
	38, 39, -11, -9, 16, 15, -11, -11, 17, -7,
	-6, -1, 13, 16, -4, 18, -11, -15, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, 16, 16, -11, -15, 16, 13, -3, -11, 12,
	18, 13, 18, -2, -1, -6, -12, -11, 17, 45,
	-13, 46, 13, -15, -11, 17, -11, 18, 17, -15,
	16, -15, 18, 18,
}
var yyDef = [...]int{

	0, -2, 0, 70, 0, 66, 0, 6, 0, 69,
	68, 0, 7, 0, 0, 0, 0, 26, 0, 0,
	9, 0, 12, 1, 2, 3, 0, 64, 0, 8,
	0, 0, 0, 0, 0, 0, 57, 24, 0, 0,
	33, 34, 35, 36, 37, 9, 0, 40, 0, 0,
	0, 27, 0, 60, 0, 10, 31, 0, 63, 57,
	17, 18, 19, 20, 21, 22, 0, 58, 6, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 39, 0, 55, 56, 6, 0,
	61, 0, 0, 30, 65, 0, 23, 0, 42, 43,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 32, 38, 0, 0, 4, 0, 59, 11, 67,
	15, 0, 28, 0, 5, 62, 13, 0, 6, 0,
	25, 0, 0, 0, 0, 6, 0, 29, 6, 0,
	41, 0, 14, 16,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:113
		{
			yyVAL.i = VBool
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:114
		{
			yyVAL.i = VInt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:115
		{
			yyVAL.i = VStr
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:119
		{
			yyVAL.i = VVoid
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:120
		{
			yyVAL.i = yyDollar[1].i
		}
	case 6:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:124
		{
			yyVAL.n = nil
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:125
		{
			yyVAL.n = yyDollar[1].n
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:126
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:130
		{
			yyVAL.n = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:131
		{
			yyVAL.n = newParam(yyDollar[1].n, yylex)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:132
		{
			yyVAL.n = addParam(yyDollar[1].n, yyDollar[3].n)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:136
		{
			yyVAL.n = newVarValue(yyDollar[1].s, yylex)
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:139
		{
			yyVAL.n = nil
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:140
		{
			yyVAL.n = yyDollar[3].n
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:144
		{
			yyVAL.n = nil
		}
	case 16:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:145
		{
			yyVAL.n = newElif(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:149
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:150
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD_ASSIGN, yylex)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:151
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB_ASSIGN, yylex)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:152
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL_ASSIGN, yylex)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:153
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV_ASSIGN, yylex)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:154
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD_ASSIGN, yylex)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:155
		{
			yyVAL.n = newBinary(newVarDecl(yyDollar[1].i, []string{yyDollar[2].s}, yylex), yyDollar[4].n, ASSIGN, yylex)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:156
		{
			yyVAL.n = newVarDecl(yyDollar[1].i, yyDollar[2].sa, yylex)
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:157
		{
			yyVAL.n = newIf(yyDollar[2].n, yyDollar[4].n, yyDollar[6].n, yyDollar[7].n, yylex)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:158
		{
			yyVAL.n = newReturn(nil, yylex)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:159
		{
			yyVAL.n = newReturn(yyDollar[2].n, yylex)
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:160
		{
			yyVAL.n = newWhile(yyDollar[2].n, yyDollar[4].n, yylex)
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:161
		{
			yyVAL.n = newFunc(yyDollar[2].s, yyDollar[3].va, yyDollar[5].i, yyDollar[7].n, yylex)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:164
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:165
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:169
		{
			yyVAL.n = yyDollar[2].n
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:170
		{
			yyVAL.n = newValue(yyDollar[1].i, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:171
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:172
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:173
		{
			yyVAL.n = newValue(true, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:174
		{
			yyVAL.n = newValue(false, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:175
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:176
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yylex)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:177
		{
			yyVAL.n = newGetVar(yyDollar[1].s, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:178
		{
			yyVAL.n = newQuestion(yyDollar[3].n, yyDollar[5].n, yyDollar[7].n, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:179
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:180
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV, yylex)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:181
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD, yylex)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:182
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB, yylex)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:183
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD, yylex)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:184
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, AND, yylex)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:185
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, OR, yylex)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:186
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, EQ, yylex)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:187
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, NOT_EQ, yylex)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:188
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LTE, yylex)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:189
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GTE, yylex)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:190
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LT, yylex)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:191
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GT, yylex)
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:193
		{
			yyVAL.n = newUnary(yyDollar[2].n, SUB, yylex)
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:194
		{
			yyVAL.n = newUnary(yyDollar[2].n, NOT, yylex)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:198
		{
			yyVAL.sa = []string{yyDollar[1].s}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:199
		{
			yyVAL.sa = append(yyDollar[1].sa, yyDollar[2].s)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:203
		{
			yyVAL.va = newVars(yyDollar[1].i, yyDollar[2].sa)
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:207
		{
			yyVAL.va = nil
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:208
		{
			yyVAL.va = yyDollar[1].va
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:209
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:213
		{
			yyVAL.va = newVars(yyDollar[1].i, yyDollar[2].sa)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:217
		{
			yyVAL.va = yyDollar[1].va
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:218
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:222
		{
			yyVAL.va = nil
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:223
		{
			yyVAL.va = yyDollar[3].va
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:227
		{
			yyVAL.n = newBlock(yyDollar[1].va, yyDollar[2].n, yylex)
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:233
		{
			yyVAL.n = newContract(yyDollar[2].s, yyDollar[4].n, yylex)
			setResult(yylex, yyVAL.n)
		}
	}
	goto yystack /* stack new state and value */
}
