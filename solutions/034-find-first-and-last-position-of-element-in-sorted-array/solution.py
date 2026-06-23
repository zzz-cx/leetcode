# 思路：两次二分——先找第一个等于 target 的下标，再找最后一个
from typing import List


class Solution:
    def search_range(self, nums: List[int], target: int) -> List[int]:
        first = self._find_bound(nums, target, True)
        if first == -1:
            return [-1, -1]
        last = self._find_bound(nums, target, False)
        return [first, last]

    def _find_bound(self, nums: List[int], target: int, find_first: bool) -> int:
        left, right = 0, len(nums) - 1
        ans = -1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                ans = mid
                if find_first:
                    right = mid - 1
                else:
                    left = mid + 1
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return ans


if __name__ == "__main__":
    sol = Solution()
    got = sol.search_range([5, 7, 7, 8, 8, 10], 8)
    status = "PASS" if got == [3, 4] else "FAIL"
    print(f"{status} | range={got} (expected [3, 4])")
