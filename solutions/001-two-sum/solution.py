from typing import List


class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        m = {}
        for i, v in enumerate(nums):
            if target - v in m:
                return [m[target - v], i]
            m[v] = i
        return []


if __name__ == "__main__":
    tests = [([2, 7, 11, 15], 9, [0, 1]), ([3, 2, 4], 6, [1, 2]), ([3, 3], 6, [0, 1])]
    sol = Solution()
    for nums, target, expected in tests:
        result = sol.two_sum(nums, target)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | nums={nums}, target={target} => {result} (expected {expected})")
