class Solution:
    def is_isomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        s_to_t: dict[str, str] = {}
        t_to_s: dict[str, str] = {}

        for a, b in zip(s, t):
            if a in s_to_t:
                if s_to_t[a] != b:
                    return False
            else:
                if b in t_to_s:
                    return False
                s_to_t[a] = b
                t_to_s[b] = a

        return True


if __name__ == "__main__":
    tests = [
        ("egg", "add", True),
        ("f11", "b23", False),
        ("paper", "title", True),
        ("badc", "baba", False),
    ]
    sol = Solution()
    for s, t, expected in tests:
        result = sol.is_isomorphic(s, t)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | s={s!r}, t={t!r} => {result} (expected {expected})")
