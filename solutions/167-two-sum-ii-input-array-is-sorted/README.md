# 167. 两数之和 II - 输入有序数组

**难度：** 中等  
**标签：** 数组、双指针、二分查找

> [two-sum-ii-input-array-is-sorted](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/)

## 题目描述

给你一个下标从 1 开始的整数数组 `numbers`，该数组已按**非递减顺序**排列，请你从数组中找出满足相加之和等于目标数 `target` 的两个数。

设这两个数分别是 `numbers[index1]` 和 `numbers[index2]`，则 `1 <= index1 < index2 <= numbers.length`。

以长度为 2 的整数数组 `[index1, index2]` 的形式返回这两个整数的下标（**从 1 开始**）。

你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。

### 示例 1

```
输入: numbers = [2,7,11,15], target = 9
输出: [1,2]
```

解释：2 与 7 之和等于 9，下标为 1 和 2。

### 示例 2

```
输入: numbers = [2,3,4], target = 6
输出: [1,3]
```

解释：2 与 4 之和等于 6，下标为 1 和 3。

### 示例 3

```
输入: numbers = [-1,0], target = -1
输出: [1,2]
```

### 提示

- `2 <= numbers.length <= 3 * 10^4`
- `-1000 <= numbers[i] <= 1000`
- `-1000 <= target <= 1000`
- 仅存在一个有效答案

---

## 思路说明

### 问题转化

数组已有序，若固定较小数 `numbers[left]`，较大数只能出现在其右侧。因此不必枚举所有 `(i, j)` 对，可用双指针从两端向中间收缩。

### 双指针策略

| 指针 | 初始位置 | 移动规则 |
|------|----------|----------|
| `left` | 0 | 和太小则右移 |
| `right` | n - 1 | 和太大则左移 |

计算 `total = numbers[left] + numbers[right]`：

- `total == target` → 返回 `[left + 1, right + 1]`（题目要求 1 下标）
- `total < target` → `left++`（需要更大的和）
- `total > target` → `right--`（需要更小的和）

**正确性：** 数组升序，当 `left` 固定时，增大 `right` 只会使和变大；减小 `right` 只会使和变小。每次移动都排除不可能成为答案的一对，且不会漏解。

### 算法流程图

```
开始
  ↓
left = 0, right = n - 1
  ↓
left < right ?
  ├─ 否 → 结束
  └─ 是 → total = numbers[left] + numbers[right]
         ↓
         total == target ? → 返回 [left+1, right+1]
         total < target  ? → left++
         total > target  ? → right--
         ↓
         继续循环
```

### 复杂度

- **时间：** O(n)
- **空间：** O(1)

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def two_sum(self, numbers: List[int], target: int) -> List[int]:
        left, right = 0, len(numbers) - 1
        while left < right:
            total = numbers[left] + numbers[right]
            if total == target:
                return [left + 1, right + 1]
            if total < target:
                left += 1
            else:
                right -= 1
        return []
```

**代码解析**

- **双指针** — 利用有序性，O(n) 找一对数。
- **`left + 1, right + 1`** — 题目下标从 1 开始。
- **和偏小/偏大** — 分别移动左、右指针收缩搜索区间。

### Golang

文件：[solution.go](./solution.go)

```go
func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    for left < right {
        total := numbers[left] + numbers[right]
        if total == target {
            return []int{left + 1, right + 1}
        }
        if total < target {
            left++
        } else {
            right--
        }
    }
    return nil
}
```

**代码解析**

- **与 Python 逻辑一致** — 左右夹逼找目标和。
- **返回 1 下标** — `left + 1, right + 1`。
- **O(1) 额外空间** — 满足题目空间要求。

---

## 运行

```bash
python scripts/run.py 167
python scripts/run.py 167 --lang go
```

**预期输出**

```
PASS | numbers=[2, 7, 11, 15], target=9 => [1, 2] (expected [1, 2])
PASS | numbers=[2, 3, 4], target=6 => [1, 3] (expected [1, 3])
PASS | numbers=[-1, 0], target=-1 => [1, 2] (expected [1, 2])
```
