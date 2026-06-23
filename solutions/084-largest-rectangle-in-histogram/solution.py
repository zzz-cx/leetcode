# largest_rectangle_area 柱状图中最大的矩形（LeetCode 84）
# 思路：单调递增栈（存下标，栈内对应高度严格递增）
from typing import List


class Solution:
    def largest_rectangle_area(self, heights: List[int]) -> int:
        stack = []
        max_area = 0

        for i in range(len(heights) + 1):
            h = heights[i] if i < len(heights) else 0

            while stack and heights[stack[-1]] > h:
                mid = stack.pop()
                width = i if not stack else i - stack[-1] - 1
                max_area = max(max_area, heights[mid] * width)

            if i < len(heights):
                stack.append(i)
        return max_area


if __name__ == "__main__":
    sol = Solution()
    got = sol.largest_rectangle_area([2, 1, 5, 6, 2, 3])
    status = "PASS" if got == 10 else "FAIL"
    print(f"{status} | => {got} (expected 10)")
