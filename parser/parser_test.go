package parser

import (
	"testing"

	"github.com/ryan-berger/monkey/lexer/naive"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 1212;
`
	l := naive.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) == 3 {
		t.Fatalf("Expected 3 statements, received %d", len(program.Statements))
	}


}
