package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	defaultStartingCells = 1024 * 1024
	dirLeft              = -1
	dirRight             = 1
	opIncrement          = '+'
	opDecrement          = '-'
	opMoveRight          = '>'
	opMoveLeft           = '<'
	opPrint              = '.'
	opGetInput           = ','
	opLoopBegin          = '['
	opLoopEnd            = ']'
)

func main() {
	numCells := flag.Uint("cells", defaultStartingCells, "number of memory cells to start with")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("no input file specified\n")
		os.Exit(1)
	}
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Printf("error opening source file: %s\n", err)
		os.Exit(1)
	}
	program, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error reading from source file: %s\n", err)
		file.Close()
		os.Exit(1)
	}
	env := newEnv(int(*numCells))
	exec(program, env)
	file.Close()
}

type environment struct {
	numCells int       // starting number of cells
	cells    []byte    // memory cells
	pos      int       // current cell position
	in       io.Reader // stdin
	out      io.Writer // stdout
}

func newEnv(numCells int) *environment {
	return &environment{
		numCells: defaultStartingCells,
		cells:    make([]byte, 1, numCells),
		in:       os.Stdin,
		out:      os.Stdout,
	}
}

func exec(program []byte, env *environment) {
	inBuf := make([]byte, 1)
	var pc int
	for {
		if pc == len(program) {
			break
		}
		switch program[pc] {
		case opIncrement:
			env.cells[env.pos]++
		case opDecrement:
			env.cells[env.pos]--
		case opMoveRight:
			if len(env.cells)-1 == env.pos {
				env.cells = append(env.cells, 0)
			}
			env.pos++
		case opMoveLeft:
			if env.pos == 0 {
				env.pos = len(env.cells) - 1
			} else {
				env.pos--
			}
		case opPrint:
			fmt.Fprint(env.out, string(env.cells[env.pos]))
		case opGetInput:
			env.in.Read(inBuf)
			env.cells[env.pos] = inBuf[0]
		case opLoopBegin:
			if env.cells[env.pos] == 0 {
				bracket(dirRight, program, &pc)
			}
		case opLoopEnd:
			if env.cells[env.pos] != 0 {
				bracket(dirLeft, program, &pc)
			}
		}
		pc++
	}
}

func bracket(dir int, program []byte, pc *int) {
	for nesting := dir; dir*nesting > 0; *pc += dir {
		switch program[*pc+dir] {
		case opLoopEnd:
			nesting--
		case opLoopBegin:
			nesting++
		}
	}
}
