package main

import "fmt"

// 思路：荷兰国旗三指针。left 左侧全是 0，right 右侧全是 2，i 扫描中间未分区段
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		switch nums[i] {
		case 0:
			nums[left], nums[i] = nums[i], nums[left]
			left++
		case 2:
			nums[right], nums[i] = nums[i], nums[right]
			right--
			i-- // 换过来的是未检查元素，下一轮继续看 i
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Printf("PASS | sorted=%v\n", nums)
}
