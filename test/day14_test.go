package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day14"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay14(t *testing.T) {
	input := ""
	expected := 0
	actual := day14.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay14Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day14.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
