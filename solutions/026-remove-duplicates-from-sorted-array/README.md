# 删除有序数组中的重复项

> LeetCode 26 · [remove-duplicates-from-sorted-array](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)

## 题目

给定升序数组 `nums`，原地删除重复项，使每个元素只出现一次，返回删除后数组的新长度。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func removeDuplicates(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1

}
```

### Python

```python
from typing import List


class Solution:
    def remove_duplicates(self, nums: List[int]) -> int:
        k = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[k]:
                k += 1
                nums[k] = nums[i]
        return k + 1
```
