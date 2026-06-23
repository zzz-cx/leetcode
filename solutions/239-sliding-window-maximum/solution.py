from typing import List


class Solution:
    def max_sliding_window(self, nums: List[int], k: int) -> List[int]:
        ans = []
        n = len(nums)
        if n == 0:
            return ans
        deque = []
        for i in range(n):
            if i >= k and deque[0] <= i - k:
                deque = deque[1:]
            while deque and nums[deque[-1]] < nums[i]:
                deque.pop()
            deque.append(i)
            if i >= k - 1:
                ans.append(nums[deque[0]])
        return ans


if __name__ == "__main__":
    sol = Solution()
    got = sol.max_sliding_window([1, 3, -1, -3, 5, 3, 6, 7], 3)
    status = "PASS" if got == [3, 3, 5, 5, 6, 7] else "FAIL"
    print(f"{status} | => {got}")
