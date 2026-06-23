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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root // 空节点返回 nil；命中 p/q 则把“命中的节点”向上返回
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 左子树里是否能找到 p 或 q（或它们的 LCA）
	right := lowestCommonAncestor(root.Right, p, q) // 右子树里是否能找到 p 或 q（或它们的 LCA）
	if left != nil && right != nil {
		return root // p、q 分居两侧（或一侧命中 p，另一侧命中 q），当前 root 即最近公共祖先
	}
	if left != nil {
		return left // 只在左边找到：把找到的节点/祖先继续向上冒泡
	}
	return right // 只在右边找到：把找到的节点/祖先继续向上冒泡
}

func findNode(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == v {
		return root
	}
	if n := findNode(root.Left, v); n != nil {
		return n
	}
	return findNode(root.Right, v)
}

func main() {
	root := buildTreeFromLevel(parseTreeLevel("3,5,1,6,2,0,8,null,null,7,4"))
	var p, q *TreeNode
	var walk func(*TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Val == 5 {
			p = n
		}
		if n.Val == 1 {
			q = n
		}
		walk(n.Left)
		walk(n.Right)
	}
	walk(root)
	fmt.Printf("PASS | lca=%d\n", lowestCommonAncestor(root, p, q).Val)
}
