package hoc

import (
	"math"
)

var commonsyms symbols = newSymbols() // global symbol table

type keyword struct {
	name string
	kval int
}

var keywords = [...]keyword{
	{"proc", PROC},
	{"func", FUNC},
	{"return", RETURN},
	{"if", IF},
	{"else", ELSE},
	{"while", WHILE},
	{"for", FOR},
	{"print", PRINT},
	{"printf", PRINTF},
	{"read", READ},
	{"local", LOCAL},
	{"countof", COUNTOF},
}


type _const struct { // Constants 
	name string
	t    int
	val  symval
}

var consts = [...]_const{
	{"PI", NUMBER, float64(math.Pi)},
	{"E", NUMBER, float64(math.E)},
	{"GAMMA", NUMBER, float64(0.57721566490153286060)}, // Euler 
	{"DEG", NUMBER, float64(57.29577951308232087680)},  // degradian 
	{"PHI", NUMBER, float64(math.Phi)},                 // golden ratio 
	{"null", NULL, nil},
	{"true", BOOL, true},
	{"false", BOOL, false},
}

type builtin struct { // Built-ins 
	name string
	f    bltin
}

var builtins = [...]builtin{
	{"sin", sin},
	{"cos", cos},
	{"tan", tan},
	{"atan", atan},
	{"asin", asin}, // checks range 
	{"acos", acos}, // checks range 
	{"sinh", sinh}, // checks range 
	{"cosh", cosh}, // checks range 
	{"tanh", tanh},
	{"log", log},     // checks range 
	{"log2", log2},     // checks range 
	{"log10", log10}, // checks range 
	{"exp", exp},     // checks range 
	{"sqrt", sqrt},   // checks range 
	{"int", integer},
	{"abs", abs},
}


func init() { // install constants and built-ins in table 
	defer tracer.Exit(tracer.Enter())
	var i int
	for i = 0; i < len(keywords); i++ {
		tracer.Trace(debInit, "before commonsyms %x", commonsyms)
		sym := newSymbol(keywords[i].name, keywords[i].kval, nil)
		commonsyms = install(commonsyms, sym)
		tracer.Trace(debInit, "after commonsyms %x", commonsyms)
	}
	for i = 0; i < len(consts); i++ {
		sym := newSymbol(consts[i].name, consts[i].t, consts[i].val)
		commonsyms = install(commonsyms, sym)
	}
	for i = 0; i < len(builtins); i++ {
		sym := newSymbol(builtins[i].name, BLTIN, builtins[i].f)
		commonsyms = install(commonsyms, sym)
	}
}
