class Solution:
    def length_of_longest_substring(self, s: str) -> int:
        n = len(s)
        if n == 0:
            return 0
        max_len = 1
        left = 0
        for right in range(1, n):
            for i in range(left, right):
                if s[i] == s[right]:
                    left = i + 1
                    break
            cur_len = right - left + 1
            if cur_len > max_len:
                max_len = cur_len
        return max_len


if __name__ == "__main__":
    tests = [("abcabcbb", 3), ("bbbbb", 1), ("pwwkew", 3)]
    sol = Solution()
    for s, expected in tests:
        got = sol.length_of_longest_substring(s)
        status = "PASS" if got == expected else "FAIL"
        print(f"{status} | s={s!r} => {got} (expected {expected})")
