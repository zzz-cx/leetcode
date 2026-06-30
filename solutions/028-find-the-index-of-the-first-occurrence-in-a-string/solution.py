class Solution:
    def str_str(self, haystack: str, needle: str) -> int:
        n, m = len(haystack), len(needle)
        if m == 0:
            return 0

        for i in range(n - m + 1):
            if haystack[i : i + m] == needle:
                return i
        return -1


if __name__ == "__main__":
    tests = [
        ("sadbutsad", "sad", 0),
        ("leetcode", "leeto", -1),
        ("hello", "ll", 2),
        ("a", "a", 0),
    ]
    sol = Solution()
    for haystack, needle, expected in tests:
        result = sol.str_str(haystack, needle)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | haystack={haystack!r}, needle={needle!r} => {result} (expected {expected})")
