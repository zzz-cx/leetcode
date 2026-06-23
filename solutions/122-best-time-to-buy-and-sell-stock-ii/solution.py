# max_profit_ii 买卖股票 II（LeetCode 122）：可多次买卖，同一天可先卖再买
# 思路：贪心。把所有「上涨段」的利润都吃掉 —— 等价于累加每天比前一天高的差价
from typing import List


class Solution:
    def max_profit_ii(self, prices: List[int]) -> int:
        profit = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i - 1]:
                profit += prices[i] - prices[i - 1]
        return profit


if __name__ == "__main__":
    sol = Solution()
    got = sol.max_profit_ii([7, 1, 5, 3, 6, 4])
    status = "PASS" if got == 7 else "FAIL"
    print(f"{status} | profit={got}")
