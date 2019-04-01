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
const FLOAT = 57351
const STRING = 57352
const QSTRING = 57353
const TRUE = 57354
const FALSE = 57355
const NEWLINE = 57356
const COMMA = 57357
const COLON = 57358
const LPAREN = 57359
const RPAREN = 57360
const LBRACE = 57361
const RBRACE = 57362
const LBRACKET = 57363
const RBRACKET = 57364
const QUESTION = 57365
const DOUBLEDOT = 57366
const DOT = 57367
const ADD = 57368
const SUB = 57369
const MUL = 57370
const DIV = 57371
const MOD = 57372
const ADD_ASSIGN = 57373
const SUB_ASSIGN = 57374
const MUL_ASSIGN = 57375
const DIV_ASSIGN = 57376
const MOD_ASSIGN = 57377
const ASSIGN = 57378
const AND = 57379
const OR = 57380
const EQ = 57381
const NOT_EQ = 57382
const NOT = 57383
const LT = 57384
const GT = 57385
const LTE = 57386
const GTE = 57387
const BREAK = 57388
const CONTINUE = 57389
const DATA = 57390
const CONTRACT = 57391
const IF = 57392
const ELIF = 57393
const ELSE = 57394
const RETURN = 57395
const WHILE = 57396
const FUNC = 57397
const FOR = 57398
const IN = 57399
const T_INT = 57400
const T_BOOL = 57401
const T_STR = 57402
const T_ARR = 57403
const T_MAP = 57404
const T_FLOAT = 57405
const UNARYMINUS = 57406
const UNARYNOT = 57407

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"CALL",
	"CALLCONTRACT",
	"INDEX",
	"INT",
	"FLOAT",
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
	"DOUBLEDOT",
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
	"BREAK",
	"CONTINUE",
	"DATA",
	"CONTRACT",
	"IF",
	"ELIF",
	"ELSE",
	"RETURN",
	"WHILE",
	"FUNC",
	"FOR",
	"IN",
	"T_INT",
	"T_BOOL",
	"T_STR",
	"T_ARR",
	"T_MAP",
	"T_FLOAT",
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

const yyLast = 900

var yyAct = [...]int{

	69, 59, 114, 75, 48, 68, 122, 70, 175, 9,
	17, 185, 187, 16, 162, 2, 11, 86, 46, 49,
	29, 44, 65, 66, 91, 92, 89, 90, 93, 73,
	31, 30, 32, 33, 34, 35, 45, 44, 10, 77,
	78, 79, 80, 81, 82, 83, 84, 76, 121, 109,
	196, 102, 31, 30, 32, 33, 34, 35, 89, 90,
	93, 126, 107, 103, 110, 111, 104, 85, 39, 40,
	41, 42, 43, 38, 148, 146, 176, 180, 115, 149,
	147, 125, 46, 47, 12, 154, 118, 128, 153, 145,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 141, 142, 116, 46, 129, 144, 178, 118, 116,
	151, 119, 117, 172, 5, 150, 120, 156, 124, 37,
	155, 158, 159, 36, 6, 3, 165, 67, 161, 152,
	87, 16, 160, 157, 71, 76, 72, 60, 57, 58,
	28, 51, 52, 53, 54, 55, 56, 164, 4, 106,
	50, 166, 61, 105, 16, 1, 62, 171, 7, 14,
	63, 186, 177, 15, 170, 115, 8, 113, 179, 74,
	169, 0, 0, 181, 64, 183, 184, 0, 0, 189,
	0, 0, 0, 0, 16, 0, 195, 0, 0, 197,
	182, 198, 0, 16, 0, 0, 0, 191, 27, 24,
	25, 28, 16, 16, 0, 16, 0, 0, 13, 16,
	200, 201, 0, 203, 209, 0, 0, 0, 0, 207,
	27, 24, 25, 28, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 0, 0, 208, 0, 0, 0,
	19, 20, 0, 0, 18, 0, 0, 21, 22, 23,
	26, 0, 31, 30, 32, 33, 34, 35, 0, 0,
	0, 0, 19, 20, 0, 0, 18, 0, 0, 21,
	22, 23, 26, 0, 31, 30, 32, 33, 34, 35,
	27, 24, 25, 28, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 0, 0, 206, 0, 27, 24,
	25, 28, 0, 0, 0, 0, 0, 0, 13, 0,
	0, 0, 0, 0, 205, 0, 0, 27, 24, 25,
	28, 0, 19, 20, 0, 0, 18, 13, 0, 21,
	22, 23, 26, 199, 31, 30, 32, 33, 34, 35,
	19, 20, 0, 0, 18, 0, 0, 21, 22, 23,
	26, 0, 31, 30, 32, 33, 34, 35, 0, 19,
	20, 0, 0, 18, 0, 0, 21, 22, 23, 26,
	0, 31, 30, 32, 33, 34, 35, 27, 24, 25,
	28, 0, 0, 0, 0, 0, 0, 13, 0, 0,
	0, 0, 0, 192, 0, 27, 24, 25, 28, 0,
	0, 0, 0, 0, 0, 13, 0, 0, 0, 0,
	0, 168, 0, 0, 27, 24, 25, 28, 0, 19,
	20, 0, 0, 18, 13, 0, 21, 22, 23, 26,
	163, 31, 30, 32, 33, 34, 35, 19, 20, 0,
	0, 18, 0, 0, 21, 22, 23, 26, 0, 31,
	30, 32, 33, 34, 35, 0, 19, 20, 0, 0,
	18, 0, 0, 21, 22, 23, 26, 0, 31, 30,
	32, 33, 34, 35, 27, 24, 25, 28, 173, 0,
	0, 0, 0, 174, 13, 91, 92, 89, 90, 93,
	0, 0, 0, 0, 0, 0, 94, 95, 96, 97,
	0, 100, 101, 98, 99, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 19, 20, 0, 0,
	18, 0, 0, 21, 22, 23, 26, 204, 31, 30,
	32, 33, 34, 35, 0, 91, 92, 89, 90, 93,
	0, 0, 0, 0, 0, 0, 94, 95, 96, 97,
	202, 100, 101, 98, 99, 0, 0, 91, 92, 89,
	90, 93, 0, 0, 0, 0, 0, 0, 94, 95,
	96, 97, 194, 100, 101, 98, 99, 0, 0, 91,
	92, 89, 90, 93, 0, 0, 0, 0, 0, 0,
	94, 95, 96, 97, 193, 100, 101, 98, 99, 0,
	0, 91, 92, 89, 90, 93, 0, 0, 0, 0,
	190, 0, 94, 95, 96, 97, 0, 100, 101, 98,
	99, 91, 92, 89, 90, 93, 0, 0, 0, 0,
	167, 0, 94, 95, 96, 97, 0, 100, 101, 98,
	99, 91, 92, 89, 90, 93, 0, 0, 0, 0,
	0, 0, 94, 95, 96, 97, 143, 100, 101, 98,
	99, 0, 0, 0, 91, 92, 89, 90, 93, 0,
	0, 0, 0, 0, 0, 94, 95, 96, 97, 0,
	100, 101, 98, 99, 127, 0, 0, 0, 91, 92,
	89, 90, 93, 0, 0, 0, 0, 0, 0, 94,
	95, 96, 97, 0, 100, 101, 98, 99, 123, 0,
	0, 0, 91, 92, 89, 90, 93, 0, 0, 0,
	0, 0, 0, 94, 95, 96, 97, 112, 100, 101,
	98, 99, 0, 0, 91, 92, 89, 90, 93, 0,
	0, 0, 0, 0, 0, 94, 95, 96, 97, 88,
	100, 101, 98, 99, 0, 0, 91, 92, 89, 90,
	93, 0, 0, 0, 0, 0, 0, 94, 95, 96,
	97, 0, 100, 101, 98, 99, 91, 92, 89, 90,
	93, 0, 0, 0, 0, 0, 0, 94, 95, 96,
	97, 0, 100, 101, 98, 99, 60, 57, 58, 28,
	51, 52, 53, 54, 55, 56, 188, 0, 0, 50,
	0, 61, 0, 0, 0, 62, 0, 0, 0, 63,
	0, 0, 91, 92, 89, 90, 93, 0, 0, 0,
	0, 0, 0, 64, 95, 96, 97, 0, 100, 101,
	98, 99, 91, 92, 89, 90, 93, 0, 0, 0,
	0, 0, 0, 0, 0, 96, 97, 0, 100, 101,
	98, 99, 60, 57, 58, 28, 51, 52, 108, 54,
	55, 56, 0, 0, 0, 50, 0, 61, 0, 0,
	0, 62, 0, 0, 0, 63, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
}
var yyPact = [...]int{

	-34, 111, 144, -1000, 95, 110, -39, 18, -1000, 65,
	-1000, 470, 109, -1000, 105, 37, 0, 79, 133, -1000,
	-1000, 133, 133, 122, 133, 130, 132, -1000, 133, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -28, -1000, 133, 133,
	133, 133, 133, 133, 133, 133, -28, -19, 126, 730,
	133, -1000, -1000, -1000, -1000, -1000, -1000, 133, 130, 16,
	-1000, 858, 32, 133, 133, 750, 708, -28, 94, 750,
	93, 100, -9, 686, 104, -1000, 57, 750, 750, 750,
	750, 750, 750, 662, 750, -1000, 133, -1000, -1000, 133,
	133, 133, 133, 133, 133, 133, 133, 133, 133, 133,
	133, 133, 638, 88, 71, 60, 59, 750, 99, 133,
	-1000, -1000, -1000, 70, -1000, 57, 133, -1000, 129, -1000,
	133, 133, 128, -1000, -6, 126, -1000, -1000, 750, 410,
	-1000, -1000, 30, 30, -1000, 796, 816, -2, -2, -2,
	-2, -2, -2, -1000, -1000, -1000, 133, -1000, 116, -1000,
	133, 615, 391, -28, -28, 126, 750, 97, 750, 459,
	-49, -1000, 62, -1000, 750, 91, 750, 133, -1000, 58,
	-7, -1000, 133, -1000, 133, 133, -1000, -40, 792, 595,
	-1000, 750, 373, 575, 553, 133, -1000, 31, 133, 750,
	133, 313, -1000, -1000, -1000, 531, -1000, 750, 509, -1000,
	294, 276, -1000, 216, -1000, -1000, -1000, 194, -1000, -1000,
}
var yyPgo = [...]int{

	0, 20, 10, 170, 4, 3, 169, 2, 167, 166,
	5, 163, 0, 162, 161, 159, 16, 7, 158, 155,
	1, 153, 149,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 3,
	3, 16, 16, 16, 10, 10, 10, 17, 17, 17,
	11, 20, 20, 14, 14, 13, 13, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 21, 21,
	22, 22, 22, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 4, 4, 7, 8, 8, 8, 5, 6,
	6, 9, 9, 18, 19, 19,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 3, 0,
	1, 0, 2, 3, 0, 1, 3, 0, 3, 5,
	1, 3, 4, 0, 4, 0, 6, 3, 3, 3,
	3, 3, 3, 3, 4, 2, 7, 1, 1, 1,
	2, 5, 8, 3, 3, 7, 9, 9, 1, 3,
	3, 6, 5, 3, 1, 1, 1, 1, 1, 1,
	3, 3, 1, 1, 3, 3, 8, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 1, 2, 2, 0, 1, 3, 2, 1,
	3, 0, 7, 2, 6, 2,
}
var yyChk = [...]int{

	-1000, -19, 49, 14, 4, 19, 14, -18, -9, 48,
	20, -16, 19, 14, -15, -11, -20, -2, 50, 46,
	47, 53, 54, 55, 5, 6, 56, 4, 7, -1,
	59, 58, 60, 61, 62, 63, 14, 14, 36, 31,
	32, 33, 34, 35, 21, 36, 25, 4, -4, -12,
	17, 8, 9, 10, 11, 12, 13, 5, 6, -20,
	4, 19, 23, 27, 41, -12, -12, 5, -10, -12,
	-17, 4, 4, -12, -6, -5, -2, -12, -12, -12,
	-12, -12, -12, -12, -12, -1, 36, 4, 19, 28,
	29, 26, 27, 30, 37, 38, 39, 40, 44, 45,
	42, 43, -12, -10, -17, -21, -22, -12, 10, 17,
	-12, -12, 19, -8, -7, -2, 15, 18, 15, 18,
	16, 57, 15, 22, 14, -4, 4, 22, -12, -16,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, 18, 18, 18, 15, 20, 15, 20,
	16, -12, -16, 18, 15, -4, -12, 4, -12, -12,
	4, -5, 20, 20, -12, 10, -12, 15, 20, -3,
	-2, -7, 16, 19, 24, 57, 14, -13, 16, -12,
	19, -12, -16, -12, -12, 51, -14, 52, 14, -12,
	15, -16, 20, 19, 19, -12, 19, -12, -12, 20,
	-16, -16, 19, -16, 18, 20, 20, -16, 20, 20,
}
var yyDef = [...]int{

	0, -2, 0, 95, 0, 0, 91, 0, 11, 0,
	94, 93, 0, 12, 0, 0, 0, 0, 0, 37,
	38, 39, 0, 0, 14, 17, 0, 20, 0, 7,
	1, 2, 3, 4, 5, 6, 0, 13, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 82, 35, 0,
	0, 54, 55, 56, 57, 58, 59, 14, 17, 62,
	63, 0, 0, 0, 0, 40, 0, 85, 0, 15,
	0, 0, 0, 0, 0, 89, 0, 27, 28, 29,
	30, 31, 32, 0, 33, 8, 0, 83, 11, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 48, 56, 0,
	80, 81, 11, 0, 86, 0, 0, 43, 0, 44,
	0, 0, 0, 21, 0, 88, 82, 22, 34, 0,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	77, 78, 79, 53, 60, 61, 0, 64, 0, 65,
	0, 0, 0, 9, 0, 84, 16, 0, 18, 0,
	0, 90, 0, 25, 49, 0, 50, 0, 41, 0,
	10, 87, 0, 11, 0, 0, 92, 23, 0, 0,
	11, 19, 0, 0, 0, 0, 36, 0, 0, 52,
	0, 0, 45, 11, 11, 0, 11, 51, 0, 42,
	0, 0, 11, 0, 66, 47, 46, 0, 24, 26,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65,
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
//line parser.y:129
		{
			yyVAL.i = VBool
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:130
		{
			yyVAL.i = VInt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:131
		{
			yyVAL.i = VStr
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:132
		{
			yyVAL.i = VArr
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:133
		{
			yyVAL.i = VMap
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:134
		{
			yyVAL.i = VFloat
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:138
		{
			yyVAL.n = newType(yyDollar[1].i, yylex)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:139
		{
			yyVAL.n = addSubtype(yyDollar[1].n, yyDollar[3].i, yylex)
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:143
		{
			yyVAL.n = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:144
		{
			yyVAL.n = yyDollar[1].n
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:148
		{
			yyVAL.n = nil
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:149
		{
			yyVAL.n = yyDollar[1].n
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:150
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:154
		{
			yyVAL.n = nil
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:155
		{
			yyVAL.n = newParam(yyDollar[1].n, yylex)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:156
		{
			yyVAL.n = addParam(yyDollar[1].n, yyDollar[3].n)
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:160
		{
			yyVAL.n = nil
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:161
		{
			yyVAL.n = newContractParam(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:162
		{
			yyVAL.n = addContractParam(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:166
		{
			yyVAL.n = newVarValue(yyDollar[1].s, yylex)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:169
		{
			yyVAL.n = newIndex(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:170
		{
			yyVAL.n = addIndex(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:173
		{
			yyVAL.n = nil
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:174
		{
			yyVAL.n = yyDollar[3].n
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:178
		{
			yyVAL.n = nil
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:179
		{
			yyVAL.n = newElif(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:183
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:184
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD_ASSIGN, yylex)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:185
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB_ASSIGN, yylex)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:186
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL_ASSIGN, yylex)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:187
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV_ASSIGN, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:188
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD_ASSIGN, yylex)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:189
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:190
		{
			yyVAL.n = newBinary(newVarDecl(yyDollar[1].n, []string{yyDollar[2].s}, yylex), yyDollar[4].n, ASSIGN, yylex)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:191
		{
			yyVAL.n = newVarDecl(yyDollar[1].n, yyDollar[2].sa, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:192
		{
			yyVAL.n = newIf(yyDollar[2].n, yyDollar[4].n, yyDollar[6].n, yyDollar[7].n, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:193
		{
			yyVAL.n = newBreak(yylex)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:194
		{
			yyVAL.n = newContinue(yylex)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:195
		{
			yyVAL.n = newReturn(nil, yylex)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:196
		{
			yyVAL.n = newReturn(yyDollar[2].n, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:197
		{
			yyVAL.n = newWhile(yyDollar[2].n, yyDollar[4].n, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:198
		{
			yyVAL.n = newFunc(yyDollar[2].s, yyDollar[3].va, yyDollar[5].n, yyDollar[7].n, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:201
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:202
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 45:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:203
		{
			yyVAL.n = newFor(yyDollar[2].s, yyDollar[4].n, yyDollar[6].n, yylex)
		}
	case 46:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.y:204
		{
			yyVAL.n = newForAll(yyDollar[2].s, yyDollar[4].s, yyDollar[6].n, yyDollar[8].n, yylex)
		}
	case 47:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.y:205
		{
			yyVAL.n = newForInt(yyDollar[2].s, yyDollar[4].n, yyDollar[6].n, yyDollar[8].n, yylex)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:209
		{
			yyVAL.n = newArray(yyDollar[1].n, yylex)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:210
		{
			yyVAL.n = appendArray(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:214
		{
			yyVAL.n = newMap(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:215
		{
			yyVAL.n = appendMap(yyDollar[1].n, yyDollar[3].s, yyDollar[6].n, yylex)
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:216
		{
			yyVAL.n = appendMap(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n, yylex)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:220
		{
			yyVAL.n = yyDollar[2].n
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:221
		{
			yyVAL.n = newValue(yyDollar[1].i, yylex)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:222
		{
			yyVAL.n = newValue(yyDollar[1].f, yylex)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:223
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:224
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:225
		{
			yyVAL.n = newValue(true, yylex)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:226
		{
			yyVAL.n = newValue(false, yylex)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:227
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:228
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:229
		{
			yyVAL.n = yyDollar[1].n
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:230
		{
			yyVAL.n = newGetVar(yyDollar[1].s, yylex)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:231
		{
			yyVAL.n = yyDollar[2].n
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:232
		{
			yyVAL.n = yyDollar[2].n
		}
	case 66:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:233
		{
			yyVAL.n = newQuestion(yyDollar[3].n, yyDollar[5].n, yyDollar[7].n, yylex)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:234
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL, yylex)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:235
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV, yylex)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:236
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD, yylex)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:237
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB, yylex)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:238
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD, yylex)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:239
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, AND, yylex)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:240
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, OR, yylex)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:241
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, EQ, yylex)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:242
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, NOT_EQ, yylex)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:243
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LTE, yylex)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:244
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GTE, yylex)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:245
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LT, yylex)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:246
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GT, yylex)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:248
		{
			yyVAL.n = newUnary(yyDollar[2].n, SUB, yylex)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:249
		{
			yyVAL.n = newUnary(yyDollar[2].n, NOT, yylex)
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:253
		{
			yyVAL.sa = []string{yyDollar[1].s}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:254
		{
			yyVAL.sa = append(yyDollar[1].sa, yyDollar[2].s)
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:258
		{
			yyVAL.va = newVars(yyDollar[1].n, yyDollar[2].sa)
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:262
		{
			yyVAL.va = nil
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:263
		{
			yyVAL.va = yyDollar[1].va
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:264
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:268
		{
			yyVAL.va = newVars(yyDollar[1].n, yyDollar[2].sa)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:272
		{
			yyVAL.va = yyDollar[1].va
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:273
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:277
		{
			yyVAL.va = nil
		}
	case 92:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:278
		{
			yyVAL.va = yyDollar[4].va
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:282
		{
			yyVAL.n = newBlock(yyDollar[1].va, yyDollar[2].n, yylex)
		}
	case 94:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:288
		{
			yyVAL.n = newContract(yyDollar[2].s, yyDollar[5].n, yylex)
			setResult(yylex, yyVAL.n)
		}
	}
	goto yystack /* stack new state and value */
}
