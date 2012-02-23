package hoc

func (this *Hoc) lookup(s string) *symbol { // find s in symbol table 
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debSymbol, "looking for %s in commonsyms %v", s, commonsyms)

	if sym := lookup(commonsyms, s); sym != nil {
		return sym
	}

	if this.locsyms != nil {
		tracer.Trace(debSymbol, " looking for %s in Hoc.locsyms %v", s, this.locsyms)
		if sym := lookup(this.locsyms, s); sym != nil {
			return sym
		}
	}

	tracer.Trace(debSymbol, "looking for %s in Hoc.symbols %v", s, this.symbols)

	return lookup(this.symbols, s)
}

func newSymbol(s string, t int, v symval) *symbol {
	defer tracer.Exit(tracer.Enter())
	sp := new(symbol)
	sp.name = s
	sp.t = t
	sp.val = v
	return sp
}

func formallist(frml *symbol, list *formal) *formal { // add formal to list 
	defer tracer.Exit(tracer.Enter())

	f := new(formal)
	f.sym = frml
	f.next = list
	return f
}


// Symbols based on 
type symbols []*symbol

func lookup(syms symbols, s string) *symbol { // find s in symbol table 
	defer tracer.Exit(tracer.Enter())

	tracer.Trace(debSymbol, "looking for %s in %v", s, syms)
	for v := range syms {
		tracer.Trace(debSymbol, "v : %v", v)

		if syms[v].name == s {
			return syms[v]
		}
	}
	tracer.Trace(debSymbol, "%s not found", s)
	return nil
}

func install(syms symbols, sym *symbol) symbols { // install s in symbol table 
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debSymbol, "installing %s in %v (len:%v, cap: %v  type %T)", sym.name, syms, len(syms), cap(syms), syms)	
	syms = append(syms, sym)
	tracer.Trace(debSymbol, "syms %v", syms)
	return syms
}


func newSymbols() symbols {
	return make(symbols, 0, nSYMBOLS)
}

/*

// symbols based on map
type symbols map[string]*symbol

func lookup(syms symbols, s string) *symbol {	// find s in symbol table 
	defer tracer.Exit(tracer.Enter()))

	tracer.Trace(debSymbol, "looking for %s in %v", s, syms)
	sym := syms[s]
	if sym != nil {
		return sym
	}
	tracer.Trace(debSymbol, "%s not found", s)
	return nil
}


func install(syms symbols, sym *symbol) symbols {	// install s in symbol table 	
	syms[sym.name] = sym, true
	return syms
}


func newSymbols() symbols {
	return make(symbols)
}


*/
