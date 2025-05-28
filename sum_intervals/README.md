
# README.md

````markdown
# sumIntervals

A Go function that calculates the total length covered by a set of intervals, ensuring that overlapping intervals are only counted once.

## Description

Given a slice of intervals represented as pairs of integers (e.g., `[start, end]`), the `sumIntervals` function returns the sum of all unique interval lengths. Intervals that overlap are merged before calculating the total length to avoid double counting.

## How It Works

1. Sort intervals by their start times.
2. Merge overlapping intervals into a combined interval.
3. Sum the lengths of the merged intervals.

## Usage

```go
package main

import (
	"fmt"
)

func main() {
	intervals := [][]int{
		{1, 4},
		{7, 10},
		{3, 5},
	}

	total := sumIntervals(intervals)
	fmt.Println("Total length:", total) // Output: Total length: 7
}
````

## Function Signature

```go
func sumIntervals(intervals [][]int) int
```

* **Input:** `[][]int` — A slice of intervals, each interval is a slice of two integers `[start, end]` where `start < end`.
* **Output:** `int` — The sum of lengths of all intervals, with overlapping intervals merged.

## Examples

| Input                                           | Output    | Explanation                                                                                     |
| ----------------------------------------------- | --------- | ----------------------------------------------------------------------------------------------- |
| `[[1, 2], [6, 10], [11, 15]]`                   | 9         | Intervals don't overlap; total length is `(2-1)+(10-6)+(15-11) = 1+4+4 = 9`.                    |
| `[[1, 4], [7, 10], [3, 5]]`                     | 7         | Intervals `[1,4]` and `[3,5]` overlap; merged into `[1,5]`. Total length is `(5-1)+(10-7)`.     |
| `[[1, 5], [10, 20], [1, 6], [16, 19], [5, 11]]` | 19        | Complex overlaps merge into `[1, 20]`. Total length is `20 - 1 = 19`.                           |
| `[ [0, 20], [-100000000, 10], [30, 40] ]`       | 100000030 | Large intervals handled gracefully, overlapping merge from `-100000000` to `20` plus `[30,40]`. |

## Edge Cases

* Empty input slice returns `0`.
* Intervals fully contained within others count only once.
* Handles large ranges efficiently.

---

# Testing

Below are some example test cases you can use to validate your implementation.

```go
package main

import "testing"

func TestSumIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  int
	}{
		{[][]int{}, 0},
		{[][]int{{1, 2}}, 1},
		{[][]int{{1, 2}, {3, 4}}, 2},
		{[][]int{{1, 5}, {2, 6}}, 5},
		{[][]int{{1, 4}, {7, 10}, {3, 5}}, 7},
		{[][]int{{1, 5}, {10, 20}, {1, 6}, {16, 19}, {5, 11}}, 19},
		{[][]int{{0, 20}, {-100000000, 10}, {30, 40}}, 100000030},
	}

	for _, test := range tests {
		result := sumIntervals(test.intervals)
		if result != test.expected {
			t.Errorf("sumIntervals(%v) = %d; want %d", test.intervals, result, test.expected)
		}
	}
}
```

---

## License

This project is released under the MIT License. Feel free to use and adapt it!



