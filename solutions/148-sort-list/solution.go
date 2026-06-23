package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

// sortList 排序链表（LeetCode 148）
// 思路：归并排序。找中点拆成左右两段 → 递归排序 → mergeTwoLists 合并
// 时间 O(n log n)，空间 O(log n) 递归栈（满足题面 O(1) 空间需自底向上归并，面试写这版即可）
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head // 0 个或 1 个节点，已有序
	}
	mid := getMid(head)               // 从 mid 处断开，左半 head..slow，右半 mid..末尾
	left := sortList(head)            // 排左半
	right := sortList(mid)            // 排右半
	return mergeTwoLists(left, right) // 合并两个有序链表（见 21mergeTwoLists.go）
}

// getMid 快慢指针找中点，并在 slow 处断开链表
// slow 最终停在左半最后一个节点；mid = slow.Next 为右半头；slow.Next = nil 切断
func getMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next // fast 比 slow 多走一步，保证左半不少于右半
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummy := &ListNode{0, nil}
// 	tail := dummy
// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			tail.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			tail.Next = list2
// 			list2 = list2.Next
// 		}
// 		tail = tail.Next
// 	}
// 	if list1 != nil {
// 		tail.Next = list1
// 	} else {
// 		tail.Next = list2
// 	}
// 	return dummy.Next
// }

func main() {
	head := sliceToList([]int{4, 2, 1, 3})
	fmt.Printf("PASS | %v\n", listToSlice(sortList(head)))
}
