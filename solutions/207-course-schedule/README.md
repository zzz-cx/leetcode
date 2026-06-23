# 课程表

> LeetCode 207 · [course-schedule](https://leetcode.cn/problems/course-schedule/)

## 题目

课程表共 numCourses 门，先修关系 prerequisites，判断能否完成所有课程。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
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
```

### Python

```python
from typing import List


class Solution:
    def can_finish(self, num_courses: int, prerequisites: List[List[int]]) -> bool:
        indegree = [0] * num_courses
        for prerequisite in prerequisites:
            indegree[prerequisite[0]] += 1
        for i in range(num_courses):
            if indegree[i] == 0:
                return True
        return False
```
