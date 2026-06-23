# top_k_frequent 思路一：统计频次后按频次排序，取前 k 个。O(n log n)
import heapq
from typing import List


class Solution:
    def top_k_frequent(self, nums: List[int], k: int) -> List[int]:
        freq = {}
        for num in nums:
            freq[num] = freq.get(num, 0) + 1
        freq_list = list(freq.keys())
        freq_list.sort(key=lambda x: freq[x], reverse=True)
        return freq_list[:k]

    def top_k_frequent2(self, nums: List[int], k: int) -> List[int]:
        freq = {}
        for num in nums:
            freq[num] = freq.get(num, 0) + 1
        h = []
        for num, cnt in freq.items():
            heapq.heappush(h, (cnt, num))
            if len(h) > k:
                heapq.heappop(h)
        return [p[1] for p in h]


if __name__ == "__main__":
    sol = Solution()
    got = sol.top_k_frequent([1, 1, 1, 2, 2, 3], 2)
    status = "PASS" if sorted(got) == [1, 2] else "FAIL"
    print(f"{status} | => {got}")
