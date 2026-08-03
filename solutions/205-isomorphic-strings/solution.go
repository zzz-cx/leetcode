package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if mapped, ok := sToT[a]; ok {
			if mapped != b {
				return false
			}
		} else {
			if _, ok := tToS[b]; ok {
				return false
			}
			sToT[a] = b
			tToS[b] = a
		}
	}
	return true
}

func main() {
	tests := []struct {
		s, t     string
		expected bool
	}{
		{"egg", "add", true},
		{"f11", "b23", false},
		{"paper", "title", true},
		{"badc", "baba", false},
	}

	for _, tc := range tests {
		result := isIsomorphic(tc.s, tc.t)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | s=%q, t=%q => %v (expected %v)\n", status, tc.s, tc.t, result, tc.expected)
	}
}
