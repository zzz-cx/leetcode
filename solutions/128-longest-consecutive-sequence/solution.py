from typing import List


class Solution:
    def longest_consecutive(self, nums: List[int]) -> int:
        num_set = set(nums)
        best = 0

        for num in num_set:
            if num - 1 in num_set:
                continue

            length = 1
            while num + length in num_set:
                length += 1
            best = max(best, length)

        return best


if __name__ == "__main__":
    tests = [
        ([100, 4, 200, 1, 3, 2], 4),
        ([0, 3, 7, 2, 5, 8, 4, 6, 0, 1], 9),
        ([1, 0, 1, 2], 3),
        ([], 0),
    ]
    sol = Solution()
    for nums, expected in tests:
        result = sol.longest_consecutive(nums)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | nums={nums!r} => {result} (expected {expected})")
