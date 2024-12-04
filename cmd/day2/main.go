package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ryboe/q"
)

func main() {
	// reports := [][]int64{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{1, 3, 6, 7, 9},
	// }
	reports := loadInput("./input")

	diffInRange := func(report []int64) bool {
		for i := 1; i < len(report); i++ {
			diff := math.Abs(float64(report[i]) - float64(report[i-1]))
			if diff >= 1 && diff <= 3 {
				continue // all good, continue
			}
			return false
		}
		return true
	}

	isMonotonic := func(levels []int64) bool {
		increasing, decreasing := true, true

		for i := 1; i < len(levels); i++ {
			previous := levels[i-1]
			current := levels[i]

			if current > previous {
				decreasing = false
			} else if current < previous {
				increasing = false
			}
		}

		if increasing || decreasing {
			return true
		}

		return false
	}

	safeReports := 0
	for _, report := range reports {
		if diffInRange(report) && isMonotonic(report) {
			safeReports += 1
		}
	}

	fmt.Println("There are", safeReports, "safe reports")

	// problem dampener...
	safeReports = 0
	for _, report := range reports {
		if diffInRange(report) && isMonotonic(report) {
			safeReports += 1
		} else {

			// lol, just brute-force our way out....
			for i := range len(report) {
				tmpSlice := make([]int64, 0, len(report)-1)
				tmpSlice = append(tmpSlice, report[:i]...)
				tmpSlice = append(tmpSlice, report[i+1:]...)

				if diffInRange(tmpSlice) && isMonotonic(tmpSlice) {
					q.Q(tmpSlice, "removed", report[i])
					safeReports += 1
					break // dont look at this report again
				}
			}
		}
	}

	fmt.Println("There are", safeReports, "safe reports after problem dampener was applied")
}

func loadInput(inputFile string) [][]int64 {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	stringData := strings.TrimSpace(string(data))
	lines := strings.Split(stringData, "\n")

	out := make([][]int64, len(lines))

	for i, line := range lines {
		cols := strings.Fields(line)
		var nums []int64

		for _, stringNumber := range cols {
			num, _ := strconv.Atoi(stringNumber)
			nums = append(nums, int64(num))
		}

		out[i] = nums
	}

	return out
}
func remove(s []int64, i int) []int64 {
	// s[i] = s[len(s)-1]
	// return s[:len(s)-1]
	return s[:i+copy(s[i:], s[i+1:])]
}
