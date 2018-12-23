package check_complete_binary_tree


/**
 * Definition for a binary tree node.
*/
type TreeNode struct {
	Val int
	Left *TreeNode
    Right *TreeNode
}

func numerOfNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + numerOfNodes(root.Left) + numerOfNodes(root.Right)
}

func isCompleteTreeWithIndex(root *TreeNode, index int, numOfNodes int) bool {
	if root == nil {
		return true
	}

	if index >= numOfNodes {
		return false
	}

	return isCompleteTreeWithIndex(root.Left, 2 * index + 1, numOfNodes) && isCompleteTreeWithIndex(root.Right, 2 * index + 2, numOfNodes)
}

func IsCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isCompleteTreeWithIndex(root, 0, numerOfNodes(root))
}