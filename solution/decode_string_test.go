package solution

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			input:    "abc3[cd]xyz",
			expected: "abccdcdcdxyz",
		},
		{
			input:    "3[a2[bd]]",
			expected: "abdbdabdbdabdbd",
		},
		{
			input:    "abc3[de1[fg]]",
			expected: "abcdefgdefgdefg",
		},
	}

	for _, test := range tests {
		got := decodeString(test.input)
		if got != test.expected {
			t.Errorf("decodeString(%q) = %q; expected %q", test.input, got, test.expected)
		}
	}
}
