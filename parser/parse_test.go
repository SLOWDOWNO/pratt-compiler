package parser_test

import (
	"testing"

	"ppc/lexer"
	"ppc/parser"
)

func TestParseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"3 + 5", 8},
		{"10 - 2", 8},
		{"2 * 3", 6},
		{"8 / 2", 4},
		{"(1 + 2) * 3", 9},
		{"3 + 4 * 2 / ( 1 - 5 )", 1}, // 3 + ((4 * 2) / (1 - 5)) = 3 + (-2) = 1
		{"-3 + 4", 1},
		{"-3 * (2 + 1)", -9},
		{"-(-3)", 3},
		{"2 * (-3 + 5)", 4},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		parser := parser.New(lex)
		result := parser.ParseExpression(0)

		if result != tt.expected {
			t.Errorf("expected %d, got %d for input %q", tt.expected, result, tt.input)
		}
	}
}
