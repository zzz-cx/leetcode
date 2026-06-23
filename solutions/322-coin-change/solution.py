from typing import List


class Solution:
    def coin_change(self, coins: List[int], amount: int) -> int:
        dp = [amount + 1] * (amount + 1)
        dp[0] = 0
        for i in range(1, amount + 1):
            for coin in coins:
                if coin <= i:
                    dp[i] = min(dp[i], dp[i - coin] + 1)
        if dp[amount] > amount:
            return -1
        return dp[amount]


if __name__ == "__main__":
    sol = Solution()
    got = sol.coin_change([1, 2, 5], 11)
    status = "PASS" if got == 3 else "FAIL"
    print(f"{status} | => {got}")
