package _17

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	mt := new(TreeNode)

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	mt.Val = t1.Val + t2.Val

	mt.Left = mergeTrees(t1.Left, t2.Left)
	mt.Right = mergeTrees(t1.Right, t2.Right)

	return mt
}
