package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day14"
)

func main() {
	inputFile := "../../input/day14.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day14.Solve(string(input))
	fmt.Println("Solution for Day 14:", result)

	result2 := day14.SolvePart2(string(input))
	fmt.Println("Solution for Day 14 part 2:", result2)
}
