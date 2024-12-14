package day02

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func SolveDay2(input string) int {
	reports := parseInput(input)
	total := 0
	for _, report := range reports {
		if report.IsSafe() {
			total += 1
		}
	}

	return total
}

func SolveDay2Part2(input string) int {
	reports := parseInput(input)
	total := 0
	for _, report := range reports {
		if report.IsDampenedSafe() {
			total += 1
		}
	}

	return total
}

func parseInput(input string) (reports []Report) {
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		report := Report{}
		for _, num := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(num)
			report.numbers = append(report.numbers, num)
		}
		reports = append(reports, report)
	}
	return reports
}

type Report struct {
	numbers []int
}

func (r Report) IsDampenedSafe() bool {
	if r.IsSafe() {
		return true
	}

	for i, _ := range r.numbers {
		copy := append([]int{}, r.numbers...)
		reduced := Report{append(copy[:i], copy[i+1:]...)}
		if reduced.IsSafe() {
			return true
		}
	}
	return false
}

func (r Report) IsSafe() bool {
	sorted := append([]int{}, r.numbers...)
	sort.Ints(sorted)
	revSorted := append([]int{}, r.numbers...)
	sort.Sort(sort.Reverse(sort.IntSlice(revSorted)))

	return distanceSafe(r.numbers) && (reflect.DeepEqual(r.numbers, sorted) || reflect.DeepEqual(r.numbers, revSorted))
}

func distanceSafe(nums []int) bool {
	for i, num := range nums {
		if i == len(nums) - 1 {
			continue
		}

		distance := num - nums[i + 1]

		if distance < 0 {
			distance = -distance
		}
		if distance > 3 || distance < 1 {
			return false
		}
	}

	return true
}
