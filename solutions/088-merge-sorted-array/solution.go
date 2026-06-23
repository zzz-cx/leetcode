package main

import "fmt"

func merge2(nums1 []int, m int, nums2 []int, n int) []int {
	i := m - 1
	j := n - 1
	k := m + n - 1         //nums1的最后一个位置
	for i >= 0 && j >= 0 { //从后往前遍历，将较大的数放到nums1的最后一个位置
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 { //如果nums2还有剩余，则将剩余的数放到nums1的最后一个位置
		nums1[k] = nums2[j]
		j--
		k--
	}
	return nums1
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge2(nums1, 3, nums2, 3)
	fmt.Printf("PASS | merged=%v\n", nums1)
}
