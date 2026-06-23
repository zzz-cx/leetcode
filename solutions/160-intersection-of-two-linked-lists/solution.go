package main

import "fmt"

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	out := []int{}
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路 哈希集合，如果节点不在哈希集合中，继续遍历到下一个节点；如果在，则后面的节点都在哈希集合中，从当前节点往后就都是相交节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for node := headA; node != nil; node = node.Next {
		vis[node] = true
	}
	for node := headB; node != nil; node = node.Next {
		if vis[node] {
			return node
		}
	}
	return nil
}

func main() {
	a := sliceToList([]int{4, 1})
	b := sliceToList([]int{5, 6, 1})
	shared := sliceToList([]int{8, 4, 5})
	pa, pb := a, b
	for pa.Next != nil {
		pa = pa.Next
	}
	pa.Next = shared
	for pb.Next != nil {
		pb = pb.Next
	}
	pb.Next = shared
	got := getIntersectionNode(a, b)
	status := "PASS"
	if got == nil || got.Val != 8 {
		status = "FAIL"
	}
	fmt.Printf("%s | intersection=%v\n", status, got)
}
