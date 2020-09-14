package level4

//二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

/*func pathSum1(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var path []int
	var res [][]int
	var resQ [][]int
	dfsHelper(root, path, &res)
	for _, val := range res{
		sumTmp := 0
		for _, item := range val {
			sumTmp += item
		}
		if sumTmp == sum {
			resQ = append(resQ, val)
		}
	}
	return resQ
}

func dfsHelper(root *TreeNode,  path []int, res *[][]int)  {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}
	dfsHelper(root.Left, path, res)
	dfsHelper(root.Right, path, res)
}*/

var  resAnswer[][]int
func pathSum(root *TreeNode, sum int) [][]int {
	resAnswer = make([][]int, 0)
	dfsHelper(root, sum, []int{})
	return resAnswer
}

func dfsHelper(root *TreeNode, sum int, stack []int) {
	if root == nil {
		return
	}

	stack = append(stack ,root.Val)
	if root.Right == nil && root.Left == nil {
		if sum == root.Val {
			resTmp := make([]int, len(stack))
			copy(resTmp, stack)
			resAnswer = append(resAnswer, resTmp)
		}
	}
	dfsHelper(root.Left, sum - root.Val, stack)
	dfsHelper(root.Right, sum - root.Val, stack)
}


