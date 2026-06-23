from typing import List


class Solution:
    def max_sub_array(self, nums: List[int]) -> int:
        max_sum = nums[0]
        for i in range(1, len(nums)):
            if nums[i - 1] > 0:
                nums[i] += nums[i - 1]
            if nums[i] > max_sum:
                max_sum = nums[i]
        return max_sum


if __name__ == "__main__":
    sol = Solution()
    got = sol.max_sub_array([-2, 1, -3, 4, -1, 2, 1, -5, 4])
    status = "PASS" if got == 6 else "FAIL"
    print(f"{status} | sum={got} (expected 6)")
