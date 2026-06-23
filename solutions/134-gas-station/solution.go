package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	totalTank := 0
	currentTank := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		totalTank += diff
		currentTank += diff

		if currentTank < 0 {
			start = i + 1
			currentTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}
	return start
}

func main() {
	tests := []struct {
		gas, cost []int
		expected  int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}, 4},
		{[]int{3, 1, 1}, []int{1, 2, 2}, 0},
	}
	for _, tc := range tests {
		result := canCompleteCircuit(tc.gas, tc.cost)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | gas=%v, cost=%v => %d (expected %d)\n", status, tc.gas, tc.cost, result, tc.expected)
	}
}
