package day01

import (
	"fmt"
	"strconv"
	"strings"
	"slices"
)

func SolveDay1(input string) int {
	lines := strings.Split(input, "\n")
	first := []int{}
	second := []int{}
	for _, line := range lines {
		results := strings.Split(line, "   ")
		left, err := strconv.Atoi(results[0])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		first = append(first, left)

		right, err := strconv.Atoi(results[1])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		second = append(second, right)
	}

	slices.Sort(first)
	slices.Sort(second)
	total := 0
	for index, value := range first {
		distance := value - second[index]
		if distance < 0 {
			total += -distance
		} else {
			total += distance
		}
	}

	return total
}

func SolveDay1Part2(input string) int {
	lines := strings.Split(input, "\n")
	first := []int{}
	second := []int{}
	for _, line := range lines {
		results := strings.Split(line, "   ")
		left, err := strconv.Atoi(results[0])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		first = append(first, left)

		right, err := strconv.Atoi(results[1])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		second = append(second, right)
	}

	total := 0
	for _, value := range first {
		similarity := SimilarityScore(second, value)
		total += similarity
	}

	return total
}

func SimilarityScore(arr []int, elem int) int {
	count := 0
	for _, num := range arr {
		if num == elem {
			count++
		}
	}
	return count * elem
}

