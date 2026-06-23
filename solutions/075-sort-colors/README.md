# 颜色分类

> LeetCode 75 · [sort-colors](https://leetcode.cn/problems/sort-colors/)

## 题目

给定包含红(0)、白(1)、蓝(2) 的数组，原地排序使相同颜色相邻（荷兰国旗问题）。

## 题解思路与解析

- 思路：荷兰国旗三指针。left 左侧全是 0，right 右侧全是 2，i 扫描中间未分区段

## 解答

### Golang

```go
// 思路：荷兰国旗三指针。left 左侧全是 0，right 右侧全是 2，i 扫描中间未分区段
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		switch nums[i] {
		case 0:
			nums[left], nums[i] = nums[i], nums[left]
			left++
		case 2:
			nums[right], nums[i] = nums[i], nums[right]
			right--
			i-- // 换过来的是未检查元素，下一轮继续看 i
		}
	}
}
```

### Python

```python
# 思路：荷兰国旗三指针。left 左侧全是 0，right 右侧全是 2，i 扫描中间未分区段
from typing import List


class Solution:
    def sort_colors(self, nums: List[int]) -> None:
        left, right = 0, len(nums) - 1
        i = 0
        while i <= right:
            if nums[i] == 0:
                nums[left], nums[i] = nums[i], nums[left]
                left += 1
                i += 1
            elif nums[i] == 2:
                nums[right], nums[i] = nums[i], nums[right]
                right -= 1
                # 换过来的是未检查元素，下一轮继续看 i
            else:
                i += 1
```
