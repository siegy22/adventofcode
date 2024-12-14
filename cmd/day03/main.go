package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day03"
)

func main() {
	inputFile := "../../input/day03.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day03.SolveDay3(string(input))
	fmt.Println("Solution for Day 3:", result)

	result2 := day03.SolveDay3Part2(string(input))
	fmt.Println("Solution for Day 3 part 2:", result2)
}
