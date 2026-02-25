# LLM From Scratch

> Work in progress — self-learning project

A Go implementation of a Large Language Model built from scratch, following the book
[Build a Large Language Model (From Scratch)](https://www.manning.com/books/build-a-large-language-model-from-scratch) by Sebastian Raschka.

## Current state

Chapter 2.5 — BPE Tokenizer

Implemented so far:
- Vocabulary loading from `r50k_base.tiktoken` (GPT-2 vocabulary, ~50k tokens)
- GPT-2 regex pre-tokenizer
- Byte Pair Encoding (BPE) encoder
- Decoder (token IDs back to text)

## Usage

```go
t := NewTokenizer()

ids := t.Encode("Hello, world!")
fmt.Println(ids) // [15496 11 995 0]

text := t.Decode(ids)
fmt.Println(text) // Hello, world!
```
