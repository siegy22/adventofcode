package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day24"
)

func main() {
	inputFile := "../../input/day24.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day24.Solve(string(input))
	fmt.Println("Solution for Day 24:", result)

	result2 := day24.SolvePart2(string(input))
	fmt.Println("Solution for Day 24 part 2:", result2)
}
