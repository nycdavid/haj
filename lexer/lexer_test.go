package lexer

import (
	"testing"

	"token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;
		class Foo
			def initialize(arg1, arg2)
			end
		end
	`

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
		{token.CLASS, "class"},
		{token.IDENT, "Foo"},
		{token.DEF, "def"},
		{token.IDENT, "initialize"},
		{token.LPAREN, "("},
		{token.IDENT, "arg1"},
		{token.IDENT, "arg2"},
		{token.RPAREN, ")"},
		{token.END, "end"},
		{token.END, "end"},
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

func Test_methodIdentifiers(t *testing.T) {
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

func Test_numbersInArithmeticOp(t *testing.T) {
	input := "55 + 5"
	l := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "55"},
		{token.PLUS, "+"},
		{token.INT, "5"},
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
