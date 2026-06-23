# 思路：手写大根堆 + 堆排序思想。建堆后每轮把堆顶（当前最大）换到「未排序区」末尾并下滤
from typing import List


class Solution:
    def find_kth_largest(self, nums: List[int], k: int) -> int:
        a = nums[:]
        n = len(a)
        self._build_max_heap(a, n)
        for i in range(n - 1, n - k, -1):
            a[0], a[i] = a[i], a[0]
            self._heapify_down(a, 0, i)
        return a[0]

    def _build_max_heap(self, a: List[int], n: int) -> None:
        for i in range(n // 2 - 1, -1, -1):
            self._heapify_down(a, i, n)

    def _heapify_down(self, a: List[int], i: int, n: int) -> None:
        while True:
            largest = i
            l, r = 2 * i + 1, 2 * i + 2
            if l < n and a[l] > a[largest]:
                largest = l
            if r < n and a[r] > a[largest]:
                largest = r
            if largest == i:
                return
            a[i], a[largest] = a[largest], a[i]
            i = largest


if __name__ == "__main__":
    sol = Solution()
    got = sol.find_kth_largest([3, 2, 1, 5, 6, 4], 2)
    status = "PASS" if got == 5 else "FAIL"
    print(f"{status} | => {got}")
