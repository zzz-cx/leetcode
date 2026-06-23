from typing import List


class Solution:
    def rorate(self, nums: List[int], k: int) -> None:
        n = len(nums)
        k = k % n
        self._reverse(nums, 0, n - 1)
        self._reverse(nums, 0, k - 1)
        self._reverse(nums, k, n - 1)

    def _reverse(self, nums: List[int], start: int, end: int) -> None:
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1


if __name__ == "__main__":
    sol = Solution()
    nums = [1, 2, 3, 4, 5, 6, 7]
    sol.rorate(nums, 3)
    status = "PASS" if nums == [5, 6, 7, 1, 2, 3, 4] else "FAIL"
    print(f"{status} | rotated={nums}")
