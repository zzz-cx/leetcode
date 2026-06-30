# 6. Z 字形变换

**难度：** 中等  
**标签：** 字符串

> [zigzag-conversion](https://leetcode.cn/problems/zigzag-conversion/)

## 题目描述

将一个给定字符串 `s` 根据给定的行数 `numRows`，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 `"PAYPALISHIRING"`、行数为 `3` 时，排列如下：

```
P   A   H   N
A P L S I I G
Y   I   R
```

之后，输出需要从左往右逐行读取，产生一个新的字符串，例如：`"PAHNAPLSIIGYIR"`。

请你实现这个将字符串进行指定行数变换的函数：

```cpp
string convert(string s, int numRows);
```

### 示例 1

```
输入: s = "PAYPALISHIRING", numRows = 3
输出: "PAHNAPLSIIGYIR"
```

### 示例 2

```
输入: s = "PAYPALISHIRING", numRows = 4
输出: "PINALSIGYAHRPI"
```

解释：

```
P     I    N
A   L S  I G
Y A   H R
P     I
```

### 示例 3

```
输入: s = "A", numRows = 1
输出: "A"
```

---

## 思路说明

### 问题转化

Z 字形排列的本质是：字符按「向下走 → 向上走 → 向下走……」的折线依次填入 `numRows` 行，最后按行拼接。

### 模拟走法策略

维护三个变量：

| 变量 | 含义 |
|------|------|
| `cur_row` | 当前字符应写入的行号 |
| `going_down` | 当前是否向下移动 |
| `rows` | 每一行收集到的字符 |

遍历字符串 `s` 的每个字符：

1. 将字符追加到 `rows[cur_row]`。
2. 若到达第 `0` 行或第 `numRows - 1` 行，则反转方向。
3. 根据方向更新 `cur_row`（向下 `+1`，向上 `-1`）。

遍历结束后，将各行字符串拼接即为答案。

**边界：** 当 `numRows == 1` 或 `numRows >= len(s)` 时，无法形成 Z 字，直接返回原串。

### 算法流程图

```
开始
  ↓
numRows == 1 或 numRows >= len(s)？
  ├─ 是 → 返回 s
  └─ 否 → 初始化 rows、cur_row=0、going_down=false
  ↓
遍历 s 中每个字符 ch
  ↓
rows[cur_row] += ch
  ↓
cur_row 在首行或末行？
  ├─ 是 → 反转 going_down
  └─ 否 → 继续
  ↓
going_down ? cur_row++ : cur_row--
  ↓
拼接 rows 各行
  ↓
返回结果
```

### 复杂度

- **时间：** O(n)，n 为字符串长度
- **空间：** O(n)，存储各行字符

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def convert(self, s: str, num_rows: int) -> str:
        if num_rows == 1 or num_rows >= len(s):
            return s

        rows = [""] * num_rows
        cur_row = 0
        going_down = False

        for ch in s:
            rows[cur_row] += ch
            if cur_row == 0 or cur_row == num_rows - 1:
                going_down = not going_down
            cur_row += 1 if going_down else -1

        return "".join(rows)
```

**代码解析**

- **`rows = [""] * num_rows`** — 为每一行准备一个字符串容器。
- **`going_down`** — 在首行和末行反转方向，模拟 Z 字折返。
- **`cur_row += 1 if going_down else -1`** — 根据方向移动到下一行。
- **`"".join(rows)`** — 按行从左到右拼接最终结果。

### Golang

文件：[solution.go](./solution.go)

```go
func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    rows := make([][]byte, numRows)
    curRow := 0
    goingDown := false

    for i := 0; i < len(s); i++ {
        rows[curRow] = append(rows[curRow], s[i])
        if curRow == 0 || curRow == numRows-1 {
            goingDown = !goingDown
        }
        if goingDown {
            curRow++
        } else {
            curRow--
        }
    }

    out := make([]byte, 0, len(s))
    for _, row := range rows {
        out = append(out, row...)
    }
    return string(out)
}
```

**代码解析**

- **`rows := make([][]byte, numRows)`** — 每行用 `[]byte` 收集字符，避免频繁字符串拼接。
- **方向反转逻辑** — 与 Python 版本完全一致。
- **最后合并各行** — 将 `rows` 中所有字节拼成结果字符串。

---

## 运行

```bash
# Python
python scripts/run.py 6

# Go
python scripts/run.py 6 --lang go
go run solutions/006-zigzag-conversion/solution.go
```

**预期输出**

```
PASS | s='PAYPALISHIRING', numRows=3 => 'PAHNAPLSIIGYIR' (expected 'PAHNAPLSIIGYIR')
PASS | s='PAYPALISHIRING', numRows=4 => 'PINALSIGYAHRPI' (expected 'PINALSIGYAHRPI')
PASS | s='A', numRows=1 => 'A' (expected 'A')
```
