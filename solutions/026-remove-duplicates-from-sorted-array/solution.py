from typing import List


class Solution:
    def remove_duplicates(self, nums: List[int]) -> int:
        k = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[k]:
                k += 1
                nums[k] = nums[i]
        return k + 1


if __name__ == "__main__":
    sol = Solution()
    nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
    k = sol.remove_duplicates(nums)
    status = "PASS" if k == 5 and nums[:k] == [0, 1, 2, 3, 4] else "FAIL"
    print(f"{status} | k={k}, nums[:k]={nums[:k]}")
