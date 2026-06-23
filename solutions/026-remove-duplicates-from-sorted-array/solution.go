package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1

}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	status := "PASS"
	if k != 5 {
		status = "FAIL"
	}
	fmt.Printf("%s | k=%d nums[:k]=%v\n", status, k, nums[:k])
}
