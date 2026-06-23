package main

import "fmt"

// 思路：二维 DP。dp[i][j] = word1 前 i 个字符变成 word2 前 j 个字符的最少操作数
// 相等则 dp[i][j]=dp[i-1][j-1]；否则取 删/增/替 三者最小 +1
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 删 i 个
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 插 j 个
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
				//              删 word1[i-1]  插 word2[j-1]  替换
			}
		}
	}
	return dp[m][n]
}

func main() {
	got := minDistance("horse", "ros")
	fmt.Printf("PASS | %d\n", got)
}
