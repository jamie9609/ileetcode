package level1

//输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

//方法一：后序遍历（DFS）
//树的后序遍历 / 深度优先搜索往往利用 递归 或 栈 实现，本文使用递归实现。
//关键点： 此树的深度和其左（右）子树的深度之间的关系。显然，此树的深度 等于 左子树的深度 与 右子树的深度 中的 最大值 +1+1 。
//
//
//作者：jyd
//链接：https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/solution/mian-shi-ti-55-i-er-cha-shu-de-shen-du-xian-xu-bia/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
type TreeNode struct {
	Val 	int
	Left  	*TreeNode
	Right	*TreeNode
}

func max(a int , b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left 	:= maxDepth(root.Left)
	right 	:= maxDepth(root.Right)
	return max(left, right) + 1
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a

}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(maxDepth(root.Left) - maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}





func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return  nil
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Right)
	invertTree(root.Left)
	return root
}

