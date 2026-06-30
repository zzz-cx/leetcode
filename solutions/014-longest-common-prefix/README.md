# 14. 最长公共前缀

**难度：** 简单  
**标签：** 字符串、字典树

> [longest-common-prefix](https://leetcode.cn/problems/longest-common-prefix/)

## 题目描述

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 `""`。

### 示例 1

```
输入: strs = ["flower","flow","flight"]
输出: "fl"
```

### 示例 2

```
输入: strs = ["dog","racecar","car"]
输出: ""
```

解释：输入不存在公共前缀。

### 提示

- `1 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` 如果非空，则仅由小写英文字母组成

---

## 思路说明

### 问题转化

最长公共前缀，等价于从第 0 列开始，逐列检查所有字符串在该位置的字符是否相同；一旦出现不匹配或某个字符串已结束，就停止扩展。

### 纵向扫描策略

以第一个字符串 `strs[0]` 为基准，用下标 `i` 表示当前检查的列：

1. 取 `ch = strs[0][i]` 作为这一列的候选字符。
2. 遍历其余字符串 `strs[1..]`，若某个字符串长度 ≤ `i`，或 `strs[j][i] != ch`，则公共前缀为 `strs[0][:i]`。
3. 若所有字符串在第 `i` 列都匹配，则 `i++` 继续下一列。
4. 若 `strs[0]` 全部字符都匹配成功，则返回 `strs[0]`。

### 算法流程图

```
开始
  ↓
strs 为空？ → 是 → 返回 ""
  ↓ 否
i = 0
  ↓
i < len(strs[0])？
  ├─ 否 → 返回 strs[0]
  └─ 是 → ch = strs[0][i]
         ↓
         遍历 j = 1 .. n-1
         ↓
         i >= len(strs[j]) 或 strs[j][i] != ch？
           ├─ 是 → 返回 strs[0][:i]
           └─ 否 → 继续
         ↓
         i++，回到循环
```

### 复杂度

- **时间：** O(S)，S 为所有字符串字符总数；最坏情况下每个字符比较一次
- **空间：** O(1)，仅使用常数额外变量（不计返回字符串）

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
from typing import List


class Solution:
    def longest_common_prefix(self, strs: List[str]) -> str:
        if not strs:
            return ""

        for i in range(len(strs[0])):
            ch = strs[0][i]
            for s in strs[1:]:
                if i >= len(s) or s[i] != ch:
                    return strs[0][:i]

        return strs[0]
```

**代码解析**

- **`for i in range(len(strs[0]))`** — 以第一个字符串的长度为上限逐列扫描。
- **`i >= len(s)`** — 某个字符串提前结束，说明公共前缀不可能再向后延伸。
- **`s[i] != ch`** — 同一列字符不一致，立即返回当前已匹配的前缀 `strs[0][:i]`。
- **循环正常结束** — 说明 `strs[0]` 整个字符串都是公共前缀，直接返回 `strs[0]`。

### Golang

文件：[solution.go](./solution.go)

```go
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for i := 0; i < len(strs[0]); i++ {
        ch := strs[0][i]
        for j := 1; j < len(strs); j++ {
            if i >= len(strs[j]) || strs[j][i] != ch {
                return strs[0][:i]
            }
        }
    }

    return strs[0]
}
```

**代码解析**

- **与 Python 逻辑完全一致** — 纵向扫描，逐列比较字符。
- **`strs[0][:i]`** — Go 切片截取前缀，语义与 Python 切片相同。
- **`ch := strs[0][i]`** — Go 字符串可按字节索引；题目保证为小写字母，无需处理 UTF-8 多字节问题。

---

## 运行

```bash
# Python
python scripts/run.py 14

# Go
python scripts/run.py 14 --lang go
go run solutions/014-longest-common-prefix/solution.go
```

**预期输出**

```
PASS | strs=['flower', 'flow', 'flight'] => 'fl' (expected 'fl')
PASS | strs=['dog', 'racecar', 'car'] => '' (expected '')
PASS | strs=['ab', 'a'] => 'a' (expected 'a')
PASS | strs=[''] => '' (expected '')
```
