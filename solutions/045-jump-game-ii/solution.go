package main

import "fmt"

func jump(nums []int) int { //反向找
	n := len(nums) - 1
	step := 0
	for n > 0 {
		for i := 0; i < n; i++ {
			if i+nums[i] >= n { //找到能跳到目标的位置，然后这就是下一个目标的位置
				n = i
				step++
				break
			}
		}
	}
	return step
}
func jump2(nums []int) int { //正向找
	maxPos := 0
	end := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
		}
		if i == end { //表示已经到达了当前能跳到的最远距离，需要更新下一个能跳到的最远距离
			end = maxPos
			step++
		}
	}
	return step
}

func main() {
	got := jump([]int{2, 3, 1, 1, 4})
	fmt.Printf("PASS | jumps=%d\n", got)
}
