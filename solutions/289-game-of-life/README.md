# 289. 生命游戏

**难度：** 中等  
**标签：** 数组、矩阵、模拟

> [game-of-life](https://leetcode.cn/problems/game-of-life/)

## 题目描述

根据 [生命游戏](https://baike.baidu.com/item/%E7%94%9F%E5%91%BD%E6%B8%B8%E6%88%8F)（Conway's Game of Life），给定 `m × n` 面板，每个格子为活细胞 `1` 或死细胞 `0`。

每个细胞与**八个相邻格子**（水平、垂直、对角线）同时遵循以下规则：

| 当前状态 | 周围活细胞数 | 下一状态 |
|----------|:------------:|----------|
| 活 | < 2 | 死（孤独） |
| 活 | 2 或 3 | 活 |
| 活 | > 3 | 死（拥挤） |
| 死 | = 3 | 活（繁殖） |

下一状态由所有细胞**同时**更新得到。给定当前 `board`，**原地**更新到下一状态（无需返回值）。

### 示例 1

```
输入: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
输出: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
```

### 示例 2

```
输入: board = [[1,1],[1,0]]
输出: [[1,1],[1,1]]
```

### 提示

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 25`
- `board[i][j]` 为 `0` 或 `1`

---

## 思路说明

### 问题分析

关键约束：**所有细胞同时更新**。若在第一次扫描时就直接把活细胞改成 `0`，会影响尚未统计到的邻居数量，造成错误连锁。

```
错误示例: 三个活细胞排成一行 1 1 1
中间格邻居数为 2，应存活
若先改左边 1→0，中间格统计时邻居变少 → 误判为死亡
```

因此需要：**第一遍只记录下一状态，第二遍再统一落地**。

### 方法一：复制矩阵（O(m·n) 空间）

先复制 `board` 到 `copy`，在 `copy` 上统计邻居，结果写回 `board`。

- 思路最简单
- 额外 O(m·n) 空间

### 方法二：状态编码（推荐，O(1) 空间）

用额外数值在 `board` 中**同时存当前态和下一态**：

| 值 | 含义 |
|:--:|------|
| `0` | 死 → 死 |
| `1` | 活 → 活 |
| `2` | 活 → 死（当前仍算活） |
| `3` | 死 → 活（下一回合变活） |

统计邻居时，**`1` 和 `2` 都视为当前活细胞**（`2` 只是标记「下回合会死」）。

```
第一遍: 统计邻居 → 按规则写入 2 或 3
第二遍: 2 → 0,  3 → 1
```

### 算法流程

```
开始
  ↓
遍历每个格子 (i, j)
  ↓
统计 8 邻域中 board[nx][ny] ∈ {1, 2} 的个数
  ↓
当前为活(1) 且邻居<2 或 >3 → 标记为 2
当前为死(0) 且邻居==3      → 标记为 3
  ↓
第二遍: 2→0, 3→1
  ↓
结束
```

### 示例 1 简析

```
初始:          下一状态:
0 1 0          0 0 0
0 0 1    →     1 0 1
1 1 1          0 1 1
0 0 0          0 1 0
```

中间行 `1 1 1` 中，左右两端的 `1` 各有 3 个活邻居而复活/存活，中间的 `1` 因邻居过多（3 个）在某些位置会死亡——需按完整 8 邻域逐格计算。

### 复杂度

| 方法 | 时间 | 额外空间 |
|------|------|----------|
| 复制矩阵 | O(m·n) | O(m·n) |
| 状态编码 | O(m·n) | **O(1)** |

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **289** | **生命游戏** | 状态编码，避免同时更新覆盖 |
| [73](../073-set-matrix-zeroes/) | 矩阵置零 | 首行首列作标记，原地修改 |
| [48](../048-rotate-image/) | 旋转图像 | 四元素环交换 |
| [54](../054-spiral-matrix/) | 螺旋矩阵 | 矩阵方向遍历 |
| [36](../036-valid-sudoku/) | 有效的数独 | 8 邻域 / 矩阵坐标 |

**核心模式：** 矩阵**原地**且**不能边算边改**时，用额外编码位或多遍扫描分离「读」与「写」。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def game_of_life(self, board: List[List[int]]) -> None:
        m, n = len(board), len(board[0])
        directions = (
            (-1, -1), (-1, 0), (-1, 1),
            (0, -1),           (0, 1),
            (1, -1),  (1, 0),  (1, 1),
        )

        for i in range(m):
            for j in range(n):
                live_neighbors = 0
                for di, dj in directions:
                    ni, nj = i + di, j + dj
                    if 0 <= ni < m and 0 <= nj < n and board[ni][nj] in (1, 2):
                        live_neighbors += 1

                if board[i][j] == 1:
                    if live_neighbors < 2 or live_neighbors > 3:
                        board[i][j] = 2
                elif live_neighbors == 3:
                    board[i][j] = 3

        for i in range(m):
            for j in range(n):
                if board[i][j] == 2:
                    board[i][j] = 0
                elif board[i][j] == 3:
                    board[i][j] = 1
```

**代码解析**

- **`board[nx][ny] in (1, 2)`** — `2` 表示「当前活、下回合死」，统计邻居时仍算活。
- **`2` / `3` 标记** — 第一遍只写标记，不破坏当前态的可读性。
- **第二遍归一化** — 将编码还原为 `0` / `1`。

### Golang

文件：[solution.go](./solution.go)

```go
func gameOfLife(board [][]int) {
    m, n := len(board), len(board[0])

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            liveNeighbors := 0
            for _, d := range directions {
                ni, nj := i+d[0], j+d[1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n &&
                    (board[ni][nj] == 1 || board[ni][nj] == 2) {
                    liveNeighbors++
                }
            }

            if board[i][j] == 1 {
                if liveNeighbors < 2 || liveNeighbors > 3 {
                    board[i][j] = 2
                }
            } else if liveNeighbors == 3 {
                board[i][j] = 3
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 2 {
                board[i][j] = 0
            } else if board[i][j] == 3 {
                board[i][j] = 1
            }
        }
    }
}
```

**代码解析**

- **`directions` 数组** — 8 个方向的偏移量，避免 8 层 if。
- **逻辑与 Python 一致** — 两遍扫描 + 状态编码。
- **O(1) 额外空间** — 满足原地更新要求。

---

## 运行

```bash
python scripts/run.py 289
python scripts/run.py 289 --lang go
```

**预期输出**

```
PASS | board=[[0, 0, 0], [1, 0, 1], [0, 1, 1], [0, 1, 0]]
PASS | board=[[1, 1], [1, 1]]
```
