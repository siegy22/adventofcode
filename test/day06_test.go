package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day06"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay6(t *testing.T) {
	input := ""
	expected := 0
	actual := day06.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay6Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day06.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
