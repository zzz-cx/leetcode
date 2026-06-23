# 全排列

> LeetCode 46 · [permutations](https://leetcode.cn/problems/permutations/)

## 题目

给定不含重复数字的数组 `nums`，返回其所有可能的全排列。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func permute(nums []int) [][]int { //回溯算法
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}

	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums)) // used[i]=true 表示 nums[i] 已放入当前排列

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path) // 必须拷贝，否则后续回溯会修改同一底层数组
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
```

### Python

```python
from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        if len(nums) == 0:
            return res

        path = []
        used = [False] * len(nums)

        def backtrack() -> None:
            if len(path) == len(nums):
                res.append(path[:])
                return
            for i in range(len(nums)):
                if used[i]:
                    continue
                used[i] = True
                path.append(nums[i])
                backtrack()
                path.pop()
                used[i] = False

        backtrack()
        return res
```
