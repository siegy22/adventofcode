package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day24"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay24(t *testing.T) {
	input := ""
	expected := 0
	actual := day24.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay24Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day24.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
