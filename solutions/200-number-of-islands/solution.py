from typing import List


class Solution:
    def num_islands(self, grid: List[List[str]]) -> int:
        m, n = len(grid), len(grid[0])
        visited = [[False] * n for _ in range(m)]
        count = 0
        for i in range(m):
            for j in range(n):
                if visited[i][j] or grid[i][j] == "0":
                    continue
                self._dfs(grid, visited, i, j)
                count += 1
        return count

    def _dfs(self, grid: List[List[str]], visited: List[List[bool]], i: int, j: int) -> None:
        if (
            i < 0
            or i >= len(grid)
            or j < 0
            or j >= len(grid[0])
            or visited[i][j]
            or grid[i][j] == "0"
        ):
            return
        visited[i][j] = True
        self._dfs(grid, visited, i + 1, j)
        self._dfs(grid, visited, i - 1, j)
        self._dfs(grid, visited, i, j + 1)
        self._dfs(grid, visited, i, j - 1)


if __name__ == "__main__":
    sol = Solution()
    grid = [["1", "1", "1", "1", "0"], ["1", "1", "0", "1", "0"], ["1", "1", "0", "0", "0"], ["0", "0", "0", "0", "0"]]
    got = sol.num_islands([list(r) for r in grid])
    status = "PASS" if got == 1 else "FAIL"
    print(f"{status} | islands={got}")
