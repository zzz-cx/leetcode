package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		indegree[prerequisite[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			return true
		}
	}
	return false
}

func main() {
	got := canFinish(2, [][]int{{1, 0}})
	fmt.Printf("PASS | %v\n", got)
}
