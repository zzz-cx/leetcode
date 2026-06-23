package main

import "fmt"

// TwoSum 两数之和 - 在数组中找出和为目标值的两个数的下标
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tc := range tests {
		got := TwoSum(tc.nums, tc.target)
		status := "PASS"
		if len(got) != len(tc.want) || got[0] != tc.want[0] || got[1] != tc.want[1] {
			status = "FAIL"
		}
		fmt.Printf("%s | nums=%v target=%d => %v (expected %v)\n", status, tc.nums, tc.target, got, tc.want)
	}
}
