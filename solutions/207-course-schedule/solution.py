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


if __name__ == "__main__":
    sol = Solution()
    got = sol.can_finish(2, [[1, 0]])
    status = "PASS" if got is True else "FAIL"
    print(f"{status} | => {got}")
