# 思路：中心扩展。每个位置作为回文中心，向两侧扩展；分奇长（i,i）和偶长（i,i+1）两种情况
class Solution:
    def longest_palindrome(self, s: str) -> str:
        if len(s) == 0:
            return ""
        start, max_len = 0, 1

        def expand(left: int, right: int) -> None:
            nonlocal start, max_len
            while left >= 0 and right < len(s) and s[left] == s[right]:
                left -= 1
                right += 1
            l = right - left - 1
            if l > max_len:
                max_len = l
                start = left + 1

        for i in range(len(s)):
            expand(i, i)
            expand(i, i + 1)
        return s[start : start + max_len]


if __name__ == "__main__":
    tests = [("babad", "bab"), ("cbbd", "bb")]
    sol = Solution()
    for s, expected in tests:
        got = sol.longest_palindrome(s)
        status = "PASS" if got == expected else "FAIL"
        print(f"{status} | s={s!r} => {got!r} (expected {expected!r})")
