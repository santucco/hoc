
//line hoc.y:2
package hoc
import __yyfmt__ "fmt"
//line hoc.y:2
		
type _pair struct{
	inst	int
	nargs	int
}


//line hoc.y:10
type yySymType struct {
	yys int
	sym			*symbol		// symbol pointer
	syms		symbols		// symbols table
	val			symval		// value 
	inst		int			// machine instruction position
	narg		int 		// number of arguments 
	formals 	*formal		// list of formal parameters 
	pair		_pair
	
}

const NUMBER = 57346
const STRING = 57347
const BOOL = 57348
const NULL = 57349
const VAR = 57350
const PRINT = 57351
const PRINTF = 57352
const BLTIN = 57353
const UNDEF = 57354
const WHILE = 57355
const FOR = 57356
const IF = 57357
const ELSE = 57358
const LOCAL = 57359
const FUNCTION = 57360
const PROCEDURE = 57361
const RETURN = 57362
const FUNC = 57363
const PROC = 57364
const READ = 57365
const COUNTOF = 57366
const ADDEQ = 57367
const SUBEQ = 57368
const MULEQ = 57369
const DIVEQ = 57370
const MODEQ = 57371
const OR = 57372
const AND = 57373
const GT = 57374
const GE = 57375
const LT = 57376
const LE = 57377
const EQ = 57378
const NE = 57379
const UNARYMINUS = 57380
const NOT = 57381
const INC = 57382
const DEC = 57383

var yyToknames = []string{
	"NUMBER",
	"STRING",
	"BOOL",
	"NULL",
	"VAR",
	"PRINT",
	"PRINTF",
	"BLTIN",
	"UNDEF",
	"WHILE",
	"FOR",
	"IF",
	"ELSE",
	"LOCAL",
	"FUNCTION",
	"PROCEDURE",
	"RETURN",
	"FUNC",
	"PROC",
	"READ",
	"COUNTOF",
	" =",
	"ADDEQ",
	"SUBEQ",
	"MULEQ",
	"DIVEQ",
	"MODEQ",
	" ,",
	"OR",
	"AND",
	"GT",
	"GE",
	"LT",
	"LE",
	"EQ",
	"NE",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARYMINUS",
	"NOT",
	"INC",
	"DEC",
	" ^",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line hoc.y:868

	// end of grammar


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 113
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 482

var yyAct = []int{

	171, 181, 7, 168, 188, 63, 13, 2, 127, 112,
	152, 39, 40, 41, 42, 43, 75, 124, 79, 189,
	139, 141, 201, 155, 199, 81, 177, 175, 174, 88,
	89, 90, 85, 165, 148, 146, 136, 135, 91, 92,
	131, 122, 87, 84, 176, 94, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 105, 106, 107, 66,
	67, 68, 69, 70, 71, 65, 125, 115, 116, 117,
	118, 119, 120, 121, 83, 82, 65, 159, 158, 160,
	161, 73, 74, 128, 128, 128, 65, 65, 132, 8,
	9, 72, 129, 130, 64, 183, 147, 49, 198, 126,
	134, 57, 56, 50, 51, 52, 53, 54, 55, 44,
	45, 46, 47, 48, 166, 138, 137, 125, 49, 123,
	194, 58, 77, 143, 144, 33, 145, 140, 12, 114,
	164, 80, 143, 62, 33, 151, 44, 45, 46, 47,
	48, 93, 153, 153, 156, 49, 113, 154, 128, 86,
	59, 187, 173, 149, 109, 162, 172, 4, 185, 157,
	60, 61, 76, 167, 5, 108, 178, 180, 34, 3,
	179, 46, 47, 48, 169, 184, 6, 1, 49, 111,
	110, 163, 142, 153, 128, 21, 192, 186, 20, 195,
	191, 196, 193, 19, 197, 14, 182, 78, 200, 0,
	0, 0, 204, 203, 202, 206, 0, 205, 22, 24,
	23, 0, 33, 17, 18, 27, 0, 35, 36, 37,
	0, 38, 26, 16, 15, 0, 0, 0, 25, 0,
	0, 57, 56, 50, 51, 52, 53, 54, 55, 44,
	45, 46, 47, 48, 0, 29, 0, 0, 49, 0,
	30, 31, 32, 150, 8, 9, 0, 28, 0, 170,
	190, 22, 24, 23, 0, 33, 17, 18, 27, 0,
	35, 36, 37, 0, 38, 26, 16, 15, 10, 11,
	0, 25, 0, 0, 0, 0, 22, 24, 23, 0,
	33, 17, 18, 27, 0, 35, 36, 37, 29, 38,
	26, 16, 15, 30, 31, 32, 25, 8, 9, 0,
	28, 0, 12, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 29, 0, 0, 0, 0, 30, 31,
	32, 0, 0, 0, 0, 28, 0, 170, 57, 56,
	50, 51, 52, 53, 54, 55, 44, 45, 46, 47,
	48, 0, 0, 0, 0, 49, 0, 0, 0, 0,
	133, 57, 56, 50, 51, 52, 53, 54, 55, 44,
	45, 46, 47, 48, 0, 0, 0, 0, 49, 8,
	9, 22, 24, 23, 0, 33, 0, 0, 27, 0,
	0, 0, 0, 0, 0, 26, 0, 0, 0, 0,
	0, 25, 125, 57, 56, 50, 51, 52, 53, 54,
	55, 44, 45, 46, 47, 48, 0, 0, 29, 0,
	49, 0, 0, 30, 31, 32, 0, 0, 0, 0,
	28, 57, 56, 50, 51, 52, 53, 54, 55, 44,
	45, 46, 47, 48, 0, 0, 0, 0, 49, 56,
	50, 51, 52, 53, 54, 55, 44, 45, 46, 47,
	48, 0, 0, 0, 0, 49, 50, 51, 52, 53,
	54, 55, 44, 45, 46, 47, 48, 0, 0, 0,
	0, 49,
}
var yyPact = []int{

	-1000, 257, -1000, 39, 39, 39, 39, 329, -1000, -1000,
	142, 142, 43, 34, -1000, 377, -1000, 377, 126, 22,
	21, -10, -1000, -1000, -1000, 117, -1000, -11, 377, 377,
	377, 117, 117, -1000, 133, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 377, 377, 377, 377, 377, 377,
	377, 377, 377, 377, 377, 377, 377, 377, -1000, -1000,
	-1000, -1000, -1000, 141, -1000, 121, 377, 377, 377, 377,
	377, 377, 377, -1000, -1000, 399, -1000, -12, 88, 399,
	86, 35, 377, 377, 377, 24, -13, 377, 306, 48,
	48, 24, 24, 75, 129, 129, 48, 48, 48, 48,
	96, 96, 96, 96, 96, 96, 432, 416, -16, -17,
	43, 84, -1000, -39, -1000, 399, 399, 399, 399, 399,
	399, 69, 377, 377, -1000, 377, -1000, -19, 399, 46,
	-20, 377, 199, -1000, 377, 117, 117, -33, 43, 73,
	-1000, -21, 83, 399, 399, 371, 282, 377, 282, -26,
	-1000, 399, -27, 13, -28, -1000, 141, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 43, -1000, 377, -1000, -1000, -1000,
	-1000, 399, 45, -1000, -1000, -1000, 117, -1000, -1000, 73,
	399, -1000, 204, 377, 104, 282, -1000, 282, 43, 67,
	-1000, -1000, -1000, -30, 282, -1000, -1000, -36, 43, 282,
	-1000, -1000, 73, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 10, 0, 174, 162, 197, 196, 195, 3, 8,
	193, 188, 185, 122, 1, 121, 6, 9, 182, 21,
	17, 19, 181, 155, 4, 180, 179, 177, 7, 169,
	168, 165, 158, 154, 151, 5,
}
var yyR1 = []int{

	0, 27, 27, 27, 27, 27, 27, 27, 28, 28,
	30, 7, 7, 16, 16, 4, 4, 4, 4, 4,
	4, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 8, 8, 9, 10, 11, 12,
	13, 14, 6, 6, 6, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 5, 5, 31, 32, 29,
	33, 34, 29, 1, 1, 1, 15, 15, 15, 19,
	19, 18, 18, 20, 20, 35, 35, 23, 25, 25,
	26, 26, 17, 22, 24, 24, 24, 21, 21, 21,
	21, 21, 21,
}
var yyR2 = []int{

	0, 0, 2, 3, 3, 3, 3, 3, 1, 1,
	1, 2, 4, 1, 3, 3, 3, 3, 3, 3,
	3, 1, 1, 1, 2, 5, 1, 2, 3, 3,
	6, 10, 9, 6, 1, 3, 1, 1, 1, 1,
	0, 0, 0, 2, 2, 1, 1, 1, 1, 4,
	2, 1, 5, 4, 3, 3, 3, 3, 3, 3,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 1, 3, 0, 0, 8,
	0, 0, 8, 0, 1, 3, 1, 1, 1, 0,
	1, 1, 3, 0, 3, 0, 1, 5, 0, 1,
	1, 4, 3, 5, 0, 1, 4, 1, 1, 1,
	1, 1, 1,
}
var yyChk = []int{

	-1000, -27, -28, -29, -23, -4, -3, -2, 50, 51,
	21, 22, 55, -16, -7, 20, 19, 9, 10, -10,
	-11, -12, 4, 6, 5, 24, 18, 11, 53, 41,
	46, 47, 48, 8, -30, 13, 14, 15, 17, -28,
	-28, -28, -28, -28, 40, 41, 42, 43, 44, 49,
	34, 35, 36, 37, 38, 39, 33, 32, -15, 8,
	18, 19, -15, -35, 51, 52, 25, 26, 27, 28,
	29, 30, 57, 47, 48, -2, -4, -13, -5, -2,
	5, -16, 53, 53, 53, -16, -13, 53, -2, -2,
	-2, -16, -16, 8, -2, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, -31, -33,
	-25, -26, -17, 5, 8, -2, -2, -2, -2, -2,
	-2, -2, 53, 31, -20, 31, -20, -9, -2, -9,
	-9, 53, -2, 54, 25, 53, 53, -35, 31, 59,
	58, -19, -18, -2, -2, -2, 54, 50, 54, -19,
	54, -2, -1, -16, -1, 56, -35, -21, 5, 4,
	6, 7, -23, -22, 57, 54, 31, -20, -8, -3,
	55, -2, -9, -8, 54, 54, 31, 54, -17, -35,
	-2, -14, -6, 50, -14, -32, -1, -34, -24, -21,
	56, -28, -8, -9, 16, -8, -8, -35, 31, 54,
	-8, 58, -35, -8, -14, -24, -14,
}
var yyDef = []int{

	1, -2, 2, 0, 0, 51, 0, 0, 8, 9,
	0, 0, 95, 48, 21, 23, 40, 26, 0, 0,
	0, 0, 45, 46, 47, 0, 40, 0, 0, 0,
	0, 0, 0, 13, 0, 37, 38, 39, 10, 3,
	4, 5, 6, 7, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 77, 86,
	87, 88, 80, 98, 96, 0, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 24, 51, 0, 27, 75,
	93, 93, 0, 0, 0, 50, 0, 0, 0, 61,
	70, 71, 72, 11, 55, 56, 57, 58, 59, 60,
	62, 63, 64, 65, 66, 67, 68, 69, 0, 0,
	95, 99, 100, 0, 14, 15, 16, 17, 18, 19,
	20, 0, 89, 0, 28, 0, 29, 0, 36, 0,
	0, 89, 0, 54, 0, 83, 83, 0, 95, 0,
	49, 0, 90, 91, 76, 93, 0, 0, 0, 0,
	53, 12, 0, 84, 0, 97, 0, 102, 107, 108,
	109, 110, 111, 112, 95, 25, 0, 94, 41, 34,
	42, 22, 0, 41, 52, 78, 83, 81, 101, 104,
	92, 30, 0, 0, 33, 0, 85, 0, 95, 105,
	35, 43, 44, 0, 0, 79, 82, 0, 95, 0,
	41, 103, 104, 41, 32, 106, 31,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	51, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 44, 3, 3,
	53, 54, 42, 40, 31, 41, 52, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 50,
	3, 25, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 57, 3, 58, 49, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 56,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 26, 27, 28, 29, 30, 32, 33,
	34, 35, 36, 37, 38, 39, 45, 46, 47, 48,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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

	case 4:
		//line hoc.y:44
		{
			this := yylex.(*Hoc)
			for  k, v := range yyS[yypt-1].val.(symbols) {
				tracer.Trace(debLex, "list object: installing %v = %v", k, v)
				if sym := lookup(this.symbols, v.name); sym != nil {
					sym = v
				} else {
					this.symbols = install(this.symbols, v)
				}
			}
		
		}
	case 5:
		//line hoc.y:57
		{ 
			this := yylex.(*Hoc)
			this.code((*Hoc).xpop); 
			this.code(nil); 
			return 1; 
		}
	case 6:
		//line hoc.y:64
		{ 
			this := yylex.(*Hoc)
			this.code(nil); 
			return 1; 
		}
	case 7:
		//line hoc.y:70
		{ 
			this := yylex.(*Hoc)
			this.code((*Hoc).printtop); 
			this.code(nil); 
			return 1; 
		}
	case 10:
		//line hoc.y:85
		{
			tracer.Trace(debLex, "localdef::LOCAL")
	
			this := yylex.(*Hoc)
			if this.locsyms == nil {
				yylex.Error("using local definition outside of function of procedure")
			}
		}
	case 11:
		//line hoc.y:97
		{
			tracer.Trace(debLex, "localvar:: localdef varname")
			this := yylex.(*Hoc)
			sym := lookup(this.locsyms, yyS[yypt-0].sym.name)
			if sym == nil {
				this.locsyms = install(this.locsyms, yyS[yypt-0].sym)
			}	
		}
	case 12:
		//line hoc.y:106
		{
			tracer.Trace(debLex, "localvar::localdev varname '=' expr" )
			this := yylex.(*Hoc)
			sym := lookup(this.locsyms, yyS[yypt-2].sym.name)
			if sym == nil {
				this.locsyms = install(this.locsyms, yyS[yypt-2].sym)
				sym = yyS[yypt-2].sym
			}
			this.code((*Hoc).vpush)
			this.code(sym)
			this.code((*Hoc).assign)		
		}
	case 13:
		//line hoc.y:122
		{
			tracer.Trace(debLex, "VAR - %s", yyS[yypt-0].sym.name)
			this := yylex.(*Hoc)
			sym := this.lookup(yyS[yypt-0].sym.name)
			if sym == nil {
				this.symbols = install(this.symbols, yyS[yypt-0].sym)
				sym = yyS[yypt-0].sym
			}
			yyVAL.sym = sym
		}
	case 14:
		//line hoc.y:133
		{		
			tracer.Trace(debLex, "%s '.' %s ", yyS[yypt-2].sym.name, yyS[yypt-0].sym.name )
			if yyS[yypt-2].sym.val == nil {
				yyS[yypt-2].sym.val = newSymbols()
				yyS[yypt-2].sym.t = VAR
			}
			syms, ok := yyS[yypt-2].sym.val.(symbols)
			if !ok {
				yylex.Error( yyS[yypt-2].sym.name + " is not compound type")
			}
			sym := lookup(syms, yyS[yypt-0].sym.name)
			if sym == nil {
				syms = install(syms, yyS[yypt-0].sym)
				sym = yyS[yypt-0].sym
			}
			yyVAL.sym = sym
		}
	case 15:
		//line hoc.y:154
		{ 
			tracer.Trace(debLex, "varname '=' expr" )
			this := yylex.(*Hoc)
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-2].sym)
			this.code((*Hoc).assign)
			yyVAL.inst=yyS[yypt-0].inst 
		}
	case 16:
		//line hoc.y:163
		{ 
			tracer.Trace(debLex, "%s ADDEQ expr ", yyS[yypt-2].sym.name)
			this := yylex.(*Hoc)
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-2].sym)
			this.code((*Hoc).addeq) 
			yyVAL.inst=yyS[yypt-0].inst 
		}
	case 17:
		//line hoc.y:172
		{ 
			tracer.Trace(debLex, "%s SUBEQ expr ", yyS[yypt-2].sym.name)
			this := yylex.(*Hoc)	
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-2].sym)
			this.code((*Hoc).subeq) 
			yyVAL.inst=yyS[yypt-0].inst 
		}
	case 18:
		//line hoc.y:181
		{ 
			tracer.Trace(debLex, "%s MULEQ expr ", yyS[yypt-2].sym.name)
			this := yylex.(*Hoc)
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-2].sym)
			this.code((*Hoc).muleq) 
			yyVAL.inst=yyS[yypt-0].inst 
		}
	case 19:
		//line hoc.y:190
		{ 
			tracer.Trace(debLex, "%s DIVEQ expr ", yyS[yypt-2].sym.name)
			this := yylex.(*Hoc)
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-2].sym)
			this.code((*Hoc).diveq)
			yyVAL.inst=yyS[yypt-0].inst
		}
	case 20:
		//line hoc.y:199
		{ 
			tracer.Trace(debLex, "%s MODEQ expr ", yyS[yypt-2].sym.name)
			this := yylex.(*Hoc)
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-2].sym)
			this.code((*Hoc).modeq)
			yyVAL.inst=yyS[yypt-0].inst
		}
	case 21:
		//line hoc.y:211
		{
			tracer.Trace(debLex, "stmt::localvar")
			yyVAL.inst=yyS[yypt-0].inst
		}
	case 22:
		//line hoc.y:216
		{ 
			tracer.Trace(debLex, "stmt::expr")
			this := yylex.(*Hoc)
			this.code((*Hoc).xpop)
		}
	case 23:
		//line hoc.y:222
		{ 	
			tracer.Trace(debLex, "stmt::RETURN")	
			this := yylex.(*Hoc)
			this.defnonly("return") 
			this.code((*Hoc).procret) 
		}
	case 24:
		//line hoc.y:229
		{ 
			tracer.Trace(debLex, "stmt::RETURN expr")	
			this := yylex.(*Hoc)
			this.defnonly("return") 
			yyVAL.inst=yyS[yypt-0].inst 
			this.code((*Hoc).funcret)
		}
	case 25:
		//line hoc.y:237
		{ 
			tracer.Trace(debLex, "stmt::PROCEDURE begin")
			yyVAL.inst = yyS[yypt-3].inst 
			this := yylex.(*Hoc)	
			this.code((*Hoc).call)
			this.code(yyS[yypt-4].sym)
			this.code(yyS[yypt-1].narg)
		}
	case 26:
		//line hoc.y:246
		{
			tracer.Trace(debLex, "stmt::PRINT")	
			this := yylex.(*Hoc)	
			this.code((*Hoc).printall) 
		}
	case 27:
		//line hoc.y:252
		{ 
			tracer.Trace(debLex, "stmt::PRINT prlist")
			yyVAL.inst = yyS[yypt-0].inst
		}
	case 28:
		//line hoc.y:257
		{ 
			tracer.Trace(debLex, "stmt::PRINTF STRING paramlist")
			this := yylex.(*Hoc)
			this.code((*Hoc).cpush)
			this.code(yyS[yypt-0].pair.nargs) 
			this.code((*Hoc).cpush)
			this.code(yyS[yypt-1].val.(string)) 
			this.code((*Hoc).prfstr)
			yyVAL.inst = yyS[yypt-0].pair.inst
		}
	case 29:
		//line hoc.y:268
		{ 
			tracer.Trace(debLex, "stmt::PRINTF STRING paramlist")
			this := yylex.(*Hoc)
			this.code((*Hoc).cpush)
			this.code(yyS[yypt-0].pair.nargs) 
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-1].sym) 
			this.code((*Hoc).prfstr)
			yyVAL.inst = yyS[yypt-0].pair.inst
		}
	case 30:
		//line hoc.y:278
		{
			tracer.Trace(debLex, "stmt::while")
			tracer.Trace(debLex, "while '(' cond ')' stmts end ")
			this := yylex.(*Hoc)	
			this.prog[yyS[yypt-5].inst + 1] = yyS[yypt-1].inst	// body of loop 
		this.prog[yyS[yypt-5].inst + 2] = yyS[yypt-0].inst 	// end, if cond fails 
	}
	case 31:
		//line hoc.y:286
		{
			tracer.Trace(debLex, "stmt::for")
			this := yylex.(*Hoc)	
			this.prog[yyS[yypt-9].inst + 1] = yyS[yypt-5].inst	// condition 
		this.prog[yyS[yypt-9].inst + 2] = yyS[yypt-3].inst	// post loop 
		this.prog[yyS[yypt-9].inst + 3] = yyS[yypt-1].inst	// body of loop 
		this.prog[yyS[yypt-9].inst + 4] = yyS[yypt-0].inst	// end, if cond fails 
	}
	case 32:
		//line hoc.y:295
		{	// if with else 
		tracer.Trace(debLex, "stmt::if else") 
			tracer.Trace(debCode, "$1 %v $5 %v $8 %v $9 %v ", yyS[yypt-8].inst, yyS[yypt-4].inst, yyS[yypt-1].inst,yyS[yypt-0].inst)
			this := yylex.(*Hoc)	
			this.prog[yyS[yypt-8].inst + 1] = yyS[yypt-4].inst	// thenpart 
		this.prog[yyS[yypt-8].inst + 2] = yyS[yypt-1].inst	//elsepart 
		this.prog[yyS[yypt-8].inst + 3] = yyS[yypt-0].inst 	// end, if cond fails 
	}
	case 33:
		//line hoc.y:304
		{	// else-less if 
		tracer.Trace(debLex, "stmt::if") 
			this := yylex.(*Hoc)	
			this.prog[yyS[yypt-5].inst + 1] = yyS[yypt-1].inst	// thenpart  
		this.prog[yyS[yypt-5].inst + 3] = yyS[yypt-0].inst 	// end, if cond fails 
	}
	case 34:
		//line hoc.y:314
		{
			tracer.Trace(debLex, "stmts::stmt")
		}
	case 35:
		//line hoc.y:318
		{ 
			tracer.Trace(debLex, "stmt::stmtlist")
			yyVAL.inst = yyS[yypt-1].inst
		}
	case 36:
		//line hoc.y:326
		{ 
			tracer.Trace(debLex, "cond")
			this := yylex.(*Hoc)	
			this.code(nil)
		}
	case 37:
		//line hoc.y:335
		{ 
			tracer.Trace(debLex, "WHILE")
			this := yylex.(*Hoc)
			yyVAL.inst = this.code((*Hoc).whilecode)
			this.code(nil)
			this.code(nil) 
		}
	case 38:
		//line hoc.y:346
		{ 
			tracer.Trace(debLex, "FOR")
			this := yylex.(*Hoc)
			yyVAL.inst = this.code((*Hoc).forcode)
			this.code(nil)
			this.code(nil)
			this.code(nil) 
			this.code(nil)
		}
	case 39:
		//line hoc.y:359
		{ 
			tracer.Trace(debLex, "IF")
			this := yylex.(*Hoc)
			yyVAL.inst = this.code((*Hoc).ifcode) 
			this.code(nil)
			this.code(nil)
			this.code(nil) 
		}
	case 40:
		//line hoc.y:370
		{ 
			tracer.Trace(debLex, "begin")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.progp
		}
	case 41:
		//line hoc.y:378
		{ 
			tracer.Trace(debLex, "end")
			this := yylex.(*Hoc)	
			this.code(nil) 
			yyVAL.inst = this.progp
		}
	case 42:
		//line hoc.y:387
		{ 
			tracer.Trace(debLex, "stmtlist")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.progp 
		}
	case 43:
		yyVAL.inst = yyS[yypt-0].inst
	case 44:
		yyVAL.inst = yyS[yypt-0].inst
	case 45:
		//line hoc.y:398
		{ 
			tracer.Trace(debLex, "expr::NUMBER")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).cpush)
			switch yyS[yypt-0].val.(type) {
			case int:
				this.code(yyS[yypt-0].val.(int))
			case float64:
				this.code(yyS[yypt-0].val.(float64))
			default:
				tracer.Trace(debLex, "type of %v : %T", yyS[yypt-0].val, yyS[yypt-0].val)
				yylex.Error("wrong type for number constant") 
			}
		}
	case 46:
		//line hoc.y:413
		{
			tracer.Trace(debLex, "expr::NUMBER")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).cpush)
			this.code(yyS[yypt-0].val.(bool))	
		}
	case 47:
		//line hoc.y:420
		{ 
			tracer.Trace(debLex, "expr::STRING")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).cpush)
			this.code(yyS[yypt-0].val.(string)) 
		}
	case 48:
		//line hoc.y:427
		{ 
			tracer.Trace(debLex, "expr::VARNAME %s", yyS[yypt-0].sym.name)
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).vpush)
			this.code(yyS[yypt-0].sym)
			this.code((*Hoc).eval) 
		}
	case 49:
		//line hoc.y:435
		{
			tracer.Trace(debLex, "expr::VARNAME %s '[' expr ']'", yyS[yypt-3].sym.name)
			this := yylex.(*Hoc)	
			this.code((*Hoc).vpush)
			this.code(yyS[yypt-3].sym)
			this.code((*Hoc).elemeval) 
			yyVAL.inst = yyS[yypt-1].inst
		}
	case 50:
		//line hoc.y:444
		{
			tracer.Trace(debLex, "expr::COUNTOF %s", yyS[yypt-1].sym.name)
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).vpush)
			this.code(yyS[yypt-0].sym)
			this.code((*Hoc).countof) 
		}
	case 51:
		yyVAL.inst = yyS[yypt-0].inst
	case 52:
		//line hoc.y:453
		{ 
			tracer.Trace(debLex, "expr::FUNCTION")
			yyVAL.inst = yyS[yypt-3].inst; 
			this := yylex.(*Hoc)	
			this.code((*Hoc).call)
			this.code(yyS[yypt-4].sym)
			this.code(yyS[yypt-1].narg) 
		}
	case 53:
		//line hoc.y:469
		{ 
			tracer.Trace(debLex, "expr::BLTIN")
			yyVAL.inst=yyS[yypt-1].inst 
			this := yylex.(*Hoc)	
			this.code((*Hoc).bltin)
			this.code(yyS[yypt-3].sym.val.(bltin))
		}
	case 54:
		//line hoc.y:477
		{ 		tracer.Trace(debLex, "expr::'(' expr ')'")
			yyVAL.inst = yyS[yypt-1].inst; 
		}
	case 55:
		//line hoc.y:481
		{ 
			tracer.Trace(debLex, "expr::expr '+' expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).add) 
		}
	case 56:
		//line hoc.y:487
		{ 
			tracer.Trace(debLex, "expr::expr '-' expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).sub) 
		}
	case 57:
		//line hoc.y:493
		{ 
			tracer.Trace(debLex, "expr::expr '*' expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).mul) 
		}
	case 58:
		//line hoc.y:499
		{ 
			tracer.Trace(debLex, "expr::expr '/' expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).div) 
		}
	case 59:
		//line hoc.y:505
		{ 
			tracer.Trace(debLex, "expr::expr '%' expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).mod) 
		}
	case 60:
		//line hoc.y:511
		{ 
			tracer.Trace(debLex, "expr::expr '^' expr")
			this := yylex.(*Hoc)	
			this.code ((*Hoc).power) 
		}
	case 61:
		//line hoc.y:517
		{ 
			tracer.Trace(debLex, "expr::'-' expr")
			yyVAL.inst=yyS[yypt-0].inst; 
			this := yylex.(*Hoc)	
			this.code((*Hoc).negate) 
		}
	case 62:
		//line hoc.y:524
		{ 
			tracer.Trace(debLex, "expr::expr 'GT expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).gt)
		}
	case 63:
		//line hoc.y:530
		{ 
			tracer.Trace(debLex, "expr::expr GE expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).ge) 
		}
	case 64:
		//line hoc.y:536
		{ 		
			tracer.Trace(debLex, "expr::expr LT expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).lt)
		}
	case 65:
		//line hoc.y:542
		{ 
			tracer.Trace(debLex, "expr::expr LE expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).le)
		}
	case 66:
		//line hoc.y:548
		{ 
			tracer.Trace(debLex, "expr::expr EQ expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).eq) 
		}
	case 67:
		//line hoc.y:554
		{ 
			tracer.Trace(debLex, "expr::expr NE expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).ne) 
		}
	case 68:
		//line hoc.y:560
		{ 
			tracer.Trace(debLex, "expr::expr AND expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).and)
		}
	case 69:
		//line hoc.y:566
		{ 
			tracer.Trace(debLex, "expr::expr OR expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).or)
		}
	case 70:
		//line hoc.y:572
		{ 
			tracer.Trace(debLex, "expr::NOT expr")
			yyVAL.inst = yyS[yypt-0].inst; 
			this := yylex.(*Hoc)	
			this.code((*Hoc).not) 
		}
	case 71:
		//line hoc.y:579
		{ 
			tracer.Trace(debLex, "expr::INC expr")
			this := yylex.(*Hoc)
			yyVAL.inst = this.code((*Hoc).vpush)
			this.code(yyS[yypt-0].sym)	
			yyVAL.inst = this.code((*Hoc).preinc)
		}
	case 72:
		//line hoc.y:587
		{ 
			tracer.Trace(debLex, "expr::DEC expr")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).vpush)
			this.code(yyS[yypt-0].sym)		
			this.code((*Hoc).predec)
		}
	case 73:
		//line hoc.y:595
		{ 
			tracer.Trace(debLex, "expr::expr INC")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).vpush)
			this.code(yyS[yypt-1].sym)
			this.code((*Hoc).postinc)
		}
	case 74:
		//line hoc.y:603
		{ 
			tracer.Trace(debLex, "expr::expr DEC")
			this := yylex.(*Hoc)	
			yyVAL.inst = this.code((*Hoc).vpush)
			this.code(yyS[yypt-1].sym)
			this.code((*Hoc).postdec)
		}
	case 75:
		//line hoc.y:614
		{
			tracer.Trace(debLex, "prlist::expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).prexpr) 
		}
	case 76:
		//line hoc.y:620
		{ 
			tracer.Trace(debLex, "prlist::prlist ',' expr")
			this := yylex.(*Hoc)	
			this.code((*Hoc).prexpr)
		}
	case 77:
		//line hoc.y:629
		{ 
			tracer.Trace(debLex, "defn::FUNC %s", yyS[yypt-0].sym.name)
			this := yylex.(*Hoc)	
			sym := lookup(this.symbols, yyS[yypt-0].sym.name)
			if sym != nil {
				yylex.Error("symbol " + yyS[yypt-0].sym.name + " already defined")
			}
			yyS[yypt-0].sym.t = FUNCTION
			this.symbols = install(this.symbols, yyS[yypt-0].sym)
			this.locsyms = newSymbols()
		}
	case 78:
		//line hoc.y:641
		{ 
			tracer.Trace(debLex, "defn::FUNC %s formals", yyS[yypt-4].sym.name)	
		}
	case 79:
		//line hoc.y:645
		{
			tracer.Trace(debLex, "defn::FUNC %s stmt", yyS[yypt-6].sym.name)
			this := yylex.(*Hoc)	
			this.code((*Hoc).procret); 
			this.define(yyS[yypt-6].sym, yyS[yypt-3].formals) 
			this.locsyms = nil
		}
	case 80:
		//line hoc.y:653
		{ 
			tracer.Trace(debLex, "defn::PROC %s", yyS[yypt-0].sym.name)
			this := yylex.(*Hoc)	
			sym := lookup(this.symbols, yyS[yypt-0].sym.name)
			if sym != nil {
				yylex.Error("symbol " + yyS[yypt-0].sym.name + " already defined")
			}
			yyS[yypt-0].sym.t = PROCEDURE
			this.symbols = install(this.symbols, yyS[yypt-0].sym)
			this.locsyms = newSymbols()
		}
	case 81:
		//line hoc.y:665
		{
			tracer.Trace(debLex, "defn::PROC %s formals", yyS[yypt-4].sym.name)	
		}
	case 82:
		//line hoc.y:669
		{ 
			tracer.Trace(debLex, "defn::PROC %s formals", yyS[yypt-6].sym.name)
			this := yylex.(*Hoc)	
			this.code((*Hoc).procret) 
			this.define(yyS[yypt-6].sym, yyS[yypt-3].formals) 
			this.locsyms = nil
		}
	case 83:
		//line hoc.y:678
		{ yyVAL.formals = nil; }
	case 84:
		//line hoc.y:680
		{ 
			tracer.Trace(debLex, "formals::%s", yyS[yypt-0].sym.name)
			tracer.Trace(debLex, "formal parameter %s", yyS[yypt-0].sym.name)
			yyVAL.formals = formallist(yyS[yypt-0].sym, nil) 
		}
	case 85:
		//line hoc.y:686
		{ 
			tracer.Trace(debLex, "formals::%s ',' formals", yyS[yypt-2].sym.name)
			yyVAL.formals = formallist(yyS[yypt-2].sym, yyS[yypt-0].formals) 
		}
	case 86:
		yyVAL.sym = yyS[yypt-0].sym
	case 87:
		yyVAL.sym = yyS[yypt-0].sym
	case 88:
		yyVAL.sym = yyS[yypt-0].sym
	case 89:
		//line hoc.y:698
		{ 
			tracer.Trace(debLex, "args:: no arguments")
			yyVAL.narg = 0 
		}
	case 90:
		//line hoc.y:704
		{ 
			tracer.Trace(debLex, "args::arglist")	}
	case 91:
		//line hoc.y:709
		{ 
			tracer.Trace(debLex, "arglist:: one argument")
			yyVAL.narg = 1 
		}
	case 92:
		//line hoc.y:714
		{ 
			tracer.Trace(debLex, "arglist:: yet another argument")
			yyVAL.narg = yyS[yypt-2].narg + 1 
		}
	case 93:
		//line hoc.y:721
		{ 
			tracer.Trace(debLex, "paramlist")
			yyVAL.pair.nargs = 0
		}
	case 94:
		//line hoc.y:726
		{ 
			tracer.Trace(debLex, "paramlist:: expr")
			yyVAL.pair.nargs = yyS[yypt-0].pair.nargs + 1
			yyVAL.pair.inst = yyS[yypt-1].inst
		}
	case 97:
		//line hoc.y:739
		{
			tracer.Trace(debLex, "object::members %v", yyS[yypt-2].val)
			yyVAL.val = yyS[yypt-2].val
		}
	case 98:
		//line hoc.y:746
		{
			yyVAL.val = newSymbols()
		}
	case 99:
		//line hoc.y:750
		{}
	case 100:
		//line hoc.y:755
		{
			syms := newSymbols()
			syms = install(syms, yyS[yypt-0].sym)
			tracer.Trace(debLex, "members::pair %v", syms)
			yyVAL.val = syms
		}
	case 101:
		//line hoc.y:762
		{
			tracer.Trace(debLex, "members::pair ',' members")
			syms := yyS[yypt-3].val.(symbols)
	
			tracer.Trace(debLex, "syms: %v", syms)
		
			if s := lookup(syms, yyS[yypt-0].sym.name); s != nil {
					yylex.Error("dublicated value name: " + yyS[yypt-0].sym.name)
			}
			tracer.Trace(debLex, "pair %v", yyS[yypt-0].sym)
			syms = install(syms, yyS[yypt-0].sym)
			tracer.Trace(debLex, "members::pair ',' members %v", syms)
			yyVAL.val = syms	
		}
	case 102:
		//line hoc.y:780
		{
			tracer.Trace(debLex, "pair")
			sym := newSymbol(yyS[yypt-2].val.(string), VAR, yyS[yypt-0].val)
			tracer.Trace(debLex, "new symbol %v", sym)
			yyVAL.sym = sym
		}
	case 103:
		//line hoc.y:790
		{
			tracer.Trace(debLex, "array::elements %v", yyS[yypt-2].val)
			yyVAL.val = yyS[yypt-2].val
		}
	case 104:
		//line hoc.y:797
		{
			yyVAL.val = make([]symval, 0)
		}
	case 105:
		//line hoc.y:801
		{
			tracer.Trace(debLex, "elements::value")
			e := make([]symval, 1)
			e[0] = yyS[yypt-0].val
			yyVAL.val = e		
		}
	case 106:
		//line hoc.y:808
		{
			tracer.Trace(debLex, "elements::value ',' elements")
			e := make([]symval, len(yyS[yypt-0].val.([]symval)) + 1)
			e[0] = yyS[yypt-3].val
			copy(e[1:], yyS[yypt-0].val.([]symval))
			yyVAL.val = e
		}
	case 107:
		yyVAL.val = yyS[yypt-0].val
	case 108:
		yyVAL.val = yyS[yypt-0].val
	case 109:
		yyVAL.val = yyS[yypt-0].val
	case 110:
		yyVAL.val = yyS[yypt-0].val
	case 111:
		yyVAL.val = yyS[yypt-0].val
	case 112:
		yyVAL.val = yyS[yypt-0].val
	}
	goto yystack /* stack new state and value */
}
