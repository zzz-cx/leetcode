package main

import (
	"fmt"
	"strings"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeFromLevel(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func parseTreeLevel(s string) []interface{} {
	s = strings.Trim(s, "[]")
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]interface{}, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "null" || p == "" {
			out = append(out, nil)
		} else {
			v, _ := strconv.Atoi(p)
			out = append(out, v)
		}
	}
	return out
}

// Pair 队列元素：节点 + 它在「满二叉树编号」下的下标
// 根编号 1，左子 2*idx，右子 2*idx+1（与层序下标一致，便于算层宽）
type Pair struct {
	node *TreeNode
	idx  int
}

// widthOfBinaryTree 二叉树最大宽度（LeetCode 662）
//
// 思路：BFS 按层遍历。每一层用队首、队尾节点的 idx 算宽度 = right - left + 1，再取全局最大。
// 层宽定义：该层最左、最右「仍存在节点」之间的间隔（含中间空格，按满二叉树位置算）。
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []Pair{{root, 1}} // 根下标从 1 开始
	ans := 0

	for len(queue) > 0 {
		size := len(queue) // 当前层节点数，固定本轮只处理这么多

		// 本层第一个、最后一个在队列里的下标 → 即该层最左、最右节点的编号
		left := queue[0].idx
		right := queue[size-1].idx
		ans = max(ans, right-left+1)

		// 弹出本层所有节点，把下一层子节点入队
		for i := 0; i < size; i++ {
			pair := queue[0]
			queue = queue[1:]

			node, idx := pair.node, pair.idx
			if node.Left != nil {
				queue = append(queue, Pair{node.Left, idx * 2})
			}
			if node.Right != nil {
				queue = append(queue, Pair{node.Right, idx*2 + 1})
			}
		}
	}

	return ans
}

func main() {
	root := buildTreeFromLevel(parseTreeLevel("1,2,3,4,5,6,7"))
	fmt.Printf("PASS | width=%d\n", widthOfBinaryTree(root))
}
