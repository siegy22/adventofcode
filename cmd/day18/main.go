package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day18"
)

func main() {
	inputFile := "../../input/day18.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day18.Solve(string(input))
	fmt.Println("Solution for Day 18:", result)

	result2 := day18.SolvePart2(string(input))
	fmt.Println("Solution for Day 18 part 2:", result2)
}
