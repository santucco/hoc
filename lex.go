package hoc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func (this *Hoc) Lex(yylval *yySymType) int {
	defer tracer.Exit(tracer.Enter())
	c := this.getrune()
	for c == ' ' || c == '\t' {
		c = this.getrune()
	}
	tracer.Trace(debLex, "c = '%c' (0x%x)", c, c)
	if c == 0 {
		return 0
	}
	if c == '\\' {
		tracer.Trace(debLex, "escape?")

		c = this.getrune()
		if c == '\n' {
			tracer.Trace(debLex, "new line")
			this.lineno++
			return this.Lex(yylval)
		}
	}
	if c == '#' { // comment
		for c != '\n' && c > 0 {
			c = this.getrune()
		}
		if c == '\n' {
			this.lineno++
		}
		tracer.Trace(debLex, "c = '%c' (0x%x)", c, c)
		return int(c)
	}
	if unicode.IsDigit(c) { //number
		this.ungetrune(c)
		p := this.linepos
		for c = this.getrune(); unicode.IsDigit(c); c = this.getrune() {
		}
		if c != '.' {
			this.ungetrune(c)
			tracer.Trace(debLex, "test for int %s", this.line[p:this.linepos])
			d, err := strconv.Atoi(this.line[p:this.linepos])
			if err != nil {
				panic(err.Error())
			}
			yylval.val = d
			tracer.Trace(debLex, "Lex -> NUMBER(int=%v)", d)
			return NUMBER
		}
		for c = this.getrune(); unicode.IsDigit(c); c = this.getrune() {
		}
		this.ungetrune(c)
		tracer.Trace(debLex, "test for float %s", this.line[p:this.linepos])
		d, err := strconv.ParseFloat(this.line[p:this.linepos], 64)
		if err != nil {
			panic(err.Error())
		}
		yylval.val = d
		tracer.Trace(debLex, "Lex -> NUMBER(float=%v)", d)
		return NUMBER
	}
	if unicode.IsLetter(c) || c == '_' {
		this.ungetrune(c)
		p := this.linepos
		for c = this.getrune(); unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_'; c = this.getrune() {
		}
		this.ungetrune(c)
		name := this.line[p:this.linepos]
		s := this.lookup(name)
		tracer.Trace(debLex, "%d %d %s", p, this.linepos, name)
		if s != nil {
			tracer.Trace(debLex, "symbol %v has found", s)
			switch s.t {
			case NUMBER, BOOL, NULL:
				tracer.Trace(debLex, "result is constant %v", s)
				yylval.val = s.val
				return s.t
			case VAR, UNDEF: // ignoring 
			default:
				tracer.Trace(debLex, "result is %v", s)
				yylval.sym = s
				return s.t
			}
		}
		tracer.Trace(debLex, "consider %s like a new VAR", name)
		yylval.sym = newSymbol(name, UNDEF, nil)
		return VAR

	}
	if c == '"' { //quoted string
		p := this.linepos
		for c = this.checkrune(); c != '"'; c = this.checkrune() {
			if c == 0 {
				panic("missing quote")
			}
		}
		this.ungetrune(c)
		yylval.val = this.line[p:this.linepos]
		c = this.getrune()
		tracer.Trace(debLex, "Lex -> STRING(%v)", yylval.val)
		return STRING
	}
	switch c {
	case '+':
		return int(this.follow('+', INC, this.follow('=', ADDEQ, '+')))
	case '-':
		return int(this.follow('-', DEC, this.follow('=', SUBEQ, '-')))
	case '*':
		return int(this.follow('=', MULEQ, '*'))
	case '/':
		return int(this.follow('=', DIVEQ, '/'))
	case '%':
		return int(this.follow('=', MODEQ, '%'))
	case '>':
		return int(this.follow('=', GE, GT))
	case '<':
		return int(this.follow('=', LE, LT))
	case '=':
		return int(this.follow('=', EQ, '='))
	case '!':
		return int(this.follow('=', NE, NOT))
	case '|':
		return int(this.follow('|', OR, '|'))
	case '&':
		return int(this.follow('&', AND, '&'))
	case '\n':
		this.lineno++
		return '\n'
	}
	tracer.Trace(debLex, "Lex ->'%c'", c)
	return int(c)
}

func (this *Hoc) Error(s string) {
	defer tracer.Exit(tracer.Enter())
	panic(s)
}

func (this *Hoc) checkrune() rune { // get next char with \'s interpreted
	defer tracer.Exit(tracer.Enter())

	c1 := this.getrune()
	if c1 != '\\' {
		return c1
	}
	c2 := this.getrune()
	if c2 == '\\' {
		p := this.linepos
		this.ungetrune(c2)
		this.line = this.line[0:this.linepos] + this.line[p:len(this.line)]
		return c1
	}
	if !unicode.IsLower(c2) {
		this.ungetrune(c2)
		return c1
	}
	transtab := "b\bf\fn\nr\rt\t"

	if i := strings.IndexRune(transtab, c2); i != -1 {
		tracer.Trace(debLex, "transtab[%d]=%d transtab[%d]=%d", i, transtab[i], i+1, transtab[i+1])
		p := this.linepos
		this.ungetrune(c2)
		this.ungetrune(c1)
		this.line = this.line[0:this.linepos] + transtab[i+1:i+2] + this.line[p:len(this.line)]
		return this.getrune()
	}
	this.ungetrune(c2)
	return c1
}

func (this *Hoc) follow(expect rune, ifyes rune, ifno rune) rune { // look ahead for >=, etc
	defer tracer.Exit(tracer.Enter())
	c := this.getrune()
	if c == expect {
		tracer.Trace(debLex, "Lex ->'%c'", ifno)
		return ifyes
	}
	this.ungetrune(c)
	tracer.Trace(debLex, "Lex ->'%c'", ifno)
	return ifno
}

func (this *Hoc) printState() {
	defer tracer.Exit(tracer.Enter())
	if ((tracer.TraceLevel & debState) == 0) || this.prog == nil {
		return
	}
	fmt.Fprint(os.Stderr, "\n************************************************************\n")
	if this.prog != nil && this.pc < len(this.prog) {
		fmt.Fprintf(os.Stderr, "pc %v\nprog[%v] %v\nprog[pc] %v\nprogp %v\nprogbase %v\n",
			this.pc, len(this.prog), this.prog,
			this.prog[this.pc], this.progp, this.progbase)
	}
	fmt.Fprint(os.Stderr, "\n************************************************************\n")
}

func (this *Hoc) Process(in chan string, out chan string) { // execute until EOF
	defer tracer.Exit(tracer.Enter())
	defer func() {
		defer tracer.Exit(tracer.Enter())
		if x := recover(); x != nil {
			this.restoreall()
			this.printState()
			//close(this.in)
			//close(this.out)
			var err = HocError{this.line, this.linepos, this.lineno, x}
			panic(err)
		}
	}()
	this.in = in
	this.out = out
	this.lineno = 1

	for this.initcode(); yyParse(this) != 0; this.initcode() {
		tracer.Trace(debState, "executing")
		this.execute(this.progbase)
	}
	close(this.out)
}

func (this *Hoc) defnonly(s string) { // warn if illegal definition
	defer tracer.Exit(tracer.Enter())

	if this.locsyms == nil {
		panic(s + " used outside definition")
	}
}

func (this *Hoc) readline() bool {
	defer tracer.Exit(tracer.Enter())
	var str string
	for true {
		s, ok := <-this.in
		if !ok {
			break
		}
		if len(s) == 0 {
			tracer.Trace(debLex, "ignoring empty input")
			continue
		}
		str += s

		if i := strings.IndexRune(str, '\n'); i != -1 {
			this.line = str[:i+1]
			this.linepos = 0
			str = str[i+1:]
			return true
		}
	}
	if len(str) != 0 {
		this.line = str + "\n"
		this.linepos = 0
		str = ""
		return true
	}
	return false
}

func (this *Hoc) getrune() rune {
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debLex, "linep %d, len %d", this.linepos, len(this.line))
	if this.linepos == len(this.line) && !this.readline() {
		tracer.Trace(debLex, "end of input")
		return 0
	}
	c, n := utf8.DecodeRuneInString(this.line[this.linepos:len(this.line)])
	this.linepos += n
	tracer.Trace(debLex, "got '%c' (%d) line[%d:%d]: %s", c, c, this.linepos, len(this.line), this.line[this.linepos:len(this.line)])
	return c
}

func (this *Hoc) replacerune() rune {
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debLex, "linep %d, len %d", this.linepos, len(this.line))
	if this.linepos == len(this.line) && !this.readline() {
		tracer.Trace(debLex, "end of input")
		return 0
	}
	c, n := utf8.DecodeRuneInString(this.line[this.linepos:len(this.line)])
	this.linepos += n
	tracer.Trace(debLex, "got '%c' (%d) line[%d:%d]: %s", c, c, this.linepos, len(this.line), this.line[this.linepos:len(this.line)])
	return c
}

func (this *Hoc) ungetrune(c rune) {
	if c == 0 {
		return
	}
	defer tracer.Exit(tracer.Enter())
	n := utf8.RuneLen(c)

	if this.linepos-n < 0 {
		panic("can't unget rune")
	}
	this.linepos -= n
	tracer.Trace(debLex, " ungot '%c' line[%d:%d]: %s", c, this.linepos, len(this.line), this.line[this.linepos:len(this.line)])
}
