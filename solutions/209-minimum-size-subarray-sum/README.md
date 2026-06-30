# 209. 长度最小的子数组

**难度：** 中等  
**标签：** 数组、二分查找、前缀和、滑动窗口

> [minimum-size-subarray-sum](https://leetcode.cn/problems/minimum-size-subarray-sum/)

## 题目描述

给定一个含有 `n` 个正整数的数组和一个正整数 `target`。

找出该数组中满足其总和大于等于 `target` 的长度最小的子数组 `[numsl, numsl+1, ..., numsr-1, numsr]`，并返回其长度。如果不存在符合条件的子数组，返回 `0`。

### 示例 1

```
输入: target = 7, nums = [2,3,1,2,4,3]
输出: 2
```

解释：子数组 `[4,3]` 是满足条件的长度最小的子数组。

### 示例 2

```
输入: target = 4, nums = [1,4,4]
输出: 1
```

### 示例 3

```
输入: target = 11, nums = [1,1,1,1,1,1,1,1]
输出: 0
```

### 提示

- `1 <= target <= 10^9`
- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^4`

---

## 思路说明

### 问题转化

在正整数数组中，找**最短**连续子数组，使其元素和 ≥ `target`。

数组元素均为正数：当窗口右边界右移时和增大，左边界右移时和减小。因此可用**滑动窗口**在线性时间内求解。

### 滑动窗口策略

维护窗口 `[left, right]` 及其和 `total`：

1. `right` 右移，将 `nums[right]` 加入窗口。
2. 当 `total >= target` 时，当前窗口合法，更新最小长度 `ans`。
3. 不断右移 `left` 缩小窗口，直到 `total < target`。
4. 继续扩大 `right`。

### 算法流程图

```
开始
  ↓
left = 0, total = 0, ans = +∞
  ↓
right = 0 .. n-1
  ↓
total += nums[right]
  ↓
total >= target ?
  ├─ 是 → ans = min(ans, right-left+1)
  │       total -= nums[left], left++
  │       （循环直到 total < target）
  └─ 否 → 继续
  ↓
ans 仍为 +∞ ? → 是 → 返回 0
  └─ 否 → 返回 ans
```

### 复杂度

- **时间：** O(n)，每个元素最多进出窗口各一次
- **空间：** O(1)

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def min_sub_array_len(self, target: int, nums: List[int]) -> int:
        left = 0
        total = 0
        ans = float("inf")

        for right in range(len(nums)):
            total += nums[right]
            while total >= target:
                ans = min(ans, right - left + 1)
                total -= nums[left]
                left += 1

        return 0 if ans == float("inf") else ans
```

**代码解析**

- **`for right`** — 扩展窗口右边界。
- **`while total >= target`** — 窗口合法时尝试缩小，寻找更短长度。
- **`ans = min(...)`** — 记录满足条件的最小窗口长度。
- **返回 0** — 从未找到合法窗口时。

### Golang

文件：[solution.go](./solution.go)

```go
func minSubArrayLen(target int, nums []int) int {
    left := 0
    total := 0
    ans := len(nums) + 1

    for right := 0; right < len(nums); right++ {
        total += nums[right]
        for total >= target {
            if right-left+1 < ans {
                ans = right - left + 1
            }
            total -= nums[left]
            left++
        }
    }

    if ans == len(nums)+1 {
        return 0
    }
    return ans
}
```

**代码解析**

- **`ans := len(nums) + 1`** — 用不可能的大值代替 Python 的 `inf`。
- **内层 `for`** — 与 Python 的 `while total >= target` 等价。
- **最终判断** — `ans` 未更新则返回 0。

---

## 运行

```bash
python scripts/run.py 209
python scripts/run.py 209 --lang go
```

**预期输出**

```
PASS | target=7, nums=[2, 3, 1, 2, 4, 3] => 2 (expected 2)
PASS | target=4, nums=[1, 4, 4] => 1 (expected 1)
PASS | target=11, nums=[1, 1, 1, 1, 1, 1, 1, 1] => 0 (expected 0)
```
