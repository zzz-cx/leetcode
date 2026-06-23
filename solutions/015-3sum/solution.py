# 三数之和（LeetCode 15）
# 思路：排序 + 固定 i + 双指针找两数之和为 -nums[i]
# 易错：找到一组后必须 left++、right--，否则会无限 append 导致 OOM
from typing import List


class Solution:
    def three_sum(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        if n < 3:
            return []
        nums.sort()
        ans = []

        for i in range(n - 2):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            if nums[i] > 0:
                break
            target = -nums[i]
            left, right = i + 1, n - 1
            while left < right:
                s = nums[left] + nums[right]
                if s == target:
                    ans.append([nums[i], nums[left], nums[right]])
                    while left < right and nums[left] == nums[left + 1]:
                        left += 1
                    while left < right and nums[right] == nums[right - 1]:
                        right -= 1
                    left += 1
                    right -= 1
                elif s < target:
                    left += 1
                else:
                    right -= 1
        return ans


if __name__ == "__main__":
    sol = Solution()
    got = sol.three_sum([-1, 0, 1, 2, -1, -4])
    expected = [[-1, -1, 2], [-1, 0, 1]]
    status = "PASS" if sorted(map(tuple, got)) == sorted(map(tuple, expected)) else "FAIL"
    print(f"{status} | triplets={got} (expected {expected})")
