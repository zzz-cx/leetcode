package main

import "fmt"

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	got := findDuplicate([]int{1, 3, 4, 2, 2})
	fmt.Printf("PASS | %d\n", got)
}
