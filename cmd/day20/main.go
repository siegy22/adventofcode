package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day20"
)

func main() {
	inputFile := "../../input/day20.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day20.Solve(string(input))
	fmt.Println("Solution for Day 20:", result)

	result2 := day20.SolvePart2(string(input))
	fmt.Println("Solution for Day 20 part 2:", result2)
}
