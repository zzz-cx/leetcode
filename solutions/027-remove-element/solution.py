from typing import List


class Solution:
    def remove_element(self, nums: List[int], val: int) -> int:
        k = 0
        for i in range(len(nums)):
            if nums[i] != val:
                nums[k] = nums[i]
                k += 1
        return k


if __name__ == "__main__":
    sol = Solution()
    nums = [0, 1, 2, 2, 3, 0, 4, 2]
    k = sol.remove_element(nums, 2)
    status = "PASS" if k == 5 else "FAIL"
    print(f"{status} | k={k}, nums[:k]={sorted(nums[:k])}")
