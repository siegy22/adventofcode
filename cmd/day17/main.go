package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day17"
)

func main() {
	inputFile := "../../input/day17.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day17.Solve(string(input))
	fmt.Println("Solution for Day 17:", result)

	result2 := day17.SolvePart2(string(input))
	fmt.Println("Solution for Day 17 part 2:", result2)
}
