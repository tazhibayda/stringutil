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
