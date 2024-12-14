package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day07"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay7(t *testing.T) {
	input := ""
	expected := 0
	actual := day07.Solve(input)

	assert.Equal(t, expected, actual)
}

func TestSolveDay7Part2(t *testing.T) {
	input := ""
	expected := 0
	actual := day07.SolvePart2(input)

	assert.Equal(t, expected, actual)
}
