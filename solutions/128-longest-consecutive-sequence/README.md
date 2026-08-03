# 128. 最长连续序列

**难度：** 中等  
**标签：** 并查集、数组、哈希表

> [longest-consecutive-sequence](https://leetcode.cn/problems/longest-consecutive-sequence/)

## 题目描述

给定一个未排序的整数数组 `nums`，找出数字连续的最长序列（**不要求**序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 **O(n)** 的算法。

### 示例 1

```
输入: nums = [100,4,200,1,3,2]
输出: 4
```

解释：最长数字连续序列是 `[1, 2, 3, 4]`，长度为 4。

### 示例 2

```
输入: nums = [0,3,7,2,5,8,4,6,0,1]
输出: 9
```

解释：最长序列为 `[0, 1, 2, 3, 4, 5, 6, 7, 8]`。

### 示例 3

```
输入: nums = [1,0,1,2]
输出: 3
```

解释：最长序列为 `[0, 1, 2]`（重复元素只计一次）。

### 提示

- `0 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

---

## 思路说明

### 问题转化

在集合中寻找最长的「连续整数段」。例如 `[100,4,200,1,3,2]` 的集合里，连续段有 `{1,2,3,4}` 和 `{100}`、`{200}`，最长为 4。

### 方法一：排序（不符合 O(n) 要求）

排序后线性扫描统计连续段，时间 **O(n log n)**，不满足题意。

### 方法二：哈希集合 + 只从序列起点扩展（推荐）

1. 将所有数放入 `set`（去重，O(n)）
2. 对每个 `num`，**仅当 `num - 1` 不在 set 中**时，说明 `num` 是某连续段的起点
3. 从 `num` 开始不断检查 `num+1, num+2, ...` 是否在 set 中，统计长度
4. 取全局最大值

```
nums = [100, 4, 200, 1, 3, 2]
set  = {1, 2, 3, 4, 100, 200}

num=1: 0∉set → 起点，扩展 1→2→3→4，长度 4
num=2: 1∈set → 跳过（不是起点）
num=3: 2∈set → 跳过
num=4: 3∈set → 跳过
num=100: 99∉set → 起点，扩展长度 1
...
```

### 为什么是 O(n)

看似双重循环，但每个数最多被「扩展内层 while」访问**一次**：

- 只有序列起点才进入 while
- 内层每步 `length++` 对应一个数被计数一次

所有 while 的总步数 ≤ n，整体 **O(n)**。

### 算法流程图

```
开始
  ↓
num_set = set(nums)
best = 0
  ↓
遍历 num in num_set
  ↓
num - 1 in num_set ? → 是 → continue（非起点）
  ↓ 否
length = 1,  while num+length in num_set: length++
  ↓
best = max(best, length)
  ↓
返回 best
```

### 复杂度

- **时间：** O(n)
- **空间：** O(n) — 哈希集合

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **128** | **最长连续序列** | 哈希 set + 只从起点扩展 |
| [1](../001-two-sum/) | 两数之和 | 哈希表 O(1) 查找 |
| [49](../049-group-anagrams/) | 字母异位词分组 | 哈希分组 |
| [202](../202-happy-number/) | 快乐数 | 哈希判环 |
| [298](https://leetcode.cn/problems/find-the-missing-number/) | 找缺失数 | 连续整数 / 集合思维 |

**核心模式：** 用 **哈希集合 O(1) 查找** 替代排序，配合「只从关键位置出发」保证总工作量 O(n)。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def longest_consecutive(self, nums: List[int]) -> int:
        num_set = set(nums)
        best = 0

        for num in num_set:
            if num - 1 in num_set:
                continue

            length = 1
            while num + length in num_set:
                length += 1
            best = max(best, length)

        return best
```

**代码解析**

- **`set(nums)`** — 去重，O(n) 构建。
- **`num - 1 in num_set`** — 过滤非起点，避免重复计算同一段。
- **`while num + length in num_set`** — 从起点向右扩展，统计连续长度。

### Golang

文件：[solution.go](./solution.go)

```go
func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{}, len(nums))
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    best := 0
    for num := range numSet {
        if _, ok := numSet[num-1]; ok {
            continue
        }

        length := 1
        for {
            if _, ok := numSet[num+length]; !ok {
                break
            }
            length++
        }
        if length > best {
            best = length
        }
    }
    return best
}
```

**代码解析**

- **`map[int]struct{}`** — 空 struct 不占额外内存，作集合用。
- **逻辑与 Python 一致** — 起点判定 + 向右扩展。
- **空数组** — 循环不执行，返回 0。

---

## 运行

```bash
python scripts/run.py 128
python scripts/run.py 128 --lang go
```

**预期输出**

```
PASS | nums=[100, 4, 200, 1, 3, 2] => 4 (expected 4)
PASS | nums=[0, 3, 7, 2, 5, 8, 4, 6, 0, 1] => 9 (expected 9)
PASS | nums=[1, 0, 1, 2] => 3 (expected 3)
PASS | nums=[] => 0 (expected 0)
```
