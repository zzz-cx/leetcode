# 452. 用最少数量的箭引爆气球

**难度：** 中等  
**标签：** 贪心、数组、排序

> [minimum-number-of-arrows-to-burst-balloons](https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/)

## 题目描述

有一些球形气球贴在 XY 平面的墙面上。`points[i] = [xstart, xend]` 表示气球的水平直径范围。

一支箭沿 x 轴垂直射出，在坐标 `x` 处射箭，若 `xstart ≤ x ≤ xend`，则该气球被引爆。箭可以无限前进，弓箭数量不限。

返回引爆所有气球的**最少**箭数。

### 示例 1

```
输入: points = [[10,16],[2,8],[1,6],[7,12]]
输出: 2
```

解释：

- 在 `x = 6` 射箭，击破 `[2,8]` 和 `[1,6]`
- 在 `x = 11` 射箭，击破 `[10,16]` 和 `[7,12]`

### 示例 2

```
输入: points = [[1,2],[3,4],[5,6],[7,8]]
输出: 4
```

解释：区间互不重叠，每气球需一箭。

### 示例 3

```
输入: points = [[1,2],[2,3],[3,4],[4,5]]
输出: 2
```

解释：

- 在 `x = 2` 射箭，击破 `[1,2]` 和 `[2,3]`
- 在 `x = 4` 射箭，击破 `[3,4]` 和 `[4,5]`

### 提示

- `1 <= points.length <= 10^5`
- `points[i].length == 2`
- `-2^31 <= xstart < xend <= 2^31 - 1`

---

## 思路说明

### 问题转化

每个气球是一个区间 `[start, end]`。在 `x` 处射一箭能引爆所有满足 `start ≤ x ≤ end` 的气球。

求最少箭数 ⟺ 在数轴上选最少的点，使每个区间至少包含一个被选中的点。

这是经典的**区间点覆盖**问题，用贪心解决。

### 贪心策略：按右端点排序

1. 将所有区间按 **end 升序**排序
2. 第一支箭射在第一个区间的右端点 `end`（尽量靠右，覆盖更多后续区间）
3. 遍历后续区间：
   - 若 `start > arrow_pos`，当前箭覆盖不到，需要新箭，射在该区间 `end`
   - 否则当前箭已覆盖，跳过

```
points = [[10,16],[2,8],[1,6],[7,12]]
按 end 排序 → [[1,6],[2,8],[7,12],[10,16]]

箭1: 射在 x=6，覆盖 [1,6] 和 [2,8]
箭2: [7,12] 的 start=7 > 6，新箭射在 x=12，覆盖 [7,12] 和 [10,16]

答案: 2
```

### 为什么射在右端点最优

按 end 排序后，对于当前能覆盖的一组重叠区间，把箭放在**最靠右**的合法位置（当前区间 end），留给后面区间的覆盖空间最大。

若射在左端点，可能错过后续与当前组「刚好衔接」的区间。

### 重叠判定

排序后，当前箭在 `arrow_pos`，下一区间 `[start, end]`：

```
start > arrow_pos  →  不相交，需要新箭
start <= arrow_pos →  相交（箭在 end 处，start 可达），共用
```

### 算法流程图

```
开始
  ↓
按 end 升序排序 points
  ↓
arrows = 1, arrow_pos = points[0].end
  ↓
遍历 i = 1 .. n-1
  ↓
points[i].start > arrow_pos ?
  ├─ 是 → arrows++, arrow_pos = points[i].end
  └─ 否 → 继续
  ↓
返回 arrows
```

### 复杂度

- **时间：** O(n log n) — 排序 dominates
- **空间：** O(log n) — 排序栈空间

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **452** | **用最少数量的箭引爆气球** | 按 end 排序 + 贪心 |
| [435](https://leetcode.cn/problems/non-overlapping-intervals/) | 无重叠区间 | 按 end 排序，贪心选区间 |
| [56](../056-merge-intervals/) | 合并区间 | 区间合并 |
| [57](../057-insert-interval/) | 插入区间 | 三阶段插入合并 |
| [253](https://leetcode.cn/problems/meeting-rooms-ii/) | 会议室 II | 区间调度 |

**452 vs 435：** 452 求覆盖所有区间的最少点数；435 求移除最少区间使剩余不重叠。都按 end 排序贪心，但目标不同。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def find_min_arrow_shots(self, points: List[List[int]]) -> int:
        points.sort(key=lambda p: p[1])
        arrows = 1
        arrow_pos = points[0][1]

        for start, end in points[1:]:
            if start > arrow_pos:
                arrows += 1
                arrow_pos = end

        return arrows
```

**代码解析**

- **`sort(key=lambda p: p[1])`** — 按右端点升序。
- **`arrow_pos`** — 当前箭的 x 坐标，取被射区间的 end。
- **`start > arrow_pos`** — 新区间与当前箭无交集，需新箭。

### Golang

文件：[solution.go](./solution.go)

```go
func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })

    arrows := 1
    arrowPos := points[0][1]

    for i := 1; i < len(points); i++ {
        if points[i][0] > arrowPos {
            arrows++
            arrowPos = points[i][1]
        }
    }
    return arrows
}
```

**代码解析**

- **`sort.Slice` 按 end 排序** — 与 Python 一致。
- **逻辑完全相同** — 贪心选右端点射箭。

---

## 运行

```bash
python scripts/run.py 452
python scripts/run.py 452 --lang go
```

**预期输出**

```
PASS | points=[[10, 16], [2, 8], [1, 6], [7, 12]] => 2 (expected 2)
PASS | points=[[1, 2], [3, 4], [5, 6], [7, 8]] => 4 (expected 4)
PASS | points=[[1, 2], [2, 3], [3, 4], [4, 5]] => 2 (expected 2)
```
