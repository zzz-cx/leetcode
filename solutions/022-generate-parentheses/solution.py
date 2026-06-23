# 思路：回溯算法，每次添加一个左括号或右括号，如果左括号数量小于n，则添加左括号，如果右括号数量小于左括号数量，则添加右括号
from typing import List


class Solution:
    def generate_parenthesis(self, n: int) -> List[str]:
        res = []
        if n == 0:
            return res

        def backtrack(cur: str, left: int, right: int) -> None:
            if len(cur) == 2 * n:
                res.append(cur)
                return
            if left < n:
                backtrack(cur + "(", left + 1, right)
            if right < left:
                backtrack(cur + ")", left, right + 1)

        backtrack("", 0, 0)
        return res


if __name__ == "__main__":
    sol = Solution()
    got = sol.generate_parenthesis(3)
    status = "PASS" if len(got) == 5 else "FAIL"
    print(f"{status} | n=3 => {len(got)} combinations")
