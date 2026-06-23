# 初始位置是矩阵的左上角，初始方向是向右，当路径超出界限或者进入之前访问过的位置时，顺时针旋转，进入下一个方向。
from typing import List


class Solution:
    def spiral_order(self, matrix: List[List[int]]) -> List[int]:
        m, n = len(matrix), len(matrix[0])
        out: List[int] = []
        direction = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        visited = [[False] * n for _ in range(m)]
        row, col, dir_idx = 0, 0, 0
        for _ in range(m * n):
            out.append(matrix[row][col])
            visited[row][col] = True
            next_row, next_col = row + direction[dir_idx][0], col + direction[dir_idx][1]
            if (
                next_row < 0
                or next_row >= m
                or next_col < 0
                or next_col >= n
                or visited[next_row][next_col]
            ):
                dir_idx = (dir_idx + 1) % 4
                next_row, next_col = row + direction[dir_idx][0], col + direction[dir_idx][1]
            row, col = next_row, next_col
        return out


if __name__ == "__main__":
    sol = Solution()
    got = sol.spiral_order([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
    status = "PASS" if got == [1, 2, 3, 6, 9, 8, 7, 4, 5] else "FAIL"
    print(f"{status} | => {got}")
