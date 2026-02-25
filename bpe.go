package main

import (
	"math"
)

// BPE byte pair encoding
type BPE struct{}

func (b *BPE) Encode(chunk []byte) []int {
	parts := make([][]byte, len(chunk))
	for i, b := range chunk {
		parts[i] = []byte{b}
	}

	for {
		bestRank := math.MaxInt
		bestIdx := -1

		for i := 0; i < len(parts)-1; i++ {
			merged := append(append([]byte{}, parts[i]...), parts[i+1]...)
			rank, ok := Vocabulary[string(merged)]
			if ok && rank < bestRank {
				bestRank = rank
				bestIdx = i
			}
		}
		// no valid merge found, let's break
		if bestIdx == -1 {
			break
		}

		merged := append(append([]byte{}, parts[bestIdx]...), parts[bestIdx+1]...)
		parts = append(parts[:bestIdx+1], parts[bestIdx+2:]...)
		parts[bestIdx] = merged
	}
	result := make([]int, len(parts))
	for i, part := range parts {
		result[i] = Vocabulary[string(part)]
	}

	return result
}
