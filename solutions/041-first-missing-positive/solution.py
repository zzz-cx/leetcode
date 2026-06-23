# firstMissingPositive 将 [1,n] 内的数通过交换放到下标 i 上应满足 nums[i]==i+1，
# 再线性扫描第一个不符合的位置。时间 O(n)，额外空间 O(1)。
from typing import List


class Solution:
    def first_missing_positive(self, nums: List[int]) -> int:
        n = len(nums)
        for i in range(n):
            v = nums[i]
            while 1 <= v <= n and nums[v - 1] != v:
                nums[v - 1], nums[i] = nums[i], nums[v - 1]
                v = nums[i]
        for i in range(n):
            if nums[i] != i + 1:
                return i + 1
        return n + 1


if __name__ == "__main__":
    sol = Solution()
    got = sol.first_missing_positive([3, 4, -1, 1])
    status = "PASS" if got == 2 else "FAIL"
    print(f"{status} | => {got} (expected 2)")
