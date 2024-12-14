package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day04"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay4(t *testing.T) {
	input := ""
	expected := 0
	actual := day04.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay4Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day04.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
