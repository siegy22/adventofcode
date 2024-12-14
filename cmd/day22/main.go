package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day22"
)

func main() {
	inputFile := "../../input/day22.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day22.Solve(string(input))
	fmt.Println("Solution for Day 22:", result)

	result2 := day22.SolvePart2(string(input))
	fmt.Println("Solution for Day 22 part 2:", result2)
}
