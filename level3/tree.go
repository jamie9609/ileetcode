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

//二叉搜索树的后序遍历序列
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	f := len(postorder) -1
	for i := 0; i < len(postorder) -1 ; i ++{
		if postorder[i] >= postorder[len(postorder) - 1]{
			f = i
		}
		if i > f && postorder[i] < postorder[len(postorder) - 1]{
			return false
		}
	}
	return verifyPostorder(postorder[:f]) && verifyPostorder(postorder[f:len(postorder)-1])
}


