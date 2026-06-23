from typing import List


class Solution:
    def word_break(self, s: str, word_dict: List[str]) -> bool:
        dp = [False] * (len(s) + 1)
        dp[0] = True
        for i in range(1, len(s) + 1):
            for word in word_dict:
                if i >= len(word) and s[i - len(word) : i] == word:
                    dp[i] = dp[i] or dp[i - len(word)]
        return dp[len(s)]


if __name__ == "__main__":
    sol = Solution()
    got = sol.word_break("leetcode", ["leet", "code"])
    status = "PASS" if got is True else "FAIL"
    print(f"{status} | => {got}")
