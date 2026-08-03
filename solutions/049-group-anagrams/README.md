# 49. 字母异位词分组

**难度：** 中等  
**标签：** 哈希表、字符串、排序

> [group-anagrams](https://leetcode.cn/problems/group-anagrams/)

## 题目描述

给你一个字符串数组，请你将**字母异位词**组合在一起。可以按任意顺序返回结果列表。

字母异位词：字母相同、顺序不同的字符串（如 `"eat"` 与 `"tea"`）。

### 示例 1

```
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

### 示例 2

```
输入: strs = [""]
输出: [[""]]
```

### 示例 3

```
输入: strs = ["a"]
输出: [["a"]]
```

### 提示

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` 由小写英文字母组成

---

## 思路说明

### 问题转化

异位词拥有**相同的字符频次**。把每个字符串映射到同一个「签名（key）」，签名相同的归入同一组。

### 方法一：排序作为 key

将每个字符串排序后作为哈希表的 key：

```
"eat" → "aet"
"tea" → "aet"   ← 同一组
"bat" → "abt"   ← 另一组
```

- **时间：** O(n · k log k)，n 为字符串个数，k 为单串最大长度
- **空间：** O(n · k)

实现简单，适合快速写码。

### 方法二：字符计数作为 key（推荐）

用长度为 26 的计数数组表示每个字母出现次数，再转为 tuple / 数组作为 key：

```
"eat" → (1,0,0,0,1,0,...,1,0)  // a:1, e:1, t:1
"tea" → (1,0,0,0,1,0,...,1,0)  // 相同 key
```

- **时间：** O(n · k)
- **空间：** O(n · k)

本题采用**方法二**，与 [438 找到字符串中所有字母异位词](../438-find-all-anagrams-in-a-string/) 的频次统计思路一致。

### 算法流程图

```
开始
  ↓
初始化 groups = {}
  ↓
遍历每个字符串 s
  ↓
统计 s 中 26 个字母频次 → key
  ↓
groups[key].append(s)
  ↓
返回 groups 的所有 value
```

### 复杂度

- **时间：** O(n · k)
- **空间：** O(n · k)

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **49** | **字母异位词分组** | 哈希表 + 排序/计数 key |
| [438](../438-find-all-anagrams-in-a-string/) | 找到字符串中所有字母异位词 | 固定窗口 + 26 位频次比较 |
| [242](https://leetcode.cn/problems/valid-anagram/) | 有效的字母异位词 | 判断两个串是否异位词 |
| [567](https://leetcode.cn/problems/permutation-in-string/) | 字符串的排列 | 滑动窗口判断 s1 排列是否出现在 s2 |
| [30](../030-substring-with-concatenation-of-all-words/) | 串联所有单词的子串 | 多块滑动窗口 + 单词频次哈希 |

**核心模式：** 用**频次签名**（排序串或计数数组）作为哈希 key，把「本质相同」的字符串归为一类。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
from collections import defaultdict
from typing import List


class Solution:
    def group_anagrams(self, strs: List[str]) -> List[List[str]]:
        groups: dict[tuple[int, ...], list[str]] = defaultdict(list)

        for s in strs:
            count = [0] * 26
            for ch in s:
                count[ord(ch) - ord("a")] += 1
            groups[tuple(count)].append(s)

        return list(groups.values())
```

**代码解析**

- **`count[26]`** — 统计每个小写字母出现次数。
- **`tuple(count)`** — list 不可哈希，转 tuple 作为 dict key。
- **`defaultdict(list)`** — 同 key 的字符串自动追加到同一组。

### Golang

文件：[solution.go](./solution.go)

```go
func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)

    for _, s := range strs {
        var count [26]int
        for i := 0; i < len(s); i++ {
            count[s[i]-'a']++
        }
        groups[count] = append(groups[count], s)
    }

    ans := make([][]string, 0, len(groups))
    for _, g := range groups {
        ans = append(ans, g)
    }
    return ans
}
```

**代码解析**

- **`[26]int` 数组** — Go 中数组是值类型，可直接作为 map key（切片不行）。
- **遍历 map values** — 分组顺序任意，符合题意。
- **O(n·k) 计数** — 比排序 key 更高效。

---

## 运行

```bash
python scripts/run.py 49
python scripts/run.py 49 --lang go
```

**预期输出**

```
PASS | strs=['eat', 'tea', 'tan', 'ate', 'nat', 'bat'] => ...
PASS | strs=[''] => [['']]
PASS | strs=['a'] => [['a']]
```

（分组内及组间顺序均可不同，测试按排序后比较。）
