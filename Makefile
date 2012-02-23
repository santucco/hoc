all:hoc

hoc : y.go
	go build

clean: 
	rm y.go y.output 

test:
	go test

install:
	go install

y.go: hoc.y
	go tool yacc hoc.y

