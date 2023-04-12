package lexer_test

import (
	"testing"

	"github.com/cesarFuhr/monkey/lexer"
	"github.com/cesarFuhr/monkey/token"
	"github.com/matryer/is"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for _, tt := range tests {
		is := is.New(t)
		tok := l.NextToken()

		is.Equal(tok.Type, tt.expectedType)       // Token type is wrong
		is.Equal(tok.Literal, tt.expectedLiteral) // Literal is wrong
	}
}
