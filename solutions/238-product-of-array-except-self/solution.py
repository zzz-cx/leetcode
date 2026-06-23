from typing import List


class Solution:
    def product_except_self(self, nums: List[int]) -> List[int]:
        n = len(nums)
        L = [0] * n
        R = [0] * n
        answer = [0] * n
        L[0] = 1
        R[n - 1] = 1
        for i in range(1, n):
            L[i] = L[i - 1] * nums[i - 1]
        for i in range(n - 2, -1, -1):
            R[i] = R[i + 1] * nums[i + 1]
        for i in range(n):
            answer[i] = L[i] * R[i]
        return answer


if __name__ == "__main__":
    sol = Solution()
    got = sol.product_except_self([1, 2, 3, 4])
    status = "PASS" if got == [24, 12, 8, 6] else "FAIL"
    print(f"{status} | => {got}")
