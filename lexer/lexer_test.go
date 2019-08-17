package lexer

import (
	"github.com/ryan-berger/monkey/lexer/naive"
	"github.com/ryan-berger/monkey/token"
	"testing"
)

var builder Builder = naive.NewLexer

func TestLexer(t *testing.T) {
	input := `=+(){},;`

	lexer := builder(input)

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
		}
	}
}
