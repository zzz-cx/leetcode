package main

import "fmt"

//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//先遍历一下数组，查看0的位置，对应的行和列标记一下，最后将对应行和列置零

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	m := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m)
	fmt.Printf("PASS | %v\n", m)
}
