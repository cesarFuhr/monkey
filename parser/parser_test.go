package parser_test

import (
	"testing"

	"github.com/cesarFuhr/monkey/ast"
	"github.com/cesarFuhr/monkey/lexer"
	"github.com/cesarFuhr/monkey/parser"
	"github.com/matryer/is"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	is := is.New(t)
	is.Equal(s.TokenLiteral(), "let")

	letStmt, ok := s.(*ast.LetStatement)
	is.True(ok) // s not *ast.LetStatement

	is.Equal(letStmt.Name.Value, name)
	is.Equal(letStmt.Name.TokenLiteral(), name)
}
