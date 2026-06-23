from typing import List


class Solution:
    def find_anagrams(self, s: str, p: str) -> List[int]:
        ans = []
        s_len, p_len = len(s), len(p)
        if s_len < p_len:
            return ans
        s_count = [0] * 26
        p_count = [0] * 26
        for i, ch in enumerate(p):
            s_count[ord(s[i]) - ord("a")] += 1
            p_count[ord(ch) - ord("a")] += 1
        if s_count == p_count:
            ans.append(0)

        for i, ch in enumerate(s[: s_len - p_len]):
            s_count[ord(ch) - ord("a")] -= 1
            s_count[ord(s[i + p_len]) - ord("a")] += 1
            if s_count == p_count:
                ans.append(i + 1)
        return ans


if __name__ == "__main__":
    sol = Solution()
    got = sol.find_anagrams("cbaebabacd", "abc")
    status = "PASS" if got == [0, 6] else "FAIL"
    print(f"{status} | => {got}")
