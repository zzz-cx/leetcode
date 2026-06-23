# 滑动窗口最大值

> LeetCode 239 · [sliding-window-maximum](https://leetcode.cn/problems/sliding-window-maximum/)

## 题目

给定数组和滑动窗口大小 k，返回每个窗口中的最大值。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	n := len(nums)
	if n == 0 {
		return ans
	}
	deque := make([]int, 0)
	for i := 0; i < n; i++ {
		if i >= k && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}
	return ans
}
```

### Python

```python
from typing import List


class Solution:
    def max_sliding_window(self, nums: List[int], k: int) -> List[int]:
        ans = []
        n = len(nums)
        if n == 0:
            return ans
        deque = []
        for i in range(n):
            if i >= k and deque[0] <= i - k:
                deque = deque[1:]
            while deque and nums[deque[-1]] < nums[i]:
                deque.pop()
            deque.append(i)
            if i >= k - 1:
                ans.append(nums[deque[0]])
        return ans
```
