from typing import List


class Solution:
    def find_duplicate(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1
        while left < right:
            mid = (left + right) // 2
            count = sum(1 for num in nums if num <= mid)
            if count > mid:
                right = mid
            else:
                left = mid + 1
        return left


if __name__ == "__main__":
    sol = Solution()
    got = sol.find_duplicate([1, 3, 4, 2, 2])
    status = "PASS" if got == 2 else "FAIL"
    print(f"{status} | => {got}")
