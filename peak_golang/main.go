package main

import "fmt"

func main() {
	fmt.Println(pickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3})) // => Pos: [3 7], Peaks: [6 3]
	fmt.Println(pickPeaks([]int{1, 2, 2, 2, 1}))                      // => Pos: [1], Peaks: [2]
	fmt.Println(pickPeaks([]int{1, 2, 2, 2, 3}))                      // => Pos: [], Peaks: []
}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func pickPeaks(array []int) PosPeaks {
	result := PosPeaks{Pos: []int{}, Peaks: []int{}}

	if len(array) < 3 {
		return result
	}

	var position int
	posiblePeak := false

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			position = i
			posiblePeak = true
		} else if array[i] < array[i-1] && posiblePeak {

			result.Pos = append(result.Pos, position)
			result.Peaks = append(result.Peaks, array[position])
			posiblePeak = false
		}
	}

	return result
}
