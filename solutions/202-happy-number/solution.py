class Solution:
    def is_happy(self, n: int) -> bool:
        seen: set[int] = set()

        while n != 1 and n not in seen:
            seen.add(n)
            n = sum(int(d) * int(d) for d in str(n))

        return n == 1


if __name__ == "__main__":
    tests = [
        (19, True),
        (2, False),
        (1, True),
        (7, True),  # 7 → 49 → 97 → 130 → 10 → 1
    ]
    sol = Solution()
    for n, expected in tests:
        result = sol.is_happy(n)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | n={n} => {result} (expected {expected})")
