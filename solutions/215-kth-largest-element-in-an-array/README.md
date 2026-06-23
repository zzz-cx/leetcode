# 第K大的元素

> LeetCode 215 · [kth-largest-element-in-an-array](https://leetcode.cn/problems/kth-largest-element-in-an-array/)

## 题目

在未排序数组中找到第 k 大的元素。

## 题解思路与解析

- 思路：手写大根堆 + 堆排序思想。建堆后每轮把堆顶（当前最大）换到「未排序区」末尾并下滤，
- 执行 k-1 轮后，堆顶即为原数组第 k 大。时间 O(n + k log n)，空间 O(n)（拷贝避免改原切片）
- 每轮把当前最大值沉到 a[i]，共 k-1 轮后 a[0] 为第 k 大
- buildMaxHeap 自底向上建堆，O(n)
- heapifyDown 大根堆下滤：保证以 i 为根的子树在 [0, n) 内满足堆性质

## 解答

### Golang

```go
// 思路：手写大根堆 + 堆排序思想。建堆后每轮把堆顶（当前最大）换到「未排序区」末尾并下滤，
// 执行 k-1 轮后，堆顶即为原数组第 k 大。时间 O(n + k log n)，空间 O(n)（拷贝避免改原切片）

func findKthLargest(nums []int, k int) int {
	a := make([]int, len(nums))
	copy(a, nums)
	n := len(a)
	buildMaxHeap(a, n)
	// 每轮把当前最大值沉到 a[i]，共 k-1 轮后 a[0] 为第 k 大
	for i := n - 1; i > n-k; i-- {
		a[0], a[i] = a[i], a[0]
		heapifyDown(a, 0, i) // 只维护 [0, i) 这一段堆
	}
	return a[0]
}

// buildMaxHeap 自底向上建堆，O(n)
func buildMaxHeap(a []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(a, i, n)
	}
}

// heapifyDown 大根堆下滤：保证以 i 为根的子树在 [0, n) 内满足堆性质
func heapifyDown(a []int, i, n int) {
	for {
		largest := i
		l, r := 2*i+1, 2*i+2
		if l < n && a[l] > a[largest] {
			largest = l
		}
		if r < n && a[r] > a[largest] {
			largest = r
		}
		if largest == i {
			return
		}
		a[i], a[largest] = a[largest], a[i]
		i = largest
	}
}
```

### Python

```python
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
```
