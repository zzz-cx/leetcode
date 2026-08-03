package main

import "fmt"

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := 0
			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && (board[ni][nj] == 1 || board[ni][nj] == 2) {
					liveNeighbors++
				}
			}

			if board[i][j] == 1 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					board[i][j] = 2
				}
			} else if liveNeighbors == 3 {
				board[i][j] = 3
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	tests := []struct {
		board    [][]int
		expected [][]int
	}{
		{
			[][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
			[][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
		},
		{
			[][]int{{1, 1}, {1, 0}},
			[][]int{{1, 1}, {1, 1}},
		},
	}

	for _, tc := range tests {
		gameOfLife(tc.board)
		status := "PASS"
		if !equal(tc.board, tc.expected) {
			status = "FAIL"
		}
		fmt.Printf("%s | board=%v\n", status, tc.board)
	}
}
