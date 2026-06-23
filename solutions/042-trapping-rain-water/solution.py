# Trap 接雨水 - 按双指针找区间的方式实现
# 思路：不动指针作为区间左边界，动指针向右移动，当发现下一个数小于当前数时停止（找到峰顶）
# 此时 [不动指针, 动指针] 构成一个可接雨水的区间，计算区间内雨水后，将不动指针移到动指针位置继续
from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        n = len(height)
        if n < 3:
            return 0

        total = 0
        left = 0

        while left < n - 1:
            right = left + 1
            while right < n and height[right] < height[left]:
                right += 1

            if right == n:
                max_idx = left + 1
                for i in range(left + 2, n):
                    if height[i] >= height[max_idx]:
                        max_idx = i
                right = max_idx

            if right > left + 1:
                water_level = min(height[left], height[right])
                for i in range(left + 1, right):
                    if water_level > height[i]:
                        total += water_level - height[i]

            left = right

        return total

    def trap2(self, height: List[int]) -> int:
        n = len(height)
        if n < 3:
            return 0
        left_max = [0] * n
        left_max[0] = height[0]
        for i in range(1, n):
            left_max[i] = max(left_max[i - 1], height[i])
        right_max = [0] * n
        right_max[n - 1] = height[n - 1]
        for i in range(n - 2, -1, -1):
            right_max[i] = max(right_max[i + 1], height[i])
        total = 0
        for i in range(n):
            total += min(left_max[i], right_max[i]) - height[i]
        return total


if __name__ == "__main__":
    sol = Solution()
    h = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
    got = sol.trap2(h)
    status = "PASS" if got == 6 else "FAIL"
    print(f"{status} | => {got} (expected 6)")
