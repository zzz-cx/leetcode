# 除自身以外数组的乘积

> LeetCode 238 · [product-of-array-except-self](https://leetcode.cn/problems/product-of-array-except-self/)

## 题目

给定整数数组 `nums`，返回数组 `answer`，其中 `answer[i]` 等于 `nums` 中除 `nums[i]` 外其余各元素的乘积。

## 题解思路与解析

- 思路，分为左右两部分，分别计算左边的乘积和右边的乘积，然后相乘——》将O(n^2)的复杂度降低到O(n)

## 解答

### Golang

```go
func productExceptSelf(nums []int) []int {
	// 思路，分为左右两部分，分别计算左边的乘积和右边的乘积，然后相乘——》将O(n^2)的复杂度降低到O(n)
	n := len(nums)
	L, R, answer := make([]int, n), make([]int, n), make([]int, n)
	L[0] = 1
	R[n-1] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}
```

### Python

```python
from typing import List


class Solution:
    def product_except_self(self, nums: List[int]) -> List[int]:
        n = len(nums)
        L = [0] * n
        R = [0] * n
        answer = [0] * n
        L[0] = 1
        R[n - 1] = 1
        for i in range(1, n):
            L[i] = L[i - 1] * nums[i - 1]
        for i in range(n - 2, -1, -1):
            R[i] = R[i + 1] * nums[i + 1]
        for i in range(n):
            answer[i] = L[i] * R[i]
        return answer
```
