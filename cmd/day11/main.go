package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day11"
)

func main() {
	inputFile := "../../input/day11.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day11.Solve(string(input))
	fmt.Println("Solution for Day 11:", result)

	result2 := day11.SolvePart2(string(input))
	fmt.Println("Solution for Day 11 part 2:", result2)
}
