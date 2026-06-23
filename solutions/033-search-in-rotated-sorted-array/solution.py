# 思路：每次 mid 必有一侧仍是有序段，先判断 target 是否落在该有序段内，再缩区间
from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                return mid
            if nums[left] <= nums[mid]:
                if nums[left] <= target < nums[mid]:
                    right = mid - 1
                else:
                    left = mid + 1
            else:
                if nums[mid] < target <= nums[right]:
                    left = mid + 1
                else:
                    right = mid - 1
        return -1


if __name__ == "__main__":
    sol = Solution()
    got = sol.search([4, 5, 6, 7, 0, 1, 2], 0)
    status = "PASS" if got == 4 else "FAIL"
    print(f"{status} | index={got} (expected 4)")
