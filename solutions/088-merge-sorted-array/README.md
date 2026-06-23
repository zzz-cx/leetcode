# 合并两个有序数组

> LeetCode 88 · [merge-sorted-array](https://leetcode.cn/problems/merge-sorted-array/)

## 题目

将两个升序数组合并到 `nums1` 中，使 `nums1` 成为升序数组。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func merge2(nums1 []int, m int, nums2 []int, n int) []int {
	i := m - 1
	j := n - 1
	k := m + n - 1         //nums1的最后一个位置
	for i >= 0 && j >= 0 { //从后往前遍历，将较大的数放到nums1的最后一个位置
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 { //如果nums2还有剩余，则将剩余的数放到nums1的最后一个位置
		nums1[k] = nums2[j]
		j--
		k--
	}
	return nums1
}
```

### Python

```python
from typing import List


class Solution:
    def merge2(self, nums1: List[int], m: int, nums2: List[int], n: int) -> List[int]:
        i = m - 1
        j = n - 1
        k = m + n - 1
        while i >= 0 and j >= 0:
            if nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
            k -= 1
        while j >= 0:
            nums1[k] = nums2[j]
            j -= 1
            k -= 1
        return nums1
```
