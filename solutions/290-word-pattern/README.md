# 290. 单词规律

**难度：** 简单  
**标签：** 哈希表、字符串

> [word-pattern](https://leetcode.cn/problems/word-pattern/)

## 题目描述

给定一种规律 `pattern` 和一个字符串 `s`，判断 `s` 是否遵循相同的规律。

这里的**遵循**指完全匹配：`pattern` 里的每个字母和 `s` 中的每个非空单词之间存在**双向一一对应**：

- `pattern` 中的每个字母都恰好映射到 `s` 中的一个唯一单词
- `s` 中的每个唯一单词都恰好映射到 `pattern` 中的一个字母
- 没有两个字母映射到同一个单词，也没有两个单词映射到同一个字母

### 示例 1

```
输入: pattern = "abba", s = "dog cat cat dog"
输出: true
```

### 示例 2

```
输入: pattern = "abba", s = "dog cat cat fish"
输出: false
```

### 示例 3

```
输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false
```

### 提示

- `1 <= pattern.length <= 300`
- `pattern` 只包含小写英文字母
- `1 <= s.length <= 3000`
- `s` 只包含小写英文字母和 `' '`
- `s` **不包含** 任何前导或尾随空格
- `s` 中每个单词之间由单个空格分隔

---

## 思路说明

### 问题转化

把 `s` 按空格拆成单词列表 `words`，问题等价于判断两个序列是否**同构**：

```
pattern:  a   b   b   a
words:    dog cat cat dog
```

这与 [205. 同构字符串](../205-isomorphic-strings/) 本质相同，只是把「字符 ↔ 字符」换成了「字符 ↔ 单词」。

### 双向哈希映射

维护两个映射：

| 映射 | 含义 |
|------|------|
| `char_to_word` | `pattern` 中字母 → 单词 |
| `word_to_char` | 单词 → `pattern` 中字母 |

遍历每一对 `(ch, word)`：

1. 若 `ch` 已映射，则必须满足 `char_to_word[ch] == word`。
2. 若 `ch` 未映射，则 `word` 也不能已被其他字母占用（`word ∉ word_to_char`）。
3. 建立 `char_to_word[ch] = word` 和 `word_to_char[word] = ch`。

任一条件不满足则返回 `false`。

**前置检查：** 若 `len(pattern) != len(words)`，直接返回 `false`。

### 算法流程图

```
开始
  ↓
words = split(s)
  ↓
len(pattern) != len(words) ? → 是 → 返回 false
  ↓ 否
初始化 char_to_word, word_to_char
  ↓
遍历 zip(pattern, words)
  ↓
ch 已在 char_to_word 中？
  ├─ 是 → char_to_word[ch] == word ? 否 → 返回 false
  └─ 否 → word 已在 word_to_char 中？ 是 → 返回 false
         建立双向映射
  ↓
全部通过 → 返回 true
```

### 复杂度

- **时间：** O(n + m)，n 为 `pattern` 长度，m 为 `s` 长度（拆分单词）
- **空间：** O(k)，k 为不同字母/单词的数量

---

## 同类型题目

这类题的核心都是：**判断两个序列之间是否存在一一映射（双射）**，常用**双向哈希表**在线维护。

| 题号 | 题目 | 映射关系 | 思路要点 |
|:----:|------|----------|----------|
| [205](../205-isomorphic-strings/) | 同构字符串 | 字符 ↔ 字符 | 本题的直接前身，逻辑完全一致 |
| **290** | **单词规律** | 字母 ↔ 单词 | 先 `split` 再双向映射 |
| [291](https://leetcode.cn/problems/word-pattern-ii/) | 单词规律 II | 字母 ↔ 子串 | 单词不必等长，需 **DFS + 回溯** 尝试分割 |
| [890](https://leetcode.cn/problems/find-and-replace-pattern/) | 查找和替换模式 | 字符 ↔ 字符（模式组） | 对每个单词跑 205 同构判断，再分组 |
| [444](https://leetcode.cn/problems/sequence-reconstruction/) | 序列重建 | 顺序约束 | 哈希 + 拓扑，判断唯一重建（进阶） |

### 通用模板

```python
def is_bijection(seq_a, seq_b) -> bool:
    if len(seq_a) != len(seq_b):
        return False
    a_to_b, b_to_a = {}, {}
    for a, b in zip(seq_a, seq_b):
        if a in a_to_b:
            if a_to_b[a] != b:
                return False
        else:
            if b in b_to_a:
                return False
            a_to_b[a] = b
            b_to_a[b] = a
    return True
```

- **205**：`is_bijection(list(s), list(t))`
- **290**：`is_bijection(list(pattern), s.split())`

### 205 vs 290 vs 291

```
205  字符序列  ←→  字符序列     双向哈希，O(n)
290  字符序列  ←→  单词序列     双向哈希，O(n)
291  字符序列  ←→  可变长子串   回溯枚举分割点
```

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def word_pattern(self, pattern: str, s: str) -> bool:
        words = s.split()
        if len(pattern) != len(words):
            return False

        char_to_word: dict[str, str] = {}
        word_to_char: dict[str, str] = {}

        for ch, word in zip(pattern, words):
            if ch in char_to_word:
                if char_to_word[ch] != word:
                    return False
            else:
                if word in word_to_char:
                    return False
                char_to_word[ch] = word
                word_to_char[word] = ch

        return True
```

**代码解析**

- **`s.split()`** — 按空格拆单词；题目保证单词间只有一个空格。
- **长度不等直接 false** — 字母数与单词数必须一一对应。
- **双向映射** — 与 205 题相同，防止「一字母多单词」或「一单词多字母」。

### Golang

文件：[solution.go](./solution.go)

```go
func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)
    if len(pattern) != len(words) {
        return false
    }

    charToWord := make(map[byte]string)
    wordToChar := make(map[string]byte)

    for i := 0; i < len(pattern); i++ {
        ch := pattern[i]
        word := words[i]
        if mapped, ok := charToWord[ch]; ok {
            if mapped != word {
                return false
            }
        } else {
            if _, ok := wordToChar[word]; ok {
                return false
            }
            charToWord[ch] = word
            wordToChar[word] = ch
        }
    }
    return true
}
```

**代码解析**

- **`strings.Fields(s)`** — 按空白符拆分，效果与 Python `split()` 一致。
- **`map[byte]string` + `map[string]byte`** — 字母用 byte，单词用 string。
- **逻辑与 205 题同构** — 只是映射值的类型从 byte 变成了 string。

---

## 运行

```bash
python scripts/run.py 290
python scripts/run.py 290 --lang go
```

**预期输出**

```
PASS | pattern='abba', s='dog cat cat dog' => True (expected True)
PASS | pattern='abba', s='dog cat cat fish' => False (expected False)
PASS | pattern='aaaa', s='dog cat cat dog' => False (expected False)
PASS | pattern='abba', s='dog dog dog dog' => False (expected False)
PASS | pattern='a', s='dog' => True (expected True)
```
