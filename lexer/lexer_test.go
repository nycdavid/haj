package lexer

import (
	"testing"

	"token"
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

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}

func TestDef(t *testing.T) {
	input := "def foo; end"
	l := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEF, "def"},
		{token.IDENT, "foo"},
		{token.SEMICOLON, ";"},
		{token.END, "end"},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		checkType(t, i, tok.Type, tt.expectedType)
		checkLiteral(t, i, tok.Literal, tt.expectedLiteral)
	}
}

func checkType(t *testing.T, i int, actual token.TokenType, expected token.TokenType) {
	if actual != expected {
		t.Fatalf(
			"tests[%d] - type wrong. expected=%q, got=%q",
			i,
			expected,
			actual,
		)
	}
}

func checkLiteral(t *testing.T, i int, actual string, expected string) {
	if actual != expected {
		t.Fatalf(
			"tests[%d] - literal wrong. expected=%q, got=%q",
			i,
			expected,
			actual,
		)
	}
}
