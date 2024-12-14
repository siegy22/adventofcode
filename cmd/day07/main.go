package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day07"
)

func main() {
	inputFile := "../../input/day07.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day07.Solve(string(input))
	fmt.Println("Solution for Day 7:", result)

	result2 := day07.SolvePart2(string(input))
	fmt.Println("Solution for Day 7 part 2:", result2)
}
