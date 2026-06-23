# 思路：二维 DP。dp[i][j] = word1 前 i 个字符变成 word2 前 j 个字符的最少操作数
# 相等则 dp[i][j]=dp[i-1][j-1]；否则取 删/增/替 三者最小 +1
class Solution:
    def min_distance(self, word1: str, word2: str) -> int:
        m, n = len(word1), len(word2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m + 1):
            dp[i][0] = i
        for j in range(n + 1):
            dp[0][j] = j
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if word1[i - 1] == word2[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    dp[i][j] = 1 + min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1])
        return dp[m][n]


if __name__ == "__main__":
    sol = Solution()
    got = sol.min_distance("horse", "ros")
    status = "PASS" if got == 3 else "FAIL"
    print(f"{status} | => {got} (expected 3)")
