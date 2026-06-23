from typing import List


class Solution:
    def next_permutation(self, nums: List[int]) -> None:
        i = len(nums) - 2
        while i >= 0 and nums[i] >= nums[i + 1]:
            i -= 1
        if i >= 0:
            j = len(nums) - 1
            while j >= 0 and nums[j] <= nums[i]:
                j -= 1
            nums[i], nums[j] = nums[j], nums[i]
        self._reverse2(nums, i + 1, len(nums) - 1)

    def _reverse2(self, nums: List[int], start: int, end: int) -> None:
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1


if __name__ == "__main__":
    sol = Solution()
    nums = [1, 2, 3]
    sol.next_permutation(nums)
    status = "PASS" if nums == [1, 3, 2] else "FAIL"
    print(f"{status} | => {nums} (expected [1, 3, 2])")
