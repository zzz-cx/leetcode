package main

import "fmt"

// 思路：贪心思想，记录每个字符最后出现的位置，然后遍历字符串，记录当前字符最后出现的位置，就是end，如果后续有字符的最后位置大于end了，就说明可以分割，小于就说明这个字符都在start到end的区间里
func partitionLabels(s string) []int {
	var result []int
	lastIndex := make(map[rune]int)
	for i, ch := range s {
		lastIndex[ch] = i
	}
	start := 0
	end := 0
	for i, ch := range s {
		if lastIndex[ch] > end {
			end = lastIndex[ch]
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}

func main() {
	got := partitionLabels("ababcbacadefegdehijhklij")
	fmt.Printf("PASS | %v\n", got)
}
