from typing import List


class Solution:
    def min_path_sum(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = [[0] * n for _ in range(m)]
        dp[0][0] = grid[0][0]
        for i in range(1, m):
            dp[i][0] = dp[i - 1][0] + grid[i][0]
        for j in range(1, n):
            dp[0][j] = dp[0][j - 1] + grid[0][j]
        for i in range(1, m):
            for j in range(1, n):
                dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
        return dp[m - 1][n - 1]


if __name__ == "__main__":
    sol = Solution()
    got = sol.min_path_sum([[1, 3, 1], [1, 5, 1], [4, 2, 1]])
    status = "PASS" if got == 7 else "FAIL"
    print(f"{status} | => {got} (expected 7)")
