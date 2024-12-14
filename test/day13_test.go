package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day13"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay13(t *testing.T) {
	input := ""
	expected := 0
	actual := day13.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay13Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day13.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
