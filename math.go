package hoc

import (
	"math"
)
func sin(x float64) symval {
	return errcheck(math.Sin(x), "sin")
}

func cos(x float64) symval {
	return errcheck(math.Cos(x), "cos")
}

func tan(x float64) symval {
	return errcheck(math.Tan(x), "tan")
}

func atan(x float64) symval {
	return errcheck(math.Atan(x), "atan")
}

func tanh(x float64) symval {
	return errcheck(math.Tanh(x), "tanh")
}

func abs(x float64) symval {	
	return errcheck(math.Abs(x), "abs")
}

func log(x float64) symval {
	return errcheck(math.Log(x), "log")
}

func log2(x float64) symval {
	return errcheck(math.Log2(x), "log")
}

func log10(x float64) symval {
	return errcheck(math.Log10(x), "log10")
}

func sqrt(x float64) symval {
	return errcheck(math.Sqrt(x), "sqrt")
}

func exp(x float64) symval {
	return errcheck(math.Exp(x), "exp")
}

func asin(x float64) symval {
	return errcheck(math.Asin(x), "asin")
}

func acos(x float64) symval {
	return errcheck(math.Acos(x), "acos")
}

func sinh(x float64) symval {
	return errcheck(math.Sinh(x), "sinh")
}

func cosh(x float64) symval {
	return errcheck(math.Cosh(x), "cosh")
}

func pow(x float64, y float64) symval {
	return errcheck(math.Pow(x, y), "exponentiation")
}

func integer(x float64) symval {
	if x < -2147483648.0 || x > 2147483647.0 {
		panic("argument out of domain")
	}
	return int(math.Trunc(x))
}


func errcheck(d float64, s string) symval { // check result of library call 
	if math.IsNaN(d) {
		panic(s + ": argument out of domain")
	}
	if math.IsInf(d, 0) {
		panic(s + ": result out of range")
	}
	return d
}
