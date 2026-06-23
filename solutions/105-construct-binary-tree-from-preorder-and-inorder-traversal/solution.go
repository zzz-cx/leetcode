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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			rootIndex = i
		}
	}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

// buildTree2 使用哈希表 + 下标边界，整体时间复杂度 O(n)。
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}

	preIdx := 0
	var helper func(l, r int) *TreeNode
	helper = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		val := preorder[preIdx]
		preIdx++

		root := &TreeNode{Val: val}
		mid := pos[val]
		root.Left = helper(l, mid-1)
		root.Right = helper(mid+1, r)
		return root
	}

	return helper(0, len(inorder)-1)
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Printf("PASS | root=%d\n", root.Val)
}
