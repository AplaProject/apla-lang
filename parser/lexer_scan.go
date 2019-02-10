// Code generated by golex. DO NOT EDIT.

package parser

import (
	"strconv"

	"modernc.org/golex/lex"
)

func (l *lexer) scan(lval *yySymType) lex.Char {
	c := l.Enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.Rule0()

	goto yystart1

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	case 21:
		goto yyrule21
	case 22:
		goto yyrule22
	case 23:
		goto yyrule23
	case 24:
		goto yyrule24
	case 25:
		goto yyrule25
	case 26:
		goto yyrule26
	case 27:
		goto yyrule27
	case 28:
		goto yyrule28
	case 29:
		goto yyrule29
	case 30:
		goto yyrule30
	case 31:
		goto yyrule31
	case 32:
		goto yyrule32
	case 33:
		goto yyrule33
	case 34:
		goto yyrule34
	case 35:
		goto yyrule35
	case 36:
		goto yyrule36
	case 37:
		goto yyrule37
	case 38:
		goto yyrule38
	case 39:
		goto yyrule39
	case 40:
		goto yyrule40
	case 41:
		goto yyrule41
	case 42:
		goto yyrule42
	case 43:
		goto yyrule43
	case 44:
		goto yyrule44
	case 45:
		goto yyrule45
	case 46:
		goto yyrule46
	case 47:
		goto yyrule47
	case 48:
		goto yyrule48
	case 49:
		goto yyrule49
	case 50:
		goto yyrule50
	}
	goto yystate1 // silence unused label error
yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate6
	case c == '%':
		goto yystate9
	case c == '&':
		goto yystate13
	case c == '(':
		goto yystate15
	case c == ')':
		goto yystate17
	case c == '*':
		goto yystate18
	case c == '+':
		goto yystate22
	case c == ',':
		goto yystate26
	case c == '-':
		goto yystate28
	case c == '/':
		goto yystate32
	case c == '0':
		goto yystate41
	case c == ':':
		goto yystate45
	case c == ';':
		goto yystate46
	case c == '<':
		goto yystate47
	case c == '=':
		goto yystate52
	case c == '>':
		goto yystate55
	case c == '?':
		goto yystate60
	case c == '[':
		goto yystate62
	case c == '\n':
		goto yystate3
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c == ']':
		goto yystate64
	case c == 'b':
		goto yystate65
	case c == 'c':
		goto yystate69
	case c == 'd':
		goto yystate77
	case c == 'e':
		goto yystate81
	case c == 'f':
		goto yystate87
	case c == 'h':
		goto yystate92
	case c == 'i':
		goto yystate98
	case c == 'r':
		goto yystate102
	case c == 't':
		goto yystate108
	case c == 'w':
		goto yystate112
	case c == '{':
		goto yystate117
	case c == '|':
		goto yystate119
	case c == '}':
		goto yystate121
	case c >= '1' && c <= '9':
		goto yystate42
	case c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'g' || c >= 'j' && c <= 'q' || c == 's' || c == 'u' || c == 'v' || c >= 'x' && c <= 'z' || c == '\u0080':
		goto yystate61
	}

yystate2:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.Next()
	yyrule = 20
	l.Mark()
	switch {
	default:
		goto yyrule20
	case c == ')':
		goto yystate5
	case c == '\t' || c == ' ':
		goto yystate4
	}

yystate4:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == ')':
		goto yystate5
	case c == '\t' || c == ' ':
		goto yystate4
	}

yystate5:
	c = l.Next()
	yyrule = 14
	l.Mark()
	goto yyrule14

yystate6:
	c = l.Next()
	yyrule = 31
	l.Mark()
	switch {
	default:
		goto yyrule31
	case c == '=':
		goto yystate7
	}

yystate7:
	c = l.Next()
	yyrule = 30
	l.Mark()
	switch {
	default:
		goto yyrule30
	case c == '\n':
		goto yystate8
	case c == '\t' || c == ' ':
		goto yystate7
	}

yystate8:
	c = l.Next()
	yyrule = 30
	l.Mark()
	goto yyrule30

yystate9:
	c = l.Next()
	yyrule = 28
	l.Mark()
	switch {
	default:
		goto yyrule28
	case c == '=':
		goto yystate12
	case c == '\n':
		goto yystate11
	case c == '\t' || c == ' ':
		goto yystate10
	}

yystate10:
	c = l.Next()
	yyrule = 28
	l.Mark()
	switch {
	default:
		goto yyrule28
	case c == '\n':
		goto yystate11
	case c == '\t' || c == ' ':
		goto yystate10
	}

yystate11:
	c = l.Next()
	yyrule = 28
	l.Mark()
	goto yyrule28

yystate12:
	c = l.Next()
	yyrule = 8
	l.Mark()
	goto yyrule8

yystate13:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate14
	}

yystate14:
	c = l.Next()
	yyrule = 22
	l.Mark()
	goto yyrule22

yystate15:
	c = l.Next()
	yyrule = 13
	l.Mark()
	switch {
	default:
		goto yyrule13
	case c == '\n':
		goto yystate16
	case c == '\t' || c == ' ':
		goto yystate15
	}

yystate16:
	c = l.Next()
	yyrule = 13
	l.Mark()
	goto yyrule13

yystate17:
	c = l.Next()
	yyrule = 15
	l.Mark()
	goto yyrule15

yystate18:
	c = l.Next()
	yyrule = 26
	l.Mark()
	switch {
	default:
		goto yyrule26
	case c == '=':
		goto yystate21
	case c == '\n':
		goto yystate20
	case c == '\t' || c == ' ':
		goto yystate19
	}

yystate19:
	c = l.Next()
	yyrule = 26
	l.Mark()
	switch {
	default:
		goto yyrule26
	case c == '\n':
		goto yystate20
	case c == '\t' || c == ' ':
		goto yystate19
	}

yystate20:
	c = l.Next()
	yyrule = 26
	l.Mark()
	goto yyrule26

yystate21:
	c = l.Next()
	yyrule = 6
	l.Mark()
	goto yyrule6

yystate22:
	c = l.Next()
	yyrule = 24
	l.Mark()
	switch {
	default:
		goto yyrule24
	case c == '=':
		goto yystate25
	case c == '\n':
		goto yystate24
	case c == '\t' || c == ' ':
		goto yystate23
	}

yystate23:
	c = l.Next()
	yyrule = 24
	l.Mark()
	switch {
	default:
		goto yyrule24
	case c == '\n':
		goto yystate24
	case c == '\t' || c == ' ':
		goto yystate23
	}

yystate24:
	c = l.Next()
	yyrule = 24
	l.Mark()
	goto yyrule24

yystate25:
	c = l.Next()
	yyrule = 4
	l.Mark()
	goto yyrule4

yystate26:
	c = l.Next()
	yyrule = 10
	l.Mark()
	switch {
	default:
		goto yyrule10
	case c == '\n':
		goto yystate27
	case c == '\t' || c == ' ':
		goto yystate26
	}

yystate27:
	c = l.Next()
	yyrule = 10
	l.Mark()
	goto yyrule10

yystate28:
	c = l.Next()
	yyrule = 25
	l.Mark()
	switch {
	default:
		goto yyrule25
	case c == '=':
		goto yystate31
	case c == '\n':
		goto yystate30
	case c == '\t' || c == ' ':
		goto yystate29
	}

yystate29:
	c = l.Next()
	yyrule = 25
	l.Mark()
	switch {
	default:
		goto yyrule25
	case c == '\n':
		goto yystate30
	case c == '\t' || c == ' ':
		goto yystate29
	}

yystate30:
	c = l.Next()
	yyrule = 25
	l.Mark()
	goto yyrule25

yystate31:
	c = l.Next()
	yyrule = 5
	l.Mark()
	goto yyrule5

yystate32:
	c = l.Next()
	yyrule = 27
	l.Mark()
	switch {
	default:
		goto yyrule27
	case c == '*':
		goto yystate35
	case c == '/':
		goto yystate38
	case c == '=':
		goto yystate40
	case c == '\n':
		goto yystate34
	case c == '\t' || c == ' ':
		goto yystate33
	}

yystate33:
	c = l.Next()
	yyrule = 27
	l.Mark()
	switch {
	default:
		goto yyrule27
	case c == '\n':
		goto yystate34
	case c == '\t' || c == ' ':
		goto yystate33
	}

yystate34:
	c = l.Next()
	yyrule = 27
	l.Mark()
	goto yyrule27

yystate35:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate36
	case c >= '\x01' && c <= ')' || c >= '+' && c <= 'ÿ':
		goto yystate35
	}

yystate36:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate36
	case c == '/':
		goto yystate37
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= 'ÿ':
		goto yystate35
	}

yystate37:
	c = l.Next()
	yyrule = 2
	l.Mark()
	switch {
	default:
		goto yyrule2
	case c == '*':
		goto yystate36
	case c >= '\x01' && c <= ')' || c >= '+' && c <= 'ÿ':
		goto yystate35
	}

yystate38:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate39
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate38
	}

yystate39:
	c = l.Next()
	yyrule = 3
	l.Mark()
	goto yyrule3

yystate40:
	c = l.Next()
	yyrule = 7
	l.Mark()
	goto yyrule7

yystate41:
	c = l.Next()
	yyrule = 49
	l.Mark()
	switch {
	default:
		goto yyrule49
	case c == 'X' || c == 'x':
		goto yystate43
	case c >= '0' && c <= '9':
		goto yystate42
	}

yystate42:
	c = l.Next()
	yyrule = 49
	l.Mark()
	switch {
	default:
		goto yyrule49
	case c >= '0' && c <= '9':
		goto yystate42
	}

yystate43:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate44
	}

yystate44:
	c = l.Next()
	yyrule = 48
	l.Mark()
	switch {
	default:
		goto yyrule48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate44
	}

yystate45:
	c = l.Next()
	yyrule = 11
	l.Mark()
	goto yyrule11

yystate46:
	c = l.Next()
	yyrule = 21
	l.Mark()
	goto yyrule21

yystate47:
	c = l.Next()
	yyrule = 34
	l.Mark()
	switch {
	default:
		goto yyrule34
	case c == '=':
		goto yystate50
	case c == '\n':
		goto yystate49
	case c == '\t' || c == ' ':
		goto yystate48
	}

yystate48:
	c = l.Next()
	yyrule = 34
	l.Mark()
	switch {
	default:
		goto yyrule34
	case c == '\n':
		goto yystate49
	case c == '\t' || c == ' ':
		goto yystate48
	}

yystate49:
	c = l.Next()
	yyrule = 34
	l.Mark()
	goto yyrule34

yystate50:
	c = l.Next()
	yyrule = 32
	l.Mark()
	switch {
	default:
		goto yyrule32
	case c == '\n':
		goto yystate51
	case c == '\t' || c == ' ':
		goto yystate50
	}

yystate51:
	c = l.Next()
	yyrule = 32
	l.Mark()
	goto yyrule32

yystate52:
	c = l.Next()
	yyrule = 9
	l.Mark()
	switch {
	default:
		goto yyrule9
	case c == '=':
		goto yystate53
	}

yystate53:
	c = l.Next()
	yyrule = 29
	l.Mark()
	switch {
	default:
		goto yyrule29
	case c == '\n':
		goto yystate54
	case c == '\t' || c == ' ':
		goto yystate53
	}

yystate54:
	c = l.Next()
	yyrule = 29
	l.Mark()
	goto yyrule29

yystate55:
	c = l.Next()
	yyrule = 35
	l.Mark()
	switch {
	default:
		goto yyrule35
	case c == '=':
		goto yystate58
	case c == '\n':
		goto yystate57
	case c == '\t' || c == ' ':
		goto yystate56
	}

yystate56:
	c = l.Next()
	yyrule = 35
	l.Mark()
	switch {
	default:
		goto yyrule35
	case c == '\n':
		goto yystate57
	case c == '\t' || c == ' ':
		goto yystate56
	}

yystate57:
	c = l.Next()
	yyrule = 35
	l.Mark()
	goto yyrule35

yystate58:
	c = l.Next()
	yyrule = 33
	l.Mark()
	switch {
	default:
		goto yyrule33
	case c == '\n':
		goto yystate59
	case c == '\t' || c == ' ':
		goto yystate58
	}

yystate59:
	c = l.Next()
	yyrule = 33
	l.Mark()
	goto yyrule33

yystate60:
	c = l.Next()
	yyrule = 12
	l.Mark()
	goto yyrule12

yystate61:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate62:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c == '\n':
		goto yystate63
	case c == '\t' || c == ' ':
		goto yystate62
	}

yystate63:
	c = l.Next()
	yyrule = 18
	l.Mark()
	goto yyrule18

yystate64:
	c = l.Next()
	yyrule = 19
	l.Mark()
	goto yyrule19

yystate65:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'o':
		goto yystate66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate66:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'o':
		goto yystate67
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate67:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'l':
		goto yystate68
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate68:
	c = l.Next()
	yyrule = 45
	l.Mark()
	switch {
	default:
		goto yyrule45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate69:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'o':
		goto yystate70
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate70:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'n':
		goto yystate71
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate71:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 't':
		goto yystate72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate72:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'r':
		goto yystate73
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate73:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'a':
		goto yystate74
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate74:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'c':
		goto yystate75
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate75:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 't':
		goto yystate76
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate76:
	c = l.Next()
	yyrule = 37
	l.Mark()
	switch {
	default:
		goto yyrule37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate77:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'a':
		goto yystate78
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate78:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 't':
		goto yystate79
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate79:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'a':
		goto yystate80
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate80:
	c = l.Next()
	yyrule = 36
	l.Mark()
	switch {
	default:
		goto yyrule36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate81:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'l':
		goto yystate82
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate82:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'i':
		goto yystate83
	case c == 's':
		goto yystate85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate83:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'f':
		goto yystate84
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate84:
	c = l.Next()
	yyrule = 40
	l.Mark()
	switch {
	default:
		goto yyrule40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate85:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'e':
		goto yystate86
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate86:
	c = l.Next()
	yyrule = 41
	l.Mark()
	switch {
	default:
		goto yyrule41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate87:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'a':
		goto yystate88
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate88:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'l':
		goto yystate89
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate89:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 's':
		goto yystate90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate90:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'e':
		goto yystate91
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate91:
	c = l.Next()
	yyrule = 44
	l.Mark()
	switch {
	default:
		goto yyrule44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate92:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'e':
		goto yystate93
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate93:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'x':
		goto yystate94
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate94:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'i':
		goto yystate95
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate95:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'n':
		goto yystate96
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate96:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 't':
		goto yystate97
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate97:
	c = l.Next()
	yyrule = 47
	l.Mark()
	switch {
	default:
		goto yyrule47
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate98:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'f':
		goto yystate99
	case c == 'n':
		goto yystate100
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate99:
	c = l.Next()
	yyrule = 39
	l.Mark()
	switch {
	default:
		goto yyrule39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate100:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 't':
		goto yystate101
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate101:
	c = l.Next()
	yyrule = 46
	l.Mark()
	switch {
	default:
		goto yyrule46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate102:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'e':
		goto yystate103
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate103:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 't':
		goto yystate104
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate104:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'u':
		goto yystate105
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate105:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'r':
		goto yystate106
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate106:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'n':
		goto yystate107
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate107:
	c = l.Next()
	yyrule = 42
	l.Mark()
	switch {
	default:
		goto yyrule42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate108:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'r':
		goto yystate109
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate109:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'u':
		goto yystate110
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate110:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'e':
		goto yystate111
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate111:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate112:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'h':
		goto yystate113
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate113:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'i':
		goto yystate114
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate114:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'l':
		goto yystate115
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate115:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c == 'e':
		goto yystate116
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate116:
	c = l.Next()
	yyrule = 38
	l.Mark()
	switch {
	default:
		goto yyrule38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate61
	}

yystate117:
	c = l.Next()
	yyrule = 16
	l.Mark()
	switch {
	default:
		goto yyrule16
	case c == '\n':
		goto yystate118
	case c == '\t' || c == ' ':
		goto yystate117
	}

yystate118:
	c = l.Next()
	yyrule = 16
	l.Mark()
	goto yyrule16

yystate119:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate120
	}

yystate120:
	c = l.Next()
	yyrule = 23
	l.Mark()
	goto yyrule23

yystate121:
	c = l.Next()
	yyrule = 17
	l.Mark()
	goto yyrule17

yyrule1: // [ \t\r ]+
	{
		// ignore all whitespace
		goto yystate0
	}
yyrule2: // \/\*(.|\n)*\*\/
	{
		// ignore comments
		goto yystate0
	}
yyrule3: // \/\/(.)*\n
	{
		return l.char(NEWLINE)
	}
yyrule4: // \+=
	{
		return l.char(ADD_ASSIGN)
	}
yyrule5: // -=
	{
		return l.char(SUB_ASSIGN)
	}
yyrule6: // \*=
	{
		return l.char(MUL_ASSIGN)
	}
yyrule7: // \/=
	{
		return l.char(DIV_ASSIGN)
	}
yyrule8: // %=
	{
		return l.char(MOD_ASSIGN)
	}
yyrule9: // =
	{
		return l.char(ASSIGN)
	}
yyrule10: // ,[ \t]*\n?
	{
		return l.char(COMMA)
	}
yyrule11: // :
	{
		return l.char(COLON)
	}
yyrule12: // \?
	{
		return l.char(QUESTION)
	}
yyrule13: // \([ \t]*\n?
	{
		return l.char(LPAREN)
	}
yyrule14: // \n[ \t]*\)
	{
		return l.char(RPAREN)
	}
yyrule15: // \)
	{
		return l.char(RPAREN)
	}
yyrule16: // \{[ \t]*\n?
	{
		return l.char(LBRACE)
	}
yyrule17: // \}
	{
		return l.char(RBRACE)
	}
yyrule18: // \[[ \t]*\n?
	{
		return l.char(LBRAKET)
	}
yyrule19: // \]
	{
		return l.char(RBRAKET)
	}
yyrule20: // \n
	{
		return l.char(NEWLINE)
	}
yyrule21: // ;
	{
		return l.char(NEWLINE)
	}
yyrule22: // &&
	{
		return l.char(AND)
	}
yyrule23: // \|\|
	{
		return l.char(OR)
	}
yyrule24: // \+[ \t]*\n?
	{
		return l.char(ADD)
	}
yyrule25: // -[ \t]*\n?
	{
		return l.char(SUB)
	}
yyrule26: // \*[ \t]*\n?
	{
		return l.char(MUL)
	}
yyrule27: // \/[ \t]*\n?
	{
		return l.char(DIV)
	}
yyrule28: // %[ \t]*\n?
	{
		return l.char(MOD)
	}
yyrule29: // ==[ \t]*\n?
	{
		return l.char(EQ)
	}
yyrule30: // !=[ \t]*\n?
	{
		return l.char(NOT_EQ)
	}
yyrule31: // !
	{
		return l.char(NOT)
	}
yyrule32: // \<=[ \t]*\n?
	{
		return l.char(LTE)
	}
yyrule33: // >=[ \t]*\n?
	{
		return l.char(GTE)
	}
yyrule34: // \<[ \t]*\n?
	{
		return l.char(LT)
	}
yyrule35: // >[ \t]*\n?
	{
		return l.char(GT)
	}
yyrule36: // data
	{
		return l.char(DATA)
	}
yyrule37: // contract
	{
		return l.char(CONTRACT)
	}
yyrule38: // while
	{
		return l.char(WHILE)
	}
yyrule39: // if
	{
		return l.char(IF)
	}
yyrule40: // elif
	{
		return l.char(ELIF)
	}
yyrule41: // else
	{
		return l.char(ELSE)
	}
yyrule42: // return
	{
		return l.char(RETURN)
	}
yyrule43: // true
	{
		return l.char(TRUE)
	}
yyrule44: // false
	{
		return l.char(FALSE)
	}
yyrule45: // bool
	{
		return l.char(T_BOOL)
	}
yyrule46: // int
	{
		return l.char(T_INT)
	}
yyrule47: // hexint
	{
		return l.char(T_INT)
	}
yyrule48: // {hexint}
	{
		{
			val, _ := strconv.ParseInt(string(l.TokenBytes(nil)), 0, 64)
			lval.i = int(val)
			return l.char(INT)
		}
		goto yystate0
	}
yyrule49: // {int}
	{
		{
			lval.i, _ = strconv.Atoi(string(l.TokenBytes(nil)))
			return l.char(INT)
		}
		goto yystate0
	}
yyrule50: // {identifier}
	{
		{
			lval.s = string(l.TokenBytes(nil))
			return l.char(IDENT)
		}
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := l.Abort(); ok {
		return l.char(c)
	}

	goto yyAction
}
