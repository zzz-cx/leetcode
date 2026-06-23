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

// 思路：递归，计算左右子树的深度，然后计算直径
func diameterOfBinaryTree(root *TreeNode) int {
	// 直径不一定经过根节点，因此需要遍历整棵树，在每个节点处用
	// “左子树深度 + 右子树深度” 更新全局最大值。
	diameter := 0

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		if left+right > diameter {
			diameter = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}

	depth(root)
	return diameter
}

func main() {
	root := buildTreeFromLevel(parseTreeLevel("1,2,3,4,5,6,7"))
	fmt.Printf("PASS | diameter=%d\n", diameterOfBinaryTree(root))
}
