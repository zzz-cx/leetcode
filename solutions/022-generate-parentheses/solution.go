package main

import "fmt"

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

func main() {
	got := generateParenthesis(3)
	status := "PASS"
	if len(got) != 5 {
		status = "FAIL"
	}
	fmt.Printf("%s | %d combinations\n", status, len(got))
}
