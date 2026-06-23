from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        n = len(nums) - 1
        step = 0
        while n > 0:
            for i in range(n):
                if i + nums[i] >= n:
                    n = i
                    step += 1
                    break
        return step

    def jump2(self, nums: List[int]) -> int:
        max_pos = 0
        end = 0
        step = 0
        for i in range(len(nums) - 1):
            if i + nums[i] > max_pos:
                max_pos = i + nums[i]
            if i == end:
                end = max_pos
                step += 1
        return step


if __name__ == "__main__":
    sol = Solution()
    got = sol.jump([2, 3, 1, 1, 4])
    status = "PASS" if got == 2 else "FAIL"
    print(f"{status} | jumps={got} (expected 2)")
