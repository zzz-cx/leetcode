class Solution:
    def word_pattern(self, pattern: str, s: str) -> bool:
        words = s.split()
        if len(pattern) != len(words):
            return False

        char_to_word: dict[str, str] = {}
        word_to_char: dict[str, str] = {}

        for ch, word in zip(pattern, words):
            if ch in char_to_word:
                if char_to_word[ch] != word:
                    return False
            else:
                if word in word_to_char:
                    return False
                char_to_word[ch] = word
                word_to_char[word] = ch

        return True


if __name__ == "__main__":
    tests = [
        ("abba", "dog cat cat dog", True),
        ("abba", "dog cat cat fish", False),
        ("aaaa", "dog cat cat dog", False),
        ("abba", "dog dog dog dog", False),
        ("a", "dog", True),
    ]
    sol = Solution()
    for pattern, s, expected in tests:
        result = sol.word_pattern(pattern, s)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | pattern={pattern!r}, s={s!r} => {result} (expected {expected})")
