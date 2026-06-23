# 括号生成

> LeetCode 22 · [generate-parentheses](https://leetcode.cn/problems/generate-parentheses/)

## 题目

数字 n 代表生成括号的对数，请生成所有可能并且有效的括号组合。

## 题解思路与解析

- 思路：回溯算法，每次添加一个左括号或右括号，如果左括号数量小于n，则添加左括号，如果右括号数量小于左括号数量，则添加右括号

## 解答

### Golang

```go
// 思路：回溯算法，每次添加一个左括号或右括号，如果左括号数量小于n，则添加左括号，如果右括号数量小于左括号数量，则添加右括号
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}

	var backtrack func(cur string, left, right int)
	backtrack = func(cur string, left, right int) {
		if len(cur) == 2*n {
			res = append(res, cur) // 闭包捕获 res，append 才能写回外层切片
			return
		}
		if left < n {
			backtrack(cur+"(", left+1, right)
		}
		if right < left {
			backtrack(cur+")", left, right+1)
		}
	}

	backtrack("", 0, 0)
	return res
}
```

### Python

```python
# 思路：回溯算法，每次添加一个左括号或右括号，如果左括号数量小于n，则添加左括号，如果右括号数量小于左括号数量，则添加右括号
from typing import List


class Solution:
    def generate_parenthesis(self, n: int) -> List[str]:
        res = []
        if n == 0:
            return res

        def backtrack(cur: str, left: int, right: int) -> None:
            if len(cur) == 2 * n:
                res.append(cur)
                return
            if left < n:
                backtrack(cur + "(", left + 1, right)
            if right < left:
                backtrack(cur + ")", left, right + 1)

        backtrack("", 0, 0)
        return res
```
