// Code generated by goyacc -l -o v5/out.go v5/parser.y. DO NOT EDIT.
package v5

import __yyfmt__ "fmt"

var ftm map[string]int //field type map
func init() {
	ftm = map[string]int{
		"tinyint":   tpTinyint,
		"smallint":  tpSmallint,
		"mediumint": tpMediumint,
		"int":       tpInt,
		"bigint":    tpBigint,
		"float":     tpFloat,
		"doublee":   tpDouble,
		"char":      tpChar,
		"varchar":   tpVarchar_,
		"text":      tpText,
		"date":      tpDate,
		"datetime":  tpDatetime,
		"timestamp": tpTimestamp,
	}
}

type yySymType struct {
	yys       int
	pos       int
	s         string
	i         int64
	express   ExprNode
	statement StmtNode
}

const tString = 57346
const tNumber = 57347
const LPAREN = 57348
const LBRACK = 57349
const LBRACE = 57350
const COMMA = 57351
const PERIOD = 57352
const RPAREN = 57353
const RBRACK = 57354
const RBRACE = 57355
const SEMICOLON = 57356
const EQ = 57357
const QUOTE = 57358
const DQUOTE = 57359
const TILDE = 57360
const kwUnsigned = 57361
const kwNOT = 57362
const kwNull = 57363
const kwAutoIncrement = 57364
const kwCharacter = 57365
const kwCharset = 57366
const kwSet = 57367
const kwComment = 57368
const kwDefault = 57369
const kwTable = 57370
const kwCreate = 57371
const kwCollate = 57372
const kwCurrentTimestamp = 57373
const kwOn = 57374
const kwUpdate = 57375
const kwEngine = 57376
const kwInnoDB = 57377
const kwPrimary = 57378
const kwKey = 57379
const kwUnique = 57380
const tpTinyint = 57381
const tpSmallint = 57382
const tpMediumint = 57383
const tpInt = 57384
const tpBigint = 57385
const tpFloat = 57386
const tpDouble = 57387
const tpChar = 57388
const tpVarchar_ = 57389
const tpText = 57390
const tpDate = 57391
const tpDatetime = 57392
const tpTimestamp = 57393

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tString",
	"tNumber",
	"LPAREN",
	"'('",
	"LBRACK",
	"\"[\"",
	"LBRACE",
	"\"{\"",
	"COMMA",
	"\",\"",
	"PERIOD",
	"\".\"",
	"RPAREN",
	"')'",
	"RBRACK",
	"\"]\"",
	"RBRACE",
	"\"}\"",
	"SEMICOLON",
	"\";\"",
	"EQ",
	"\"=\"",
	"QUOTE",
	"\"'\"",
	"DQUOTE",
	"\"\\\"\"",
	"TILDE",
	"\"`\"",
	"kwUnsigned",
	"kwNOT",
	"kwNull",
	"kwAutoIncrement",
	"kwCharacter",
	"kwCharset",
	"kwSet",
	"kwComment",
	"kwDefault",
	"kwTable",
	"kwCreate",
	"kwCollate",
	"kwCurrentTimestamp",
	"kwOn",
	"kwUpdate",
	"kwEngine",
	"kwInnoDB",
	"kwPrimary",
	"kwKey",
	"kwUnique",
	"tpTinyint",
	"tpSmallint",
	"tpMediumint",
	"tpInt",
	"tpBigint",
	"tpFloat",
	"tpDouble",
	"tpChar",
	"tpVarchar_",
	"tpText",
	"tpDate",
	"tpDatetime",
	"tpTimestamp",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 123

var yyAct = [...]int8{
	89, 49, 38, 25, 27, 28, 26, 29, 30, 31,
	37, 36, 35, 32, 33, 34, 62, 18, 20, 19,
	40, 45, 63, 44, 40, 41, 15, 90, 64, 41,
	83, 97, 39, 6, 111, 114, 39, 13, 78, 50,
	52, 51, 53, 54, 16, 65, 3, 57, 46, 13,
	55, 63, 70, 71, 5, 73, 21, 64, 76, 18,
	20, 19, 74, 72, 77, 85, 110, 8, 79, 14,
	92, 88, 87, 68, 93, 61, 60, 95, 80, 86,
	4, 11, 113, 23, 108, 105, 100, 101, 66, 102,
	99, 98, 107, 7, 42, 106, 109, 43, 91, 69,
	67, 59, 58, 47, 9, 84, 82, 112, 81, 104,
	103, 96, 94, 75, 48, 10, 1, 17, 12, 2,
	24, 56, 22,
}

var yyPact = [...]int16{
	4, -1000, 58, 13, -1000, 63, 98, 111, -1000, 63,
	39, 10, 44, -49, -1000, -11, 78, 85, -27, -29,
	63, 63, -1000, 110, 7, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 96, 95, -1000, 52,
	51, 14, -11, -32, 94, 63, 93, -1000, -1000, -1000,
	7, 7, 29, 7, 24, 109, 7, 34, 103, 101,
	-18, 100, -15, 48, 47, -1000, -1000, 63, 92, 63,
	-1000, -1000, 7, -1000, 108, 7, -1000, 107, -1000, -1000,
	-14, 75, 74, -11, -11, -1000, -11, 106, 105, 69,
	83, 63, 68, -1000, 7, -1000, 36, -12, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 63, 66, -1000, -1000,
	-1000, -9, -1000, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 27, 122, 121, 120, 16, 119, 1, 81, 2,
	118, 44, 117, 0, 116,
}

var yyR1 = [...]int8{
	0, 14, 6, 6, 6, 8, 8, 8, 10, 10,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 7, 7, 7, 7, 7, 7, 7,
	7, 3, 3, 3, 3, 3, 11, 11, 11, 12,
	12, 12, 13, 13, 9, 9, 9, 9, 9, 5,
	5, 1, 1, 2,
}

var yyR2 = [...]int8{
	0, 1, 2, 7, 8, 1, 2, 3, 2, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 4, 0, 2, 2, 3, 2, 4, 3,
	2, 4, 2, 2, 2, 5, 1, 2, 3, 5,
	6, 5, 1, 3, 4, 4, 3, 4, 0, 3,
	3, 3, 1, 2,
}

var yyChk = [...]int16{
	-1000, -14, -6, 42, 22, 41, -1, 30, 4, 6,
	4, -8, -10, -1, 30, 16, -11, -12, 49, 51,
	50, 12, -2, 39, -4, 52, 55, 53, 54, 56,
	57, 58, 62, 63, 64, 61, 60, 59, -9, 47,
	35, 40, 16, 12, 50, 50, -1, -8, 4, -7,
	32, 34, 33, 35, 36, 43, -3, 40, 6, 6,
	24, 24, -5, 37, 43, -9, -11, 6, -1, 6,
	-7, -7, 34, -7, 38, 4, -7, 30, 4, 34,
	44, 5, 5, 48, 5, -9, -5, 24, 24, -13,
	-1, 6, -13, -7, 4, -7, 4, 45, 16, 16,
	-9, -9, -9, 4, 4, 16, 12, -13, 16, -7,
	30, 46, -13, 16, 44,
}

var yyDef = [...]int8{
	0, -2, 1, 0, 2, 0, 0, 0, 52, 0,
	0, 0, 5, 0, 51, 48, 0, 36, 0, 0,
	0, 6, 8, 0, 23, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 0, 0, 3, 0,
	0, 0, 48, 37, 0, 0, 0, 7, 53, 9,
	23, 23, 0, 23, 0, 0, 23, 0, 0, 0,
	0, 0, 48, 0, 0, 4, 38, 0, 0, 0,
	24, 25, 23, 27, 0, 23, 30, 0, 32, 33,
	34, 0, 0, 48, 48, 46, 48, 0, 0, 0,
	42, 0, 0, 26, 23, 29, 0, 0, 21, 22,
	44, 45, 47, 49, 50, 39, 0, 0, 41, 28,
	31, 0, 43, 40, 35,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 29, 3, 3, 3, 3, 27,
	7, 17, 3, 3, 13, 3, 15, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 23,
	3, 25, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 9, 3, 19, 3, 3, 31, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 11, 3, 21,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

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
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
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
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
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
	yyn = int(yyPact[yystate])
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
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
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
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
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
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yylex.(*lex).Stmt = *yyDollar[1].statement.(*CreateTableStmt)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 3:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.statement = &CreateTableStmt{}
			yyVAL.statement.Receive(yyDollar[3].express, yylex)
			yyVAL.statement.Receive(yyDollar[5].statement, yylex)
			yyVAL.statement.Receive(yyDollar[7].statement, yylex)

		}
	case 4:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.statement = &CreateTableStmt{}
			yyVAL.statement.Receive(yyDollar[3].express, yylex)
			yyVAL.statement.Receive(yyDollar[5].statement, yylex)
			yyVAL.statement.Receive(yyDollar[6].statement, yylex)
			yyVAL.statement.Receive(yyDollar[8].statement, yylex)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.statement = new(ColumnsStmt)
			yyVAL.statement.Receive(yyDollar[1].statement, yylex)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.statement = new(ColumnsStmt)
			yyVAL.statement.Receive(yyDollar[1].statement, yylex)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.statement = yyDollar[3].statement
			yyVAL.statement.Receive(yyDollar[1].statement, yylex)
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.statement = yyDollar[1].statement
			yyVAL.statement.Receive(yyDollar[2].express, yylex)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.statement = &ColumnStmt{pos: yyVAL.pos}
			yyVAL.statement.Receive(yyDollar[1].express, yylex)
			yyVAL.statement.Receive(yyDollar[2].express, yylex)
			yyVAL.statement.Receive(yyDollar[3].statement, yylex)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tTINYINT}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tINT}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tSMALLINT}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tMEDIUMINT}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tBIGINT}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tFLOAT}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tDOUBLE}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tDATE}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tDATETIME}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tTIMESTAMP}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tTEXT}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tVARCHAR, length: yyDollar[3].i}

		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.express = FieldTypeNode{t: tCHAR, length: yyDollar[3].i}
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.statement = new(FieldOptsStmt)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[2].statement.Receive(FieldOptNode{opt: optUnsigned, pos: yyVAL.pos}, yylex)
			yyVAL.statement = yyDollar[2].statement
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[2].statement.Receive(FieldOptNode{opt: optNull, pos: yyVAL.pos}, yylex)
			yyVAL.statement = yyDollar[2].statement
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].statement.Receive(FieldOptNode{opt: optNotNull, pos: yyVAL.pos}, yylex)
			yyVAL.statement = yyDollar[3].statement
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[2].statement.Receive(FieldOptNode{opt: optAutoIncrement, pos: yyVAL.pos}, yylex)
			yyVAL.statement = yyDollar[2].statement
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[4].statement.Receive(FieldOptNode{opt: optCharacterSet, s: yyDollar[3].s, pos: yyVAL.pos}, yylex)

			yyVAL.statement = yyDollar[4].statement
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].statement.Receive(FieldOptNode{opt: optCollate, s: yyDollar[2].s, pos: yyVAL.pos}, yylex)
			yyVAL.statement = yyDollar[3].statement
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		{

			yyDollar[2].statement.Receive(yyDollar[1].express, yylex)
			yyVAL.statement = yyDollar[2].statement
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.express = FieldOptNode{opt: optDefault, s: yyDollar[3].s, pos: yyVAL.pos}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.express = FieldOptNode{opt: optDefault, s: yyDollar[2].s, pos: yyVAL.pos}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.express = FieldOptNode{s: "null", opt: optDefaultNull, pos: yyVAL.pos}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.express = FieldOptNode{s: "kwCurrentTimestamp", opt: optDefaultCurrentTimestamp, pos: yyVAL.pos}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.express = FieldOptNode{s: "kwCurrentTimestamp on UPDATE", opt: optOnUpdCurrentTs, pos: yyVAL.pos}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.statement = new(TableIndexListStmt)
			yyVAL.statement.Receive(yyDollar[1].statement, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.statement = new(TableIndexListStmt)
			yyVAL.statement.Receive(yyDollar[1].statement, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].statement.Receive(yyDollar[1].statement, yylex)
			yyVAL.statement = yyDollar[3].statement
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.statement = &TableIndexStmt{t: idxPrimary}
			yyVAL.statement.Receive(yyDollar[4].statement, yylex)

		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.statement = &TableIndexStmt{t: idxUnique}
			yyVAL.statement.Receive(yyDollar[3].express, yylex)
			yyVAL.statement.Receive(yyDollar[5].statement, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.statement = &TableIndexStmt{t: idxKey}
			yyVAL.statement.Receive(yyDollar[2].express, yylex)
			yyVAL.statement.Receive(yyDollar[4].statement, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.statement = new(NameListStmt)
			yyVAL.statement.Receive(yyDollar[1].express, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].statement.Receive(yyDollar[1].express, yylex)
			yyVAL.statement = yyDollar[3].statement
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[4].statement.Receive(TableOptNode{opt: optEngine, s: "InnoDB", pos: yyVAL.pos}, yylex)
			yyVAL.statement = yyDollar[4].statement
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[4].statement.Receive(TableOptNode{opt: optAutoIncrement, i: yyDollar[3].i, pos: yyVAL.pos}, yylex)
			yyVAL.statement = yyDollar[4].statement
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].statement.Receive(yyDollar[2].express, yylex)
			yyVAL.statement = yyDollar[3].statement
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[4].statement.Receive(yyDollar[2].express, yylex)
			yyDollar[4].statement.Receive(yyDollar[3].express, yylex)
			yyVAL.statement = yyDollar[4].statement
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.statement = &TableOptsStmt{}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.express = TableOptNode{opt: optCharSet, s: yyDollar[3].s, pos: yyVAL.pos}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.express = TableOptNode{opt: optCollate, s: yyDollar[3].s, pos: yyVAL.pos}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.express = NameNode{s: yyDollar[2].s, pos: yyVAL.pos}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.express = NameNode{s: yyDollar[1].s, pos: yyVAL.pos}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.express = CommentNode{s: yyDollar[2].s}
		}
	}
	goto yystack /* stack new state and value */
}
