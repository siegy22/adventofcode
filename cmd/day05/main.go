package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day05"
)

func main() {
	inputFile := "../../input/day05.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day05.Solve(string(input))
	fmt.Println("Solution for Day 5:", result)

	result2 := day05.SolvePart2(string(input))
	fmt.Println("Solution for Day 5 part 2:", result2)
}
