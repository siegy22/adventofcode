package day04

import (
	"strings"
	"fmt"
)

func Solve(input string) (total int) {
	lineLength := len(strings.Split(input, "\n")[0])

	input = strings.Replace(input, "\n", "", -1)
	words := []string{}
	for i, _ := range input {
		words = append(words, horizontal(i, input, lineLength)...)
		words = append(words, vertical(i, input, lineLength)...)
		words = append(words, diagonal(i, input, lineLength)...)
	}
	for _, word := range words {
		if word == "XMAS" {
			total += 1
		}
	}
	return total
}

func SolvePart2(input string) (total int) {

	return total
}

func vertical(index int, input string, lineLength int) (results []string) {
	if index >= 3 * lineLength {
		// Up
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index - lineLength]),
			string(input[index - 2 * lineLength]),
			string(input[index - 3 * lineLength]),
		))
	}

	if index <= len(input) - 3 * lineLength - 1 {
		// Down
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index + lineLength]),
			string(input[index + 2 * lineLength]),
			string(input[index + 3 * lineLength]),
		))
	}
	return results
}

func horizontal(index int, input string, lineLength int) (results []string) {
	if index % lineLength >= 3 {
		// Left
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index - 1]),
			string(input[index - 2]),
			string(input[index - 3]),
		))
	}

	if index % lineLength <= lineLength - 4 {
		// Right
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index + 1]),
			string(input[index + 2]),
			string(input[index + 3]),
		))
	}
	return results
}

func diagonal(index int, input string, lineLength int) (results []string) {
	if index % lineLength >= 3 && index >= 3 * lineLength {
		// Left Up
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index - lineLength - 1]),
			string(input[index - 2 * lineLength - 2]),
			string(input[index - 3 * lineLength - 3]),
		))
	}

	if index % lineLength <= lineLength - 4 && index >= 3 * lineLength {
		// Right up
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index - lineLength + 1]),
			string(input[index - 2 * lineLength + 2]),
			string(input[index - 3 * lineLength + 3]),
		))
	}

	if index % lineLength >= 3 && index <= len(input) - 3 * lineLength - 1 {
		// Left down
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index + lineLength - 1]),
			string(input[index + 2 * lineLength - 2]),
			string(input[index + 3 * lineLength - 3]),
		))
	}

	if index % lineLength <= lineLength - 4 && index <= len(input) - 3 * lineLength - 1 {
		// Right down
		results = append(results, fmt.Sprintf("%s%s%s%s",
			string(input[index]),
			string(input[index + lineLength + 1]),
			string(input[index + 2 * lineLength + 2]),
			string(input[index + 3 * lineLength + 3]),
		))
	}
 	return results
}
