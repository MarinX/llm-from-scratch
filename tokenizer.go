package main

import (
	"github.com/dlclark/regexp2"
)

// GPT-2 fixed regex pre-tokenization pattern
var expr = regexp2.MustCompile(`'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+`, 0)

type Tokenizer struct {
	bpe *BPE
}

func NewTokenizer() *Tokenizer {
	if err := LoadVocabulary(); err != nil {
		panic(err)
	}
	return &Tokenizer{
		bpe: new(BPE),
	}
}

func (t *Tokenizer) Encode(input string) []int {
	match, err := expr.FindStringMatch(input)
	if err != nil {
		panic(err)
	}
	var result []int
	for match != nil {
		ids := t.bpe.Encode([]byte(match.String()))
		result = append(result, ids...)
		match, err = expr.FindNextMatch(match)
		if err != nil {
			panic(err)
		}
	}

	return result
}

func (t *Tokenizer) Decode(input []int) string {
	var result []byte
	for _, id := range input {
		result = append(result, []byte(VacabularyInv[id])...)
	}
	return string(result)
}
