package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day08"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay8(t *testing.T) {
	input := ""
	expected := 0
	actual := day08.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay8Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day08.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
