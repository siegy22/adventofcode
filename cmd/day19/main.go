package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day19"
)

func main() {
	inputFile := "../../input/day19.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day19.Solve(string(input))
	fmt.Println("Solution for Day 19:", result)

	result2 := day19.SolvePart2(string(input))
	fmt.Println("Solution for Day 19 part 2:", result2)
}
