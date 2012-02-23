package hoc

import (
	"math"
	"fmt"
)

func (this *Hoc) initcode() {
	defer tracer.Exit(tracer.Enter())
	if this.stack == nil {
		this.stack = make([]datum, 1, nSTACK)
	} else {
		this.stack = this.stack[0:1]
	}
	if this.prog == nil {
		this.prog = make([]inst, 0, nPROG)
	}
	if this.frame == nil {
		this.frame = make([]frame, 1, nFRAME)
	} else {
		this.frame = this.frame[0:1]
	}
	if this.symbols == nil {
		this.symbols = newSymbols()
	}
	this.progp = this.progbase
	
	this.returning = false
}

func (this *Hoc) nop() {
	defer tracer.Exit(tracer.Enter())
}

func (this *Hoc) push(d datum) {
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debCode, "push %v", d)
	if len(this.stack)+1 == cap(this.stack) {
		panic("stack too deep")
	}
	this.stack = append(this.stack, d)
}

func (this *Hoc) pop() datum {
	defer tracer.Exit(tracer.Enter())

	if len(this.stack) == 0 {
		panic("stack underflow")
	}

	d := this.stack[len(this.stack) - 1]
	this.stack = this.stack[0 : len(this.stack)-1]
	tracer.Trace(debCode, "pop %v", d)
	return d
}

func (this *Hoc) xpop() { // pop when no value is wanted
	defer tracer.Exit(tracer.Enter())

	if len(this.stack) == 0 {
		panic("this.stack underflow")
	}
	this.stack = this.stack[0 : len(this.stack)-1]

}

func (this *Hoc) cpush() {
	defer tracer.Exit(tracer.Enter())
	s := this.prog[this.pc]
	this.pc++
	tracer.Trace(debCode, "cpush: %v (%T)", s, s)
	switch s.(type) {
	case int:
		this.push(s.(int))
	case float64:
		this.push(s.(float64))
	case string:
		this.push(s.(string))
	case bool:
		this.push(s.(bool))
	default:
		panic("unexpected type for 'cpush' operation")
	}
}

func (this *Hoc) vpush() {
	defer tracer.Exit(tracer.Enter())
	val := this.prog[this.pc]
	this.pc++
	tracer.Trace(debCode, "vpush: %v (%T)", val, val)
	switch val.(type) {
	case int:
		this.push(val.(int))
	case float64:
		this.push(val.(float64))
	case string:
		this.push(val.(string))
	case *symbol:
		this.push(val.(*symbol))
	case bool:
		this.push(val.(bool))
	default:
		panic("unexpected type for 'vpush' operation")
	}
}

func (this *Hoc) checkcond() bool {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	switch d.(type) {
	case int:
		return d.(int) != 0
	case bool:
		return d.(bool)
	}
	panic("unexpected type in condition expression")
	return false
}

func (this *Hoc) whilecode() {
	defer tracer.Exit(tracer.Enter())
	savepc := this.pc
	this.execute(savepc + 2) // condition
	for this.checkcond() {
		this.execute(this.prog[savepc].(int)) // body
		if this.returning {
			break
		}
		this.execute(savepc + 2) // condition
	}
	if !this.returning {
		this.pc = this.prog[savepc+1].(int) // next stmt
	}
}

func (this *Hoc) forcode() {
	defer tracer.Exit(tracer.Enter())

	savepc := this.pc

	this.execute(savepc + 4) // precharge
	this.pop()
	this.execute(this.prog[savepc].(int)) // condition

	for this.checkcond() {
		this.execute(this.prog[savepc+2].(int)) // body
		if this.returning {
			break
		}
		this.execute(this.prog[savepc+1].(int)) // post loop
		this.pop()
		this.execute(this.prog[savepc].(int)) // condition
	}
	if !this.returning {
		this.pc = this.prog[savepc+3].(int) // next stmt
	}
}

func (this *Hoc) ifcode() {
	defer tracer.Exit(tracer.Enter())

	savepc := this.pc // then part

	this.execute(savepc + 3) // condition

	if this.checkcond() {
		this.execute(this.prog[savepc].(int))
	} else if this.prog[savepc+1] != nil { // else part?
		this.execute(this.prog[savepc+1].(int))
	}
	if !this.returning {
		this.pc = this.prog[savepc+2].(int) // next stmt
	}
}

func (this *Hoc) define(sp *symbol, f *formal) { // put func proc in symbol table
	defer tracer.Exit(tracer.Enter())

	fd := new(fndefn)
	fd.code = this.progbase    // start of code
	this.progbase = this.progp // next code starts here
	fd.formals = f
	var n int
	for n = 0; f != nil; f = f.next {
		n++
	}
	fd.nargs = n
	sp.val = fd
}

func (this *Hoc) call() { // call a function
	defer tracer.Exit(tracer.Enter())

	sp := this.prog[this.pc].(*symbol) // symbol table entry
	// for function
	if len(this.frame) + 2 == cap(this.frame) {
		panic(sp.name + " call nested too deeply")
	}
	var fr frame
	fr.sp = sp
	fr.nargs = this.prog[this.pc+1].(int)
	fr.retpc = this.pc + 2
	fr.argn = len(this.stack) - 1 // last argument
	this.frame = append(this.frame, fr)

	if this.frame[len(this.frame) - 1].nargs != sp.val.(*fndefn).nargs {
		panic(sp.name + " called with wrong number of arguments")
	}
	// bind formals
	f := sp.val.(*fndefn).formals
	arg := len(this.stack) - this.frame[len(this.frame) - 1].nargs
	for f != nil {
		s := new(saveval)
		s.val = f.sym.val
		s.t = f.sym.t
		s.next = f.save
		f.save = s

		switch this.stack[arg].(type) {
		case int:
			f.sym.val = this.stack[arg].(int)
		case float64:
			f.sym.val = this.stack[arg].(float64)
		default:
			panic("unexpected type for formal operand")
		}
		f.sym.t = VAR
		tracer.Trace(debCode, "formal VAR %d f(%v, %v) %s %v save %v",
			VAR, f, f.sym, f.sym.name, this.stack[arg], s)
		f = f.next
		arg++
	}
	for i := 0; i < this.frame[len(this.frame) - 1].nargs; i++ {
		this.pop() // this.pop arguments; no longer needed
	}
	tracer.Trace(debCode, "formals %v", sp.val.(*fndefn).formals)
	this.execute(sp.val.(*fndefn).code)
	this.returning = false
}

func (this *Hoc) restore(sp *symbol) { // restore formals associated with symbol
	defer tracer.Exit(tracer.Enter())

	var f *formal
	var s *saveval

	f = sp.val.(*fndefn).formals
	for f != nil {
		s = f.save
		if s == nil { // more actuals than formals
			break
		}
		f.sym.val = s.val
		f.sym.t = s.t
		f.save = s.next
		s = nil
		f = f.next
	}
}


func (this *Hoc) restoreall() { // restore all variables in case of error
	defer tracer.Exit(tracer.Enter())
	for len(this.frame) > 1 && this.frame[len(this.frame) - 1].sp != nil {
		this.restore(this.frame[len(this.frame) - 1].sp)
		this.frame = this.frame[0:len(this.frame) - 1]
	}	
}

func (this *Hoc) ret() { // common return from func or proc
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debCode, "this before %v", this.frame)
	// restore formals
	this.restore(this.frame[len(this.frame) - 1].sp)
	this.pc = this.frame[len(this.frame) - 1].retpc
	this.frame = this.frame[0:len(this.frame) - 1]
	tracer.Trace(debCode, "this before %v", this.frame)
	this.returning = true
}

func (this *Hoc) funcret() { // return from a function
	defer tracer.Exit(tracer.Enter())

	var d datum
	if this.frame[len(this.frame) - 1].sp.t == PROCEDURE {
		panic(this.frame[len(this.frame) - 1].sp.name + " (proc) returns value")
	}
	d = this.pop() // preserve function return value
	this.ret()
	this.push(d)
}


func (this *Hoc) procret() { // return from a procedure
	defer tracer.Exit(tracer.Enter())

	if this.frame[len(this.frame) - 1].sp.t == FUNCTION {
		panic(this.frame[len(this.frame) - 1].sp.name + " func returns no value")
	}
	this.ret()
}

func (this *Hoc) bltin() {
	defer tracer.Exit(tracer.Enter())

	d := this.pop()
	f := this.prog[this.pc].(bltin)
	this.pc++
	tracer.Trace(debCode, "builtin: %v (%T)", d, d)
	switch d.(type) {
	case int:
		d = f(float64(d.(int))).(datum)
	case float64:
		d = f(d.(float64)).(datum)
	default:
		panic("unexpected type for builtin operation")
	}
	this.push(d)
}

func (this *Hoc) add() {
	defer tracer.Exit(tracer.Enter())

	d2 := this.pop()
	d1 := this.pop()
	tracer.Trace(debCode, "%v + %v", d1, d2)

	switch d1.(type) {
	case int:
		if v, ok := d2.(int); ok {
			this.push(d1.(int) + v)
			return
		} else if v, ok := d2.(float64); ok {
			this.push(float64(d1.(int)) + v)
			return
		}
	case float64:
		if v, ok := d2.(float64); ok {
			this.push(d1.(float64) + v)
			return
		} else if v, ok := d2.(int); ok {
			this.push(d1.(float64) + float64(v))
			return
		}
	}
	panic("type mismatch for 'add' operation")
}

func (this *Hoc) sub() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	tracer.Trace(debCode, "%v - %v", d1, d2)

	switch d1.(type) {
	case int:
		if v, ok := d2.(int); ok {
			this.push(d1.(int) - v)
			return
		} else if v, ok := d2.(float64); ok {
			this.push(float64(d1.(int)) - v)
			return
		}
	case float64:
		if v, ok := d2.(float64); ok {
			this.push(d1.(float64) - v)
			return
		} else if v, ok := d2.(int); ok {
			this.push(d1.(float64) - float64(v))
			return
		}
	}
	panic("type mismatch for 'sub' operation")
}

func (this *Hoc) mul() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	tracer.Trace(debCode, "%v * %v", d1, d2)

	switch d1.(type) {
	case int:
		if v, ok := d2.(int); ok {
			this.push(d1.(int) * v)
			return
		} else if v, ok := d2.(float64); ok {
			this.push(float64(d1.(int)) * v)
			return
		}
	case float64:
		if v, ok := d2.(float64); ok {
			this.push(d1.(float64) * v)
			return
		} else if v, ok := d2.(int); ok {
			this.push(d1.(float64) * float64(v))
			return
		}
	}
	panic("type mismatch for 'mul' operation")
}

func (this *Hoc) div() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	tracer.Trace(debCode, "%v / %v", d1, d2)
	switch d1.(type) {
	case int:
		if v, ok := d2.(int); ok {
			if v == 0 {
				panic("division by zero")
			}
			this.push(d1.(int) / v)
			return
		} else if v, ok := d2.(float64); ok {
			if v == 0.0 {
				panic("division by zero")
			}
			this.push(float64(d1.(int)) / v)
			return
		}
	case float64:
		if v, ok := d2.(float64); ok {
			if v == 0.0 {
				panic("division by zero")
			}
			this.push(d1.(float64) / v)
			return
		} else if v, ok := d2.(int); ok {
			if v == 0 {
				panic("division by zero")
			}
			this.push(d1.(float64) / float64(v))
			return
		}
	}
	panic("type mismatch for 'div' operation")
}

func (this *Hoc) mod() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	tracer.Trace(debCode, "%v %% %v", d1, d2)

	v1, ok1 := d1.(int)
	v2, ok2 := d2.(int)
	if !ok1 || !ok2 {
		panic("type mismatch for 'mod' operation")
	}
	if v2 == 0 {
		panic("division by zero")
	}
	this.push(v1 % v2)
}

func (this *Hoc) negate() {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	switch d.(type) {
	case int:
		this.push(-d.(int))
	case float64:
		this.push(-d.(float64))
	default:
		panic("unexpected type for 'negate' operation")
	}
}

func (this *Hoc) verify(s *symbol) *symbol {
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debCode, "%v", s)
	if s.t != VAR && s.t != UNDEF {
		panic("attempt to evaluate non-variable " + s.name)
	}
	if s.t == UNDEF {
		tracer.Trace(debCode, "variable %v", s)
		panic("undefined variable " + s.name)
	}
	return s
}

func (this *Hoc) eval() { // evaluate variable on stack
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	s := d.(*symbol)
	tracer.Trace(debCode, "%v", s)
	this.verify(s)
	switch s.val.(type) {
	case int:
		this.push(s.val.(int))
	case float64:
		this.push(s.val.(float64))
	case string:
		this.push(s.val.(string))
	case symbols:
		this.push(s.val.(symbols))
	case bool:
		this.push(s.val.(bool))
	case []symval:
		this.push(s.val.([]symval))
	default:
		panic("unexpected type for 'eval' operation")
	}
}

func (this *Hoc) elemeval() { // evaluate element of array on stack
	defer tracer.Exit(tracer.Enter())

	d := this.pop()
	arr := d.(*symbol).val.([]symval)
	tracer.Trace(debCode, "%v", arr)
	d = this.pop()
	i, ok := d.(int)
	if !ok {
		panic("index of array must be int")
	}
	if i < 0 || i >= len(arr) {
		panic("index of array is out of range")
	}
	this.push(arr[i].(datum))
}


func (this *Hoc) countof() { // count of array elements
	defer tracer.Exit(tracer.Enter())

	d := this.pop()
	tracer.Trace(debCode, "countof %v", d)
	arr, ok := d.(*symbol).val.([]symval)
	if !ok {
		panic("unexpected type for 'countof' operation")
	}
	this.push(len(arr))
}


func (this *Hoc) preinc() {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	s := d.(*symbol)
	tracer.Trace(debCode, "%v", s)
	this.verify(s)
	switch s.val.(type) {
	case int:
		s.val = s.val.(int) + 1
	case float64:
		s.val = s.val.(float64) + 1.0
	default:
		panic("unexpected type for 'preinc' operation")
	}
	this.push(s.val.(datum))
}

func (this *Hoc) predec() {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	s := d.(*symbol)
	tracer.Trace(debCode, "%v", s)
	this.verify(s)
	switch s.val.(type) {
	case int:
		s.val = s.val.(int) - 1
	case float64:
		s.val = s.val.(float64) - 1.0
	default:
		panic("unexpected type for 'predec' operation")
	}
	this.push(s.val.(datum))
}

func (this *Hoc) postinc() {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	s := d.(*symbol)
	tracer.Trace(debCode, "%v", s)
	this.verify(s)
	this.push(s.val.(datum))

	switch s.val.(type) {
	case int:
		s.val = s.val.(int) + 1
	case float64:
		s.val = s.val.(float64) + 1.0
	default:
		panic("unexpected type for 'postinc' operation")
	}
}

func (this *Hoc) postdec() {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	s := d.(*symbol)
	tracer.Trace(debCode, "%v", s)
	this.verify(s)
	this.push(s.val.(datum))
	switch s.val.(type) {
	case int:
		s.val = s.val.(int) - 1
	case float64:
		s.val = s.val.(float64) - 1.0
	default:
		panic("unexpected type for 'postdec' operation")
	}
}

func (this *Hoc) gt() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	var f1, f2 float64
	switch d1.(type) {
	case int:
		f1 = float64(d1.(int))
	case float64:
		f1 = d1.(float64)
	default:
		panic("unexpected type for 'qt' operation")
	}
	switch d2.(type) {
	case int:
		f2 = float64(d2.(int))
	case float64:
		f2 = d2.(float64)
	default:
		panic("unexpected type for 'gt' operation")
	}
	tracer.Trace(debCode, "gt? %d %d", f1, f2)
	this.push(f1 > f2)
}

func (this *Hoc) lt() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	var f1, f2 float64
	switch d1.(type) {
	case int:
		f1 = float64(d1.(int))
	case float64:
		f1 = d1.(float64)
	default:
		panic("unexpected type for 'lt' operation")
	}
	switch d2.(type) {
	case int:
		f2 = float64(d2.(int))
	case float64:
		f2 = d2.(float64)
	default:
		panic("unexpected type for 'lt' operation")
	}
	tracer.Trace(debCode, "lt? %d %d", f1, f2)
	this.push(f1 < f2)
}

func (this *Hoc) ge() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	var f1, f2 float64
	switch d1.(type) {
	case int:
		f1 = float64(d1.(int))
	case float64:
		f1 = d1.(float64)
	default:
		panic("unexpected type for 'ge' operation")
	}
	switch d2.(type) {
	case int:
		f2 = float64(d2.(int))
	case float64:
		f2 = d2.(float64)
	default:
		panic("unexpected type for 'ge' operation")
	}
	tracer.Trace(debCode, "ge? %d %d", f1, f2)
	this.push(f1 >= f2)
}

func (this *Hoc) le() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	var f1, f2 float64
	switch d1.(type) {
	case int:
		f1 = float64(d1.(int))
	case float64:
		f1 = d1.(float64)
	default:
		panic("unexpected type for 'le' operation")
	}
	switch d2.(type) {
	case int:
		f2 = float64(d2.(int))
	case float64:
		f2 = d2.(float64)
	default:
		panic("unexpected type for 'le' operation")
	}
	tracer.Trace(debCode, "le? %d %d", f1, f2)
	this.push(f1 <= f2)
}

func (this *Hoc) eq() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	var f1, f2 float64
	switch d1.(type) {
	case int:
		f1 = float64(d1.(int))
	case float64:
		f1 = d1.(float64)
	case bool:
		if v, ok := d2.(bool); ok {
			this.push(d1.(bool) == v)
			return
		}
		panic("unexpected type for 'eq' operation")
	default:
		panic("unexpected type for 'eq' operation")
	}
	switch d2.(type) {
	case int:
		f2 = float64(d2.(int))
	case float64:
		f2 = d2.(float64)
	case bool:
		if v, ok := d1.(bool); ok {
			this.push(d2.(bool) == v)
			return
		}
		panic("unexpected type for 'eq' operation")
	default:
		panic("unexpected type for 'eq' operation")
	}
	tracer.Trace(debCode, "eq? %d %d", f1, f2)
	this.push(f1 == f2)
}

func (this *Hoc) ne() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	var f1, f2 float64
	switch d1.(type) {
	case int:
		f1 = float64(d1.(int))
	case float64:
		f1 = d1.(float64)
	case bool:
		if v, ok := d2.(bool); ok {
			this.push(d1.(bool) != v)
			return
		}
		panic("unexpected type for 'ne' operation")
	default:
		panic("unexpected type for 'ne' operation")
	}
	switch d2.(type) {
	case int:
		f2 = float64(d2.(int))
	case float64:
		f2 = d2.(float64)
	case bool:
		if v, ok := d1.(bool); ok {
			this.push(d2.(bool) != v)
			return
		}
		panic("unexpected type for 'ne' operation")
	default:
		panic("unexpected type for 'ne' operation")
	}
	tracer.Trace(debCode, "ne? %d %d", f1, f2)
	this.push(f1 != f2)
}

func (this *Hoc) and() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	switch d1.(type) {
	case int:
		if v, ok := d2.(int); ok {
			this.push(d1.(int) != 0 && v != 0)
			return
		}
	case float64:
		if v, ok := d2.(float64); ok {
			this.push(d1.(float64) != 0.0 && v != 0.0)
			return
		}
	case bool:
		if v, ok := d2.(bool); ok {
			this.push(d1.(bool) && v)
			return
		}
	}
	panic("type mismatch for 'and' operation")
}

func (this *Hoc) or() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	switch d1.(type) {
	case int:
		if v, ok := d2.(int); ok {
			this.push(d1.(int) != 0 || v != 0)
			return
		}
	case float64:
		if v, ok := d2.(float64); ok {
			this.push(d1.(float64) != 0.0 || v != 0.0)
			return
		}
	case bool:
		if v, ok := d2.(bool); ok {
			this.push(d1.(bool) || v)
			return
		}
	}
	panic("type mismatch for 'or' operation")
}

func (this *Hoc) not() {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	switch d.(type) {
	case int:
		this.push(d.(int) == 0)
	case float64:
		this.push(d.(int) == 0.0)
	case bool:
		this.push(!d.(bool))
	default:
		panic("unexpected type for 'not' operation")
	}
}

func (this *Hoc) power() {
	defer tracer.Exit(tracer.Enter())
	d2 := this.pop()
	d1 := this.pop()
	var f1, f2 float64
	switch d1.(type) {
	case int:
		f1 = float64(d1.(int))
	case float64:
		f1 = d1.(float64)
	default:
		panic("unexpected type for 'power' operation")
	}
	switch d2.(type) {
	case int:
		f2 = float64(d2.(int))
	case float64:
		f2 = d2.(float64)
	default:
		panic("unexpected type for 'power' operation")
	}
	this.push(math.Pow(f1, f2))
}

func (this *Hoc) assign() {
	defer tracer.Exit(tracer.Enter())
	d1 := this.pop()
	d2 := this.pop()
	s := d1.(*symbol)
	tracer.Trace(debCode, "assign %v (%T), %v (%T) ", d1, d1, d2, d2)
	if s.t != VAR && s.t != UNDEF {
		panic("assignment to non-variable " + s.name)
	}
	tracer.Trace(debCode, "before assignment %v", s.val)
	switch d2.(type) {
	case int:
		s.val = d2.(int)
	case float64:
		s.val = d2.(float64)
	case string:
		s.val = d2.(string)
	case bool:
		s.val = d2.(bool)
	default:
		panic("unexpected type for 'assign' operation")
	}
	tracer.Trace(debCode, "after assignment %v", s.val)
	s.t = VAR
	this.push(d2)
}

func (this *Hoc) addeq() {
	defer tracer.Exit(tracer.Enter())
	d1 := this.pop()
	d2 := this.pop()
	s := d1.(*symbol)
	tracer.Trace(debCode, "%v", s)
	if s.t != VAR && s.t != UNDEF {
		panic("assignment to non-variable " + s.name)
	}
	switch s.val.(type) {
	case int:
		switch d2.(type) {
		case int:
			s.val = s.val.(int) + d2.(int)
		case float64:
			s.val = s.val.(int) + int(d2.(float64))
		default:
			panic("type mismatch for 'addeq' operation")
		}
	case float64:
		switch d2.(type) {
		case int:
			s.val = s.val.(float64) + float64(d2.(int))
		case float64:
			s.val = s.val.(float64) + d2.(float64)
		default:
			panic("type mismatch for 'addeq' operation")
		}
	default:
		panic("type mismatch for 'addeq' operation")
	}
	s.t = VAR
	this.push(s.val.(datum))
}

func (this *Hoc) subeq() {
	defer tracer.Exit(tracer.Enter())
	d1 := this.pop()
	d2 := this.pop()
	s := d1.(*symbol)
	tracer.Trace(debCode, "%v", s)
	if s.t != VAR && s.t != UNDEF {
		panic("assignment to non-variable " + s.name)
	}
	switch s.val.(type) {
	case int:
		switch d2.(type) {
		case int:
			s.val = s.val.(int) - d2.(int)
		case float64:
			s.val = s.val.(int) - int(d2.(float64))
		default:
			panic("type mismatch for 'subeq' operation")
		}
	case float64:
		switch d2.(type) {
		case int:
			s.val = s.val.(float64) - float64(d2.(int))
		case float64:
			s.val = s.val.(float64) - d2.(float64)
		default:
			panic("type mismatch for 'subeq' operation")
		}
	default:
		panic("type mismatch for 'subeq' operation")
	}
	s.t = VAR
	this.push(s.val.(datum))
}

func (this *Hoc) muleq() {
	defer tracer.Exit(tracer.Enter())
	d1 := this.pop()
	d2 := this.pop()
	s := d1.(*symbol)
	tracer.Trace(debCode, "%v", s)
	if s.t != VAR && s.t != UNDEF {
		panic("assignment to non-variable " + s.name)
	}
	switch s.val.(type) {
	case int:
		switch d2.(type) {
		case int:
			s.val = s.val.(int) * d2.(int)
		case float64:
			s.val = s.val.(int) * int(d2.(float64))
		default:
			panic("type mismatch for 'muleq' operation")
		}
	case float64:
		switch d2.(type) {
		case int:
			s.val = s.val.(float64) * float64(d2.(int))
		case float64:
			s.val = s.val.(float64) * d2.(float64)
		default:
			panic("type mismatch for 'muleq' operation")
		}
	default:
		panic("type mismatch for 'muleq' operation")
	}
	s.t = VAR
	this.push(s.val.(datum))
}

func (this *Hoc) diveq() {
	defer tracer.Exit(tracer.Enter())
	d1 := this.pop()
	d2 := this.pop()
	s := d1.(*symbol)
	tracer.Trace(debCode, "%v", s)
	if s.t != VAR && s.t != UNDEF {
		panic("assignment to non-variable " + s.name)
	}
	switch s.val.(type) {
	case int:
		switch d2.(type) {
		case int:
			s.val = s.val.(int) / d2.(int)
		case float64:
			s.val = s.val.(int) / int(d2.(float64))
		default:
			panic("type mismatch for 'diveq' operation")
		}
	case float64:
		switch d2.(type) {
		case int:
			s.val = s.val.(float64) / float64(d2.(int))
		case float64:
			s.val = s.val.(float64) / d2.(float64)
		default:
			panic("type mismatch for 'diveq' operation")
		}
	default:
		panic("type mismatch for 'diveq' operation")
	}
	s.t = VAR
	this.push(s.val.(datum))
}

func (this *Hoc) modeq() {
	defer tracer.Exit(tracer.Enter())
	d1 := this.pop()
	d2 := this.pop()
	s := d1.(*symbol)
	tracer.Trace(debCode, "%v", s)
	if s.t != VAR && s.t != UNDEF {
		panic("assignment to non-variable " + s.name)
	}
	v1, ok1 := s.val.(int)
	v2, ok2 := d2.(int)
	if !ok1 || !ok2 {
		panic("type mismatch for 'modeq' operation")
	}
	v1 %= v2
	d2 = v1
	s.val = v1
	s.t = VAR
	this.push(d2)
}

func (this *Hoc) print(s string) {
	this.out <- s
}

func (this *Hoc) printValue(val symval, indent int, quote bool) {
	defer tracer.Exit(tracer.Enter())
	if val == nil {
		this.print("null")
		return
	}
	switch val.(type) {
	case string:
		if len(val.(string)) > 0 {
			if quote {
				this.print(fmt.Sprintf("\"%s\"", val.(string)))
			} else {
				this.print(val.(string))
			}
		}
	case symbols:
		this.printSymbols(val.(symbols), indent)
	case []symval:
		this.printArray(val.([]symval), indent)
	default:
		this.print(fmt.Sprintf("%v", val))
	}
}

func (this *Hoc) newLine() {
	if this.NewLine {
		this.print("\n")
	}
}


func (this *Hoc) printArray(ar []symval, indent int) {
	if len(ar) == 0 {
		this.print("[]")
		return
	}

	this.print("[")
	this.newLine()
	first := true
	for v := range ar {
		if !first {
			this.print(",")
			this.newLine()
		}
		for i := 0; i < indent+1; i++ {
			this.print(this.IndentSym)
		}
		this.printValue(ar[v], indent, true)
		first = false
	}
	this.newLine()
	for i := 0; i < indent; i++ {
		this.print(this.IndentSym)
	}
	this.print("]")
}

func (this *Hoc) printSymbols(syms symbols, indent int) {
	defer tracer.Exit(tracer.Enter())
	if len(syms) == 0 {
		this.print("{}")
		return
	}
	this.print("{")
	this.newLine()
	first := true
	for k, v := range syms {
		tracer.Trace(debCode, "printing syms[%v]=%v", k, v)
		if v.t != VAR {
			continue
		}
		if !first {
			this.print(",")
			this.newLine()
		}
		tracer.Trace(debCode, "indent %v", indent)
		for i := 0; i < indent+1; i++ {
			this.print(this.IndentSym)
		}
		this.print(fmt.Sprintf("\"%s\": ", v.name))
		switch v.val.(type) {
		case symbols:
			this.printSymbols(v.val.(symbols), indent+1)
		case []symval:
			this.printArray(v.val.([]symval), indent+1)
		default:
			this.printValue(v.val, indent, true)
		}
		first = false
	}
	this.newLine()
	for i := 0; i < indent; i++ {
		this.print(this.IndentSym)
	}
	this.print("}")
}

func (this *Hoc) printtop() { // this.pop top value from this.stack, print it
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	this.printValue(d.(symval), 0, false)
}

func (this *Hoc) printall() { // this.pop top value from this.stack, print it
	defer tracer.Exit(tracer.Enter())
	this.printSymbols(this.symbols, 0)
	this.newLine()
}

func (this *Hoc) prexpr() { // print numeric value
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	this.printValue(d.(symval), 0, false)
}

/*func (this *Hoc) prstr() {	// print string value
	defer tracer.Exit(tracer.Enter())
	s := this.prog[this.pc].(*symbol)
	this.printValue(s.val.(symval), 0, false)
	this.pc++
}
*/

// print with formatting

func (this *Hoc) prfstr() {
	defer tracer.Exit(tracer.Enter())
	d := this.pop()
	f, ok := d.(string)
	if !ok {
		f, ok = d.(*symbol).val.(string)
	}
	if !ok {
		panic("unexpected type for format paramers of 'printf' operation")
	}
	nargs := this.pop().(int)
	v := make([]interface{}, nargs)
	tracer.Trace(debCode, "format: %s arguments: %d", f, nargs)
	for i := nargs - 1; i >= 0; i-- {
		v[i] = this.pop()
	}
	this.print(fmt.Sprintf(f, v...))
}

/*func (this *Hoc) varread(){ // read into variable
	var d datum
	var bin *Biobuf
	var v *Symbol = *this.pc
	this.pc++
	var c int

Again:
	do
	c = Bgetc(bin)
	while(c == ' ' || c == '\t')
	if c == Beof {
	Iseof:
		if moreinput() {
			goto Again
		}
		d.val = 0.0
		v.u.val = 0.0

		goto Return
	}

	if strchr("+-.0123456789", c) == 0 {
		panic("non-number read into " + v.name)
	}
	Bungetc(bin)
	if Bgetd(bin, &v.u.val) == Beof {
		goto Iseof
	} else {
		d.val = 1.0
	}
Return:
	v.t = VAR
	this.push(d)
}
*/

func (this *Hoc) code(f inst) int { // install one instruction or operand
	defer tracer.Exit(tracer.Enter())
	oprogp := this.progp
	if this.progp+1 == cap(this.prog) {
		panic("program too big")
	}
	this.prog = this.prog[0 : this.progp+1]
	tracer.Trace(debCode, "pc: %d, code: %v", this.progp, f)
	this.prog[this.progp] = f
	this.progp++
	return oprogp
}

func (this *Hoc) execute(p int) {
	defer tracer.Exit(tracer.Enter())
	tracer.Trace(debCode, "program index p = %d, len %d", p, len(this.prog))
	for this.pc = p; this.prog[this.pc] != nil && !this.returning; {
		tracer.Trace(debCode, "pc: %d, code: %v", this.pc, this.prog[this.pc])

		proc := this.prog[this.pc].(func(*Hoc))
		this.pc++
		proc(this)
	}
}
