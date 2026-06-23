package main

import (
	"fmt"
	"strings"
)

// countPasswords 统计合法密码个数：a-z 字母，* 表示未知，相邻两位不能相同
//
// 思路：从左到右 DP，dp[c] = 填完当前位置且该位为字母 c 的方案数
//   - 当前位是 *：newDp[k] = sum(dp[j]) - dp[k]（上一位不能等于 k）
//   - 当前位是固定字母 c：newDp[c] = sum(dp[j], j!=c)，其余为 0
func countPasswords(pattern string) int64 {
	slots := parsePasswordPattern(pattern)
	if len(slots) == 0 {
		return 0
	}

	const letters = 26
	dp := make([]int64, letters)

	if slots[0] == '*' {
		for i := range dp {
			dp[i] = 1
		}
	} else {
		dp[slots[0]-'a'] = 1
	}

	for i := 1; i < len(slots); i++ {
		newDp := make([]int64, letters)
		if slots[i] == '*' {
			var total int64
			for _, v := range dp {
				total += v
			}
			for k := range newDp {
				newDp[k] = total - dp[k]
			}
		} else {
			c := int(slots[i] - 'a')
			var sum int64
			for j, v := range dp {
				if j != c {
					sum += v
				}
			}
			newDp[c] = sum
		}
		dp = newDp
	}

	var ans int64
	for _, v := range dp {
		ans += v
	}
	return ans
}

// parsePasswordPattern 解析输入，忽略空格；仅允许 a-z 与 *
func parsePasswordPattern(pattern string) []byte {
	var slots []byte
	for _, r := range strings.TrimSpace(pattern) {
		if r == ' ' {
			continue
		}
		switch {
		case r == '*':
			slots = append(slots, '*')
		case r >= 'a' && r <= 'z':
			slots = append(slots, byte(r))
		default:
			return nil
		}
	}
	return slots
}

func main() {
	tests := []struct {
		pattern string
		want    int64
	}{
		{"a * b", 24},
		{"*", 26},
		{"a", 1},
	}
	for _, tc := range tests {
		got := countPasswords(tc.pattern)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s | pattern=%q => %d (expected %d)\n", status, tc.pattern, got, tc.want)
	}
}
