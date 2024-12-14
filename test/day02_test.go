package aoc_test

import (
	"github.com/siegy22/adventofcode/internal/day02"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	expected := 2
	actual := day02.SolveDay2(input)

	assert.Equal(t, expected, actual)
}


func TestSolveDay2Part2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	expected := 4
	actual := day02.SolveDay2Part2(input)

	assert.Equal(t, expected, actual)
}
