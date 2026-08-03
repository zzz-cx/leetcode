package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	i, n := 0, len(nums)

	for i < n {
		start := nums[i]
		for i+1 < n && nums[i+1] == nums[i]+1 {
			i++
		}
		end := nums[i]
		if start == end {
			result = append(result, strconv.Itoa(start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, end))
		}
		i++
	}
	return result
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		nums     []int
		expected []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{}, []string{}},
		{[]int{1}, []string{"1"}},
	}

	for _, tc := range tests {
		result := summaryRanges(tc.nums)
		status := "PASS"
		if !equal(result, tc.expected) {
			status = "FAIL"
		}
		fmt.Printf("%s | nums=%v => %v\n", status, tc.nums, result)
	}
}
