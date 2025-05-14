
---
````markdown
# 🏔️ Pick Peaks (Go Implementation)

This project solves the **"Pick Peaks"** challenge: identifying local maxima ("peaks") in an integer array, including flat plateaus that eventually fall. It returns both the positions and values of such peaks.

## 📘 Problem Statement

Given an array of integers, find all the "peaks" (local maxima) such that:

- A peak is an element that is greater than its neighbors (excluding the first and last elements).
- A "plateau peak" is a flat sequence of equal numbers that eventually descends. We return the **starting index** of that plateau.

### 🧪 Examples

```go
pickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}) // → Pos: [3 7], Peaks: [6 3]
pickPeaks([]int{1, 2, 2, 2, 1})                      // → Pos: [1], Peaks: [2]
pickPeaks([]int{1, 2, 2, 2, 3})                      // → Pos: [], Peaks: []
pickPeaks([]int{0, 1, 2, 5, 1, 0})                   // → Pos: [3], Peaks: [5]
````

---

## 💡 Approach

1. Iterate through the array from the second to the last element.
2. Detect a rising edge: when `array[i] > array[i-1]`, mark it as a possible peak.
3. Track flat plateaus by continuing if `array[i] == array[i-1]`.
4. Confirm a peak if a descent is found (`array[i] < array[i-1]`) **after** a rise.
5. Return the first point of the rise (or plateau start) as the peak position.

---

## 🚀 How to Run

```bash
go run main.go
```

---

## 📁 Output Format

The function returns a struct:

```go
type PosPeaks struct {
	Pos   []int
	Peaks []int
}
```

---

## 🧠 Edge Cases

* Arrays with fewer than 3 elements have no peaks.
* Peaks must not be at the first or last index.
* Flat plateaus that do not fall are not peaks.

---

````

---

## ✅ **Go Solution (`main.go`)**

```go
package main

import (
	"fmt"
)

// PosPeaks holds the result — positions and values of detected peaks
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// pickPeaks identifies all valid peaks and plateau peaks
func pickPeaks(array []int) PosPeaks {
	result := PosPeaks{Pos: []int{}, Peaks: []int{}}

	if len(array) < 3 {
		return result
	}

	var pos int
	possiblePeak := false

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			pos = i
			possiblePeak = true
		} else if array[i] < array[i-1] && possiblePeak {
			result.Pos = append(result.Pos, pos)
			result.Peaks = append(result.Peaks, array[pos])
			possiblePeak = false
		}
	}

	return result
}

func main() {
	fmt.Println(pickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3})) // → Pos: [3 7], Peaks: [6 3]
	fmt.Println(pickPeaks([]int{1, 2, 2, 2, 1}))                      // → Pos: [1], Peaks: [2]
	fmt.Println(pickPeaks([]int{1, 2, 2, 2, 3}))                      // → Pos: [], Peaks: []
	fmt.Println(pickPeaks([]int{0, 1, 2, 5, 1, 0}))                   // → Pos: [3], Peaks: [5]
	fmt.Println(pickPeaks([]int{1, 2, 3, 3, 2, 1, 2, 2, 2, 1}))       // → Pos: [2, 6], Peaks: [3, 2]
}
````

---

