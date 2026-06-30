# 28. 找出字符串中第一个匹配项的下标

**难度：** 简单  
**标签：** 双指针、字符串、字符串匹配

> [find-the-index-of-the-first-occurrence-in-a-string](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/)

## 题目描述

给你两个字符串 `haystack` 和 `needle`，请你在 `haystack` 字符串中找出 `needle` 字符串的第一个匹配项的下标（下标从 0 开始）。如果 `needle` 不是 `haystack` 的一部分，则返回 `-1`。

### 示例 1

```
输入: haystack = "sadbutsad", needle = "sad"
输出: 0
```

解释：`"sad"` 在下标 0 和 6 处匹配，第一个匹配项的下标是 0。

### 示例 2

```
输入: haystack = "leetcode", needle = "leeto"
输出: -1
```

解释：`"leeto"` 没有在 `"leetcode"` 中出现。

### 提示

- `1 <= haystack.length, needle.length <= 10^4`
- `haystack` 和 `needle` 仅由小写英文字母组成

---

## 思路说明

### 问题转化

等价于实现 `strStr` / `indexOf`：在文本串 `haystack` 中查找模式串 `needle` 的**第一次**出现位置。

### 滑动窗口 / 暴力匹配

设 `n = len(haystack)`，`m = len(needle)`。

1. 若 `m == 0`，返回 `0`。
2. 枚举 `haystack` 中所有可能的起始下标 `i`（范围 `0 .. n-m`）。
3. 比较 `haystack[i : i+m]` 是否等于 `needle`。
4. 找到第一个相等位置即返回 `i`；枚举结束仍未找到则返回 `-1`。

### 算法流程图

```
开始
  ↓
m = len(needle), n = len(haystack)
  ↓
m == 0 ? → 是 → 返回 0
  ↓ 否
i = 0 .. n-m
  ↓
haystack[i:i+m] == needle ?
  ├─ 是 → 返回 i
  └─ 否 → 继续
  ↓
返回 -1
```

### 复杂度

- **时间：** O(n × m)，最坏情况下每个起点都要比较 m 个字符
- **空间：** O(1)

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def str_str(self, haystack: str, needle: str) -> int:
        n, m = len(haystack), len(needle)
        if m == 0:
            return 0

        for i in range(n - m + 1):
            if haystack[i : i + m] == needle:
                return i
        return -1
```

**代码解析**

- **`range(n - m + 1)`** — 枚举所有可能的起始下标，保证 `i + m` 不越界。
- **`haystack[i : i + m] == needle`** — 直接比较子串，逻辑清晰。
- **找到即返回** — 保证返回的是第一个匹配位置。

### Golang

文件：[solution.go](./solution.go)

```go
func strStr(haystack string, needle string) int {
    n, m := len(haystack), len(needle)
    if m == 0 {
        return 0
    }

    for i := 0; i <= n-m; i++ {
        if haystack[i:i+m] == needle {
            return i
        }
    }
    return -1
}
```

**代码解析**

- **`i <= n-m`** — 与 Python 的 `range(n-m+1)` 等价。
- **`haystack[i:i+m]`** — Go 切片比较，语义与 Python 子串一致。
- **与 Python 逻辑完全一致** — 便于对照学习。

---

## 运行

```bash
# Python
python scripts/run.py 28

# Go
python scripts/run.py 28 --lang go
go run solutions/028-find-the-index-of-the-first-occurrence-in-a-string/solution.go
```

**预期输出**

```
PASS | haystack='sadbutsad', needle='sad' => 0 (expected 0)
PASS | haystack='leetcode', needle='leeto' => -1 (expected -1)
PASS | haystack='hello', needle='ll' => 2 (expected 2)
PASS | haystack='a', needle='a' => 0 (expected 0)
```
