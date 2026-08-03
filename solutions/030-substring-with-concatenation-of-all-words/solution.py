from collections import Counter
from typing import List


class Solution:
    def find_substring(self, s: str, words: List[str]) -> List[int]:
        if not s or not words:
            return []

        word_len = len(words[0])
        word_count = len(words)
        target = Counter(words)
        result = []

        for offset in range(word_len):
            left = offset
            seen = Counter()
            used = 0

            for right in range(offset, len(s) - word_len + 1, word_len):
                word = s[right : right + word_len]
                seen[word] += 1
                used += 1

                while seen[word] > target.get(word, 0):
                    left_word = s[left : left + word_len]
                    seen[left_word] -= 1
                    if seen[left_word] == 0:
                        del seen[left_word]
                    left += word_len
                    used -= 1

                if used == word_count:
                    result.append(left)
                    left_word = s[left : left + word_len]
                    seen[left_word] -= 1
                    if seen[left_word] == 0:
                        del seen[left_word]
                    left += word_len
                    used -= 1

        return result


if __name__ == "__main__":
    tests = [
        ("barfoothefoobarman", ["foo", "bar"], [0, 9]),
        ("wordgoodgoodgoodbestword", ["word", "good", "best", "word"], []),
        ("barfoofoobarthefoobarman", ["bar", "foo", "the"], [6, 9, 12]),
    ]
    sol = Solution()
    for s, words, expected in tests:
        result = sorted(sol.find_substring(s, words))
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | s={s!r}, words={words} => {result} (expected {expected})")
