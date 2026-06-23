from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        if len(nums) == 0:
            return res

        path = []
        used = [False] * len(nums)

        def backtrack() -> None:
            if len(path) == len(nums):
                res.append(path[:])
                return
            for i in range(len(nums)):
                if used[i]:
                    continue
                used[i] = True
                path.append(nums[i])
                backtrack()
                path.pop()
                used[i] = False

        backtrack()
        return res


if __name__ == "__main__":
    sol = Solution()
    got = sol.permute([1, 2, 3])
    status = "PASS" if len(got) == 6 else "FAIL"
    print(f"{status} | nums=[1,2,3] => {len(got)} permutations")
