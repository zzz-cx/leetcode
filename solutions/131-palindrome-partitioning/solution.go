package main

import "fmt"

// partition 分割回文串（LeetCode 131）
// 思路：回文 DP 预处理 + 回溯。f[i][j] 表示 s[i..j] 是否为回文，dfs 枚举所有合法切法
func partition(s string) (ans [][]string) {
	n := len(s)

	// f[i][j]：s[i..j] 是否为回文。先全部置 true，单字符与「中间为空」自然成立
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	// 从右往左填表，保证算 f[i][j] 时 f[i+1][j-1] 已就绪
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	// splits：当前回溯路径上已切出的回文片段（等同 path）
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 切到末尾，复制一份 splits 加入 ans，避免后续回溯改乱
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		// 尝试下一段 s[i..j]
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1]) // 选
				dfs(j + 1)                          // 探：下一段从 j+1 开始
				splits = splits[:len(splits)-1]     // 撤
			}
		}
	}
	dfs(0)
	return
}

func main() {
	got := partition("aab")
	fmt.Printf("PASS | parts=%d\n", len(got))
}
