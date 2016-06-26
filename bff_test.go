package main

import (
	"bytes"
	"io"
	"testing"
)

func newTestEnv(in io.Reader, out io.Writer) *environment {
	env := newEnv(defaultStartingCells)
	env.in = in
	env.out = out
	return env
}

func TestHelloWorld(t *testing.T) {
	const program = `++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`
	var buf bytes.Buffer
	env := newTestEnv(nil, &buf)
	exec([]byte(program), env)
	const expected = "Hello World!\n"
	if buf.String() != expected {
		t.Errorf("expecting output: %s\n got: %s", buf.String(), expected)
	}
}

func TestReverse(t *testing.T) {
	const program = `+[->,----------]<[+++++++++++.<]`
	var buf bytes.Buffer
	env := newTestEnv(bytes.NewReader([]byte("Hello World!\n")), &buf) // newline is needed here, else it blocks
	exec([]byte(program), env)
	const expected = "!dlroW olleH"
	if buf.String() != expected {
		t.Errorf("expecting output: %s\nreceived: %s\n", expected, buf.String())
	}
}
