package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		total := numbers[left] + numbers[right]
		if total == target {
			return []int{left + 1, right + 1}
		}
		if total < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tc := range tests {
		result := twoSum(tc.numbers, tc.target)
		status := "PASS"
		if len(result) != len(tc.expected) || result[0] != tc.expected[0] || result[1] != tc.expected[1] {
			status = "FAIL"
		}
		fmt.Printf("%s | numbers=%v, target=%d => %v (expected %v)\n", status, tc.numbers, tc.target, result, tc.expected)
	}
}
