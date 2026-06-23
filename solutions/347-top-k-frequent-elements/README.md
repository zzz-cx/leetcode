# 前 K 个高频元素

> LeetCode 347 · [top-k-frequent-elements](https://leetcode.cn/problems/top-k-frequent-elements/)

## 题目

给定整数数组和整数 k，返回出现频率前 k 高的元素。

## 题解思路与解析

- topKFrequent 思路一：统计频次后按频次排序，取前 k 个。O(n log n)
- pair 堆元素：数字 + 出现次数
- minPairHeap 按 count 的小根堆。实现 heap.Interface 时方法名必须大写：Len/Less/Swap/Push/Pop
- topKFrequent2 思路二：小根堆维护 size=k。频次不够大的在堆顶被弹出。O(n log k)

## 解答

### Golang

```go
"container/heap"
	"sort"

// topKFrequent 思路一：统计频次后按频次排序，取前 k 个。O(n log n)
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	freqList := make([]int, 0, len(freq))
	for num := range freq {
		freqList = append(freqList, num)
	}
	sort.Slice(freqList, func(i, j int) bool {
		return freq[freqList[i]] > freq[freqList[j]]
	})
	return freqList[:k]
}

// pair 堆元素：数字 + 出现次数
type pair struct {
	num   int
	count int
}

// minPairHeap 按 count 的小根堆。实现 heap.Interface 时方法名必须大写：Len/Less/Swap/Push/Pop
type minPairHeap []pair

func (h minPairHeap) Len() int           { return len(h) } // 必须是 Len，不能写成 len
func (h minPairHeap) Less(i, j int) bool { return h[i].count < h[j].count } 
func (h minPairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPairHeap) Push(x any) { *h = append(*h, x.(pair)) }
func (h *minPairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// topKFrequent2 思路二：小根堆维护 size=k。频次不够大的在堆顶被弹出。O(n log k)
func topKFrequent2(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	h := &minPairHeap{}
	heap.Init(h)
	for num, cnt := range freq {
		heap.Push(h, pair{num, cnt})
		if len(*h) > k {
			heap.Pop(h) // 弹出频次最小的，堆里始终保留 k 个最高频
		}
	}
	res := make([]int, 0, k)
	for _, p := range *h {
		res = append(res, p.num)
	}
	return res
}
```

### Python

```python
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
```
