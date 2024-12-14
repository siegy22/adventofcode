package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day09"
)

func main() {
	inputFile := "../../input/day09.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day09.Solve(string(input))
	fmt.Println("Solution for Day 9:", result)

	result2 := day09.SolvePart2(string(input))
	fmt.Println("Solution for Day 9 part 2:", result2)
}
