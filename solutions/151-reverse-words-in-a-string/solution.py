class Solution:
    def reverse_words(self, s: str) -> str:
        """方法二：原地三次反转，O(1) 额外空间。"""
        return self.reverse_words_inplace(s)

    def reverse_words_split(self, s: str) -> str:
        """方法一：拆分 + 双指针反转 + 拼接。"""
        words = s.split()
        left, right = 0, len(words) - 1
        while left < right:
            words[left], words[right] = words[right], words[left]
            left += 1
            right -= 1
        return " ".join(words)

    def reverse_words_inplace(self, s: str) -> str:
        """方法二：压缩空格 → 整体反转 → 逐词反转。"""
        chars = self._trim_spaces(list(s))
        if not chars:
            return ""

        self._reverse(chars, 0, len(chars) - 1)

        i, n = 0, len(chars)
        while i < n:
            j = i
            while j < n and chars[j] != " ":
                j += 1
            self._reverse(chars, i, j - 1)
            i = j + 1

        return "".join(chars)

    @staticmethod
    def _reverse(chars: list[str], left: int, right: int) -> None:
        while left < right:
            chars[left], chars[right] = chars[right], chars[left]
            left += 1
            right -= 1

    @staticmethod
    def _trim_spaces(chars: list[str]) -> list[str]:
        slow, n = 0, len(chars)
        fast = 0
        while fast < n:
            if chars[fast] == " ":
                fast += 1
                continue
            if slow > 0:
                chars[slow] = " "
                slow += 1
            while fast < n and chars[fast] != " ":
                chars[slow] = chars[fast]
                slow += 1
                fast += 1
        return chars[:slow]


if __name__ == "__main__":
    tests = [
        ("the sky is blue", "blue is sky the"),
        ("  hello world  ", "world hello"),
        ("a good   example", "example good a"),
        (" ", ""),
    ]
    sol = Solution()
    for name, fn in [
        ("split", sol.reverse_words_split),
        ("inplace", sol.reverse_words_inplace),
    ]:
        print(f"--- {name} ---")
        for s, expected in tests:
            result = fn(s)
            status = "PASS" if result == expected else "FAIL"
            print(f"{status} | s={s!r} => {result!r} (expected {expected!r})")
