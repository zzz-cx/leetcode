package main

import "fmt"

// firstMissingPositive 将 [1,n] 内的数通过交换放到下标 i 上应满足 nums[i]==i+1，
// 再线性扫描第一个不符合的位置。时间 O(n)，额外空间 O(1)。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		v := nums[i]
		for v >= 1 && v <= n && nums[v-1] != v {
			nums[v-1], nums[i] = nums[i], nums[v-1]
			v = nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	got := firstMissingPositive([]int{3, 4, -1, 1})
	fmt.Printf("PASS | %d\n", got)
}
