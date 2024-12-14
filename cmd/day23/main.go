package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day23"
)

func main() {
	inputFile := "../../input/day23.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day23.Solve(string(input))
	fmt.Println("Solution for Day 23:", result)

	result2 := day23.SolvePart2(string(input))
	fmt.Println("Solution for Day 23 part 2:", result2)
}
