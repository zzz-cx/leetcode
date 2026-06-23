# 和为 K 的子数组

> LeetCode 560 · [subarray-sum-equals-k](https://leetcode.cn/problems/subarray-sum-equals-k/)

## 题目

给定整数数组 `nums` 和整数 k，统计和为 k 的连续子数组个数。

## 题解思路与解析

- 前缀和 + 哈希表

## 解答

### Golang

```go
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
```

### Python

```python
from typing import List


class Solution:
    def subarray_sum(self, nums: List[int], k: int) -> int:
        count = 0
        for start in range(len(nums)):
            s = 0
            for end in range(start, -1, -1):
                s += nums[end]
                if s == k:
                    count += 1
        return count

    def subarray_sum2(self, nums: List[int], k: int) -> int:
        count = 0
        s = 0
        prefix_sum = {0: 1}
        for num in nums:
            s += num
            count += prefix_sum.get(s - k, 0)
            prefix_sum[s] = prefix_sum.get(s, 0) + 1
        return count
```
