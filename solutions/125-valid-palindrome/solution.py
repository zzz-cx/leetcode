class Solution:
    def is_palindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1
        while left < right:
            while left < right and not s[left].isalnum():
                left += 1
            while left < right and not s[right].isalnum():
                right -= 1
            if s[left].lower() != s[right].lower():
                return False
            left += 1
            right -= 1
        return True


if __name__ == "__main__":
    tests = [
        ("A man, a plan, a canal: Panama", True),
        ("race a car", False),
        (" ", True),
        ("0P", False),
    ]
    sol = Solution()
    for s, expected in tests:
        result = sol.is_palindrome(s)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | s={s!r} => {result} (expected {expected})")
