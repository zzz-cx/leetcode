# 205. 同构字符串

**难度：** 简单  
**标签：** 哈希表、字符串

> [isomorphic-strings](https://leetcode.cn/problems/isomorphic-strings/)

## 题目描述

给定两个字符串 `s` 和 `t`，判断它们是否是同构的。

如果 `s` 中的字符可以按某种映射关系替换得到 `t`，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

### 示例 1

```
输入: s = "egg", t = "add"
输出: true
```

解释：将 `e → a`，`g → d`。

### 示例 2

```
输入: s = "f11", t = "b23"
输出: false
```

解释：`1` 不能同时映射到 `2` 和 `3`。

### 示例 3

```
输入: s = "paper", t = "title"
输出: true
```

### 提示

- `1 <= s.length <= 5 * 10^4`
- `t.length == s.length`
- `s` 和 `t` 由任意有效 ASCII 字符组成

---

## 思路说明

### 问题转化

同构等价于存在**一一映射** `f`，使得 `s[i]` 映射到 `t[i]`，即 `f(s[i]) = t[i]`，且 `f` 是双射（不同字符不能映到同一字符）。

### 双向哈希映射

维护两个映射：

| 映射 | 含义 |
|------|------|
| `s_to_t` | `s` 中字符 → `t` 中字符 |
| `t_to_s` | `t` 中字符 → `s` 中字符 |

遍历每一对 `(a, b)`（`a = s[i]`, `b = t[i]`）：

1. 若 `a` 已映射，则必须满足 `s_to_t[a] == b`。
2. 若 `a` 未映射，则 `b` 也不能已被其他字符占用（`b ∉ t_to_s`）。
3. 建立 `s_to_t[a] = b` 和 `t_to_s[b] = a`。

任一条件不满足则返回 `false`。

### 算法流程图

```
开始
  ↓
len(s) != len(t) ? → 是 → 返回 false
  ↓ 否
初始化 s_to_t, t_to_s
  ↓
遍历 i = 0 .. n-1，a = s[i], b = t[i]
  ↓
a 已在 s_to_t 中？
  ├─ 是 → s_to_t[a] == b ? 否 → 返回 false
  └─ 否 → b 已在 t_to_s 中？ 是 → 返回 false
         建立双向映射
  ↓
全部通过 → 返回 true
```

### 复杂度

- **时间：** O(n)
- **空间：** O(字符集大小)，最多 O(256)

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def is_isomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        s_to_t: dict[str, str] = {}
        t_to_s: dict[str, str] = {}

        for a, b in zip(s, t):
            if a in s_to_t:
                if s_to_t[a] != b:
                    return False
            else:
                if b in t_to_s:
                    return False
                s_to_t[a] = b
                t_to_s[b] = a

        return True
```

**代码解析**

- **双向映射** — 同时保证 `s → t` 和 `t → s` 都是一一对应。
- **`b in t_to_s`** — 防止两个不同 `s` 字符映射到同一个 `t` 字符。
- **`s_to_t[a] != b`** — 防止同一 `s` 字符映射到不同 `t` 字符。

### Golang

文件：[solution.go](./solution.go)

```go
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sToT := make(map[byte]byte)
    tToS := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        a, b := s[i], t[i]
        if mapped, ok := sToT[a]; ok {
            if mapped != b {
                return false
            }
        } else {
            if _, ok := tToS[b]; ok {
                return false
            }
            sToT[a] = b
            tToS[b] = a
        }
    }
    return true
}
```

**代码解析**

- **`map[byte]byte`** — ASCII 字符用 byte 存储即可。
- **逻辑与 Python 一致** — 双向校验保证同构。
- **O(n) 单次遍历** — 高效简洁。

---

## 运行

```bash
python scripts/run.py 205
python scripts/run.py 205 --lang go
```

**预期输出**

```
PASS | s='egg', t='add' => True (expected True)
PASS | s='f11', t='b23' => False (expected False)
PASS | s='paper', t='title' => True (expected True)
PASS | s='badc', t='baba' => False (expected False)
```
