package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	arrowPos := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > arrowPos {
			arrows++
			arrowPos = points[i][1]
		}
	}
	return arrows
}

func main() {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	}

	for _, tc := range tests {
		points := make([][]int, len(tc.points))
		for i, p := range tc.points {
			points[i] = []int{p[0], p[1]}
		}
		result := findMinArrowShots(points)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | points=%v => %d (expected %d)\n", status, tc.points, result, tc.expected)
	}
}
