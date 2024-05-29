package lexer

import (
	"ppc/token"
	"unicode"
)

type Lexer struct {
	input string
	pos   int
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() token.Token {
	for l.pos < len(l.input) {
		switch l.input[l.pos] {
		case ' ':
			l.pos++
		case '+':
			l.pos++
			return token.Token{Type: token.PLUS, Value: "+"}
		case '-':
			l.pos++
			return token.Token{Type: token.MINUS, Value: "-"}
		case '*':
			l.pos++
			return token.Token{Type: token.ASTERISK, Value: "*"}
		case '/':
			l.pos++
			return token.Token{Type: token.SLASH, Value: "/"}
		case '(':
			l.pos++
			return token.Token{Type: token.LPAREN, Value: "("}
		case ')':
			l.pos++
			return token.Token{Type: token.RPAREN, Value: ")"}
		default:
			if unicode.IsDigit(rune(l.input[l.pos])) {
				start := l.pos
				for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
					l.pos++
				}
				return token.Token{Type: token.NUMBER, Value: l.input[start:l.pos]}
			} else {
				panic("invalid character: " + string(l.input[l.pos]))
			}
		}
	}
	return token.Token{Type: token.EOF}
}
