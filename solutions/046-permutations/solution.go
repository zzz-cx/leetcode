package main

import "fmt"

func permute(nums []int) [][]int { //回溯算法
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}

	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums)) // used[i]=true 表示 nums[i] 已放入当前排列

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path) // 必须拷贝，否则后续回溯会修改同一底层数组
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

func main() {
	nums := []int{1, 2, 3}
	got := permute(nums)
	status := "PASS"
	if len(got) != 6 {
		status = "FAIL"
	}
	fmt.Printf("%s | nums=%v => %d permutations\n", status, nums, len(got))
}
