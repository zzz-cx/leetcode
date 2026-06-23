# 接雨水

> LeetCode 42 · [trapping-rain-water](https://leetcode.cn/problems/trapping-rain-water/)

## 题目

给定 n 个非负整数表示柱状图高度，计算按此排列的柱子能接多少雨水。

## 题解思路与解析

- Trap 接雨水 - 按双指针找区间的方式实现
- 思路：不动指针作为区间左边界，动指针向右移动，当发现下一个数小于当前数时停止（找到峰顶）
- 此时 [不动指针, 动指针] 构成一个可接雨水的区间，计算区间内雨水后，将不动指针移到动指针位置继续
- 动指针：找第一个 height[right] >= height[left] 的位置作为右边界
- 若没找到 >= left 的，说明右边都比 left 低，用剩余部分最高柱的右边界
- 区间 [left, right]，水位 = min(左, 右)，只算中间 (left, right) 的雨水
- 使用动态规划求解

## 解答

### Golang

```go
// Trap 接雨水 - 按双指针找区间的方式实现
// 思路：不动指针作为区间左边界，动指针向右移动，当发现下一个数小于当前数时停止（找到峰顶）
// 此时 [不动指针, 动指针] 构成一个可接雨水的区间，计算区间内雨水后，将不动指针移到动指针位置继续
func Trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	total := 0
	left := 0 // 不动指针，区间左边界

	for left < n-1 {
		right := left + 1
		// 动指针：找第一个 height[right] >= height[left] 的位置作为右边界
		for right < n && height[right] < height[left] {
			right++
		}

		// 若没找到 >= left 的，说明右边都比 left 低，用剩余部分最高柱的右边界
		if right == n {
			maxIdx := left + 1
			for i := left + 2; i < n; i++ {
				if height[i] >= height[maxIdx] {
					maxIdx = i
				}
			}
			right = maxIdx
		}

		// 区间 [left, right]，水位 = min(左, 右)，只算中间 (left, right) 的雨水
		if right > left+1 {
			waterLevel := min(height[left], height[right])
			for i := left + 1; i < right; i++ {
				if waterLevel > height[i] {
					total += waterLevel - height[i]
				}
			}
		}

		left = right
	}

	return total
}

func Trap2(height []int) int {
	//使用动态规划求解
	n := len(height)
	if n < 3 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	total := 0
	for i := 0; i < n; i++ {
		total += min(leftMax[i], rightMax[i]) - height[i]
	}
	return total
}
```

### Python

```python
# Trap 接雨水 - 按双指针找区间的方式实现
# 思路：不动指针作为区间左边界，动指针向右移动，当发现下一个数小于当前数时停止（找到峰顶）
# 此时 [不动指针, 动指针] 构成一个可接雨水的区间，计算区间内雨水后，将不动指针移到动指针位置继续
from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        n = len(height)
        if n < 3:
            return 0

        total = 0
        left = 0

        while left < n - 1:
            right = left + 1
            while right < n and height[right] < height[left]:
                right += 1

            if right == n:
                max_idx = left + 1
                for i in range(left + 2, n):
                    if height[i] >= height[max_idx]:
                        max_idx = i
                right = max_idx

            if right > left + 1:
                water_level = min(height[left], height[right])
                for i in range(left + 1, right):
                    if water_level > height[i]:
                        total += water_level - height[i]

            left = right

        return total

    def trap2(self, height: List[int]) -> int:
        n = len(height)
        if n < 3:
            return 0
        left_max = [0] * n
        left_max[0] = height[0]
        for i in range(1, n):
            left_max[i] = max(left_max[i - 1], height[i])
        right_max = [0] * n
        right_max[n - 1] = height[n - 1]
        for i in range(n - 2, -1, -1):
            right_max[i] = max(right_max[i + 1], height[i])
        total = 0
        for i in range(n):
            total += min(left_max[i], right_max[i]) - height[i]
        return total
```
