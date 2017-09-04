package tree

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func DFSPreOrder(root *TreeNode) {
	if root != nil {
		log.Printf("%d ", root.Val)
		DFSPreOrder(root.Left)
		DFSPreOrder(root.Right)
	}
}

// 中序遍历
func DFSInOrder(root *TreeNode) {
	if root != nil {
		DFSInOrder(root.Left)
		log.Printf("%d ", root.Val)
		DFSInOrder(root.Right)
	}
}

// 后序遍历
func DFSPostOrder(root *TreeNode) {
	if root != nil {
		DFSPostOrder(root.Left)
		DFSPostOrder(root.Right)
		log.Printf("%d ", root.Val)
	}
}

// 广度优先遍历
func BFS(root *TreeNode) {
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		tmp := []*TreeNode{}
		for _, node := range stack {
			log.Printf("%d ", node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		stack = tmp
	}
}
