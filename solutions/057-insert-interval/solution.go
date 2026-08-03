package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	i, n := 0, len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}
	return result
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != 2 || len(b[i]) != 2 || a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		intervals    [][]int
		newInterval  []int
		expected     [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{[][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
	}

	for _, tc := range tests {
		newInterval := []int{tc.newInterval[0], tc.newInterval[1]}
		result := insert(tc.intervals, newInterval)
		status := "PASS"
		if !equal(result, tc.expected) {
			status = "FAIL"
		}
		fmt.Printf("%s | intervals=%v, new=%v => %v\n", status, tc.intervals, tc.newInterval, result)
	}
}
