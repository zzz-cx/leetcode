class Solution:
    def is_valid(self, s: str) -> bool:
        pairs = {")": "(", "]": "[", "}": "{"}
        stack: list[str] = []

        for ch in s:
            if ch in pairs:
                if not stack or stack[-1] != pairs[ch]:
                    return False
                stack.pop()
            else:
                stack.append(ch)

        return not stack


if __name__ == "__main__":
    tests = [
        ("()", True),
        ("()[]{}", True),
        ("(]", False),
        ("([])", True),
        ("([)]", False),
        ("", True),
        ("(", False),
    ]
    sol = Solution()
    for s, expected in tests:
        result = sol.is_valid(s)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | s={s!r} => {result} (expected {expected})")
