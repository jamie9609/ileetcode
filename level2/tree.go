package level2

//深度优先，计算二叉树的最大深度，可以用递归
//广度优先，计算二叉树最大深度

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res []*TreeNode
	res = append(res ,root)
	depth := 0
	for len(res) > 0 {
		size := len(res)
		for  i := 0; i < size; i++ {
			s := res[0]
			res = res[1:]
			if s.Left != nil{
				res = append(res, s.Left)
			}
			if s.Right != nil{
				res = append(res, s.Right)
			}
		}
		depth ++
	}
	return depth
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}


func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

