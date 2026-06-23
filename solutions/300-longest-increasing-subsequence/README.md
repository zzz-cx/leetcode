# 最长递增子序列

> LeetCode 300 · [longest-increasing-subsequence](https://leetcode.cn/problems/longest-increasing-subsequence/)

## 题目

给定整数数组，找到最长严格递增子序列的长度。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func lengthOfLIS(nums []int) int { //贪心算法+二分查找解法
	if len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums)+1)
	lens := 1
	d[lens] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[lens] { //如果当前元素大于d[len]，则将当前元素加入d中
			d[lens+1] = nums[i]
			lens++
		} else { //如果当前元素小于等于d[len]，则二分查找d中第一个大于等于当前元素的位置
			pos := binarySearch(d, 1, lens, nums[i])
			d[pos] = nums[i]
		}
	}
	return lens
}
func binarySearch(d []int, left, right, target int) int {
	for left < right {
		mid := (left + right) / 2
		if d[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func lengthOfLIS2(nums []int) int { //动态规划解法
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### Python

```python
from typing import List


class Solution:
    def length_of_lis(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        d = [0] * (len(nums) + 1)
        lens = 1
        d[lens] = nums[0]
        for i in range(1, len(nums)):
            if nums[i] > d[lens]:
                lens += 1
                d[lens] = nums[i]
            else:
                pos = self._binary_search(d, 1, lens, nums[i])
                d[pos] = nums[i]
        return lens

    def _binary_search(self, d: List[int], left: int, right: int, target: int) -> int:
        while left < right:
            mid = (left + right) // 2
            if d[mid] < target:
                left = mid + 1
            else:
                right = mid
        return left

    def length_of_lis2(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        dp = [1] * len(nums)
        for i in range(1, len(nums)):
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return max(dp)
```
