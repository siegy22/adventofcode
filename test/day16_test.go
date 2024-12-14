package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day16"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay16(t *testing.T) {
	input := ""
	expected := 0
	actual := day16.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay16Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day16.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
