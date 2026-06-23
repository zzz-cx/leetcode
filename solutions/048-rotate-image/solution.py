from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        for i in range(n // 2):
            for j in range(i, n - i - 1):
                (
                    matrix[i][j],
                    matrix[j][n - i - 1],
                    matrix[n - i - 1][n - j - 1],
                    matrix[n - j - 1][i],
                ) = (
                    matrix[n - j - 1][i],
                    matrix[i][j],
                    matrix[j][n - i - 1],
                    matrix[n - i - 1][n - j - 1],
                )


if __name__ == "__main__":
    sol = Solution()
    m = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    sol.rotate(m)
    expected = [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
    status = "PASS" if m == expected else "FAIL"
    print(f"{status} | rotated={m}")
