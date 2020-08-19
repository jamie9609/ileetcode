package level2

import "git.xiaojukeji.com/dop/dmp-enter/dao/company"

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


func sortedListToBST(head *ListNode) *TreeNode {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return toBST(res)
}
func toBST(res []int) *TreeNode{
	l := len(res)
	if l == 0 {
		return nil
	}
	mid := l >> 1
	return &TreeNode{res[mid], toBST(res[:mid]), toBST(res[mid+1:])}

}

func toBST(vs []int) *TreeNode {
	l := len(vs)
	if l == 0 {
		return nil
	}
	mid := l >> 1
	return &TreeNode{
		Val:   vs[mid],
		Left:  toBST(vs[:mid]),
		Right: toBST(vs[mid+1:]),
	}
}


func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var pre, cur *ListNode = nil, slow
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	for pre != nil {
		if pre.Val != head.Val{
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}