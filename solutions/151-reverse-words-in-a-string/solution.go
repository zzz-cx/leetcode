package main

import (
	"fmt"
	"strings"
)

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func trimSpaces(b []byte) []byte {
	slow, n := 0, len(b)
	for fast := 0; fast < n; fast++ {
		if b[fast] == ' ' {
			continue
		}
		if slow > 0 {
			b[slow] = ' '
			slow++
		}
		for fast < n && b[fast] != ' ' {
			b[slow] = b[fast]
			slow++
			fast++
		}
	}
	return b[:slow]
}

// reverseWordsSplit 方法一：拆分 + 双指针反转 + 拼接
func reverseWordsSplit(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// reverseWordsInplace 方法二：压缩空格 → 整体反转 → 逐词反转
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

// reverseWords LeetCode 提交入口（默认方法二）
func reverseWords(s string) string {
	return reverseWordsInplace(s)
}

func runTests(name string, fn func(string) string, tests []struct {
	input, expected string
}) {
	fmt.Printf("--- %s ---\n", name)
	for _, tc := range tests {
		result := fn(tc.input)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | s=%q => %q (expected %q)\n", status, tc.input, result, tc.expected)
	}
}

func main() {
	tests := []struct {
		input    string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{" ", ""},
	}

	runTests("split", reverseWordsSplit, tests)
	runTests("inplace", reverseWordsInplace, tests)
}
