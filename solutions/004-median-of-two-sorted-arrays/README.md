# 寻找两个正序数组的中位数

> LeetCode 4 · [median-of-two-sorted-arrays](https://leetcode.cn/problems/median-of-two-sorted-arrays/)

## 题目

给定两个大小分别为 m 和 n 的正序（从小到大）数组 `nums1` 和 `nums2`，找出并返回这两个正序数组的中位数。

## 题解思路与解析

- 思路：在较短数组上二分切分，使左半共 (m+n+1)/2 个且 max(左) <= min(右)

## 解答

### Golang

```go
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
```

### Python

```python
# 思路：在较短数组上二分切分，使左半共 (m+n+1)/2 个且 max(左) <= min(右)
from typing import List


class Solution:
    def find_median_sorted_arrays(self, nums1: List[int], nums2: List[int]) -> float:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        m, n = len(nums1), len(nums2)
        left, right = 0, m
        half = (m + n + 1) // 2

        while left <= right:
            i = (left + right) // 2
            j = half - i

            if i < m and j > 0 and nums1[i] < nums2[j - 1]:
                left = i + 1
            elif i > 0 and j < n and nums2[j] < nums1[i - 1]:
                right = i - 1
            else:
                if i == 0:
                    max_left = nums2[j - 1]
                elif j == 0:
                    max_left = nums1[i - 1]
                else:
                    max_left = max(nums1[i - 1], nums2[j - 1])
                if (m + n) % 2 == 1:
                    return float(max_left)

                if i == m:
                    min_right = nums2[j]
                elif j == n:
                    min_right = nums1[i]
                else:
                    min_right = min(nums1[i], nums2[j])
                return (max_left + min_right) / 2
        return 0.0
```
