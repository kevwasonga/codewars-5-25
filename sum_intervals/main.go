package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 4}, {7, 10}, {3, 5}}
	fmt.Println(SumOfIntervals(intervals))
}

func SumOfIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	current := intervals[0]

	for _, interval := range intervals[1:] {
		if interval[0] < current[1] {
			if interval[1] > current[1] {
				current[1] = interval[1]
			}
		} else {
			merged = append(merged, current)
			current = interval
		}

	}

	merged = append(merged, current)
	total := 0
	for _, interval := range merged {
		total += interval[1] - interval[0]
	}
	return total
}
