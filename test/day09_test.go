package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day09"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay9(t *testing.T) {
	input := ""
	expected := 0
	actual := day09.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay9Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day09.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
