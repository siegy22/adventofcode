package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day05"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay5(t *testing.T) {
	input := ""
	expected := 0
	actual := day05.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay5Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day05.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
