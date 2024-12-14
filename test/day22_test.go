package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day22"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay22(t *testing.T) {
	input := ""
	expected := 0
	actual := day22.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay22Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day22.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
