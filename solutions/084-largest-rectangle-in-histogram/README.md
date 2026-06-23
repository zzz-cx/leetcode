# 柱状图中最大的矩形

> LeetCode 84 · [largest-rectangle-in-histogram](https://leetcode.cn/problems/largest-rectangle-in-histogram/)

## 题目

给定非负整数数组表示柱状图高度，求柱状图中最大矩形的面积。

## 题解思路与解析

- largestRectangleArea 柱状图中最大的矩形（LeetCode 84）
- 思路：单调递增栈（存下标，栈内对应高度严格递增）
- 对每个柱子 mid，若以其高度为矩形高，左右能延伸多远？
- - 右边界：第一个比它矮的位置 i（当前扫描到的下标）
- - 左边界：栈中 mid 下面那个下标 + 1（若栈空则左边界为 0）
- 宽度 = i - 左边界 - 1；面积 = heights[mid] * 宽度
- 末尾用虚拟高度 0 的柱子，把栈里剩余下标全部结算一遍
- stack 存下标，且 heights[stack[0]] < heights[stack[1]] < ...
- 多扫一位：i == len(heights) 时 h=0，充当哨兵，强制清空栈
- 当前柱比栈顶矮 → 栈顶柱无法再以当前及右侧延伸，结算以栈顶为高的矩形
- 右边界是 i（第一个比 mid 矮的位置）
- 左边界是 stack 新栈顶的下一位；栈空说明 mid 左侧全部可用

## 解答

### Golang

```go
// largestRectangleArea 柱状图中最大的矩形（LeetCode 84）
//
// 思路：单调递增栈（存下标，栈内对应高度严格递增）
// 对每个柱子 mid，若以其高度为矩形高，左右能延伸多远？
//   - 右边界：第一个比它矮的位置 i（当前扫描到的下标）
//   - 左边界：栈中 mid 下面那个下标 + 1（若栈空则左边界为 0）
// 宽度 = i - 左边界 - 1；面积 = heights[mid] * 宽度
// 末尾用虚拟高度 0 的柱子，把栈里剩余下标全部结算一遍
func largestRectangleArea(heights []int) int {
	// stack 存下标，且 heights[stack[0]] < heights[stack[1]] < ...
	stack := make([]int, 0, len(heights))
	maxArea := 0

	// 多扫一位：i == len(heights) 时 h=0，充当哨兵，强制清空栈
	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}

		// 当前柱比栈顶矮 → 栈顶柱无法再以当前及右侧延伸，结算以栈顶为高的矩形
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			mid := stack[len(stack)-1] // 以这根柱子的高度作为矩形高
			stack = stack[:len(stack)-1]

			// 右边界是 i（第一个比 mid 矮的位置）
			// 左边界是 stack 新栈顶的下一位；栈空说明 mid 左侧全部可用
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, heights[mid]*width)
		}

		if i < len(heights) {
			stack = append(stack, i)
		}
	}
	return maxArea
}
```

### Python

```python
# largest_rectangle_area 柱状图中最大的矩形（LeetCode 84）
# 思路：单调递增栈（存下标，栈内对应高度严格递增）
from typing import List


class Solution:
    def largest_rectangle_area(self, heights: List[int]) -> int:
        stack = []
        max_area = 0

        for i in range(len(heights) + 1):
            h = heights[i] if i < len(heights) else 0

            while stack and heights[stack[-1]] > h:
                mid = stack.pop()
                width = i if not stack else i - stack[-1] - 1
                max_area = max(max_area, heights[mid] * width)

            if i < len(heights):
                stack.append(i)
        return max_area
```
