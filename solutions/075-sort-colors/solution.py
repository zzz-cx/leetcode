# 思路：荷兰国旗三指针。left 左侧全是 0，right 右侧全是 2，i 扫描中间未分区段
from typing import List


class Solution:
    def sort_colors(self, nums: List[int]) -> None:
        left, right = 0, len(nums) - 1
        i = 0
        while i <= right:
            if nums[i] == 0:
                nums[left], nums[i] = nums[i], nums[left]
                left += 1
                i += 1
            elif nums[i] == 2:
                nums[right], nums[i] = nums[i], nums[right]
                right -= 1
                # 换过来的是未检查元素，下一轮继续看 i
            else:
                i += 1


if __name__ == "__main__":
    sol = Solution()
    nums = [2, 0, 2, 1, 1, 0]
    sol.sort_colors(nums)
    status = "PASS" if nums == [0, 0, 1, 1, 2, 2] else "FAIL"
    print(f"{status} | sorted={nums}")
