package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day12"
)

func main() {
	inputFile := "../../input/day12.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day12.Solve(string(input))
	fmt.Println("Solution for Day 12:", result)

	result2 := day12.SolvePart2(string(input))
	fmt.Println("Solution for Day 12 part 2:", result2)
}
