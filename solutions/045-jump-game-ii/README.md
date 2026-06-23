# 跳跃游戏 II

> LeetCode 45 · [jump-game-ii](https://leetcode.cn/problems/jump-game-ii/)

## 题目

给定非负整数数组 `nums`，你最初位于第一个下标，每个元素代表在该位置可跳跃的最大长度，返回到达最后一个下标的最少跳跃次数。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func jump(nums []int) int { //反向找
	n := len(nums) - 1
	step := 0
	for n > 0 {
		for i := 0; i < n; i++ {
			if i+nums[i] >= n { //找到能跳到目标的位置，然后这就是下一个目标的位置
				n = i
				step++
				break
			}
		}
	}
	return step
}
func jump2(nums []int) int { //正向找
	maxPos := 0
	end := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
		}
		if i == end { //表示已经到达了当前能跳到的最远距离，需要更新下一个能跳到的最远距离
			end = maxPos
			step++
		}
	}
	return step
}
```

### Python

```python
from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        n = len(nums) - 1
        step = 0
        while n > 0:
            for i in range(n):
                if i + nums[i] >= n:
                    n = i
                    step += 1
                    break
        return step

    def jump2(self, nums: List[int]) -> int:
        max_pos = 0
        end = 0
        step = 0
        for i in range(len(nums) - 1):
            if i + nums[i] > max_pos:
                max_pos = i + nums[i]
            if i == end:
                end = max_pos
                step += 1
        return step
```
