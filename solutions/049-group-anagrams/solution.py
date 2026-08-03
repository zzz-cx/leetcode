from collections import defaultdict
from typing import List


class Solution:
    def group_anagrams(self, strs: List[str]) -> List[List[str]]:
        groups: dict[tuple[int, ...], list[str]] = defaultdict(list)

        for s in strs:
            count = [0] * 26
            for ch in s:
                count[ord(ch) - ord("a")] += 1
            groups[tuple(count)].append(s)

        return list(groups.values())


def normalize(groups: List[List[str]]) -> List[List[str]]:
    return sorted(sorted(g) for g in groups)


if __name__ == "__main__":
    tests = [
        (
            ["eat", "tea", "tan", "ate", "nat", "bat"],
            [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
        ),
        ([""], [[""]]),
        (["a"], [["a"]]),
    ]
    sol = Solution()
    for strs, expected in tests:
        result = sol.group_anagrams(strs)
        ok = normalize(result) == normalize(expected)
        status = "PASS" if ok else "FAIL"
        print(f"{status} | strs={strs!r} => {result}")
