package main

import "fmt"

// convert 模拟 Z 字形走法，按行收集字符
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([][]byte, numRows)
	curRow := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	out := make([]byte, 0, len(s))
	for _, row := range rows {
		out = append(out, row...)
	}
	return string(out)
}

func main() {
	tests := []struct {
		s, expected string
		numRows     int
	}{
		{"PAYPALISHIRING", "PAHNAPLSIIGYIR", 3},
		{"PAYPALISHIRING", "PINALSIGYAHRPI", 4},
		{"A", "A", 1},
	}

	for _, tc := range tests {
		result := convert(tc.s, tc.numRows)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | s=%q, numRows=%d => %q (expected %q)\n", status, tc.s, tc.numRows, result, tc.expected)
	}
}
