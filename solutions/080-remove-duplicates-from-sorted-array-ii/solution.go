package main

import "fmt"

// removeDuplicates80 删除有序数组重复项 II（LeetCode 80）
// 思路：快慢指针。已写入区 [0..k) 内，同一数最多出现 2 次
// 判定：若 nums[i] != nums[k-2]，说明写入后仍不超过 2 次，可保留
func removeDuplicates80(nums []int) int {
	k := 0
	for _, x := range nums {
		if k < 2 || x != nums[k-2] {
			nums[k] = x
			k++
		}
	}
	return k
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates80(nums)
	status := "PASS"
	if k != 5 {
		status = "FAIL"
	}
	fmt.Printf("%s | k=%d nums[:k]=%v\n", status, k, nums[:k])
}
