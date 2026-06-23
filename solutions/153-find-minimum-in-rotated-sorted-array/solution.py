# 思路：二分缩区间，用 nums[mid] 与 nums[right] 比较判断最小值在左半还是右半
from typing import List


class Solution:
    def find_min(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1
        while left < right:
            mid = (left + right) // 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        return nums[left]


if __name__ == "__main__":
    sol = Solution()
    got = sol.find_min([3, 4, 5, 1, 2])
    status = "PASS" if got == 1 else "FAIL"
    print(f"{status} | => {got}")
