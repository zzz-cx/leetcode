package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i >= len(word) && s[i-len(word):i] == word {
				dp[i] = dp[i] || dp[i-len(word)]
			}
		}
	}
	return dp[len(s)]
}

func main() {
	got := wordBreak("leetcode", []string{"leet", "code"})
	fmt.Printf("PASS | %v\n", got)
}
