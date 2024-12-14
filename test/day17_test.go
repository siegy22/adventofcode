package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day17"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay17(t *testing.T) {
	input := ""
	expected := 0
	actual := day17.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay17Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day17.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
