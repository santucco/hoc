package hoc

type inst interface{}
type datum interface{}
type symval interface{}
type bltin func(float64) symval

const (
	nSTACK    = 256
	nPROG     = 2000
	nFRAME    = 100
	nMAXWIDTH = 65536
	nSYMBOLS  = 100
)

type symbol struct { // symbol table entry
	name string
	t    int
	val  symval
}

type saveval struct { // saved value of variable
	val  symval
	t    int
	next *saveval
}

type formal struct { // formal parameter
	sym  *symbol
	save *saveval
	next *formal
}

type fndefn struct { // formal parameter
	code    int
	formals *formal
	nargs   int
}

type frame struct { // procfunc call stack frame
	sp    *symbol // symbol table entry
	retpc int     // where to resume after return
	argn  int     // n-th argument on stack
	nargs int     // number of arguments
}

type lexer struct {
	// lexical analyser data
	line    string
	linepos int
	lineno  int
	in      chan string
	out     chan string
}

type interpreter struct {
	//Interpreter data
	stack     []datum // the this.stack
	prog      []inst  // the machine
	progp     int     // next free spot for code generation
	pc        int     // this.program counter during execution
	progbase  int     // start of current subprogram
	returning bool    // 1 if return stmt seen
	frame     []frame // this.frames
	symbols   symbols // symbol map
	locsyms   symbols // current symbol map 
	IndentSym string  // string for atomic indention
	NewLine   bool    // add newline after all 
}

type Hoc struct {
	lexer
	interpreter
}

type HocError struct {
	Line    string
	Linepos int
	Lineno  int
	Error   interface{}
}
