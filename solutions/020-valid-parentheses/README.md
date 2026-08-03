# 20. 有效的括号

**难度：** 简单  
**标签：** 栈、字符串

> [valid-parentheses](https://leetcode.cn/problems/valid-parentheses/)

## 题目描述

给定一个只包括 `'('`、`')'`、`'{'`、`'}'`、`'['`、`']'` 的字符串 `s`，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用**相同类型**的右括号闭合
2. 左括号必须以**正确顺序**闭合
3. 每个右括号都有对应的相同类型左括号

### 示例 1

```
输入: s = "()"
输出: true
```

### 示例 2

```
输入: s = "()[]{}"
输出: true
```

### 示例 3

```
输入: s = "(]"
输出: false
```

### 示例 4

```
输入: s = "([])"
输出: true
```

### 示例 5

```
输入: s = "([)]"
输出: false
```

### 提示

- `1 <= s.length <= 10^4`
- `s` 仅由括号 `'()[]{}'` 组成

---

## 思路说明

### 问题转化

括号匹配是经典的**栈**问题：遇到左括号入栈，遇到右括号检查栈顶是否为匹配的左括号。

```
"([])"  扫描过程:
(  → push '('
[  → push '['
]  → pop, 匹配 '[' ✓
)  → pop, 匹配 '(' ✓
栈空 → true
```

```
"([)]"  扫描到 ) 时:
栈顶是 '['，期望 '(' → 不匹配 → false
```

### 栈算法

1. 遍历每个字符 `ch`
2. 若是**右括号**：栈空或栈顶不匹配 → `false`；否则 `pop`
3. 若是**左括号**：`push`
4. 遍历结束，栈空 → `true`，否则 `false`

用哈希表维护配对关系：

```
')' → '('
']' → '['
'}' → '{'
```

### 算法流程图

```
开始
  ↓
stack = []
  ↓
遍历每个字符 ch
  ↓
ch 是右括号？
  ├─ 是 → stack 空或栈顶不匹配？ → 是 → 返回 false
  │              └─ 否 → pop
  └─ 否 → push(ch)
  ↓
stack 为空？ → 是 → true，否 → false
```

### 复杂度

- **时间：** O(n)
- **空间：** O(n) — 最坏情况全部左括号入栈

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **20** | **有效的括号** | 栈匹配 |
| [22](../022-generate-parentheses/) | 括号生成 | 回溯生成合法括号 |
| [32](https://leetcode.cn/problems/longest-valid-parentheses/) | 最长有效括号 | 栈 / DP |
| [155](../155-min-stack/) | 最小栈 | 栈的基本运用 |
| [394](../394-decode-string/) | 字符串解码 | 栈处理嵌套结构 |
| [84](../084-largest-rectangle-in-histogram/) | 柱状图中最大的矩形 | 单调栈 |

**核心模式：** 遇到「最近匹配 / 嵌套结构」问题，优先考虑**栈**。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def is_valid(self, s: str) -> bool:
        pairs = {")": "(", "]": "[", "}": "{"}
        stack: list[str] = []

        for ch in s:
            if ch in pairs:
                if not stack or stack[-1] != pairs[ch]:
                    return False
                stack.pop()
            else:
                stack.append(ch)

        return not stack
```

**代码解析**

- **`pairs` 字典** — 右括号 → 对应左括号，O(1) 查找。
- **`ch in pairs`** — 判断是否为右括号。
- **`not stack`** — 右括号无左括号可匹配。
- **`return not stack`** — 最终栈空才有效。

### Golang

文件：[solution.go](./solution.go)

```go
var pairs = map[byte]byte{
    ')': '(',
    ']': '[',
    '}': '{',
}

func isValid(s string) bool {
    stack := make([]byte, 0, len(s))

    for i := 0; i < len(s); i++ {
        ch := s[i]
        if open, ok := pairs[ch]; ok {
            if len(stack) == 0 || stack[len(stack)-1] != open {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, ch)
        }
    }
    return len(stack) == 0
}
```

**代码解析**

- **`map[byte]byte`** — 右括号查对应左括号。
- **切片作栈** — `append` 入栈，`stack[:len-1]` 出栈。
- **预分配容量** — `make([]byte, 0, len(s))` 减少扩容。

---

## 运行

```bash
python scripts/run.py 20
python scripts/run.py 20 --lang go
```

**预期输出**

```
PASS | s='()' => True (expected True)
PASS | s='()[]{}' => True (expected True)
PASS | s='(]' => False (expected False)
PASS | s='([])' => True (expected True)
PASS | s='([)]' => False (expected False)
PASS | s='' => True (expected True)
PASS | s='(' => False (expected False)
```
