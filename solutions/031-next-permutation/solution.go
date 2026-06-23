package main

import "fmt"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse2(nums, i+1, len(nums)-1)
}
func reverse2(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Printf("PASS | %v\n", nums)
}
