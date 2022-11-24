package lexer

import (
	"testing"

	"token"
)

func Test_LexNumbersInIdent(t *testing.T) {
	l := New("arg_1")

	var toks []token.Token

	for l.ch != 0 {
		toks = append(toks, l.NextToken())
	}

	got := toks
	expected := []token.Token{
		token.Token{Type: token.IDENT, Literal: "arg_1"},
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected %d, got %d", len(expected), len(got))
	}

	for i, tok := range expected {
		if tok.Type != got[i].Type || tok.Literal != got[i].Literal {
			t.Fatal("Mismatching type or literal")
		}
	}
}

func Test_peekChar(t *testing.T) {
	l := New("arg_1")

	nextCh := l.peekChar()

	if nextCh != 'r' {
		t.Fatal("doesn't match")
	}
}

func Test_IdentsWithNumbersOrUnderscores(t *testing.T) {
	input := `arg_a arg_1 arg1`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "arg_a"},
		{token.IDENT, "arg_1"},
		{token.IDENT, "arg1"},
	}

	l := New(input)

	var toks []token.Token

	for i, tt := range tests {
		tok := l.NextToken()
		toks = append(toks, tok)

		checkType(t, i, tok.Type, tt.expectedType)
		checkLiteral(t, i, tok.Literal, tt.expectedLiteral)
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;!/*<>
		class Foo
			def initialize(arga, argb)
				a == b
				a != b
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
		{token.BANG, "!"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.LT, "<"},
		{token.GT, ">"},

		{token.CLASS, "class"},
		{token.IDENT, "Foo"},

		{token.DEF, "def"},
		{token.IDENT, "initialize"},
		{token.LPAREN, "("},
		{token.IDENT, "arga"},
		{token.COMMA, ","},
		{token.IDENT, "argb"},
		{token.RPAREN, ")"},

		{token.IDENT, "a"},
		{token.EQ, "=="},
		{token.IDENT, "b"},

		{token.IDENT, "a"},
		{token.NOT_EQ, "!="},
		{token.IDENT, "b"},

		{token.END, "end"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		checkType(t, i, tok.Type, tt.expectedType)
		checkLiteral(t, i, tok.Literal, tt.expectedLiteral)
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
