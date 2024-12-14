package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day15"
)

func main() {
	inputFile := "../../input/day15.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day15.Solve(string(input))
	fmt.Println("Solution for Day 15:", result)

	result2 := day15.SolvePart2(string(input))
	fmt.Println("Solution for Day 15 part 2:", result2)
}
