package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ryboe/q"
)

type LocationList []int64

func main() {
	l1, l2 := loadInput("./input1")
	// l1 := LocationList{3, 4, 2, 1, 3, 3}
	// l2 := LocationList{4, 3, 5, 3, 9, 3}

	// slices.Sort(l1)
	// slices.Sort(l2)

	if len(l1) != len(l2) {
		panic("whoops")
	}

	var totalDistance int64
	for i := 0; i < len(l1); i++ {
		// distance := l2[i] - (l1[i])
		// distance = int64(math.Abs(float64(distance)))
		distance := l1[i] * int64(occurrences(l1[i], l2))
		totalDistance += distance
		q.Q(l1[i], l2[i], occurrences(l1[i], l2), distance)
	}

	fmt.Println(totalDistance)
}

func loadInput(inputFile string) (LocationList, LocationList) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	stringData := strings.TrimSpace(string(data))
	lines := strings.Split(stringData, "\n")

	l1, l2 := LocationList{}, LocationList{}
	for _, line := range lines {
		cols := strings.Fields(line)
		a, _ := strconv.Atoi(cols[0])
		b, _ := strconv.Atoi(cols[1])
		l1 = append(l1, int64(a))
		l2 = append(l2, int64(b))
	}

	return l1, l2
}

func occurrences(search int64, list LocationList) int {
	count := 0
	for _, num := range list {
		if num == search {
			count++
		}
	}
	return count
}
