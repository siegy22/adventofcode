package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day16"
)

func main() {
	inputFile := "../../input/day16.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day16.Solve(string(input))
	fmt.Println("Solution for Day 16:", result)

	result2 := day16.SolvePart2(string(input))
	fmt.Println("Solution for Day 16 part 2:", result2)
}
