# 移除元素

> LeetCode 27 · [remove-element](https://leetcode.cn/problems/remove-element/)

## 题目

给定数组 `nums` 和值 `val`，原地移除所有数值等于 `val` 的元素，返回移除后数组的新长度。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
```

### Python

```python
from typing import List


class Solution:
    def remove_element(self, nums: List[int], val: int) -> int:
        k = 0
        for i in range(len(nums)):
            if nums[i] != val:
                nums[k] = nums[i]
                k += 1
        return k
```
