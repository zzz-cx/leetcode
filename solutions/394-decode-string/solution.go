package main

import "fmt"

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

func main() {
	got := decodeString("3[a]2[bc]")
	fmt.Printf("PASS | %q\n", got)
}
