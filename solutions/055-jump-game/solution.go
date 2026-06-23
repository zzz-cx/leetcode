package main

import "fmt"

//思路：贪心算法，从左到右遍历数组，每次都选择能跳跃到最远的距离
//1. 初始化一个变量rightmost，表示当前能跳跃到的最远距离
//2. 遍历数组，每次都选择能跳跃到最远的距离，如果当前位置能跳跃到最远的距离大于等于数组长度，则返回true，否则返回false

func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost { //表示当前能跳到最远的距离大于等于当前位置
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}

func main() {
	got := canJump([]int{2, 3, 1, 1, 4})
	fmt.Printf("PASS | %v\n", got)
}
