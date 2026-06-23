# 字符串解码

> LeetCode 394 · [decode-string](https://leetcode.cn/problems/decode-string/)

## 题目

给定编码字符串，返回解码后的字符串。

## 题解思路与解析

- 两个栈：一个存重复次数，一个存进入当前层之前的字符串

## 解答

### Golang

```go
func decodeString(s string) string {
	// 两个栈：一个存重复次数，一个存进入当前层之前的字符串
	countStack := make([]int, 0)
	strStack := make([]string, 0)

	cur := ""
	num := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch >= '0' && ch <= '9':
			num = num*10 + int(ch-'0')
		case ch == '[':
			countStack = append(countStack, num)
			strStack = append(strStack, cur)
			num = 0
			cur = ""
		case ch == ']':
			k := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			tmp := ""
			for j := 0; j < k; j++ {
				tmp += cur
			}
			cur = prev + tmp
		default:
			cur += string(ch)
		}
	}
	return cur
}
```

### Python

```python
class Solution:
    def decode_string(self, s: str) -> str:
        count_stack = []
        str_stack = []
        cur = ""
        num = 0

        for ch in s:
            if ch.isdigit():
                num = num * 10 + int(ch)
            elif ch == "[":
                count_stack.append(num)
                str_stack.append(cur)
                num = 0
                cur = ""
            elif ch == "]":
                k = count_stack.pop()
                prev = str_stack.pop()
                cur = prev + cur * k
            else:
                cur += ch
        return cur
```
