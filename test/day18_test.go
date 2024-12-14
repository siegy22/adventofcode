package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day18"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay18(t *testing.T) {
	input := ""
	expected := 0
	actual := day18.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay18Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day18.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
