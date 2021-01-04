package lexer

import (
	"github.com/ryan-berger/monkey/lexer/naive"
	"github.com/ryan-berger/monkey/token"
	"testing"
)

type tokenTest struct {
	expectedType    token.Type
	expectedLiteral string
}

type testCase struct {
	input          string
	expectedTokens []tokenTest
}

var builder Builder = func(input string) Lexer {
	return naive.NewLexer(input)
}

var tests = []testCase{
	{
		input: `=+(){},;`,
		expectedTokens: []tokenTest{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		},
	},
	{
		input: `
let five = 5;
let ten = 10;
let add = fn(x, y) {
  x + y
};
let result = add(five, ten);
!-/*;
5 < 10 > 5;

if (5 < 10) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"`,
		expectedTokens: []tokenTest{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.SEMICOLON, ";"},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.GT, ">"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.IF, "if"},
			{token.LPAREN, "("},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.TRUE, "true"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.ELSE, "else"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.FALSE, "false"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.INT, "10"},
			{token.EQ, "=="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.INT, "10"},
			{token.NEQ, "!="},
			{token.INT, "9"},
			{token.SEMICOLON, ";"},
			{token.STRING, "foobar"},
			{token.STRING, "foo bar"},
			{token.EOF, ""},
		},
	},
}

func TestLexer(t *testing.T) {
	for _, tok := range tests {
		lexer := builder(tok.input)
		for i, tt := range tok.expectedTokens {
			tok := lexer.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - token type wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
			}
		}
	}
}
