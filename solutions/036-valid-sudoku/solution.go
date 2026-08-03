package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			digit := board[r][c] - '1'
			box := (r/3)*3 + c/3
			if rows[r][digit] || cols[c][digit] || boxes[box][digit] {
				return false
			}
			rows[r][digit] = true
			cols[c][digit] = true
			boxes[box][digit] = true
		}
	}
	return true
}

func main() {
	valid := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	invalidRow := make([][]byte, 9)
	for i := range valid {
		invalidRow[i] = append([]byte(nil), valid[i]...)
	}
	invalidRow[0][0] = '3'

	tests := []struct {
		board    [][]byte
		expected bool
	}{
		{valid, true},
		{invalidRow, false},
	}

	for _, tc := range tests {
		result := isValidSudoku(tc.board)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s => %v (expected %v)\n", status, result, tc.expected)
	}
}
