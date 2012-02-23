package main

import (
	"bufio"
	"flag"
	"fmt"
	"hoc"
	"os"
	"strings"
	"unicode/utf8"
)

var expr string

func main() {
	flag.StringVar(&expr, "e", "", "expression")
	flag.Parse()
	var h *hoc.Hoc = new(hoc.Hoc)
	h.IndentSym = "  "
	h.NewLine = true
	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			case hoc.HocError:
				execerror(x.(hoc.HocError))
			default:
				fmt.Fprintf(os.Stderr, "panic: %v\n", x)
			}
		}
	}()

	if flag.NArg() == 0 {
		in := make(chan string)
		out := make(chan string)
		go func() {
			bin := bufio.NewReader(os.Stdin)
			for bin != nil {
				s, err := bin.ReadString('\n')
				if len(s) != 0 {
					in <- s
				}
				if err != nil && len(s) == 0 {
					bin = nil
				}
			}
			close(in)
		}()
		done := make(chan bool)
		go func() {
			for true {
				s, ok := <-out
				if !ok {
					break
				}
				if len(s) == 0 {
					continue
				}
				os.Stdout.WriteString(s)
			}
			done <- true
		}()
		h.Process(in, out)
		<-done
	} else {
		for i := 0; i < flag.NArg(); i++ {
			f, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				continue
			}
			bin := bufio.NewReader(f)
			if bin == nil {
				panic("can't make reader")
			}
			in := make(chan string)
			out := make(chan string)
			go func() {
				for bin != nil {
					s, err := bin.ReadString('\n')
					if len(s) != 0 {
						in <- s
					}
					if err != nil && len(s) == 0 {
						close(in)
						bin = nil
					}
				}
			}()
			done := make(chan bool)
			go func() {
				for true {
					s, ok := <-out
					if !ok {
						break
					}
					if len(s) == 0 {
						continue
					}
					os.Stdout.WriteString(s)
				}
				done <- true
			}()
			h.Process(in, out)
			<-done
		}
	}
}

func execerror(e hoc.HocError) {
	var err string
	if e.Lineno > 0 {
		f := fmt.Sprintf("\n%%s\n%%%ds", utf8.RuneCountInString(e.Line[0:e.Linepos]))
		err = fmt.Sprintf(f, strings.TrimRight(e.Line, "\n"), "^")
	}
	fmt.Fprintf(os.Stderr, "\npanic: %v %s\n", e.Error, err)

}
