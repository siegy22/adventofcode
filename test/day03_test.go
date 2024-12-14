package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day03"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay3(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161
	actual := day03.SolveDay3(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay3Part2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := 48
	actual := day03.SolveDay3Part2(input)

	assert.Equal(t, expected, actual)
}
