# 228. 汇总区间

**难度：** 简单  
**标签：** 数组

> [summary-ranges](https://leetcode.cn/problems/summary-ranges/)

## 题目描述

给定一个**无重复元素**的**有序**整数数组 `nums`。

区间 `[a, b]` 是从 `a` 到 `b`（包含）的所有整数的集合。

返回**恰好覆盖数组中所有数字**的最小区间范围列表。即 `nums` 的每个元素都恰好被某个区间覆盖，且区间内不存在不属于 `nums` 的数。

输出格式：

- `"a->b"`，若 `a != b`
- `"a"`，若 `a == b`

### 示例 1

```
输入: nums = [0,1,2,4,5,7]
输出: ["0->2","4->5","7"]
```

解释：

| 区间 | 输出 |
|------|------|
| [0,2] | `"0->2"` |
| [4,5] | `"4->5"` |
| [7,7] | `"7"` |

### 示例 2

```
输入: nums = [0,2,3,4,6,8,9]
输出: ["0","2->4","6","8->9"]
```

### 提示

- `0 <= nums.length <= 20`
- `-2^31 <= nums[i] <= 2^31 - 1`
- `nums` 按升序排列，且无重复

---

## 思路说明

### 问题转化

数组已有序且无重复，**连续数字**在数组中必然相邻。只需找出所有「最长连续段」，每段输出为一个区间。

```
[0, 1, 2, 4, 5, 7]
 └──连续──┘  └─┘  └┘
  0->2       4->5  7
```

### 双指针 / 一次扫描

用下标 `i` 扫描数组：

1. 记录当前连续段起点 `start = nums[i]`
2. 当 `nums[i+1] == nums[i] + 1` 时，`i` 右移，扩展连续段
3. 无法扩展时，记录终点 `end = nums[i]`，格式化输出
4. `i` 继续向后，处理下一段

```
i=0: start=0, 扩展到 2, 输出 "0->2"
i=3: start=4, 扩展到 5, 输出 "4->5"
i=5: start=7, 单独成段, 输出 "7"
```

### 算法流程图

```
开始
  ↓
i = 0
  ↓
i < n ?
  ├─ 否 → 返回 result
  └─ 是 → start = nums[i]
         while i+1<n 且 nums[i+1]==nums[i]+1: i++
         end = nums[i]
         start==end ? 追加 "start" : 追加 "start->end"
         i++
         回到循环
```

### 复杂度

- **时间：** O(n) — 每个元素最多被访问两次
- **空间：** O(1) — 不计输出列表

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **228** | **汇总区间** | 有序数组找连续段 |
| [128](../128-longest-consecutive-sequence/) | 最长连续序列 | 无序数组找最长连续段 |
| [56](../056-merge-intervals/) | 合并区间 | 区间合并 |
| [163](https://leetcode.cn/problems/missing-ranges/) | 缺失的区间 | 找缺失区间（变种） |
| [57](../057-insert-interval/) | 插入区间 | 区间插入与合并 |

**228 vs 128：** 228 数组已有序，直接扫描相邻差是否为 1；128 无序，需哈希集合。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def summary_ranges(self, nums: List[int]) -> List[str]:
        result: list[str] = []
        i, n = 0, len(nums)

        while i < n:
            start = nums[i]
            while i + 1 < n and nums[i + 1] == nums[i] + 1:
                i += 1
            end = nums[i]
            if start == end:
                result.append(str(start))
            else:
                result.append(f"{start}->{end}")
            i += 1

        return result
```

**代码解析**

- **内层 while** — 扩展连续段，条件 `nums[i+1] == nums[i] + 1`。
- **`start == end`** — 单元素区间输出 `"a"`，否则 `"a->b"`。
- **外层 `i += 1`** — 进入下一段连续区间。

### Golang

文件：[solution.go](./solution.go)

```go
func summaryRanges(nums []int) []string {
    result := make([]string, 0)
    i, n := 0, len(nums)

    for i < n {
        start := nums[i]
        for i+1 < n && nums[i+1] == nums[i]+1 {
            i++
        }
        end := nums[i]
        if start == end {
            result = append(result, strconv.Itoa(start))
        } else {
            result = append(result, fmt.Sprintf("%d->%d", start, end))
        }
        i++
    }
    return result
}
```

**代码解析**

- **逻辑与 Python 完全一致** — 双 while 扩展连续段。
- **`strconv.Itoa`** — 单元素格式化；区间用 `fmt.Sprintf`。

---

## 运行

```bash
python scripts/run.py 228
python scripts/run.py 228 --lang go
```

**预期输出**

```
PASS | nums=[0, 1, 2, 4, 5, 7] => ['0->2', '4->5', '7']
PASS | nums=[0, 2, 3, 4, 6, 8, 9] => ['0', '2->4', '6', '8->9']
PASS | nums=[] => []
PASS | nums=[1] => ['1']
```
