package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day12"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay12(t *testing.T) {
	input := ""
	expected := 0
	actual := day12.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay12Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day12.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
