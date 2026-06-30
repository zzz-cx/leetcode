from typing import List


class Solution:
    def longest_common_prefix(self, strs: List[str]) -> str:
        if not strs:
            return ""

        for i in range(len(strs[0])):
            ch = strs[0][i]
            for s in strs[1:]:
                if i >= len(s) or s[i] != ch:
                    return strs[0][:i]

        return strs[0]


if __name__ == "__main__":
    tests = [
        (["flower", "flow", "flight"], "fl"),
        (["dog", "racecar", "car"], ""),
        (["ab", "a"], "a"),
        ([""], ""),
    ]
    sol = Solution()
    for strs, expected in tests:
        result = sol.longest_common_prefix(strs)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | strs={strs} => {result!r} (expected {expected!r})")
