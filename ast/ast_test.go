package ast

import (
	"fmt"
	"testing"
)

func Test_Ast(t *testing.T) {
	p := &Program{}

	str := p.TokenLiteral()
	fmt.Println(str)
}
