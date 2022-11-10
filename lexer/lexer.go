package lexer

import (
	"token"
)

type Lexer struct {
	input        string
	position     int // current position in input (current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
1. If the read cursor is at the end of the line, set l.ch to NUL
2. Else, set l.ch to the current read position char
3. Set current l.position
4. Set l.readPosition to the next char
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.consumeIf(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		} else if isInt(l.ch) {
			tok.Literal = l.consumeIf(isInt)
			tok.Type = token.INT
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	if l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == 'r' {
		l.readChar()
	}
}

func (l *Lexer) consumeIf(pred func(byte) bool) string {
	position := l.position
	for pred(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

func isInt(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
