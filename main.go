package main

import (
	"fmt"
	"ppc/lexer"
	"ppc/parser"
)

func main() {
	input := "1+ 9 * (3 - 2      )+ (9 * 6)"
	l := lexer.New(input)
	p := parser.New(l)
	res := p.ParseExpression(0)
	fmt.Println(res)
}
