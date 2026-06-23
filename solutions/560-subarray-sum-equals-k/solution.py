from typing import List


class Solution:
    def subarray_sum(self, nums: List[int], k: int) -> int:
        count = 0
        for start in range(len(nums)):
            s = 0
            for end in range(start, -1, -1):
                s += nums[end]
                if s == k:
                    count += 1
        return count

    def subarray_sum2(self, nums: List[int], k: int) -> int:
        count = 0
        s = 0
        prefix_sum = {0: 1}
        for num in nums:
            s += num
            count += prefix_sum.get(s - k, 0)
            prefix_sum[s] = prefix_sum.get(s, 0) + 1
        return count


if __name__ == "__main__":
    sol = Solution()
    got = sol.subarray_sum([1, 1, 1], 2)
    status = "PASS" if got == 2 else "FAIL"
    print(f"{status} | => {got}")
