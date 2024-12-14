package main

import (
	"fmt"
	"log"
	"os"
	"github.com/siegy22/adventofcode/internal/day02"
)

func main() {
	inputFile := "../../input/day02.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Solve the problem
	result := day02.SolveDay2(string(input))
	fmt.Println("Solution for Day 2:", result)

	result2 := day02.SolveDay2Part2(string(input))
	fmt.Println("Solution for Day 2 part 2:", result2)
}
