# 思路：贪心思想，记录每个字符最后出现的位置，然后遍历字符串，记录当前字符最后出现的位置，就是end
from typing import List


class Solution:
    def partition_labels(self, s: str) -> List[int]:
        result = []
        last_index = {}
        for i, ch in enumerate(s):
            last_index[ch] = i
        start = 0
        end = 0
        for i, ch in enumerate(s):
            if last_index[ch] > end:
                end = last_index[ch]
            if i == end:
                result.append(end - start + 1)
                start = end + 1
        return result


if __name__ == "__main__":
    sol = Solution()
    got = sol.partition_labels("ababcbacadefegdehijhklij")
    status = "PASS" if got == [9, 7, 8] else "FAIL"
    print(f"{status} | => {got}")
