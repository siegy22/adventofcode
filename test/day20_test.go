package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day20"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay20(t *testing.T) {
	input := ""
	expected := 0
	actual := day20.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay20Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day20.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
