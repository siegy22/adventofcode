package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day04"
)

func main() {
	inputFile := "../../input/day04.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day04.Solve(string(input))
	fmt.Println("Solution for Day 4:", result)

	result2 := day04.SolvePart2(string(input))
	fmt.Println("Solution for Day 4 part 2:", result2)
}
