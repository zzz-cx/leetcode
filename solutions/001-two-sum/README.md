# 1. 两数之和

**难度：** 简单  
**标签：** 数组、哈希表

> [two-sum](https://leetcode.cn/problems/two-sum/)

## 题目描述

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出**和为目标值**的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。

### 示例 1

```
输入: nums = [2,7,11,15], target = 9
输出: [0,1]
```

因为 `nums[0] + nums[1] == 9`，返回 `[0, 1]`。

### 示例 2

```
输入: nums = [3,2,4], target = 6
输出: [1,2]
```

### 提示

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- 只会存在一个有效答案

---

## 思路说明

### 问题转化

暴力做法是双重循环枚举所有 `(i, j)` 对，时间 O(n²)。关键观察：当我们遍历到 `nums[i]` 时，若存在配对元素，它一定是 `target - nums[i]`，且其下标必然出现在 `i` 之前。

因此可以用哈希表记录「值 → 下标」，在单次遍历中完成查找。

### 哈希表策略

维护哈希表 `m`，键为已遍历元素的值，值为对应下标。

遍历 `nums`，对每个位置 `i`、值 `v`：

1. 查询 `target - v` 是否在 `m` 中。
2. 若在，说明找到了配对，返回 `[m[target-v], i]`。
3. 若不在，将 `m[v] = i` 写入哈希表，继续向后。

题目保证有唯一解，因此一旦找到即可返回。

### 算法流程图

```
开始
  ↓
创建空哈希表 m
  ↓
遍历 i = 0 .. n-1，当前值 v = nums[i]
  ↓
target - v 在 m 中？
  ├─ 是 → 返回 [m[target-v], i]
  └─ 否 → m[v] = i，继续
  ↓
遍历结束（题目保证有解，通常不会到达）
```

### 复杂度

- **时间：** O(n)，每个元素最多入表、查表各一次
- **空间：** O(n)，哈希表最多存储 n 个元素

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
from typing import List


class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        m = {}
        for i, v in enumerate(nums):
            if target - v in m:
                return [m[target - v], i]
            m[v] = i
        return []
```

**代码解析**

- **`m` 哈希表** — 存储「已遍历的值 → 下标」，使查找配对元素从 O(n) 降为 O(1)。
- **`target - v in m`** — 先查后存，避免同一元素被使用两次（例如 `nums = [3, 3], target = 6` 时，第二个 3 会匹配到第一个 3 的下标）。
- **`enumerate`** — 同时获取下标和值，代码更简洁。

### Golang

文件：[solution.go](./solution.go)

```go
func TwoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, v := range nums {
        if j, ok := m[target-v]; ok {
            return []int{j, i}
        }
        m[v] = i
    }
    return nil
}
```

**代码解析**

- **`make(map[int]int)`** — 创建值到下标的映射，与 Python 字典作用相同。
- **`if j, ok := m[target-v]; ok`** — Go 惯用的 map 查询写法，`ok` 表示键是否存在。
- **先查后存** — 与 Python 版本逻辑一致，保证不会把当前元素与自身配对。

---

## 运行

`solution.py` 和 `solution.go` 为纯函数实现，可在 LeetCode 提交，也可自行添加测试：

```python
# Python 快速验证
from solution import Solution
print(Solution().two_sum([2, 7, 11, 15], 9))  # [0, 1]
```

```bash
# Go 可在 main 中调用 TwoSum 验证
```
