package main

import (
	"fmt"
	"sort"
	"container/heap"
)

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

func main() {
	got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Printf("PASS | %v\n", got)
}
