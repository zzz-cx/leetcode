from typing import List


class Solution:
    def summary_ranges(self, nums: List[int]) -> List[str]:
        result: list[str] = []
        i, n = 0, len(nums)

        while i < n:
            start = nums[i]
            while i + 1 < n and nums[i + 1] == nums[i] + 1:
                i += 1
            end = nums[i]
            if start == end:
                result.append(str(start))
            else:
                result.append(f"{start}->{end}")
            i += 1

        return result


if __name__ == "__main__":
    tests = [
        ([0, 1, 2, 4, 5, 7], ["0->2", "4->5", "7"]),
        ([0, 2, 3, 4, 6, 8, 9], ["0", "2->4", "6", "8->9"]),
        ([], []),
        ([1], ["1"]),
    ]
    sol = Solution()
    for nums, expected in tests:
        result = sol.summary_ranges(nums)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | nums={nums!r} => {result}")
