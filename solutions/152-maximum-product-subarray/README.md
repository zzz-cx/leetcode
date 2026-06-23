# 最大子数组乘积

> LeetCode 152 · [maximum-product-subarray](https://leetcode.cn/problems/maximum-product-subarray/)

## 题目

给定整数数组，找出乘积最大的连续子数组，返回乘积最大值。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func maxProduct(nums []int) int {
	maxEnding, minEnding := nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxEnding, minEnding = minEnding, maxEnding
		}
		maxEnding = max(nums[i], maxEnding*nums[i])
		minEnding = min(nums[i], minEnding*nums[i])
		ans = max(ans, maxEnding)
	}
	return ans
}
```

### Python

```python
from typing import List


class Solution:
    def max_product(self, nums: List[int]) -> int:
        max_ending, min_ending = nums[0], nums[0]
        ans = nums[0]
        for i in range(1, len(nums)):
            if nums[i] < 0:
                max_ending, min_ending = min_ending, max_ending
            max_ending = max(nums[i], max_ending * nums[i])
            min_ending = min(nums[i], min_ending * nums[i])
            ans = max(ans, max_ending)
        return ans
```
