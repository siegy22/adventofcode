package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day15"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay15(t *testing.T) {
	input := ""
	expected := 0
	actual := day15.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay15Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day15.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
