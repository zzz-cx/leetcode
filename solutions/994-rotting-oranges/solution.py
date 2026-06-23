# oranges_rotting 思路：多源 BFS（把所有腐烂橘子同时作为起点）。
from typing import List


class Solution:
    def oranges_rotting(self, grid: List[List[int]]) -> int:
        if len(grid) == 0 or len(grid[0]) == 0:
            return 0
        m, n = len(grid), len(grid[0])

        queue = []
        fresh = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 2:
                    queue.append((i, j))
                elif grid[i][j] == 1:
                    fresh += 1

        if fresh == 0:
            return 0

        dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        minutes = 0
        while queue and fresh > 0:
            level_size = len(queue)
            for _ in range(level_size):
                r, c = queue.pop(0)
                for dr, dc in dirs:
                    nr, nc = r + dr, c + dc
                    if nr < 0 or nr >= m or nc < 0 or nc >= n:
                        continue
                    if grid[nr][nc] != 1:
                        continue
                    grid[nr][nc] = 2
                    fresh -= 1
                    queue.append((nr, nc))
            minutes += 1

        if fresh > 0:
            return -1
        return minutes


if __name__ == "__main__":
    sol = Solution()
    grid = [[2, 1, 1], [1, 1, 0], [0, 1, 1]]
    got = sol.oranges_rotting(grid)
    status = "PASS" if got == 4 else "FAIL"
    print(f"{status} | minutes={got}")
