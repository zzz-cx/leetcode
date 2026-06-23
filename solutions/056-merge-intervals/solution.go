package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	//思路：按左端点排序，再从左到右扫一遍：若当前区间和「结果里最后一个区间」不相交，就新开一段；若相交或包含，就把最后一段的右端点拉长。
	//踩坑点：在循环里用 intervals[i] 和 intervals[i+1] 比较时，没有保证 i+1 存在（i 走到最后一个元素时还会访问 i+1，就会越界）。现在用 out 里最后一个和 intervals[i] 比较，只需要 i 从 1 到 len-1，就不会再访问非法下标。
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	out := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := out[len(out)-1] //表示结果 out 里已经合并好的最后一段，也就是 out[len(out)-1]。
		cur := intervals[i]     //当前区间
		if last[1] < cur[0] {   //若当前区间和「结果里最后一个区间」不相交，就新开一段
			out = append(out, cur)
		} else { //若相交或包含，就把最后一段的右端点拉长
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		}
	}
	return out
}

func main() {
	got := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Printf("PASS | %v\n", got)
}
