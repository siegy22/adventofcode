package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day01"
)

func main() {
	// Read input from file (e.g., input/day01.txt)
	inputFile := "../../input/day01.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day01.SolveDay1(string(input))
	fmt.Println("Solution for Day 1:", result)

	result2 := day01.SolveDay1Part2(string(input))
	fmt.Println("Solution for Day 1 part 2:", result2)
}
