package main

import "fmt"

// 思路：在较短数组上二分切分，使左半共 (m+n+1)/2 个且 max(左) <= min(右)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	half := (m + n + 1) / 2

	for left <= right {
		i := (left + right) / 2
		j := half - i

		if i < m && j > 0 && nums1[i] < nums2[j-1] {
			left = i + 1
		} else if i > 0 && j < n && nums2[j] < nums1[i-1] {
			right = i - 1
		} else {
			var maxLeft int
			switch {
			case i == 0:
				maxLeft = nums2[j-1]
			case j == 0:
				maxLeft = nums1[i-1]
			default:
				maxLeft = nums1[i-1]
				if nums2[j-1] > maxLeft {
					maxLeft = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			switch {
			case i == m:
				minRight = nums2[j]
			case j == n:
				minRight = nums1[i]
			default:
				minRight = nums1[i]
				if nums2[j] < minRight {
					minRight = nums2[j]
				}
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}

func main() {
	got := findMedianSortedArrays([]int{1, 3}, []int{2})
	status := "PASS"
	if got != 2.0 {
		status = "FAIL"
	}
	fmt.Printf("%s | median=%v (expected 2.0)\n", status, got)
}
