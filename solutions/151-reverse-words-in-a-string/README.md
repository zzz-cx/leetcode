# 151. 反转字符串中的单词

**难度：** 中等  
**标签：** 双指针、字符串

> [reverse-words-in-a-string](https://leetcode.cn/problems/reverse-words-in-a-string/)

## 题目描述

给你一个字符串 `s`，请你反转字符串中**单词**的顺序。

**单词**是由非空格字符组成的字符串。`s` 中使用至少一个空格将字符串中的单词分隔开。

返回单词顺序颠倒且单词之间用单个空格连接的结果字符串。

**注意：** 输入字符串 `s` 中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

### 示例 1

```
输入: s = "the sky is blue"
输出: "blue is sky the"
```

### 示例 2

```
输入: s = "  hello world  "
输出: "world hello"
```

解释：反转后的字符串中不能存在前导空格和尾随空格。

### 示例 3

```
输入: s = "a good   example"
输出: "example good a"
```

解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到单个。

### 提示

- `1 <= s.length <= 10^4`
- `s` 包含英文大小写字母、数字和空格
- `s` 中 **至少存在一个** 单词

---

## 思路说明

### 问题转化

目标可以拆成两步：

1. **规范化空格** — 去掉首尾空格，将连续多个空格压缩为一个。
2. **反转单词顺序** — 保持每个单词内部字符不变，只交换单词的前后位置。

### 方法一：拆分 + 反转 + 拼接

| 步骤 | 操作 |
|------|------|
| 1 | `split()` / `strings.Fields()` 按空白切分，自动忽略多余空格 |
| 2 | 双指针交换单词数组首尾 |
| 3 | 用单个空格 `join` 拼回字符串 |

实现简单，但需额外 O(n) 空间存储单词列表。

### 方法二：原地反转（O(1) 额外空间）

将字符串转为可变字符数组后：

1. **压缩空格** — 双指针去掉多余空格，得到 `"the sky is blue"` 形式。
2. **整体反转** — `"eulb si yks eht"`。
3. **逐词反转** — 每个单词内部再反转回来，得到 `"blue is sky the"`。

整体时间 O(n)，额外空间 O(1)（不计输出）。Go 面试中常考这种写法。

**本仓库两种实现均已保留：**

| 方法 | Python | Go |
|------|--------|-----|
| 方法一 拆分 | `reverse_words_split` | `reverseWordsSplit` |
| 方法二 原地 | `reverse_words_inplace` / `reverse_words` | `reverseWordsInplace` / `reverseWords` |

### 算法流程图（方法二）

```
开始
  ↓
转为可变字符数组
  ↓
双指针压缩空格（去首尾、合并连续空格）
  ↓
长度为 0？ → 是 → 返回 ""
  ↓ 否
整体反转整个数组
  ↓
遍历每个单词区间 [i, j)
  ↓
反转单词内部字符
  ↓
拼接返回结果
```

### 复杂度

| 方法 | 时间 | 额外空间 |
|------|------|----------|
| 方法一 拆分 | O(n) | O(n) |
| 方法二 原地 | O(n) | O(1) |

---

## 解答

### Python

文件：[solution.py](./solution.py)

**方法一：`reverse_words_split`**

```python
def reverse_words_split(self, s: str) -> str:
    words = s.split()
    left, right = 0, len(words) - 1
    while left < right:
        words[left], words[right] = words[right], words[left]
        left += 1
        right -= 1
    return " ".join(words)
```

**方法二：`reverse_words_inplace`（`reverse_words` 默认调用此方法）**

```python
def reverse_words_inplace(self, s: str) -> str:
    chars = self._trim_spaces(list(s))
    if not chars:
        return ""
    self._reverse(chars, 0, len(chars) - 1)
    i, n = 0, len(chars)
    while i < n:
        j = i
        while j < n and chars[j] != " ":
            j += 1
        self._reverse(chars, i, j - 1)
        i = j + 1
    return "".join(chars)
```

**代码解析**

- **方法一** — `split()` 自动处理多余空格，双指针反转单词列表后 `join`。
- **方法二** — `_trim_spaces` 压缩空格 → 整体 `_reverse` → 逐词 `_reverse`。
- **`reverse_words`** — LeetCode 提交入口，默认使用方法二。

### Golang

文件：[solution.go](./solution.go)

**方法一：`reverseWordsSplit`**

```go
func reverseWordsSplit(s string) string {
    words := strings.Fields(s)
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}
```

**方法二：`reverseWordsInplace`（`reverseWords` 默认调用此方法）**

```go
func reverseWordsInplace(s string) string {
    b := trimSpaces([]byte(s))
    if len(b) == 0 {
        return ""
    }
    reverse(b, 0, len(b)-1)
    for i := 0; i < len(b); {
        j := i
        for j < len(b) && b[j] != ' ' {
            j++
        }
        reverse(b, i, j-1)
        i = j + 1
    }
    return string(b)
}
```

**代码解析**

- **方法一** — `strings.Fields` + 双指针反转，代码简洁，适合快速实现。
- **方法二** — `[]byte` 原地三次反转，面试常考。
- **`reverseWords`** — LeetCode 提交入口，默认使用方法二。

---

## 运行

```bash
# Python
python scripts/run.py 151

# Go
python scripts/run.py 151 --lang go
go run solutions/151-reverse-words-in-a-string/solution.go
```

**预期输出**（两种实现各跑一遍）

```
--- split ---
PASS | ...
--- inplace ---
PASS | ...
```
