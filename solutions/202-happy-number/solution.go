package main

import "fmt"

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		n = squareSum(n)
	}
	return n == 1
}

func main() {
	tests := []struct {
		n        int
		expected bool
	}{
		{19, true},
		{2, false},
		{1, true},
		{7, true},
	}

	for _, tc := range tests {
		result := isHappy(tc.n)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | n=%d => %v (expected %v)\n", status, tc.n, result, tc.expected)
	}
}
