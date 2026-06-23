from typing import List


class Solution:
    def length_of_lis(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        d = [0] * (len(nums) + 1)
        lens = 1
        d[lens] = nums[0]
        for i in range(1, len(nums)):
            if nums[i] > d[lens]:
                lens += 1
                d[lens] = nums[i]
            else:
                pos = self._binary_search(d, 1, lens, nums[i])
                d[pos] = nums[i]
        return lens

    def _binary_search(self, d: List[int], left: int, right: int, target: int) -> int:
        while left < right:
            mid = (left + right) // 2
            if d[mid] < target:
                left = mid + 1
            else:
                right = mid
        return left

    def length_of_lis2(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        dp = [1] * len(nums)
        for i in range(1, len(nums)):
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return max(dp)


if __name__ == "__main__":
    sol = Solution()
    got = sol.length_of_lis([10, 9, 2, 5, 3, 7, 101, 18])
    status = "PASS" if got == 4 else "FAIL"
    print(f"{status} | => {got}")
