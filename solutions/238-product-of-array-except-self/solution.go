package main

import "fmt"

func productExceptSelf(nums []int) []int {
	// 思路，分为左右两部分，分别计算左边的乘积和右边的乘积，然后相乘——》将O(n^2)的复杂度降低到O(n)
	n := len(nums)
	L, R, answer := make([]int, n), make([]int, n), make([]int, n)
	L[0] = 1
	R[n-1] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

func main() {
	got := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Printf("PASS | %v\n", got)
}
