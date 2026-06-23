package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	got := coinChange([]int{1, 2, 5}, 11)
	status := "PASS"
	if got != 3 {
		status = "FAIL"
	}
	fmt.Printf("%s | coins=%d\n", status, got)
}
