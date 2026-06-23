package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	n := len(nums)
	if n == 0 {
		return ans
	}
	deque := make([]int, 0)
	for i := 0; i < n; i++ {
		if i >= k && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}
	return ans
}

func main() {
	got := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Printf("PASS | %v\n", got)
}
