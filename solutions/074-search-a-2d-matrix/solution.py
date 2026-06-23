from typing import List


class Solution:
    def search_matrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        row, col = 0, n - 1
        while row < m and col >= 0:
            if matrix[row][col] == target:
                return True
            elif matrix[row][col] > target:
                col -= 1
            else:
                row += 1
        return False


if __name__ == "__main__":
    sol = Solution()
    got = sol.search_matrix([[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], 3)
    status = "PASS" if got is True else "FAIL"
    print(f"{status} | => {got}")
