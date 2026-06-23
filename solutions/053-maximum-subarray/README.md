# 最大子数组和

> LeetCode 53 · [maximum-subarray](https://leetcode.cn/problems/maximum-subarray/)

## 题目

给定整数数组 `nums`，找出一个具有最大和的连续子数组（至少包含一个元素），返回其最大和。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func maxSubArray(nums []int) int { //思路，只要前面的数字和大于零，就把他加上，如果小于0，则当前就是最大的，重新计算

	n := len(nums)
	maxSum := nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}
```

### Python

```python
from typing import List


class Solution:
    def max_sub_array(self, nums: List[int]) -> int:
        max_sum = nums[0]
        for i in range(1, len(nums)):
            if nums[i - 1] > 0:
                nums[i] += nums[i - 1]
            if nums[i] > max_sum:
                max_sum = nums[i]
        return max_sum
```
