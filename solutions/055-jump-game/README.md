# 跳跃游戏

> LeetCode 55 · [jump-game](https://leetcode.cn/problems/jump-game/)

## 题目

判断能否从第一个下标出发，每次跳跃不超过 `nums[i]`，到达最后一个下标。

## 题解思路与解析

- 思路：贪心算法，从左到右遍历数组，每次都选择能跳跃到最远的距离
- 1. 初始化一个变量rightmost，表示当前能跳跃到的最远距离
- 2. 遍历数组，每次都选择能跳跃到最远的距离，如果当前位置能跳跃到最远的距离大于等于数组长度，则返回true，否则返回false

## 解答

### Golang

```go
//思路：贪心算法，从左到右遍历数组，每次都选择能跳跃到最远的距离
//1. 初始化一个变量rightmost，表示当前能跳跃到的最远距离
//2. 遍历数组，每次都选择能跳跃到最远的距离，如果当前位置能跳跃到最远的距离大于等于数组长度，则返回true，否则返回false

func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost { //表示当前能跳到最远的距离大于等于当前位置
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}
```

### Python

```python
# 思路：贪心算法，从左到右遍历数组，每次都选择能跳跃到最远的距离
from typing import List


class Solution:
    def can_jump(self, nums: List[int]) -> bool:
        n = len(nums)
        rightmost = 0
        for i in range(n):
            if i <= rightmost:
                rightmost = max(rightmost, i + nums[i])
                if rightmost >= n - 1:
                    return True
        return False
```
