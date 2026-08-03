package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}

func main() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	rotate(m)
	status := "PASS"
	for i := range m {
		for j := range m[i] {
			if m[i][j] != expected[i][j] {
				status = "FAIL"
			}
		}
	}
	fmt.Printf("%s | rotated=%v\n", status, m)
}
