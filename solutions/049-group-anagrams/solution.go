package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++
		}
		groups[count] = append(groups[count], s)
	}

	ans := make([][]string, 0, len(groups))
	for _, g := range groups {
		ans = append(ans, g)
	}
	return ans
}

func normalize(groups [][]string) [][]string {
	out := make([][]string, len(groups))
	for i, g := range groups {
		cp := append([]string(nil), g...)
		sort.Strings(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i][0] < out[j][0]
	})
	return out
}

func main() {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for _, tc := range tests {
		result := groupAnagrams(tc.strs)
		ok := fmt.Sprint(normalize(result)) == fmt.Sprint(normalize(tc.expected))
		status := "PASS"
		if !ok {
			status = "FAIL"
		}
		fmt.Printf("%s | strs=%v => %v\n", status, tc.strs, result)
	}
}
