package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day06"
)

func main() {
	inputFile := "../../input/day06.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day06.Solve(string(input))
	fmt.Println("Solution for Day 6:", result)

	result2 := day06.SolvePart2(string(input))
	fmt.Println("Solution for Day 6 part 2:", result2)
}
