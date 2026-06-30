from typing import List


class Solution:
    def min_sub_array_len(self, target: int, nums: List[int]) -> int:
        left = 0
        total = 0
        ans = float("inf")

        for right in range(len(nums)):
            total += nums[right]
            while total >= target:
                ans = min(ans, right - left + 1)
                total -= nums[left]
                left += 1

        return 0 if ans == float("inf") else ans


if __name__ == "__main__":
    tests = [
        (7, [2, 3, 1, 2, 4, 3], 2),
        (4, [1, 4, 4], 1),
        (11, [1, 1, 1, 1, 1, 1, 1, 1], 0),
    ]
    sol = Solution()
    for target, nums, expected in tests:
        result = sol.min_sub_array_len(target, nums)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | target={target}, nums={nums} => {result} (expected {expected})")
