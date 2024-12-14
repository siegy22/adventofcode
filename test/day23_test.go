package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day23"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay23(t *testing.T) {
	input := ""
	expected := 0
	actual := day23.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay23Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day23.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
