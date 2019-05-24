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
const ENV = 57347
const CALL = 57348
const CALLCONTRACT = 57349
const INDEX = 57350
const INT = 57351
const FLOAT = 57352
const STRING = 57353
const QSTRING = 57354
const TRUE = 57355
const FALSE = 57356
const NEWLINE = 57357
const COMMA = 57358
const COLON = 57359
const LPAREN = 57360
const RPAREN = 57361
const OBJ = 57362
const LBRACE = 57363
const RBRACE = 57364
const LBRACKET = 57365
const RBRACKET = 57366
const QUESTION = 57367
const DOUBLEDOT = 57368
const DOT = 57369
const ADD = 57370
const SUB = 57371
const MUL = 57372
const DIV = 57373
const MOD = 57374
const ADD_ASSIGN = 57375
const SUB_ASSIGN = 57376
const MUL_ASSIGN = 57377
const DIV_ASSIGN = 57378
const MOD_ASSIGN = 57379
const ASSIGN = 57380
const AND = 57381
const OR = 57382
const EQ = 57383
const NOT_EQ = 57384
const NOT = 57385
const LT = 57386
const GT = 57387
const LTE = 57388
const GTE = 57389
const BREAK = 57390
const CONTINUE = 57391
const DATA = 57392
const CONTRACT = 57393
const IF = 57394
const ELIF = 57395
const ELSE = 57396
const RETURN = 57397
const WHILE = 57398
const FUNC = 57399
const FOR = 57400
const IN = 57401
const SWITCH = 57402
const CASE = 57403
const READ = 57404
const DEFAULT = 57405
const T_INT = 57406
const T_BOOL = 57407
const T_STR = 57408
const T_ARR = 57409
const T_MAP = 57410
const T_FLOAT = 57411
const T_MONEY = 57412
const T_OBJECT = 57413
const T_BYTES = 57414
const T_FILE = 57415
const UNARYMINUS = 57416
const UNARYNOT = 57417

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"ENV",
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
	"OBJ",
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
	"SWITCH",
	"CASE",
	"READ",
	"DEFAULT",
	"T_INT",
	"T_BOOL",
	"T_STR",
	"T_ARR",
	"T_MAP",
	"T_FLOAT",
	"T_MONEY",
	"T_OBJECT",
	"T_BYTES",
	"T_FILE",
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

const yyLast = 1123

var yyAct = [...]int{

	80, 105, 81, 13, 56, 79, 21, 190, 108, 129,
	86, 33, 185, 137, 187, 6, 74, 218, 20, 46,
	256, 258, 2, 75, 11, 184, 76, 77, 124, 70,
	176, 183, 141, 84, 35, 34, 36, 37, 38, 39,
	40, 41, 42, 43, 71, 89, 90, 93, 102, 72,
	70, 87, 163, 72, 159, 72, 136, 242, 104, 12,
	103, 110, 252, 113, 114, 115, 116, 117, 118, 119,
	120, 121, 122, 35, 34, 36, 37, 38, 39, 40,
	41, 42, 43, 159, 123, 130, 269, 73, 112, 249,
	143, 144, 145, 146, 147, 148, 149, 150, 151, 152,
	153, 154, 155, 250, 140, 91, 92, 89, 90, 93,
	72, 251, 237, 168, 165, 222, 94, 95, 96, 97,
	166, 100, 101, 98, 99, 170, 14, 7, 162, 163,
	171, 172, 177, 216, 161, 164, 179, 180, 217, 159,
	91, 92, 89, 90, 93, 160, 87, 175, 234, 224,
	182, 94, 95, 96, 97, 223, 100, 101, 98, 99,
	133, 207, 131, 248, 206, 247, 200, 200, 208, 215,
	205, 91, 92, 89, 90, 93, 20, 20, 167, 174,
	213, 130, 173, 135, 214, 220, 110, 133, 131, 279,
	158, 157, 225, 219, 221, 65, 66, 67, 68, 69,
	64, 139, 227, 45, 226, 228, 230, 78, 133, 200,
	235, 134, 231, 82, 131, 44, 238, 132, 240, 241,
	239, 8, 3, 107, 125, 181, 243, 189, 200, 200,
	106, 244, 245, 178, 188, 254, 83, 4, 5, 229,
	109, 259, 1, 9, 20, 17, 263, 186, 20, 31,
	142, 28, 29, 32, 266, 200, 267, 268, 265, 16,
	15, 257, 236, 19, 20, 271, 272, 282, 20, 10,
	128, 85, 212, 276, 0, 0, 20, 20, 0, 280,
	31, 20, 28, 29, 32, 20, 0, 0, 0, 0,
	0, 15, 0, 23, 24, 0, 0, 22, 281, 0,
	25, 26, 27, 30, 0, 18, 0, 0, 0, 35,
	34, 36, 37, 38, 39, 40, 41, 42, 43, 31,
	0, 28, 29, 32, 23, 24, 0, 0, 22, 0,
	15, 25, 26, 27, 30, 0, 18, 278, 0, 0,
	35, 34, 36, 37, 38, 39, 40, 41, 42, 43,
	31, 0, 28, 29, 32, 0, 0, 0, 0, 0,
	0, 15, 0, 23, 24, 0, 0, 22, 277, 0,
	25, 26, 27, 30, 0, 18, 0, 0, 0, 35,
	34, 36, 37, 38, 39, 40, 41, 42, 43, 31,
	0, 28, 29, 32, 23, 24, 0, 0, 22, 0,
	15, 25, 26, 27, 30, 0, 18, 273, 0, 0,
	35, 34, 36, 37, 38, 39, 40, 41, 42, 43,
	31, 0, 28, 29, 32, 0, 0, 0, 0, 0,
	0, 15, 0, 23, 24, 0, 0, 22, 270, 0,
	25, 26, 27, 30, 0, 18, 0, 0, 0, 35,
	34, 36, 37, 38, 39, 40, 41, 42, 43, 31,
	0, 28, 29, 32, 23, 24, 0, 0, 22, 0,
	15, 25, 26, 27, 30, 0, 18, 264, 0, 0,
	35, 34, 36, 37, 38, 39, 40, 41, 42, 43,
	31, 0, 28, 29, 32, 0, 0, 0, 0, 0,
	0, 15, 0, 23, 24, 0, 0, 22, 260, 0,
	25, 26, 27, 30, 0, 18, 0, 0, 0, 35,
	34, 36, 37, 38, 39, 40, 41, 42, 43, 31,
	0, 28, 29, 32, 23, 24, 0, 0, 22, 0,
	15, 25, 26, 27, 30, 0, 18, 211, 0, 0,
	35, 34, 36, 37, 38, 39, 40, 41, 42, 43,
	31, 0, 28, 29, 32, 0, 0, 0, 0, 0,
	0, 15, 0, 23, 24, 0, 0, 22, 210, 0,
	25, 26, 27, 30, 0, 18, 0, 0, 0, 35,
	34, 36, 37, 38, 39, 40, 41, 42, 43, 31,
	0, 28, 29, 32, 23, 24, 0, 0, 22, 0,
	15, 25, 26, 27, 30, 0, 18, 0, 0, 0,
	35, 34, 36, 37, 38, 39, 40, 41, 42, 43,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 23, 24, 0, 0, 22, 0, 0,
	25, 26, 27, 30, 0, 18, 0, 0, 0, 35,
	34, 36, 37, 38, 39, 40, 41, 42, 43, 275,
	0, 0, 0, 0, 0, 0, 91, 92, 89, 90,
	93, 0, 0, 0, 0, 0, 0, 94, 95, 96,
	97, 274, 100, 101, 98, 99, 0, 0, 0, 0,
	91, 92, 89, 90, 93, 0, 0, 0, 0, 0,
	0, 94, 95, 96, 97, 262, 100, 101, 98, 99,
	0, 0, 91, 92, 89, 90, 93, 0, 0, 0,
	0, 0, 0, 94, 95, 96, 97, 261, 100, 101,
	98, 99, 0, 0, 91, 92, 89, 90, 93, 0,
	0, 0, 255, 0, 0, 94, 95, 96, 97, 0,
	100, 101, 98, 99, 91, 92, 89, 90, 93, 0,
	0, 0, 0, 0, 0, 94, 95, 96, 97, 0,
	100, 101, 98, 99, 58, 57, 54, 55, 32, 48,
	49, 50, 51, 52, 53, 253, 0, 0, 47, 0,
	59, 60, 0, 0, 246, 61, 0, 0, 0, 62,
	0, 0, 0, 91, 92, 89, 90, 93, 0, 0,
	0, 209, 0, 63, 94, 95, 96, 97, 0, 100,
	101, 98, 99, 91, 92, 89, 90, 93, 0, 0,
	0, 0, 0, 0, 94, 95, 96, 97, 0, 100,
	101, 98, 99, 169, 0, 0, 0, 91, 92, 89,
	90, 93, 0, 0, 0, 0, 0, 0, 94, 95,
	96, 97, 156, 100, 101, 98, 99, 0, 0, 0,
	0, 91, 92, 89, 90, 93, 0, 0, 0, 0,
	0, 0, 94, 95, 96, 97, 0, 100, 101, 98,
	99, 138, 0, 0, 0, 91, 92, 89, 90, 93,
	0, 0, 0, 0, 0, 0, 94, 95, 96, 97,
	127, 100, 101, 98, 99, 0, 0, 91, 92, 89,
	90, 93, 0, 0, 0, 0, 0, 0, 94, 95,
	96, 97, 126, 100, 101, 98, 99, 0, 0, 91,
	92, 89, 90, 93, 0, 0, 88, 0, 0, 0,
	94, 95, 96, 97, 0, 100, 101, 98, 99, 91,
	92, 89, 90, 93, 0, 0, 0, 0, 0, 0,
	94, 95, 96, 97, 0, 100, 101, 98, 99, 58,
	57, 54, 55, 32, 48, 49, 50, 51, 52, 53,
	0, 0, 0, 47, 0, 59, 60, 0, 0, 0,
	61, 0, 0, 0, 62, 58, 57, 54, 55, 32,
	48, 49, 111, 51, 52, 53, 0, 0, 63, 47,
	0, 59, 60, 0, 0, 0, 61, 0, 0, 0,
	62, 0, 0, 91, 92, 89, 90, 93, 0, 0,
	0, 0, 0, 0, 63, 95, 96, 97, 0, 100,
	101, 98, 99, 91, 92, 89, 90, 93, 0, 0,
	0, 0, 0, 0, 0, 0, 96, 97, 0, 100,
	101, 98, 99, 202, 201, 198, 199, 32, 192, 193,
	194, 195, 196, 197, 0, 0, 0, 191, 0, 0,
	203, 0, 204, 233, 201, 198, 199, 32, 192, 193,
	232, 195, 196, 197, 0, 0, 0, 191, 0, 0,
	203, 0, 204,
}
var yyPact = [...]int{

	-29, 207, 233, -1000, -47, 106, -1000, 206, -26, 37,
	-1000, 105, -1000, 595, 200, -1000, -1000, 188, 985, 162,
	6, 83, 985, -1000, -1000, 985, 985, 201, 985, 209,
	232, -1000, 985, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -30, -1000, 941, 985, -1000, -1000,
	-1000, -1000, -1000, -1000, 985, 209, 27, -1000, -1000, 219,
	1011, 70, 985, 985, 985, 985, 985, 985, 985, 985,
	985, 985, -30, -10, 220, 921, 77, 899, -30, 198,
	77, 192, 166, -3, 877, 186, -1000, 28, -1000, 985,
	985, 985, 985, 985, 985, 985, 985, 985, 985, 985,
	985, 985, 853, 172, 171, 123, 117, 111, 113, 98,
	77, 161, 985, -1000, -1000, 77, 77, 77, 77, 77,
	77, 829, 77, -1000, 985, -1000, -1000, -1000, 163, -1000,
	26, 985, -1000, 229, -1000, 985, 985, 221, -1000, 9,
	220, -13, -49, -1000, -1000, 15, 15, -1000, 1015, 1035,
	143, 143, 143, 143, 143, 143, -1000, -1000, -1000, 223,
	-1000, 1079, 1079, 985, -1000, 150, -1000, 985, 805, -1000,
	77, 556, 525, -30, -30, 220, -1000, 77, 152, 77,
	112, -42, -1000, 178, 985, 985, -1000, 94, 138, 132,
	-1000, 985, -1000, -1000, -1000, -1000, -1000, -1000, 985, 209,
	27, -1000, -1000, 219, 1099, -1000, 77, 131, 77, 985,
	-1000, -1000, 91, 22, -1000, 985, -1000, 985, 985, -1000,
	77, 36, -1000, 1079, 1079, 785, 146, 144, 67, 87,
	38, -1000, 117, 111, 780, 736, -33, -1000, 77, 486,
	716, 694, -1000, 455, -1000, -1000, -1000, -1000, -1000, -1000,
	1079, -1000, -1000, 985, 77, 985, 985, -1000, 65, 416,
	-1000, -1000, -1000, 385, -1000, -1000, 77, 672, 648, -1000,
	-1000, 346, 315, 174, -1000, -1000, 276, -1000, -1000, -1000,
	245, -1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 6, 272, 16, 10, 271, 9, 270, 269,
	5, 263, 0, 262, 261, 259, 250, 247, 245, 3,
	2, 243, 242, 4, 8, 240, 7, 1, 239, 238,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 3, 3, 19, 19, 19, 19, 10,
	10, 10, 20, 20, 20, 11, 23, 23, 14, 14,
	13, 13, 16, 16, 17, 17, 15, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 24, 24,
	25, 25, 25, 27, 27, 27, 27, 28, 28, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 4, 4, 7, 8, 8,
	8, 5, 5, 6, 6, 9, 9, 21, 29, 29,
	22, 22,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 0, 1, 0, 2, 2, 3, 0,
	1, 3, 0, 3, 5, 1, 3, 4, 0, 4,
	0, 6, 0, 7, 0, 4, 5, 3, 3, 3,
	3, 3, 3, 3, 4, 2, 7, 1, 1, 1,
	2, 5, 8, 3, 3, 7, 9, 9, 1, 3,
	3, 6, 5, 3, 3, 5, 5, 1, 3, 3,
	1, 1, 1, 1, 1, 1, 3, 3, 1, 1,
	1, 3, 3, 3, 3, 1, 1, 1, 1, 1,
	1, 3, 3, 1, 1, 1, 3, 3, 3, 8,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 1, 2, 2, 0, 1,
	3, 2, 4, 1, 3, 0, 7, 2, 0, 1,
	7, 2,
}
var yyChk = [...]int{

	-1000, -22, 51, 15, 4, -29, 62, 21, 15, -21,
	-9, 50, 22, -19, 21, 15, -15, -18, 60, -11,
	-23, -2, 52, 48, 49, 55, 56, 57, 6, 7,
	58, 4, 8, -1, 65, 64, 66, 67, 68, 69,
	70, 71, 72, 73, 15, 15, -12, 18, 9, 10,
	11, 12, 13, 14, 6, 7, -23, 5, 4, 20,
	21, 25, 29, 43, 38, 33, 34, 35, 36, 37,
	23, 38, 27, 4, -4, -12, -12, -12, 6, -10,
	-12, -20, 4, 4, -12, -6, -5, -2, 15, 30,
	31, 28, 29, 32, 39, 40, 41, 42, 46, 47,
	44, 45, -12, -10, -20, -27, 11, 4, -24, -25,
	-12, 11, 18, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -1, 38, 4, 21, 21, -8, -7,
	-2, 16, 19, 16, 19, 17, 59, 16, 24, 15,
	-4, 4, -16, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, 19, 19, 19, 16,
	22, 17, 17, 16, 22, 16, 22, 17, -12, 24,
	-12, -19, -19, 19, 16, -4, 4, -12, 4, -12,
	-12, 4, -5, 22, 38, 61, -17, 63, 11, 4,
	-26, 18, 9, 10, 11, 12, 13, 14, 6, 7,
	-23, 5, 4, 21, 23, -26, -12, 11, -12, 16,
	22, 22, -3, -2, -7, 17, 21, 26, 59, 15,
	-12, -24, 21, 17, 17, -12, -10, -20, -27, -28,
	-27, -26, 11, 4, 17, -12, -13, 21, -12, -19,
	-12, -12, 21, -19, -26, -26, 19, 19, 19, 22,
	16, 24, 24, 15, -12, 16, 53, -14, 54, -19,
	22, 21, 21, -19, 22, -26, -12, -12, -12, 21,
	22, -19, -19, 22, 19, 21, -19, 22, 22, 15,
	-19, 22, 22,
}
var yyDef = [...]int{

	0, -2, 0, 131, 128, 0, 129, 0, 125, 0,
	15, 0, 130, 127, 0, 16, 17, 0, 0, 0,
	0, 0, 0, 47, 48, 49, 0, 0, 19, 22,
	0, 25, 0, 11, 1, 2, 3, 4, 5, 6,
	7, 8, 9, 10, 0, 18, 0, 0, 85, 86,
	87, 88, 89, 90, 19, 22, 93, 94, 95, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 115, 45, 0, 50, 0, 118, 0,
	20, 0, 0, 0, 0, 0, 123, 0, 32, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	58, 87, 0, 113, 114, 37, 38, 39, 40, 41,
	42, 0, 43, 12, 0, 116, 15, 15, 0, 119,
	0, 0, 53, 0, 54, 0, 0, 0, 26, 0,
	121, 115, 34, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 84, 91, 92, 0,
	96, 0, 0, 0, 97, 0, 98, 0, 0, 27,
	44, 0, 0, 13, 0, 117, 115, 21, 0, 23,
	0, 0, 124, 0, 0, 0, 36, 0, 0, 0,
	63, 0, 70, 71, 72, 73, 74, 75, 19, 22,
	78, 79, 80, 0, 0, 64, 59, 0, 60, 0,
	30, 51, 0, 14, 120, 0, 15, 0, 0, 126,
	122, 0, 15, 0, 0, 0, 0, 0, 0, 0,
	0, 67, 72, 80, 0, 0, 28, 15, 24, 0,
	0, 0, 15, 0, 65, 66, 69, 76, 77, 81,
	0, 82, 83, 0, 62, 0, 0, 46, 0, 0,
	55, 15, 15, 0, 35, 68, 61, 0, 0, 15,
	52, 0, 0, 0, 99, 15, 0, 57, 56, 33,
	0, 29, 31,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75,
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
//line parser.y:146
		{
			yyVAL.i = VBool
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:147
		{
			yyVAL.i = VInt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:148
		{
			yyVAL.i = VStr
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:149
		{
			yyVAL.i = VArr
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:150
		{
			yyVAL.i = VMap
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:151
		{
			yyVAL.i = VFloat
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:152
		{
			yyVAL.i = VMoney
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:153
		{
			yyVAL.i = VObject
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:154
		{
			yyVAL.i = VBytes
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:155
		{
			yyVAL.i = VFile
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:159
		{
			yyVAL.n = newType(yyDollar[1].i, yylex)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:160
		{
			yyVAL.n = addSubtype(yyDollar[1].n, yyDollar[3].i, yylex)
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:164
		{
			yyVAL.n = nil
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:165
		{
			yyVAL.n = yyDollar[1].n
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:169
		{
			yyVAL.n = nil
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:170
		{
			yyVAL.n = yyDollar[1].n
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:171
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:172
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:176
		{
			yyVAL.n = nil
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:177
		{
			yyVAL.n = newParam(yyDollar[1].n, yylex)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:178
		{
			yyVAL.n = addParam(yyDollar[1].n, yyDollar[3].n)
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:182
		{
			yyVAL.n = nil
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:183
		{
			yyVAL.n = newContractParam(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:184
		{
			yyVAL.n = addContractParam(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:188
		{
			yyVAL.n = newVarValue(yyDollar[1].s, yylex)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:191
		{
			yyVAL.n = newIndex(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:192
		{
			yyVAL.n = addIndex(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:195
		{
			yyVAL.n = nil
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:196
		{
			yyVAL.n = yyDollar[3].n
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:200
		{
			yyVAL.n = nil
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:201
		{
			yyVAL.n = newElif(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:205
		{
			yyVAL.n = nil
		}
	case 33:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:206
		{
			yyVAL.n = newCase(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:210
		{
			yyVAL.n = nil
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:211
		{
			yyVAL.n = yyDollar[3].n
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:215
		{
			yyVAL.n = newSwitch(yyDollar[2].n, yyDollar[4].n, yyDollar[5].n, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:219
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:220
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD_ASSIGN, yylex)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:221
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB_ASSIGN, yylex)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:222
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL_ASSIGN, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:223
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV_ASSIGN, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:224
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD_ASSIGN, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:225
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:226
		{
			yyVAL.n = newBinary(newVarDecl(yyDollar[1].n, []string{yyDollar[2].s}, yylex), yyDollar[4].n, ASSIGN, yylex)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:227
		{
			yyVAL.n = newVarDecl(yyDollar[1].n, yyDollar[2].sa, yylex)
		}
	case 46:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:228
		{
			yyVAL.n = newIf(yyDollar[2].n, yyDollar[4].n, yyDollar[6].n, yyDollar[7].n, yylex)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:229
		{
			yyVAL.n = newBreak(yylex)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:230
		{
			yyVAL.n = newContinue(yylex)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:231
		{
			yyVAL.n = newReturn(nil, yylex)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:232
		{
			yyVAL.n = newReturn(yyDollar[2].n, yylex)
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:233
		{
			yyVAL.n = newWhile(yyDollar[2].n, yyDollar[4].n, yylex)
		}
	case 52:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:234
		{
			yyVAL.n = newFunc(yyDollar[2].s, yyDollar[3].va, yyDollar[5].n, yyDollar[7].n, yylex)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:237
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:238
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 55:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:239
		{
			yyVAL.n = newFor(yyDollar[2].s, yyDollar[4].n, yyDollar[6].n, yylex)
		}
	case 56:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.y:240
		{
			yyVAL.n = newForAll(yyDollar[2].s, yyDollar[4].s, yyDollar[6].n, yyDollar[8].n, yylex)
		}
	case 57:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.y:241
		{
			yyVAL.n = newForInt(yyDollar[2].s, yyDollar[4].n, yyDollar[6].n, yyDollar[8].n, yylex)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:245
		{
			yyVAL.n = newArray(yyDollar[1].n, yylex)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:246
		{
			yyVAL.n = appendArray(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:250
		{
			yyVAL.n = newMap(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 61:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:251
		{
			yyVAL.n = appendMap(yyDollar[1].n, yyDollar[3].s, yyDollar[6].n, yylex)
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:252
		{
			yyVAL.n = appendMap(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n, yylex)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:256
		{
			yyVAL.n = newObj(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:257
		{
			yyVAL.n = newObj(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:258
		{
			yyVAL.n = appendObj(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n, yylex)
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:259
		{
			yyVAL.n = appendObj(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n, yylex)
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:263
		{
			yyVAL.n = newObjArr(yyDollar[1].n, yylex)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:264
		{
			yyVAL.n = appendObjArr(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:268
		{
			yyVAL.n = yyDollar[2].n
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:269
		{
			yyVAL.n = newValue(yyDollar[1].i, yylex)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:270
		{
			yyVAL.n = newValue(yyDollar[1].f, yylex)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:271
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:272
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:273
		{
			yyVAL.n = newValue(true, yylex)
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:274
		{
			yyVAL.n = newValue(false, yylex)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:275
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:276
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:277
		{
			yyVAL.n = yyDollar[1].n
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:278
		{
			yyVAL.n = newEnv(yyDollar[1].s, yylex)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:279
		{
			yyVAL.n = newGetVar(yyDollar[1].s, yylex)
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:280
		{
			yyVAL.n = yyDollar[2].n
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:281
		{
			yyVAL.n = yyDollar[2].n
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:282
		{
			yyVAL.n = newObjList(yyDollar[2].n, yylex)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:286
		{
			yyVAL.n = yyDollar[2].n
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:287
		{
			yyVAL.n = newValue(yyDollar[1].i, yylex)
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:288
		{
			yyVAL.n = newValue(yyDollar[1].f, yylex)
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:289
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:290
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:291
		{
			yyVAL.n = newValue(true, yylex)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:292
		{
			yyVAL.n = newValue(false, yylex)
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:293
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:294
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:295
		{
			yyVAL.n = yyDollar[1].n
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:296
		{
			yyVAL.n = newEnv(yyDollar[1].s, yylex)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:297
		{
			yyVAL.n = newGetVar(yyDollar[1].s, yylex)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:298
		{
			yyVAL.n = yyDollar[2].n
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:299
		{
			yyVAL.n = yyDollar[2].n
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:300
		{
			yyVAL.n = yyDollar[2].n
		}
	case 99:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.y:301
		{
			yyVAL.n = newQuestion(yyDollar[3].n, yyDollar[5].n, yyDollar[7].n, yylex)
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:302
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL, yylex)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:303
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV, yylex)
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:304
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD, yylex)
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:305
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB, yylex)
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:306
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD, yylex)
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:307
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, AND, yylex)
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:308
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, OR, yylex)
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:309
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, EQ, yylex)
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:310
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, NOT_EQ, yylex)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:311
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LTE, yylex)
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:312
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GTE, yylex)
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:313
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LT, yylex)
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:314
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GT, yylex)
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:316
		{
			yyVAL.n = newUnary(yyDollar[2].n, SUB, yylex)
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:317
		{
			yyVAL.n = newUnary(yyDollar[2].n, NOT, yylex)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:321
		{
			yyVAL.sa = []string{yyDollar[1].s}
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:322
		{
			yyVAL.sa = append(yyDollar[1].sa, yyDollar[2].s)
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:326
		{
			yyVAL.va = newVars(yyDollar[1].n, yyDollar[2].sa)
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:330
		{
			yyVAL.va = nil
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:331
		{
			yyVAL.va = yyDollar[1].va
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:332
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:336
		{
			yyVAL.va = newVars(yyDollar[1].n, yyDollar[2].sa)
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:337
		{
			yyVAL.va = newVarExp(yyDollar[1].n, yyDollar[2].s, yyDollar[4].n, yylex)
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:341
		{
			yyVAL.va = yyDollar[1].va
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:342
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 125:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:346
		{
			yyVAL.va = nil
		}
	case 126:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:347
		{
			yyVAL.va = yyDollar[4].va
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:351
		{
			yyVAL.n = newBlock(yyDollar[1].va, yyDollar[2].n, yylex)
		}
	case 128:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:357
		{
			yyVAL.b = false
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:358
		{
			yyVAL.b = true
		}
	case 130:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:362
		{
			yyVAL.n = newContract(yyDollar[2].s, yyDollar[3].b, yyDollar[6].n, yylex)
			setResult(yylex, yyVAL.n)
		}
	}
	goto yystack /* stack new state and value */
}
