from typing import List


class Solution:
    def max_profit(self, prices: List[int]) -> int:
        min_price = prices[0]
        max_profit = 0
        for price in prices:
            if price < min_price:
                min_price = price
            elif price - min_price > max_profit:
                max_profit = price - min_price
        return max_profit


if __name__ == "__main__":
    sol = Solution()
    got = sol.max_profit([7, 1, 5, 3, 6, 4])
    status = "PASS" if got == 5 else "FAIL"
    print(f"{status} | profit={got}")
