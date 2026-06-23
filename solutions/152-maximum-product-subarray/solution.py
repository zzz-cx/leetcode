from typing import List


class Solution:
    def max_product(self, nums: List[int]) -> int:
        max_ending, min_ending = nums[0], nums[0]
        ans = nums[0]
        for i in range(1, len(nums)):
            if nums[i] < 0:
                max_ending, min_ending = min_ending, max_ending
            max_ending = max(nums[i], max_ending * nums[i])
            min_ending = min(nums[i], min_ending * nums[i])
            ans = max(ans, max_ending)
        return ans


if __name__ == "__main__":
    sol = Solution()
    got = sol.max_product([2, 3, -2, 4])
    status = "PASS" if got == 6 else "FAIL"
    print(f"{status} | => {got}")
