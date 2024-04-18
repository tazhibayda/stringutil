package test

import (
	"github.com/tazhibayda/stringutil/stringutil"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"12345", "54321"},
		{"こんにちは", "はちにんこ"},
	}

	for _, test := range tests {
		if output := stringutil.Reverse(test.input); output != test.expected {
			t.Errorf("Reverse(%q) = %q, want %q", test.input, output, test.expected)
		}
	}
}

func TestSymbolCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"12345", 5},
		{"こんにちは", 5},
	}

	for _, test := range tests {
		if output := stringutil.SymbolCount(test.input); output != test.expected {
			t.Errorf("SymbolCount(%q) = %d, want %d", test.input, output, test.expected)
		}
	}
}
