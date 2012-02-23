package hoc

import (
	"testing"
)


type test struct {
	test   string
	result string
}

var jsonTests = [...]test{
	{"{\"1\": 1,\"1.2\": 1.2,\"строка\": \"строка\",\"true\": true,\"false\": false,\"array\": [1,1.2,\"строка\",false],\"object\": {\"1\": 1,\"1.2\": 1.2,\"строка\": \"строка\",\"true\": true,\"false\": false,\"array\": [1,1.2,\"строка\",false]}};print",
		"{\"1\": 1,\"1.2\": 1.2,\"строка\": \"строка\",\"true\": true,\"false\": false,\"array\": [1,1.2,\"строка\",false],\"object\": {\"1\": 1,\"1.2\": 1.2,\"строка\": \"строка\",\"true\": true,\"false\": false,\"array\": [1,1.2,\"строка\",false]}}",
	},
	{"print object", "{\"1\": 1,\"1.2\": 1.2,\"строка\": \"строка\",\"true\": true,\"false\": false,\"array\": [1,1.2,\"строка\",false]}"},
	{"print object.строка", "строка"},
	{"print object.array[3]", "false"},
	{"print countof object.array", "4"},
}

var forTests = [...]test{
	{"for(i=0;i<5;i++) print i", "01234"},
	{"for(i=5;i>0;i--) {print i}", "54321"},
}

var whileTests = [...]test{
	{"i=0;while(i<5) print i++", "01234"},
	{"i=5;while(i>0) {print i--}", "54321"},
}

var ifTests = [...]test{
	{"if(true) print \"correct\" else print \"wrong\"", "correct"},
	{"if(false) print \"wrong\" else print \"correct\"", "correct"},
	{"if(1&&1) print\"correct\"", "correct"},
	{"if(1&&0) print \"wrong\" else  print \"correct\"", "correct"},
	{"if(1||1) print \"correct\" else  print \"wrong\"", "correct"},
	{"if(1||0) print \"correct\" else  print \"wrong\"", "correct"},
	{"i=0;if(i!=0) {print \"wrong\"} else print \"correct\"", "correct"},
	{"i=0;if(i==0) print \"correct\" else {print \"correct\"}", "correct"},
}

var localTests = [...]test{
	{"a=true;b=1; proc test() {local a; a = false;b=0; print a, b;}; print a,b; test(); print a,b", "true1false0true0"},
	{"a=true;b=1; func test2() {local a = false;b=0; print a, b; return a;}; print a,b; test2(); print a,b", "true1false0falsetrue0"},
}

var printTests = [...]test{
	{"print 1, 1.2, true, \"string\", \"\\string\\\"", "11.2truestring\\string\\"},
}

var printfTests = [...]test{
	{"printf \"%d %g %t %s\", 1, 1.2, true, \"string\"", "1 1.2 true string"},
	{"printf \"string\"", "string"},
	{"printf \"0809%08g\", 1234.56*100", "080900123456"},
	{"f=\"%d\"; printf f, int(PI)", "3"},
}

var builtinsTest =  [...]test{
	{"cos(0)", "1"},
	{"a=0;cos(a)", "1"},
	{"a=0.0;cos(a)", "1"},
	{"sin(0)", "0"},
	{"a=0;sin(a)", "0"},
	{"a=0.0;sin(a)", "0"},
	{"int(PI)", "3"},
	{"a=PI;int(a)", "3"},
	{"a=3;int(a)", "3"},
	{"log2(4)", "2"},
	{"a=4;log2(a)", "2"},
	{"a=4.0;log2(a)", "2"},
	{"log10(100)", "2"},
	{"a=100;log10(a)", "2"},
	{"a=100.0;log10(a)", "2"},
	{"log(E)","1"},
	{"a=E; log(a)","1"},
}

var miscTests = [...]test{
	{"a=1; a+=1; a", "2"},
	{"a=1; a+=1.0; a", "2"},
	{"a=1.0; a+=1; a", "2"},
	{"a=1.0; a+=1.0; a", "2"},
	{"a=1; a=a+1; a", "2"},
	{"a=1; a=a+1.0; a", "2"},
	{"a=1.0; a=a+1; a", "2"},
	{"a=1.0; a=a+1.0; a", "2"},
	{"a=1; a-=1; a", "0"},
	{"a=1; a-=1.0; a", "0"},
	{"a=1.0; a-=1; a", "0"},
	{"a=1.0; a-=1.0; a", "0"},
	{"a=1; a=a-1; a", "0"},
	{"a=1; a=a-1.0; a", "0"},
	{"a=1.0; a=a-1; a", "0"},
	{"a=1.0; a=a-1.0; a", "0"},
	{"a=1; a*=2; a", "2"},
	{"a=1; a*=2.0; a", "2"},
	{"a=1.0; a*=2; a", "2"},
	{"a=1.0; a*=2.0; a", "2"},
	{"a=1; a=a*2; a", "2"},
	{"a=1; a=a*2.0; a", "2"},
	{"a=1.0; a=a*2; a", "2"},
	{"a=1.0; a=a*2.0; a", "2"},
	{"a=4; a/=2; a", "2"},
	{"a=4; a/=2.0; a", "2"},
	{"a=4.0; a/=2; a", "2"},
	{"a=4.0; a/=2.0; a", "2"},
	{"a=4; a=a/2; a", "2"},
	{"a=4; a=a/2.0; a", "2"},
	{"a=4.0; a=a/2; a", "2"},
	{"a=4.0; a=a/2.0; a", "2"},
	{"a=10; b=4; a%=b;a", "2"},
	{"a=10; b=4; a=a%b;a", "2"},
	{"a=-1; abs(a)", "1"},
	{"a=10; b=4; a>b", "true"},
	{"a=10; b=4.0; a>b", "true"},
	{"a=10.0; b=4; a>b", "true"},
	{"a=10.0; b=4.0; a>b", "true"},
	{"a=10; b=4; a<b", "false"},
	{"a=10; b=4.0; a<b", "false"},
	{"a=10.0; b=4; a<b", "false"},
	{"a=10.0; b=4.0; a<b", "false"},
	{"a=10; b=10; a>=b", "true"},
	{"a=10; b=10.0; a>=b", "true"},
	{"a=10.0; b=10; a>=b", "true"},
	{"a=10.0; b=10.0; a>=b", "true"},
	{"a=10; b=10; a<=b", "true"},
	{"a=10; b=10.0; a<=b", "true"},
	{"a=10.0; b=10; a<=b", "true"},
	{"a=10.0; b=10.0; a<=b", "true"},
	{"a=10; b=10; a==b", "true"},
	{"a=10; b=10.0; a==b", "true"},
	{"a=10.0; b=10; a==b", "true"},
	{"a=10.0; b=10.0; a==b", "true"},
	{"a=true; a==true", "true"},
	{"a=true; true==a", "true"},
	{"a=10; b=10; a==b", "true"},
	{"a=10; b=10.0; a==b", "true"},
	{"a=true; a!=true", "false"},
	{"a=true; true!=a", "false"},
	{"a=10; b=10; a!=b", "false"},
	{"a=10.0; b=10; a!=b", "false"},
	{"a=10; b=10.0; a!=b", "false"},
	{"a=10.0; b=10.0; a!=b", "false"},
	{"a=1; !a", "false"},
	{"a=0; !a", "true"},
	{"a=true; !a", "false"},
	{"a=false; !a", "true"},
	{"a=1; -a", "-1"},
	{"a=1; b=0; a&&b", "false"},
	{"a=1; b=0; a||b", "true"},
	{"a=true; b=false; a&&b", "false"},
	{"a=true; b=false; a||b", "true"},
	{"a=1; b=a++; b", "1"},
	{"a=1; b=++a; b", "2"},
	{"a=1; b=a--; b", "1"},
	{"a=1; b=--a; b", "0"},
	{"a=1.0; b=a++; b", "1"},
	{"a=1.0; b=++a; b", "2"},
	{"a=1.0; b=a--; b", "1"},
	{"a=1.0; b=--a; b", "0"}}

func runTest(t *testing.T, input []test) {
	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			case HocError:
				e := x.(HocError)
				t.Errorf("\npanic: %s\nline: '%s'\nposition: %d", e.Error, e.Line, e.Linepos)
			default:
				t.Errorf("panic: %v\n", x)
			}
		}
	}()
	var h Hoc
	for i := range input {
		in := make(chan string)
		out := make(chan string)
		go func() {
			in <- input[i].test
			close(in)
		}()
		done := make(chan bool)
		go func() {
			var res string
			for true {
				s, ok := <-out
				if !ok {
					break
				}
				res += s
			}
			if res != input[i].result {
				t.Errorf("index: %v\n'%s' doesn't produce '%s', but '%s'",
					i, input[i].test, input[i].result, res)
			}
			done <- true
		}()
		h.Process(in, out)
		<-done
	}
}

func TestJSON(t *testing.T) {
	runTest(t, jsonTests[:])
}

func TestFOR(t *testing.T) {
	runTest(t, forTests[:])
}

func TestWHILE(t *testing.T) {
	runTest(t, whileTests[:])
}

func TestIF(t *testing.T) {
	runTest(t, ifTests[:])
}

func TestLOCAL(t *testing.T) {
	runTest(t, localTests[:])
}

func TestPRINT(t *testing.T) {
	runTest(t, printTests[:])
}

func TestPRINTF(t *testing.T) {
	runTest(t, printfTests[:])
}

func TestBuiltins(t *testing.T) {
	runTest(t, builtinsTest[:])
}

func TestMisc(t *testing.T) {
	runTest(t, miscTests[:])
}
