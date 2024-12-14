package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day13"
)

func main() {
	inputFile := "../../input/day13.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day13.Solve(string(input))
	fmt.Println("Solution for Day 13:", result)

	result2 := day13.SolvePart2(string(input))
	fmt.Println("Solution for Day 13 part 2:", result2)
}
