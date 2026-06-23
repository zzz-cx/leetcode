package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(nums, 2)
	status := "PASS"
	if k != 5 {
		status = "FAIL"
	}
	fmt.Printf("%s | k=%d\n", status, k)
}
