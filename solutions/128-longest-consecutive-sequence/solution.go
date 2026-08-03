package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	best := 0
	for num := range numSet {
		if _, ok := numSet[num-1]; ok {
			continue
		}

		length := 1
		for {
			if _, ok := numSet[num+length]; !ok {
				break
			}
			length++
		}
		if length > best {
			best = length
		}
	}
	return best
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
		{[]int{}, 0},
	}

	for _, tc := range tests {
		result := longestConsecutive(tc.nums)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | nums=%v => %d (expected %d)\n", status, tc.nums, result, tc.expected)
	}
}
