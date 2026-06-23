# 删除有序数组中的重复项 II

> LeetCode 80 · [remove-duplicates-from-sorted-array-ii](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/)

## 题目

给定升序数组，原地删除重复项使每个元素最多出现两次，返回新长度。

## 题解思路与解析

- removeDuplicates80 删除有序数组重复项 II（LeetCode 80）
- 思路：快慢指针。已写入区 [0..k) 内，同一数最多出现 2 次
- 判定：若 nums[i] != nums[k-2]，说明写入后仍不超过 2 次，可保留

## 解答

### Golang

```go
// removeDuplicates80 删除有序数组重复项 II（LeetCode 80）
// 思路：快慢指针。已写入区 [0..k) 内，同一数最多出现 2 次
// 判定：若 nums[i] != nums[k-2]，说明写入后仍不超过 2 次，可保留
func removeDuplicates80(nums []int) int {
	k := 0
	for _, x := range nums {
		if k < 2 || x != nums[k-2] {
			nums[k] = x
			k++
		}
	}
	return k
}
```

### Python

```python
# remove_duplicates80 删除有序数组重复项 II（LeetCode 80）
# 思路：快慢指针。已写入区 [0..k) 内，同一数最多出现 2 次
from typing import List


class Solution:
    def remove_duplicates80(self, nums: List[int]) -> int:
        k = 0
        for x in nums:
            if k < 2 or x != nums[k - 2]:
                nums[k] = x
                k += 1
        return k
```
