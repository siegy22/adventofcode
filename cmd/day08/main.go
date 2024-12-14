package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day08"
)

func main() {
	inputFile := "../../input/day08.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day08.Solve(string(input))
	fmt.Println("Solution for Day 8:", result)

	result2 := day08.SolvePart2(string(input))
	fmt.Println("Solution for Day 8 part 2:", result2)
}
