package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day21"
)

func main() {
	inputFile := "../../input/day21.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day21.Solve(string(input))
	fmt.Println("Solution for Day 21:", result)

	result2 := day21.SolvePart2(string(input))
	fmt.Println("Solution for Day 21 part 2:", result2)
}
