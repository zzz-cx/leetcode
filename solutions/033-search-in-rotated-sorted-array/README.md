# 搜索旋转排序数组

> LeetCode 33 · [search-in-rotated-sorted-array](https://leetcode.cn/problems/search-in-rotated-sorted-array/)

## 题目

整数数组 `nums` 按升序排列并经过旋转，给定目标值 `target`，在数组中搜索并返回下标，不存在则返回 -1。

## 题解思路与解析

- 思路：每次 mid 必有一侧仍是有序段，先判断 target 是否落在该有序段内，再缩区间

## 解答

### Golang

```go
// 思路：每次 mid 必有一侧仍是有序段，先判断 target 是否落在该有序段内，再缩区间
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // 左半段有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半段有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
```

### Python

```python
# 思路：每次 mid 必有一侧仍是有序段，先判断 target 是否落在该有序段内，再缩区间
from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                return mid
            if nums[left] <= nums[mid]:
                if nums[left] <= target < nums[mid]:
                    right = mid - 1
                else:
                    left = mid + 1
            else:
                if nums[mid] < target <= nums[right]:
                    left = mid + 1
                else:
                    right = mid - 1
        return -1
```
