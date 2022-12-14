package parser

import (
	"testing"

	"haj/ast"
	"haj/lexer"
	"haj/token"
)

func TestAssignStatements(t *testing.T) {
	input := `
x = 5
y = 10
foobar = 838383
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
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
		if !testAssignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()

	expectedCt := 3
	got := prog.Statements

	if expectedCt != len(got) {
		t.Errorf("expected %d statements, got %d", expectedCt, got)
	}
	for _, s := range got {
		retStmt, ok := s.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s not *ast.ReturnStatement, got=%T", s)
		}
		if retStmt.Token.Type != token.RETURN {
			t.Errorf("expected '%s' token, but got '%s'", token.RETURN, retStmt.Token.Type)
		}
	}
}

func UNUSED(x ...interface{}) {}

func testAssignStatement(t *testing.T, s ast.Statement, name string) bool {
	assignStmt, ok := s.(*ast.AssignStatement)
	if !ok {
		t.Errorf("s not *ast.AssignStatement. got=%T", s)
		return false
	}

	if assignStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, assignStmt.Name.Value)
		return false
	}

	if assignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, assignStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	errorCt := len(errors)

	if errorCt == 0 {
		return
	}

	t.Errorf("parser as %d errors", errorCt)
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
}
