package day03

import (
	"regexp"
	"strings"
	"strconv"
)

func SolveDay3(input string) (total int) {
	r := regexp.MustCompile("mul\\(\\d*,\\d*\\)")
	matches := r.FindAllString(input, -1)
	for _, match := range matches {
		parts := strings.Split(match, ",")
		a, _ := strconv.Atoi(strings.Replace(parts[0], "mul(", "", -1))
		b, _ := strconv.Atoi(strings.Replace(parts[1], ")", "", -1))
		total += a * b
	}

	return total
}

func SolveDay3Part2(input string) (total int) {
	total = 0
	r := regexp.MustCompile("(mul\\(\\d*,\\d*\\)|don\\'t\\(\\)|do\\(\\))")
	matches := r.FindAllString(input, -1)
	do := true
	for _, match := range matches {
		if match == "don't()" {
			do = false
			continue
		}
		if match == "do()" {
			do = true
			continue
		}
		if !do { continue }

		parts := strings.Split(match, ",")
		a, _ := strconv.Atoi(strings.Replace(parts[0], "mul(", "", -1))
		b, _ := strconv.Atoi(strings.Replace(parts[1], ")", "", -1))
		total += a * b
	}

	return total
}
