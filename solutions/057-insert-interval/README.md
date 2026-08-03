# 57. 插入区间

**难度：** 中等  
**标签：** 数组

> [insert-interval](https://leetcode.cn/problems/insert-interval/)

## 题目描述

给你一个**无重叠**、按起始端点**升序**排列的区间列表 `intervals`，其中 `intervals[i] = [starti, endi]`。

给定新区间 `newInterval = [start, end]`，将其插入 `intervals`，使结果仍按起始端点升序排列，且区间之间不重叠（必要时合并区间）。

返回插入后的区间列表（**无需原地修改**）。

### 示例 1

```
输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
```

### 示例 2

```
输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
```

解释：新区间 `[4,8]` 与 `[3,5]`、`[6,7]`、`[8,10]` 重叠，合并为 `[3,10]`。

### 提示

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^5`
- `intervals` 按 `starti` 升序排列，且无重叠
- `newInterval.length == 2`
- `0 <= start <= end <= 10^5`

---

## 思路说明

### 问题转化

`intervals` 已有序且无重叠，插入新区间只需处理三类区间：

| 阶段 | 条件 | 操作 |
|------|------|------|
| 左侧 | `intervals[i][1] < newInterval[0]` | 直接加入结果 |
| 重叠 | `intervals[i][0] <= newInterval[1]` | 与新区间合并 |
| 右侧 | 其余 | 直接加入结果 |

### 一次扫描三阶段

```
intervals:  [1,2]  [3,5]  [6,7]  [8,10]  [12,16]
newInterval:      [4,8]

阶段1（左侧）: [1,2]  end=2 < start=4，直接加入
阶段2（合并）: [3,5][6,7][8,10] 都与 [4,8] 重叠
               合并为 [3,10]
阶段3（右侧）: [12,16] 直接加入

结果: [[1,2],[3,10],[12,16]]
```

合并时不断更新新区间边界：

```
newInterval[0] = min(newInterval[0], intervals[i][0])
newInterval[1] = max(newInterval[1], intervals[i][1])
```

### 重叠判定

两个区间 `[a,b]` 与 `[c,d]` 重叠 ⟺ `c <= b`（在已有序前提下，等价于 `a <= d`）。

```
不相交:  b < c
重叠:    c <= b
```

### 算法流程图

```
开始
  ↓
result = [], i = 0
  ↓
阶段1: while intervals[i].end < new.start → 加入 result, i++
  ↓
阶段2: while intervals[i].start <= new.end → 合并到 new, i++
  ↓
将 new 加入 result
  ↓
阶段3: while i < n → 加入 result, i++
  ↓
返回 result
```

### 复杂度

- **时间：** O(n)
- **空间：** O(n) — 结果数组

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **57** | **插入区间** | 三阶段扫描 + 合并 |
| [56](../056-merge-intervals/) | 合并区间 | 排序后线性合并 |
| [228](../228-summary-ranges/) | 汇总区间 | 有序数组连续段 |
| [128](../128-longest-consecutive-sequence/) | 最长连续序列 | 连续整数段 |
| [435](https://leetcode.cn/problems/non-overlapping-intervals/) | 无重叠区间 | 区间调度贪心 |

**57 vs 56：** 56 对任意区间列表排序后合并；57 输入已有序，插入一个新区间，三阶段 O(n) 即可。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def insert(
        self, intervals: List[List[int]], new_interval: List[int]
    ) -> List[List[int]]:
        merged = [new_interval[0], new_interval[1]]
        result: list[list[int]] = []
        i, n = 0, len(intervals)

        while i < n and intervals[i][1] < merged[0]:
            result.append(intervals[i])
            i += 1

        while i < n and intervals[i][0] <= merged[1]:
            merged[0] = min(merged[0], intervals[i][0])
            merged[1] = max(merged[1], intervals[i][1])
            i += 1
        result.append(merged)

        while i < n:
            result.append(intervals[i])
            i += 1

        return result
```

**代码解析**

- **阶段 1** — `end < new.start`，新区间左侧的无重叠区间。
- **阶段 2** — `start <= new.end`，不断扩展 `new_interval` 边界。
- **阶段 3** — 剩余区间原样追加。

### Golang

文件：[solution.go](./solution.go)

```go
func insert(intervals [][]int, newInterval []int) [][]int {
    result := make([][]int, 0, len(intervals)+1)
    i, n := 0, len(intervals)

    for i < n && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
    }

    for i < n && intervals[i][0] <= newInterval[1] {
        if intervals[i][0] < newInterval[0] {
            newInterval[0] = intervals[i][0]
        }
        if intervals[i][1] > newInterval[1] {
            newInterval[1] = intervals[i][1]
        }
        i++
    }
    result = append(result, newInterval)

    for i < n {
        result = append(result, intervals[i])
        i++
    }
    return result
}
```

**代码解析**

- **三阶段结构与 Python 一致** — 清晰对应左、合并、右。
- **`newInterval` 原地扩展** — 合并时直接修改边界值。
- **测试时复制 `newInterval`** — 避免多次测试互相影响。

---

## 运行

```bash
python scripts/run.py 57
python scripts/run.py 57 --lang go
```

**预期输出**

```
PASS | intervals=[[1, 3], [6, 9]], new=[2, 5] => [[1, 5], [6, 9]]
PASS | intervals=[[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], new=[4, 8] => [[1, 2], [3, 10], [12, 16]]
PASS | intervals=[], new=[5, 7] => [[5, 7]]
PASS | intervals=[[1, 5]], new=[2, 3] => [[1, 5]]
```
