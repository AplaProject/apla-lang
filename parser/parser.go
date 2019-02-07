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
const INT = 57347
const TRUE = 57348
const FALSE = 57349
const NEWLINE = 57350
const COMMA = 57351
const COLON = 57352
const LPAREN = 57353
const RPAREN = 57354
const LBRACE = 57355
const RBRACE = 57356
const LBRAKET = 57357
const RBRAKET = 57358
const ADD = 57359
const SUB = 57360
const MUL = 57361
const DIV = 57362
const MOD = 57363
const ADD_ASSIGN = 57364
const SUB_ASSIGN = 57365
const MUL_ASSIGN = 57366
const DIV_ASSIGN = 57367
const MOD_ASSIGN = 57368
const ASSIGN = 57369
const AND = 57370
const OR = 57371
const EQ = 57372
const NOT_EQ = 57373
const NOT = 57374
const LT = 57375
const GT = 57376
const LTE = 57377
const GTE = 57378
const DATA = 57379
const CONTRACT = 57380
const IF = 57381
const ELIF = 57382
const ELSE = 57383
const RETURN = 57384
const WHILE = 57385
const T_INT = 57386
const T_BOOL = 57387
const UNARYMINUS = 57388
const UNARYNOT = 57389

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"INT",
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
	"T_INT",
	"T_BOOL",
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

const yyLast = 296

var yyAct = [...]int{

	34, 21, 20, 74, 23, 19, 96, 98, 2, 12,
	8, 15, 19, 53, 10, 105, 12, 9, 42, 43,
	100, 92, 104, 24, 56, 57, 60, 47, 48, 49,
	50, 51, 52, 21, 20, 11, 69, 5, 44, 25,
	16, 70, 71, 17, 18, 21, 20, 16, 33, 73,
	17, 18, 21, 20, 75, 3, 24, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
	76, 19, 54, 45, 46, 12, 32, 4, 1, 19,
	6, 94, 13, 12, 97, 95, 19, 91, 14, 93,
	12, 27, 28, 29, 30, 31, 26, 99, 58, 59,
	56, 57, 60, 7, 22, 0, 16, 0, 0, 17,
	18, 21, 20, 0, 16, 102, 103, 17, 18, 21,
	20, 16, 0, 0, 17, 18, 21, 20, 101, 0,
	0, 0, 58, 59, 56, 57, 60, 0, 0, 0,
	0, 0, 0, 61, 62, 63, 64, 90, 67, 68,
	65, 66, 58, 59, 56, 57, 60, 0, 0, 0,
	0, 0, 0, 61, 62, 63, 64, 0, 67, 68,
	65, 66, 72, 0, 0, 0, 58, 59, 56, 57,
	60, 0, 0, 0, 0, 0, 0, 61, 62, 63,
	64, 0, 67, 68, 65, 66, 55, 0, 0, 0,
	58, 59, 56, 57, 60, 0, 0, 0, 0, 0,
	0, 61, 62, 63, 64, 0, 67, 68, 65, 66,
	58, 59, 56, 57, 60, 0, 0, 0, 0, 0,
	0, 61, 62, 63, 64, 0, 67, 68, 65, 66,
	58, 59, 56, 57, 60, 0, 0, 58, 59, 56,
	57, 60, 62, 63, 64, 0, 67, 68, 65, 66,
	63, 64, 0, 67, 68, 65, 66, 39, 36, 37,
	38, 0, 0, 0, 35, 0, 0, 0, 0, 0,
	0, 40, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 41,
}
var yyPact = [...]int{

	-30, 47, 73, -1000, 24, -27, 3, -1000, 22, -1000,
	82, -43, -1000, 31, 69, 72, 263, 263, 263, -1000,
	-1000, -1000, 30, -1000, 70, -1000, 263, 263, 263, 263,
	263, 263, -14, 68, 183, 263, -1000, -1000, -1000, -1000,
	263, 263, 203, 159, -11, 68, -1000, 203, 203, 203,
	203, 203, 203, 263, -1000, -1000, 263, 263, 263, 263,
	263, 263, 263, 263, 263, 263, 263, 263, 263, 135,
	-1000, -1000, -1000, -1000, 13, 203, 75, -1000, -1000, 5,
	5, -1000, 223, 230, 81, 81, 81, 81, 81, 81,
	-1000, 67, -1000, -1000, -1000, -34, 263, -1000, 7, 115,
	-1000, -1000, 8, 1, -1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 48, 4, 104, 103, 88, 0, 85, 84,
	82, 14, 80, 78,
}
var yyR1 = [...]int{

	0, 1, 1, 11, 11, 11, 6, 9, 9, 8,
	8, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 2, 2, 3, 4, 4, 5, 5,
	12, 13, 13,
}
var yyR2 = [...]int{

	0, 1, 1, 0, 2, 3, 1, 0, 4, 0,
	6, 3, 3, 3, 3, 3, 3, 4, 2, 7,
	1, 2, 5, 3, 1, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 1, 2, 2, 1, 3, 0, 6,
	2, 5, 2,
}
var yyChk = [...]int{

	-1000, -13, 38, 8, 4, 13, -12, -5, 37, 14,
	-11, 13, 8, -10, -6, -1, 39, 42, 43, 4,
	45, 44, -4, -3, -1, 8, 27, 22, 23, 24,
	25, 26, 4, -2, -7, 11, 5, 6, 7, 4,
	18, 32, -7, -7, 8, -2, 4, -7, -7, -7,
	-7, -7, -7, 27, 4, 13, 19, 20, 17, 18,
	21, 28, 29, 30, 31, 35, 36, 33, 34, -7,
	-7, -7, 13, -3, 14, -7, -11, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	12, -11, 8, 14, 14, -8, 40, -9, 41, -7,
	13, 13, -11, -11, 14, 14,
}
var yyDef = [...]int{

	0, -2, 0, 52, 0, 48, 0, 3, 0, 51,
	50, 0, 4, 0, 0, 0, 0, 20, 0, 6,
	1, 2, 0, 46, 0, 5, 0, 0, 0, 0,
	0, 0, 43, 18, 0, 0, 24, 25, 26, 27,
	0, 0, 21, 0, 0, 45, 43, 11, 12, 13,
	14, 15, 16, 0, 44, 3, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	41, 42, 3, 47, 0, 17, 0, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	23, 0, 49, 9, 22, 7, 0, 19, 0, 0,
	3, 3, 0, 0, 8, 10,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47,
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
//line parser.y:102
		{
			yyVAL.i = VBool
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:103
		{
			yyVAL.i = VInt
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:107
		{
			yyVAL.n = nil
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:108
		{
			yyVAL.n = yyDollar[1].n
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:109
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:113
		{
			yyVAL.n = newVarValue(yyDollar[1].s, yylex)
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:116
		{
			yyVAL.n = nil
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:117
		{
			yyVAL.n = yyDollar[3].n
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:121
		{
			yyVAL.n = nil
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:122
		{
			yyVAL.n = newElif(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:126
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:127
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD_ASSIGN, yylex)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:128
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB_ASSIGN, yylex)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:129
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL_ASSIGN, yylex)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:130
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV_ASSIGN, yylex)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:131
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD_ASSIGN, yylex)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:132
		{
			yyVAL.n = newBinary(newVarDecl(yyDollar[1].i, []string{yyDollar[2].s}, yylex), yyDollar[4].n, ASSIGN, yylex)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:133
		{
			yyVAL.n = newVarDecl(yyDollar[1].i, yyDollar[2].sa, yylex)
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:134
		{
			yyVAL.n = newIf(yyDollar[2].n, yyDollar[4].n, yyDollar[6].n, yyDollar[7].n, yylex)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:135
		{
			yyVAL.n = newReturn(nil, yylex)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:136
		{
			yyVAL.n = newReturn(yyDollar[2].n, yylex)
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:137
		{
			yyVAL.n = newWhile(yyDollar[2].n, yyDollar[4].n, yylex)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:141
		{
			yyVAL.n = yyDollar[2].n
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:142
		{
			yyVAL.n = newValue(yyDollar[1].i, yylex)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:143
		{
			yyVAL.n = newValue(true, yylex)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:144
		{
			yyVAL.n = newValue(false, yylex)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:145
		{
			yyVAL.n = newGetVar(yyDollar[1].s, yylex)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:146
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL, yylex)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:147
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV, yylex)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:148
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD, yylex)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:149
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:150
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD, yylex)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:151
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, AND, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:152
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, OR, yylex)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:153
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, EQ, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:154
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, NOT_EQ, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:155
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LTE, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:156
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GTE, yylex)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:157
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LT, yylex)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:158
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GT, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:160
		{
			yyVAL.n = newUnary(yyDollar[2].n, SUB, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:161
		{
			yyVAL.n = newUnary(yyDollar[2].n, NOT, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:165
		{
			yyVAL.sa = []string{yyDollar[1].s}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:166
		{
			yyVAL.sa = append(yyDollar[1].sa, yyDollar[2].s)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:170
		{
			yyVAL.va = newVars(yyDollar[1].i, yyDollar[2].sa)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:174
		{
			yyVAL.va = yyDollar[1].va
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:175
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:179
		{
			yyVAL.va = nil
		}
	case 49:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:180
		{
			yyVAL.va = yyDollar[3].va
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:184
		{
			yyVAL.n = newBlock(yyDollar[1].va, yyDollar[2].n, yylex)
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:190
		{
			yyVAL.n = newContract(yyDollar[2].s, yyDollar[4].n, yylex)
			setResult(yylex, yyVAL.n)
		}
	}
	goto yystack /* stack new state and value */
}
