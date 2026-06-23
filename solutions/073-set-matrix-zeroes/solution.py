# 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用原地算法。
# 先遍历一下数组，查看0的位置，对应的行和列标记一下，最后将对应行和列置零
from typing import List


class Solution:
    def set_zeroes(self, matrix: List[List[int]]) -> None:
        m, n = len(matrix), len(matrix[0])
        row = [False] * m
        col = [False] * n
        for i in range(m):
            for j in range(n):
                if matrix[i][j] == 0:
                    row[i] = True
                    col[j] = True
        for i in range(m):
            for j in range(n):
                if row[i] or col[j]:
                    matrix[i][j] = 0


if __name__ == "__main__":
    sol = Solution()
    m = [[1, 1, 1], [1, 0, 1], [1, 1, 1]]
    sol.set_zeroes(m)
    expected = [[1, 0, 1], [0, 0, 0], [1, 0, 1]]
    status = "PASS" if m == expected else "FAIL"
    print(f"{status} | matrix={m}")
