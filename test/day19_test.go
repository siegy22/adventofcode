package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day19"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay19(t *testing.T) {
	input := ""
	expected := 0
	actual := day19.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay19Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day19.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
