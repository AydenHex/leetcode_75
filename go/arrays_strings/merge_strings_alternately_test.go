package arraysstrings

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestSameLengthMerge(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	expected := "apbqcr"

	result := mergeAlternately(word1, word2)

	assert.Equal(t, expected, result)
}
func TestSameLengthMerge2(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	expected := "apbqcr"

	result := mergeAlternately2(word1, word2)

	assert.Equal(t, expected, result)
}

func TestWord2LongerhMerge(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	expected := "apbqrs"

	result := mergeAlternately(word1, word2)

	assert.Equal(t, expected, result)
}
func TestWord2LongerhMerge2(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	expected := "apbqrs"

	result := mergeAlternately2(word1, word2)

	assert.Equal(t, expected, result)
}

func TestWord1LongerhMerge(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	expected := "apbqcd"

	result := mergeAlternately(word1, word2)

	assert.Equal(t, expected, result)
}
func TestWord1LongerhMerge2(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	expected := "apbqcd"

	result := mergeAlternately2(word1, word2)

	assert.Equal(t, expected, result)
}

// Benchmarks

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var table = []struct {
	input int
}{
	{input: 10},
	{input: 100},
	{input: 1000},
	{input: 10000},
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BenchmarkSimpleMerge(b *testing.B) {
	for _, v := range table {
		word1 := randStringRunes(v.input)
		word2 := randStringRunes(v.input)
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mergeAlternately(word1, word2)
			}
		})
	}
}

func BenchmarkSimpleMerge2(b *testing.B) {
	for _, v := range table {
		word1 := randStringRunes(v.input)
		word2 := randStringRunes(v.input)
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mergeAlternately2(word1, word2)
			}
		})
	}
}
