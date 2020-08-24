package level3

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSubPathHelper(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSubPathHelper(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return isSubPathHelper(head.Next, root.Right) || isSubPathHelper(head.Next, root.Left)
}

