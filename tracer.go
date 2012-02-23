package hoc

import(
	"bitbucket.org/santucco/trace"
)

const (
	debLex = trace.Next << iota
	debCode
	debSymbol
	debInit
	debMath
	debState
)

var tracer  trace.Tracer