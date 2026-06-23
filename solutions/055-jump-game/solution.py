# 思路：贪心算法，从左到右遍历数组，每次都选择能跳跃到最远的距离
from typing import List


class Solution:
    def can_jump(self, nums: List[int]) -> bool:
        n = len(nums)
        rightmost = 0
        for i in range(n):
            if i <= rightmost:
                rightmost = max(rightmost, i + nums[i])
                if rightmost >= n - 1:
                    return True
        return False


if __name__ == "__main__":
    sol = Solution()
    got = sol.can_jump([2, 3, 1, 1, 4])
    status = "PASS" if got is True else "FAIL"
    print(f"{status} | => {got}")
