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
	i   int64
	f   float64
	s   string
	sa  []string
	va  []NVar
}

const IDENT = 57346
const CALL = 57347
const CALLCONTRACT = 57348
const INDEX = 57349
const INT = 57350
const STRING = 57351
const QSTRING = 57352
const TRUE = 57353
const FALSE = 57354
const NEWLINE = 57355
const COMMA = 57356
const COLON = 57357
const LPAREN = 57358
const RPAREN = 57359
const LBRACE = 57360
const RBRACE = 57361
const LBRACKET = 57362
const RBRACKET = 57363
const QUESTION = 57364
const DOT = 57365
const ADD = 57366
const SUB = 57367
const MUL = 57368
const DIV = 57369
const MOD = 57370
const ADD_ASSIGN = 57371
const SUB_ASSIGN = 57372
const MUL_ASSIGN = 57373
const DIV_ASSIGN = 57374
const MOD_ASSIGN = 57375
const ASSIGN = 57376
const AND = 57377
const OR = 57378
const EQ = 57379
const NOT_EQ = 57380
const NOT = 57381
const LT = 57382
const GT = 57383
const LTE = 57384
const GTE = 57385
const DATA = 57386
const CONTRACT = 57387
const IF = 57388
const ELIF = 57389
const ELSE = 57390
const RETURN = 57391
const WHILE = 57392
const FUNC = 57393
const T_INT = 57394
const T_BOOL = 57395
const T_STR = 57396
const T_ARR = 57397
const UNARYMINUS = 57398
const UNARYNOT = 57399

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"CALL",
	"CALLCONTRACT",
	"INDEX",
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
	"LBRACKET",
	"RBRACKET",
	"QUESTION",
	"DOT",
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
	"T_ARR",
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

const yyLast = 566

var yyAct = [...]int{

	62, 53, 102, 67, 43, 28, 27, 29, 30, 63,
	17, 61, 2, 16, 11, 141, 154, 156, 9, 44,
	58, 59, 26, 78, 41, 39, 65, 34, 35, 36,
	37, 38, 33, 39, 69, 70, 71, 72, 73, 74,
	75, 76, 68, 81, 82, 85, 94, 40, 28, 27,
	29, 30, 83, 84, 81, 82, 85, 98, 99, 112,
	10, 42, 96, 95, 77, 160, 152, 135, 12, 5,
	134, 103, 106, 111, 104, 131, 148, 130, 41, 114,
	41, 97, 116, 117, 118, 119, 120, 121, 122, 123,
	124, 125, 126, 127, 128, 115, 106, 104, 132, 107,
	105, 108, 149, 110, 32, 137, 31, 6, 136, 139,
	3, 60, 79, 138, 140, 133, 64, 16, 4, 1,
	7, 68, 14, 155, 150, 15, 8, 101, 66, 145,
	0, 0, 0, 0, 0, 16, 0, 0, 147, 24,
	22, 23, 25, 0, 151, 146, 103, 0, 13, 153,
	24, 22, 23, 25, 168, 159, 0, 0, 161, 13,
	16, 24, 22, 23, 25, 167, 16, 158, 16, 0,
	13, 0, 0, 0, 0, 164, 162, 0, 166, 0,
	0, 18, 0, 0, 19, 20, 21, 28, 27, 29,
	30, 0, 18, 0, 0, 19, 20, 21, 28, 27,
	29, 30, 0, 18, 0, 0, 19, 20, 21, 28,
	27, 29, 30, 24, 22, 23, 25, 0, 0, 0,
	0, 0, 13, 24, 22, 23, 25, 0, 144, 0,
	0, 0, 13, 24, 22, 23, 25, 0, 142, 0,
	0, 0, 13, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 18, 0, 0, 19, 20,
	21, 28, 27, 29, 30, 18, 0, 0, 19, 20,
	21, 28, 27, 29, 30, 18, 0, 0, 19, 20,
	21, 28, 27, 29, 30, 165, 0, 0, 0, 0,
	0, 0, 83, 84, 81, 82, 85, 0, 0, 0,
	0, 0, 0, 86, 87, 88, 89, 163, 92, 93,
	90, 91, 0, 83, 84, 81, 82, 85, 0, 0,
	0, 0, 0, 157, 86, 87, 88, 89, 0, 92,
	93, 90, 91, 83, 84, 81, 82, 85, 0, 0,
	0, 0, 0, 143, 86, 87, 88, 89, 0, 92,
	93, 90, 91, 83, 84, 81, 82, 85, 0, 0,
	0, 0, 0, 0, 86, 87, 88, 89, 129, 92,
	93, 90, 91, 0, 0, 83, 84, 81, 82, 85,
	0, 0, 0, 0, 0, 0, 86, 87, 88, 89,
	0, 92, 93, 90, 91, 113, 0, 0, 83, 84,
	81, 82, 85, 0, 0, 0, 0, 0, 0, 86,
	87, 88, 89, 0, 92, 93, 90, 91, 109, 0,
	0, 83, 84, 81, 82, 85, 0, 0, 0, 0,
	0, 0, 86, 87, 88, 89, 100, 92, 93, 90,
	91, 0, 83, 84, 81, 82, 85, 0, 0, 0,
	0, 0, 0, 86, 87, 88, 89, 80, 92, 93,
	90, 91, 0, 83, 84, 81, 82, 85, 0, 0,
	0, 0, 0, 0, 86, 87, 88, 89, 0, 92,
	93, 90, 91, 83, 84, 81, 82, 85, 0, 0,
	0, 0, 0, 0, 86, 87, 88, 89, 0, 92,
	93, 90, 91, 83, 84, 81, 82, 85, 0, 0,
	83, 84, 81, 82, 85, 87, 88, 89, 0, 92,
	93, 90, 91, 88, 89, 0, 92, 93, 90, 91,
	54, 51, 52, 25, 46, 47, 48, 49, 50, 0,
	0, 0, 45, 0, 0, 0, 0, 0, 55, 0,
	0, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 57,
}
var yyPact = [...]int{

	-33, 97, 114, -1000, 51, 94, -26, 41, -1000, 50,
	-1000, 229, 93, -1000, 91, -2, 13, 57, 526, 526,
	526, 106, 526, 112, -1000, 526, -1000, -1000, -1000, -1000,
	-1000, -47, -1000, 526, 526, 526, 526, 526, 526, 526,
	526, -47, -11, 108, 439, 526, -1000, -1000, -1000, -1000,
	-1000, 526, 112, 5, -1000, 65, 526, 526, 459, 418,
	-47, 83, 459, 82, 86, 397, 90, -1000, 55, 459,
	459, 459, 459, 459, 459, 374, 459, -1000, 526, -1000,
	-1000, 526, 526, 526, 526, 526, 526, 526, 526, 526,
	526, 526, 526, 526, 351, 60, 58, 526, -1000, -1000,
	-1000, 53, -1000, 55, 526, -1000, 109, -1000, 526, -1000,
	-4, 108, -1000, -1000, 459, 219, -1000, -1000, 17, 17,
	-1000, 479, 486, 28, 28, 28, 28, 28, 28, -1000,
	-1000, -1000, 329, 209, -47, -47, 108, 459, 61, 459,
	-1000, 89, -1000, 526, -1000, 48, 1, -1000, 526, -1000,
	-31, 309, -1000, 459, 526, -1000, 47, 526, 157, 289,
	-1000, 268, -1000, -1000, 146, -1000, 135, -1000, -1000,
}
var yyPgo = [...]int{

	0, 22, 10, 129, 4, 3, 128, 2, 127, 126,
	11, 125, 0, 124, 123, 122, 14, 9, 120, 119,
	1,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 1, 2, 2, 3, 3, 16,
	16, 16, 10, 10, 10, 17, 17, 17, 11, 20,
	20, 14, 14, 13, 13, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 4, 4, 7,
	8, 8, 8, 5, 6, 6, 9, 9, 18, 19,
	19,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 3, 0, 1, 0,
	2, 3, 0, 1, 3, 0, 3, 5, 1, 3,
	4, 0, 4, 0, 6, 3, 3, 3, 3, 3,
	3, 3, 4, 2, 7, 1, 2, 5, 8, 3,
	3, 3, 1, 1, 1, 1, 1, 3, 3, 1,
	1, 8, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 1, 2, 2,
	0, 1, 3, 2, 1, 3, 0, 7, 2, 6,
	2,
}
var yyChk = [...]int{

	-1000, -19, 45, 13, 4, 18, 13, -18, -9, 44,
	19, -16, 18, 13, -15, -11, -20, -2, 46, 49,
	50, 51, 5, 6, 4, 7, -1, 53, 52, 54,
	55, 13, 13, 34, 29, 30, 31, 32, 33, 20,
	34, 23, 4, -4, -12, 16, 8, 9, 10, 11,
	12, 5, 6, -20, 4, 22, 25, 39, -12, -12,
	5, -10, -12, -17, 4, -12, -6, -5, -2, -12,
	-12, -12, -12, -12, -12, -12, -12, -1, 34, 4,
	18, 26, 27, 24, 25, 28, 35, 36, 37, 38,
	42, 43, 40, 41, -12, -10, -17, 16, -12, -12,
	18, -8, -7, -2, 14, 17, 14, 17, 15, 21,
	13, -4, 4, 21, -12, -16, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, 17,
	17, 17, -12, -16, 17, 14, -4, -12, 4, -12,
	-5, 19, 19, 14, 19, -3, -2, -7, 15, 13,
	-13, -12, 18, -12, 47, -14, 48, 14, -16, -12,
	18, -12, 19, 18, -16, 17, -16, 19, 19,
}
var yyDef = [...]int{

	0, -2, 0, 80, 0, 0, 76, 0, 9, 0,
	79, 78, 0, 10, 0, 0, 0, 0, 0, 35,
	0, 0, 12, 15, 18, 0, 5, 1, 2, 3,
	4, 0, 11, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 67, 33, 0, 0, 42, 43, 44, 45,
	46, 12, 15, 49, 50, 0, 0, 0, 36, 0,
	70, 0, 13, 0, 0, 0, 0, 74, 0, 25,
	26, 27, 28, 29, 30, 0, 31, 6, 0, 68,
	9, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 65, 66,
	9, 0, 71, 0, 0, 39, 0, 40, 0, 19,
	0, 73, 67, 20, 32, 0, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 41,
	47, 48, 0, 0, 7, 0, 69, 14, 0, 16,
	75, 0, 23, 0, 37, 0, 8, 72, 0, 77,
	21, 0, 9, 17, 0, 34, 0, 0, 0, 0,
	9, 0, 38, 9, 0, 51, 0, 22, 24,
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
	52, 53, 54, 55, 56, 57,
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
//line parser.y:120
		{
			yyVAL.i = VBool
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:121
		{
			yyVAL.i = VInt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:122
		{
			yyVAL.i = VStr
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:123
		{
			yyVAL.i = VArr
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:127
		{
			yyVAL.n = newType(yyDollar[1].i, yylex)
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:128
		{
			yyVAL.n = addSubtype(yyDollar[1].n, yyDollar[3].i, yylex)
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:132
		{
			yyVAL.n = nil
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:133
		{
			yyVAL.n = yyDollar[1].n
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:137
		{
			yyVAL.n = nil
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:138
		{
			yyVAL.n = yyDollar[1].n
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:139
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:143
		{
			yyVAL.n = nil
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:144
		{
			yyVAL.n = newParam(yyDollar[1].n, yylex)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:145
		{
			yyVAL.n = addParam(yyDollar[1].n, yyDollar[3].n)
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:149
		{
			yyVAL.n = nil
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:150
		{
			yyVAL.n = newContractParam(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:151
		{
			yyVAL.n = addContractParam(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:155
		{
			yyVAL.n = newVarValue(yyDollar[1].s, yylex)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:158
		{
			yyVAL.n = newIndex(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:159
		{
			yyVAL.n = addIndex(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:162
		{
			yyVAL.n = nil
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:163
		{
			yyVAL.n = yyDollar[3].n
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:167
		{
			yyVAL.n = nil
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:168
		{
			yyVAL.n = newElif(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:172
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:173
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD_ASSIGN, yylex)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:174
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB_ASSIGN, yylex)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:175
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL_ASSIGN, yylex)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:176
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV_ASSIGN, yylex)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:177
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD_ASSIGN, yylex)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:178
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:179
		{
			yyVAL.n = newBinary(newVarDecl(yyDollar[1].n, []string{yyDollar[2].s}, yylex), yyDollar[4].n, ASSIGN, yylex)
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:180
		{
			yyVAL.n = newVarDecl(yyDollar[1].n, yyDollar[2].sa, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:181
		{
			yyVAL.n = newIf(yyDollar[2].n, yyDollar[4].n, yyDollar[6].n, yyDollar[7].n, yylex)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:182
		{
			yyVAL.n = newReturn(nil, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:183
		{
			yyVAL.n = newReturn(yyDollar[2].n, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:184
		{
			yyVAL.n = newWhile(yyDollar[2].n, yyDollar[4].n, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:185
		{
			yyVAL.n = newFunc(yyDollar[2].s, yyDollar[3].va, yyDollar[5].n, yyDollar[7].n, yylex)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:188
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:189
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:193
		{
			yyVAL.n = yyDollar[2].n
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:194
		{
			yyVAL.n = newValue(yyDollar[1].i, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:195
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:196
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:197
		{
			yyVAL.n = newValue(true, yylex)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:198
		{
			yyVAL.n = newValue(false, yylex)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:199
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:200
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:201
		{
			yyVAL.n = yyDollar[1].n
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:202
		{
			yyVAL.n = newGetVar(yyDollar[1].s, yylex)
		}
	case 51:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:203
		{
			yyVAL.n = newQuestion(yyDollar[3].n, yyDollar[5].n, yyDollar[7].n, yylex)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:204
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL, yylex)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:205
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV, yylex)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:206
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD, yylex)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:207
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB, yylex)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:208
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD, yylex)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:209
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, AND, yylex)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:210
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, OR, yylex)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:211
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, EQ, yylex)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:212
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, NOT_EQ, yylex)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:213
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LTE, yylex)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:214
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GTE, yylex)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:215
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LT, yylex)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:216
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GT, yylex)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:218
		{
			yyVAL.n = newUnary(yyDollar[2].n, SUB, yylex)
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:219
		{
			yyVAL.n = newUnary(yyDollar[2].n, NOT, yylex)
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:223
		{
			yyVAL.sa = []string{yyDollar[1].s}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:224
		{
			yyVAL.sa = append(yyDollar[1].sa, yyDollar[2].s)
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:228
		{
			yyVAL.va = newVars(yyDollar[1].n, yyDollar[2].sa)
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:232
		{
			yyVAL.va = nil
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:233
		{
			yyVAL.va = yyDollar[1].va
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:234
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:238
		{
			yyVAL.va = newVars(yyDollar[1].n, yyDollar[2].sa)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:242
		{
			yyVAL.va = yyDollar[1].va
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:243
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:247
		{
			yyVAL.va = nil
		}
	case 77:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:248
		{
			yyVAL.va = yyDollar[4].va
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:252
		{
			yyVAL.n = newBlock(yyDollar[1].va, yyDollar[2].n, yylex)
		}
	case 79:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:258
		{
			yyVAL.n = newContract(yyDollar[2].s, yyDollar[5].n, yylex)
			setResult(yylex, yyVAL.n)
		}
	}
	goto yystack /* stack new state and value */
}
