package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day11"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay11(t *testing.T) {
	input := ""
	expected := 0
	actual := day11.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay11Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day11.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
