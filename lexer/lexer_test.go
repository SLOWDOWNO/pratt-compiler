package lexer_test

import (
	"ppc/token"
	"testing"

	"ppc/lexer"
)

func Test_Lexer(t *testing.T) {
	input := `3 + 5 * (2 - 8) / 4`

	test_1 := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.NUMBER, "3"},
		{token.PLUS, "+"},
		{token.NUMBER, "5"},
		{token.ASTERISK, "*"},
		{token.LPAREN, "("},
		{token.NUMBER, "2"},
		{token.MINUS, "-"},
		{token.NUMBER, "8"},
		{token.RPAREN, ")"},
		{token.SLASH, "/"},
		{token.NUMBER, "4"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range test_1 {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}
		if token.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - token value wrong. expected=%q, got=%q", i, tt.expectedValue, token.Value)
		}
	}
}
