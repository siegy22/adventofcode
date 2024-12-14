package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day21"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay21(t *testing.T) {
	input := ""
	expected := 0
	actual := day21.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay21Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day21.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
