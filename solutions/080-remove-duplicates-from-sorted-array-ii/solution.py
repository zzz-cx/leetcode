# remove_duplicates80 删除有序数组重复项 II（LeetCode 80）
# 思路：快慢指针。已写入区 [0..k) 内，同一数最多出现 2 次
from typing import List


class Solution:
    def remove_duplicates80(self, nums: List[int]) -> int:
        k = 0
        for x in nums:
            if k < 2 or x != nums[k - 2]:
                nums[k] = x
                k += 1
        return k


if __name__ == "__main__":
    sol = Solution()
    nums = [1, 1, 1, 2, 2, 3]
    k = sol.remove_duplicates80(nums)
    status = "PASS" if k == 5 and nums[:k] == [1, 1, 2, 2, 3] else "FAIL"
    print(f"{status} | k={k}, nums[:k]={nums[:k]}")
