# 三数之和

> LeetCode 15 · [3sum](https://leetcode.cn/problems/3sum/)

## 题目

给定整数数组 `nums`，判断是否存在三元组 `[nums[i], nums[j], nums[k]]` 满足 `i != j != k` 且 `nums[i] + nums[j] + nums[k] == 0`。返回所有和为 0 且不重复的三元组。

## 题解思路与解析

- threeSum 三数之和（LeetCode 15）
- 思路：排序 + 固定 i + 双指针找两数之和为 -nums[i]
- 易错：找到一组后必须 left++、right--，否则会无限 append 导致 OOM
- 去重后收缩指针（顺序不能省，最后必须 left++ right--）
- ThreeSum 兼容旧调用名

## 解答

### Golang

```go
"sort"

// threeSum 三数之和（LeetCode 15）
// 思路：排序 + 固定 i + 双指针找两数之和为 -nums[i]
// 易错：找到一组后必须 left++、right--，否则会无限 append 导致 OOM
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int

	for i := 0; i < n-2; i++ { // 至少留两个给 left/right
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复的 i
		}
		if nums[i] > 0 {
			break // 已排序，后面都 > 0，不可能和为 0
		}
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 去重后收缩指针（顺序不能省，最后必须 left++ right--）
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

// ThreeSum 兼容旧调用名
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
```

### Python

```python
# 三数之和（LeetCode 15）
# 思路：排序 + 固定 i + 双指针找两数之和为 -nums[i]
# 易错：找到一组后必须 left++、right--，否则会无限 append 导致 OOM
from typing import List


class Solution:
    def three_sum(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        if n < 3:
            return []
        nums.sort()
        ans = []

        for i in range(n - 2):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            if nums[i] > 0:
                break
            target = -nums[i]
            left, right = i + 1, n - 1
            while left < right:
                s = nums[left] + nums[right]
                if s == target:
                    ans.append([nums[i], nums[left], nums[right]])
                    while left < right and nums[left] == nums[left + 1]:
                        left += 1
                    while left < right and nums[right] == nums[right - 1]:
                        right -= 1
                    left += 1
                    right -= 1
                elif s < target:
                    left += 1
                else:
                    right -= 1
        return ans
```
