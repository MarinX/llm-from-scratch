package main

import "testing"

func TestTokenizer(t *testing.T) {
	cases := []struct {
		it    string
		input string
	}{
		{
			it:    "encode decode simple",
			input: "Hello World!",
		},
		{
			it:    "encode decode sentence",
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
		{
			it:    "encode decode Croatian sentence",
			input: "Kako se ti zoveš?",
		},
		{
			it:    "encode decode Korean sentence",
			input: "안녕하세요 세상",
		},
	}
	tokenizer := NewTokenizer()

	for _, c := range cases {
		t.Run(c.it, func(t *testing.T) {
			result := tokenizer.Encode(c.input)
			decoded := tokenizer.Decode(result)
			if decoded != c.input {
				t.Errorf("Expected %s but got %s", c.input, decoded)
			}
		})
	}
}
