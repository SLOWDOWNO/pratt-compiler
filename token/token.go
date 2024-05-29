package token

type TokenType int

const (
	EOF      TokenType = iota
	NUMBER             //1
	PLUS               //2
	MINUS              //3
	ASTERISK           //4
	SLASH              //5
	LPAREN             //6
	RPAREN             //7
)

type Token struct {
	Type  TokenType
	Value string
}
