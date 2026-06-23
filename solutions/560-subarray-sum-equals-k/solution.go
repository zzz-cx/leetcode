package main

import "fmt"

func subarraySum(nums []int, k int) int { // 暴力枚举
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySum2(nums []int, k int) int {
	// 前缀和 + 哈希表
	count := 0
	sum := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	for _, num := range nums {
		sum += num
		count += prefixSum[sum-k]
		prefixSum[sum]++
	}
	return count
}

func main() {
	got := subarraySum([]int{1, 1, 1}, 2)
	status := "PASS"
	if got != 2 {
		status = "FAIL"
	}
	fmt.Printf("%s | count=%d\n", status, got)
}
