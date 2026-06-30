# 125. 验证回文串

**难度：** 简单  
**标签：** 双指针、字符串

> [valid-palindrome](https://leetcode.cn/problems/valid-palindrome/)

## 题目描述

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样，则可以认为该短语是一个**回文串**。

字母和数字都属于字母数字字符。

给你一个字符串 `s`，如果它是回文串，返回 `true`；否则，返回 `false`。

### 示例 1

```
输入: s = "A man, a plan, a canal: Panama"
输出: true
```

解释：`"amanaplanacanalpanama"` 是回文串。

### 示例 2

```
输入: s = "race a car"
输出: false
```

解释：`"raceacar"` 不是回文串。

### 示例 3

```
输入: s = " "
输出: true
```

解释：移除非字母数字字符后为空串，空串是回文串。

### 提示

- `1 <= s.length <= 2 * 10^5`
- `s` 仅由可打印 ASCII 字符组成

---

## 思路说明

### 问题转化

过滤后的有效字符序列必须满足「首尾对称」。不必先构造新字符串，可用双指针直接在原串上比较。

### 双指针策略

维护 `left` 和 `right`：

1. `left` 向右跳过非字母数字字符。
2. `right` 向左跳过非字母数字字符。
3. 比较 `s[left]` 与 `s[right]`（忽略大小写）。
4. 不相等则返回 `false`；相等则 `left++`、`right--` 继续。
5. 当 `left >= right` 时，返回 `true`。

空串（如 `" "` 过滤后）会在循环开始前或跳过字符后直接满足条件。

### 算法流程图

```
开始
  ↓
left = 0, right = len(s) - 1
  ↓
left < right ?
  ├─ 否 → 返回 true
  └─ 是 → 跳过 left/right 处的非字母数字
         ↓
         s[left] 与 s[right] 忽略大小写相等？
           ├─ 否 → 返回 false
           └─ 是 → left++, right--，继续
```

### 复杂度

- **时间：** O(n)
- **空间：** O(1)

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def is_palindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1
        while left < right:
            while left < right and not s[left].isalnum():
                left += 1
            while left < right and not s[right].isalnum():
                right -= 1
            if s[left].lower() != s[right].lower():
                return False
            left += 1
            right -= 1
        return True
```

**代码解析**

- **`isalnum()`** — 判断字母或数字，自动过滤标点与空格。
- **`lower()`** — 忽略大小写比较。
- **双指针** — 无需额外数组，O(1) 空间。

### Golang

文件：[solution.go](./solution.go)

```go
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        for left < right && !isAlphaNum(s[left]) {
            left++
        }
        for left < right && !isAlphaNum(s[right]) {
            right--
        }
        if toLower(s[left]) != toLower(s[right]) {
            return false
        }
        left++
        right--
    }
    return true
}
```

**代码解析**

- **`isAlphaNum`** — 手动判断字母数字（本题字符集为 ASCII）。
- **`toLower`** — 大写转小写后比较。
- **逻辑与 Python 一致** — 双指针原地验证回文。

---

## 运行

```bash
python scripts/run.py 125
python scripts/run.py 125 --lang go
```

**预期输出**

```
PASS | s='A man, a plan, a canal: Panama' => True (expected True)
PASS | s='race a car' => False (expected False)
PASS | s=' ' => True (expected True)
PASS | s='0P' => False (expected False)
```
