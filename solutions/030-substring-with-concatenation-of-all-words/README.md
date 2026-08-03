# 30. 串联所有单词的子串

**难度：** 困难  
**标签：** 哈希表、字符串、滑动窗口

> [substring-with-concatenation-of-all-words](https://leetcode.cn/problems/substring-with-concatenation-of-all-words/)

## 题目描述

给定一个字符串 `s` 和一个字符串数组 `words`。`words` 中所有字符串长度相同。

`s` 中的**串联子串**是指一个包含 `words` 中所有字符串以任意顺序排列连接起来的子串。

返回所有串联子串在 `s` 中的开始索引。你可以以任意顺序返回答案。

### 示例 1

```
输入: s = "barfoothefoobarman", words = ["foo","bar"]
输出: [0,9]
```

解释：`"barfoo"`（下标 0）和 `"foobar"`（下标 9）都是合法串联子串。

### 示例 2

```
输入: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出: []
```

### 示例 3

```
输入: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出: [6,9,12]
```

### 提示

- `1 <= s.length <= 10^4`
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 30`
- `words[i]` 和 `s` 由小写英文字母组成

---

## 思路说明

### 问题转化

设每个单词长度为 `wordLen`，单词个数为 `wordCount`，则合法串联子串长度固定为：

```
totalLen = wordLen * wordCount
```

目标：找出所有起点 `i`，使得 `s[i:i+totalLen]` 可被切分为 `wordCount` 个长度为 `wordLen` 的块，且这些块的多重集合恰好等于 `words`。

### 滑动窗口 + 哈希计数

将所有起点按「在单词块内的偏移」分组：偏移 `offset ∈ [0, wordLen)` 的块边界为 `offset, offset+wordLen, ...`

对每个 `offset` 维护：

| 变量 | 含义 |
|------|------|
| `left` | 窗口左边界（块起点） |
| `right` | 窗口右边界（块起点） |
| `seen` | 当前窗口内各单词出现次数 |
| `used` | 当前窗口内单词块数量 |

每次将 `s[right:right+wordLen]` 作为一个单词块加入窗口：

1. `seen[word]++`，`used++`
2. 若 `seen[word]` 超过 `target[word]`，则不断从左侧缩窗，直到合法
3. 若 `used == wordCount`，说明窗口恰好包含全部单词，记录 `left`
4. 记录后左移一格，继续寻找下一个答案

### 算法流程图

```
开始
  ↓
统计 target = words 中各单词频次
  ↓
offset = 0 .. wordLen-1
  ↓
初始化 left = offset, seen, used
  ↓
right 按 wordLen 步长右移
  ↓
加入 word = s[right:right+wordLen]
  ↓
seen[word] > target[word] ? → 是 → 左缩窗
  ↓
used == wordCount ? → 是 → 记录 left，再左缩窗一格
  ↓
返回所有起点
```

### 复杂度

- **时间：** O(n × wordLen)，n 为 `s` 长度
- **空间：** O(wordCount)

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
from collections import Counter

class Solution:
    def find_substring(self, s: str, words: List[str]) -> List[int]:
        word_len = len(words[0])
        word_count = len(words)
        target = Counter(words)
        result = []

        for offset in range(word_len):
            left = offset
            seen = Counter()
            used = 0

            for right in range(offset, len(s) - word_len + 1, word_len):
                word = s[right : right + word_len]
                seen[word] += 1
                used += 1

                while seen[word] > target.get(word, 0):
                    left_word = s[left : left + word_len]
                    seen[left_word] -= 1
                    left += word_len
                    used -= 1

                if used == word_count:
                    result.append(left)
                    left += word_len
                    used -= 1

        return result
```

**代码解析**

- **按 offset 分组** — 保证窗口边界始终对齐单词块。
- **`while seen[word] > target[word]`** — 窗口内某单词过多时左缩。
- **`used == word_count`** — 窗口恰好包含全部单词块时记录答案。

### Golang

文件：[solution.go](./solution.go)

```go
func findSubstring(s string, words []string) []int {
    wordLen := len(words[0])
    wordCount := len(words)
    target := map[string]int{}
    for _, w := range words {
        target[w]++
    }

    for offset := 0; offset < wordLen; offset++ {
        left := offset
        used := 0
        seen := map[string]int{}

        for right := offset; right <= len(s)-wordLen; right += wordLen {
            word := s[right : right+wordLen]
            seen[word]++
            used++

            for seen[word] > target[word] {
                leftWord := s[left : left+wordLen]
                seen[leftWord]--
                left += wordLen
                used--
            }

            if used == wordCount {
                result = append(result, left)
                left += wordLen
                used--
            }
        }
    }
    return result
}
```

**代码解析**

- **`map[string]int`** — 统计单词频次，与 Python `Counter` 对应。
- **块对齐滑动** — 每次移动 `wordLen`，不是逐字符。
- **逻辑与 Python 一致** — 便于对照学习。

---

## 运行

```bash
python scripts/run.py 30
python scripts/run.py 30 --lang go
```

**预期输出**

```
PASS | s='barfoothefoobarman', words=['foo', 'bar'] => [0, 9] (expected [0, 9])
PASS | s='wordgoodgoodgoodbestword', words=['word', 'good', 'best', 'word'] => [] (expected [])
PASS | s='barfoofoobarthefoobarman', words=['bar', 'foo', 'the'] => [6, 9, 12] (expected [6, 9, 12])
```
