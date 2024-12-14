package day01_test

import (
	"github.com/siegy22/adventofcode/internal/day01"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolveDay1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 11
	actual := day01.SolveDay1(input)

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestSolveDay1Part2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 31
	actual := day01.SolveDay1Part2(input)

	assert.Equal(t, expected, actual, "The two words should be the same.")
}
