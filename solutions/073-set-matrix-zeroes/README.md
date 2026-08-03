# 73. 矩阵置零

**难度：** 中等  
**标签：** 数组、哈希表、矩阵

> [set-matrix-zeroes](https://leetcode.cn/problems/set-matrix-zeroes/)

## 题目描述

给定一个 `m × n` 的矩阵，如果一个元素为 `0`，则将其所在**行**和**列**的所有元素都设为 `0`。

请使用**原地**算法。

### 示例 1

```
输入: matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出: [[1,0,1],[0,0,0],[1,0,1]]
```

### 示例 2

```
输入: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
```

### 提示

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 200`
- `-2^31 <= matrix[i][j] <= 2^31 - 1`

---

## 思路说明

### 问题分析

若 `(i, j)` 为 0，则第 `i` 行、第 `j` 列全部置零。难点在于：**不能在第一次扫描时就直接置零**，否则后续扫描会把「被连带置零的位置」误判为原始 0，产生连锁错误。

```
错误做法示例:
1 1 1       看到 (1,1)=0 立刻置零
1 0 1  →    1 0 1
1 1 1       0 0 0   ← 第二行全变 0，第三行也被误伤
            0 0 0
```

正确流程：**先标记哪些行/列需要置零，再统一置零**。

### 方法一：额外数组标记（O(m + n) 空间）

用 `row[i]`、`col[j]` 记录第 `i` 行 / 第 `j` 列是否需要置零：

```
第一遍: 遇 0 → row[i]=True, col[j]=True
第二遍: row[i] 或 col[j] 为 True → matrix[i][j]=0
```

- **时间：** O(m · n)
- **空间：** O(m + n)

思路直观，但不符合题目「原地 O(1) 额外空间」的要求。

### 方法二：首行首列作标记（推荐，O(1) 空间）

用矩阵**第一行**和**第一列**代替 `row[]` / `col[]` 数组：

```
matrix[i][0]  标记第 i 行是否需要置零
matrix[0][j]  标记第 j 列是否需要置零
```

但第一行、第一列本身也可能含有 0，需要两个布尔变量单独记录：

```
first_row_zero  — 第一行是否原本就有 0
first_col_zero  — 第一列是否原本就有 0
```

### 算法流程（O(1) 空间）

```
1. 记录 first_row_zero, first_col_zero
2. 遍历 i=1..m-1, j=1..n-1（跳过首行首列）
   若 matrix[i][j]==0 → matrix[i][0]=0, matrix[0][j]=0
3. 再次遍历内部区域
   若 matrix[i][0]==0 或 matrix[0][j]==0 → matrix[i][j]=0
4. 若 first_row_zero → 第一行全置 0
5. 若 first_col_zero → 第一列全置 0
```

```
标记示意（3×3，(1,1)=0）:

步骤2 后（用首行首列存标记）:
1  1  1        col标记
1  0  1   →    0  0  1
0  1  1        0  1  1
↑
row标记

步骤3~5 后:
1  0  1
0  0  0
1  0  1
```

### 为什么不能直接在第一遍就改首行首列

若 `(0, j)` 或 `(i, 0)` 本身就是 0，第一遍扫描时就会写入标记 0，与「原本就需要置零」混在一起。因此：

- 首行首列的标记信息，在步骤 2~3 中只反映**内部区域 (1..m-1, 1..n-1)** 的 0
- 首行首列是否最终置零，由 `first_row_zero` / `first_col_zero` 在步骤 4~5 单独处理

### 复杂度

| 方法 | 时间 | 额外空间 |
|------|------|----------|
| 额外数组 | O(m·n) | O(m + n) |
| 首行首列标记 | O(m·n) | **O(1)** |

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **73** | **矩阵置零** | 首行首列作标记，O(1) 空间 |
| [48](../048-rotate-image/) | 旋转图像 | 矩阵原地变换 |
| [54](../054-spiral-matrix/) | 螺旋矩阵 | 矩阵分层遍历 |
| [36](../036-valid-sudoku/) | 有效的数独 | 矩阵坐标与分块 |
| [289](../289-game-of-life/) | 生命游戏 | 矩阵状态更新需避免覆盖（状态编码） |

**核心模式：** 矩阵原地修改时，用**已有单元格**或**边界行/列**存储辅助信息，避免 O(m+n) 额外数组。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def set_zeroes(self, matrix: List[List[int]]) -> None:
        m, n = len(matrix), len(matrix[0])
        first_row_zero = any(matrix[0][j] == 0 for j in range(n))
        first_col_zero = any(matrix[i][0] == 0 for i in range(m))

        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][j] == 0:
                    matrix[i][0] = 0
                    matrix[0][j] = 0

        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    matrix[i][j] = 0

        if first_row_zero:
            for j in range(n):
                matrix[0][j] = 0
        if first_col_zero:
            for i in range(m):
                matrix[i][0] = 0
```

**代码解析**

- **`matrix[i][0]` / `matrix[0][j]`** — 复用首列/首行存行、列标记。
- **两个 `first_*_zero`** — 单独保存首行/首列是否原本含 0，最后处理。
- **两遍扫描内部区域** — 第一遍写标记，第二遍按标记置零。

### Golang

文件：[solution.go](./solution.go)

```go
func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0])

    firstRowZero := false
    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            firstRowZero = true
            break
        }
    }
    firstColZero := false
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            firstColZero = true
            break
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    if firstRowZero {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    if firstColZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
```

**代码解析**

- 逻辑与 Python 完全一致，注意循环从 `1` 开始，跳过首行首列。
- **`break` 优化** — 找到 0 即可确定首行/首列需要置零。

---

## 运行

```bash
python scripts/run.py 73
python scripts/run.py 73 --lang go
```

**预期输出**

```
PASS | matrix=[[1, 0, 1], [0, 0, 0], [1, 0, 1]]
PASS | matrix=[[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]
```
