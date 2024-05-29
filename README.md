# Pratt Parser 四则运算解析器

## Example

```go
package main

import (
	"fmt"
	"ppc/lexer"
	"ppc/parser"
)

func main() {
    // 此处输入运算式
	input := "1 + 9 * (3 - 2)+ (9 * 6)"
	l := lexer.New(input)
	p := parser.New(l)
	res := p.ParseExpression(0)
	fmt.Println(res)
}
```