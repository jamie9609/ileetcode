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


func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := dsfHelper(root)
	return sum
}

//返回的是左叶子节点之和
func dsfHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + dsfHelper(root.Right)
	}
	lSum := dsfHelper(root.Left)
	rSum := dsfHelper(root.Right)
	return lSum + rSum
}

//最小祖父节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pVal, qVal := p.Val, q.Val
	node := root
	for node != nil {
		parentVal := node.Val
		switch  {
		case pVal > parentVal && qVal > parentVal:
			node = node.Right
		case pVal < parentVal && qVal < parentVal:
			node = node.Left
		default:
			return node
		}
	}
	return nil
}

//递归的方式去解决

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1,n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}

	for i := start; i <=end; i ++ {
		leftTree := helper(start, i - 1)
		rightTree := helper(i + 1, end)

		for _, left := range leftTree{
			for _, right := range rightTree {
				currTree := &TreeNode{i, nil , nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}

var res [][]int
func pathSum(root *TreeNode, sum int) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	helper2(root, sum, []int{})
	return res
}

func helper2(root *TreeNode, sum int, path []int) {
	if root == nil {
		return
	}

	path = append(path ,root.Val)
	sum = sum - root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0{
			// slice是一个指向底层的数组的指针结构体
			// 因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
			// 所以这里需要 copy arr 到 tmp，再添加进 ret，防止 arr 底层数据修改带来的错误
			tmp := make([]int, len(path))
			copy(tmp ,path)
			res = append(res, tmp)
		}
	}
	helper2(root.Left, sum , path)
	helper2(root.Right, sum , path)
}


func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	arr := make([]int, 0)
	find(root, &arr)
	return arr[k-1]
}

func find(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	find(root.Left, arr)
	*arr = append(*arr, root.Val)
	find(root.Right, arr)
}