package main

import "testing"

func TestBPE(t *testing.T) {
	cases := []struct {
		it       string
		input    string
		expected []int
	}{
		{
			it:       "convert Hello World!",
			input:    "Hello World!",
			expected: []int{15496, 2159, 0},
		},
		{
			it:       "should handle Croatian language",
			input:    "Bok! Kako si?",
			expected: []int{33, 482, 0, 509, 25496, 33721, 30},
		},
		{
			it:       "should handle hangul",
			input:    "안녕하세요",
			expected: []int{168, 243, 230, 167, 227, 243, 47991, 246, 168, 226, 116, 168, 248, 242},
		},
		{
			it:       "handle empty space",
			input:    " ",
			expected: []int{220},
		},
	}

	tkn := NewTokenizer() // need to load vocabulary
	for _, c := range cases {
		t.Run(c.it, func(t *testing.T) {
			output := tkn.bpe.Encode([]byte(c.input))
			if len(output) != len(c.expected) {
				t.Errorf("Length mismatch - got %d, expected %d", len(output), len(c.expected))
				return
			}
			for i, id := range output {
				if id != c.expected[i] {
					t.Errorf("ID mismatch at position %d - got %d, expected %d", i, id, c.expected[i])
					return
				}
			}
		})
	}
}
