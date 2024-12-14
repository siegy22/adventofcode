package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay10(t *testing.T) {
	input := ""
	expected := 0
	actual := day10.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay10Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day10.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
