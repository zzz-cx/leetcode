# partition 分割回文串（LeetCode 131）
# 思路：回文 DP 预处理 + 回溯。f[i][j] 表示 s[i..j] 是否为回文，dfs 枚举所有合法切法
from typing import List


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        n = len(s)
        f = [[True] * n for _ in range(n)]
        for i in range(n - 1, -1, -1):
            for j in range(i + 1, n):
                f[i][j] = s[i] == s[j] and f[i + 1][j - 1]

        ans: List[List[str]] = []
        splits: List[str] = []

        def dfs(i: int) -> None:
            if i == n:
                ans.append(splits[:])
                return
            for j in range(i, n):
                if f[i][j]:
                    splits.append(s[i : j + 1])
                    dfs(j + 1)
                    splits.pop()

        dfs(0)
        return ans


if __name__ == "__main__":
    sol = Solution()
    got = sol.partition("aab")
    status = "PASS" if len(got) == 2 else "FAIL"
    print(f"{status} | parts={got}")
