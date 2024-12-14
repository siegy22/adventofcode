package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day10"
)

func main() {
	inputFile := "../../input/day10.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day10.Solve(string(input))
	fmt.Println("Solution for Day 10:", result)

	result2 := day10.SolvePart2(string(input))
	fmt.Println("Solution for Day 10 part 2:", result2)
}
