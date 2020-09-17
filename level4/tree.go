package level4

import (
	"fmt"
	"github.com/pingcap/parser/ast"
	"math"
	"math/big"
)

type TreeNode struct {
	Val  	int
	Left    *TreeNode
	Right   *TreeNode
}

//层序遍历，使用队列。
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var tmpQueue []*TreeNode
		var tmpRes []int
		for j := 0; j < len(queue); j ++{
			node := queue[j]
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = tmpQueue
		res = append(res, tmpRes)
	}
	return res
}


//N叉树的前序遍历
/*
type Node struct {
	Val 		int
	Children 	[]*Node
}*/
var res []int

func preorder(root *Node) []int{
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *Node){
	if root != nil {
		res = append(res, root.Val)
		for _, index := range root.Children {
			dfs(index)
		}
	}
}

/*func preorder(root *Node) []int {
	var res []int
	stack := []*Node{root}

	for len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			if len(root.Children) == 0 {
				break
			}
			for i := len(root.Children) - 1; i >0 ; i++{
				stack = append(stack, root.Children[i])
			}
		}
		//出栈
		root = stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]

	}
	return res
}
*/


//迭代实现前序遍历。
func preorder(root *Node) []int {
	var res []int
	stack := []*Node{root}

	if root == nil {
		return nil
	}

	for len(stack) >0	{
		cur := stack[len(stack) - 1]
		stack = stack[ :len(stack) -1]
		res = append(res, cur.Val)
		for i := len(cur.Children) - 1; i >=0; i-- {
			stack = append(stack, cur.Children[i])
		}
	}
	return res
}

//非迭代实现中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) != 0 || root != nil {
		if root == nil {
			tmpNum := stack[len(stack) - 1]
			res = append(res, tmpNum.Val)
			stack = stack[ : len(stack) - 1]
			root = tmpNum.Right
		}else {
			stack = append(stack, root)
			root = root.Left
		}
	}
	return res
}

//二叉树的后序遍历

func postorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)
	var stack []*TreeNode
	var visitRoot *TreeNode
	var node *TreeNode
	node = root
	if root == nil {
		return res
	}
	for len(stack) > 0 || node != nil {
		for  node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		tmpNode := stack[len(stack) - 1]
		if tmpNode.Right == nil || tmpNode.Right == visitRoot{
			res = append(res, tmpNode.Val)
			stack = stack[ :len(stack) -1]
			visitRoot, tmpNode = tmpNode, nil
		} else {
			node = tmpNode.Right
		}
	}
	return res
}

//二叉树前序遍历

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var node *TreeNode
	node = root

	if root == nil {
		return res
	}
	stack = append(stack, node)
	for len(stack) > 0 || node != nil {

		res = append(res, node.Val)
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if len(stack) == 0 {
			return res
		}else {
			node = stack[len(stack)-1]
		}
	}
	return res
}

//n叉树的层序遍历

/*type Node struct {
	Val   			int
	Children  		[]*Node
}*/

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		var tmpQueue []*Node
		var tmpRes []int
		for j := 0; j < len(queue); j ++{
			node := queue[j]
			tmpRes = append(tmpRes, node.Val)
			if len(node.Children) != 0 {
				for i := 0; i < len(node.Children); i ++ {
					tmpQueue = append(tmpQueue, node.Children[i])
				}
			}
		}
		queue = tmpQueue
		res = append(res, tmpRes)
	}
	return res
}

//N叉树的后序遍历
func postorder(root *Node) []int {
	var res []int
	var stack []*Node
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		tmpNode := stack[len(stack) -1]
		stack = stack[ :len(stack) -1 ]
		//在切片开头添加元素
		res = append([]int{tmpNode.Val}, res...)
		for i := 0; i < len(tmpNode.Children); i++{
			stack = append(stack, tmpNode.Children[i])
		}
	}
	return res
}

//合并二叉树

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	queue := []*TreeNode{t1, t2}
	for len(queue) > 0{
		n1 := queue[0]
		n2 := queue[1]

		n1.Val = n1.Val + n2.Val

		if n1.Left != nil && n2.Left != nil {
			queue = append(queue, n1.Left, n2.Left)
		} else if n1.Left == nil && n2.Left != nil {
			n1.Left = n2.Left
		}
		if n1.Right != nil && n2.Right != nil {
			queue = append(queue, n1.Right, n2.Right)
		}else if n1.Right == nil && n2.Right != nil{
			n1.Right = n2.Right
		}

		queue = queue[2:]
	}
	return t1
}

//对称的二叉树

func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	var node *TreeNode
	node = root

	if node == nil || node.Right== nil && node.Left == nil {
		return true
	}

	if node.Left == nil && node.Right != nil || node.Left != nil && node.Right == nil {
		return false
	}

	if node.Left.Val != node.Right.Val {
		return false
	}
	l := node.Left
	r := node.Right

	queue = append(queue, l.Left, r.Right, l.Right, r.Left)
	for len(queue) > 0{
		q := queue[0]
		p := queue[1]
		queue = queue[2:]
		if q == nil && p == nil {
			continue
		}
		if q != nil && p == nil || q ==nil && p != nil {
			return false
		}else if q.Val != p.Val {
			return false
		}
		queue = append(queue, q.Left, p.Right, q.Right, p.Left)
	}
	return true
}

//二叉树的最小深度

func minDepth(root *TreeNode) int {
	var node *TreeNode
	node = root
	queue := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	queue = append(queue, node)
	level:=1

	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i ++ {
			tmpNode := queue[i]
			if tmpNode.Right == nil && tmpNode.Left == nil {
				return level
			}
			if tmpNode.Left != nil {
				tmpQueue = append(tmpQueue, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				tmpQueue = append(tmpQueue, tmpNode.Right)
			}
		}
		queue = tmpQueue
		level ++
	}
	return level
}


//n叉树的最大深度

func maxDepth(root *Node) int {
	queue := make([]*Node, 0)
	if root == nil {
		return 0
	}
	level := 1
	queue = append(queue, root)

	for len(queue) > 0 {
		tmpQueue := make([]*Node, 0)

		for i := 0; i < len(queue); i ++ {
			tmpNode := queue[i]

			if len(tmpNode.Children) == 0{
				continue
			}

			for j := 0; j <len(tmpNode.Children); j ++ {
				tmpQueue = append(tmpQueue, tmpNode.Children[j])
			}

		}
		level ++
		queue = tmpQueue
	}
	return level - 1
}

//最小高度树 ：给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	root := &TreeNode{nums[mid], nil,nil}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

//二叉搜索树的范围和
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if root.Val > R{
		return rangeSumBST(root.Left, L, R)
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	return root.Val + rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R)
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	res := 0
	var node *TreeNode
	node = root
	var stack []*TreeNode
	stack = append(stack, node)
	for node != nil || len(stack) >0 {
		tmpNode := stack[len(stack) -1]
		stack = stack[ : len(stack) -1]
		if tmpNode != nil {
			if tmpNode.Val >= L && tmpNode.Val <= R{
				res += tmpNode.Val
			}
			if tmpNode.Val > L {
				stack = append(stack, tmpNode.Left)
			}
			if tmpNode.Val < R {
				stack = append(stack, tmpNode.Right)
			}
		}
	}
	return res
}


func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil || s != nil && t == nil{
		return false
	}
	if helper(s, t) {
		return true
	} else {
		return isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}

}

//判断俩树是否相等
func helper(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil || s != nil && t == nil{
		return false
	}
	if s.Val == t.Val {
		return helper(s.Right, t.Right) && helper(s.Left, t.Left)
	}
	return false
}
//二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return nil
	}
	for root != nil {
		if root.Val == val {
			break
		}
		if root.Val < val {
			root = root.Right
		} else if root.Val > val {
			root = root.Left
		}
	}
	return root
}

//将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	node := &TreeNode{nums[mid], nil ,nil }

	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}

//路径总和：给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
//用层序遍历，新增一个数组用来记录和
func hasPathSum(root *TreeNode, sum int) bool {
	var sumQueue []int
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	sumQueue  = append(sumQueue, root.Val)
	for len(queue) > 0 {

		tmpNode := queue[0]
		tmpSumNode := sumQueue[0]
		queue = queue[1:]
		sumQueue = sumQueue[1:]
		if tmpNode.Left != nil {
			queue = append(queue, tmpNode.Left)
			sumQueue = append(sumQueue, tmpNode.Left.Val + tmpSumNode)
		}
		if tmpNode.Right != nil {
			queue = append(queue, tmpNode.Right)
			sumQueue = append(sumQueue, tmpNode.Right.Val + tmpSumNode)
		}
		if tmpNode.Right == nil && tmpNode.Left == nil {
			if tmpSumNode == sum {
				return true
			}
		}
	}
	return false
}

//递增顺序查找树

func increasingBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	res := dsf(root)

	newRoot := &TreeNode{res[0],nil,nil}

	resRoot := newRoot
	for _, val := range res[1:]{
		next := &TreeNode{val,nil,nil}

		newRoot.Right = next
		newRoot = newRoot.Right
	}
	return resRoot
}

func dsf(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	resLeft := dsf(root.Left)
	resRight := dsf(root.Right)
	return append(append(resLeft, root.Val), resRight...)
}
// ... 作为语法糖，会把元素分解后添加进去

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	num := root.Val
	return helpers(root, num)
}

func helpers(root *TreeNode, num int) bool {
	if root == nil  {
		return true
	}
	if root.Val != num {
		return false
	}
	return helpers(root.Right, num) && helpers(root.Left, num)
}


func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	res = append(res, []int{root.Val})
	for len(queue) > 0 {
		tmpRes := make([]int, 0)
		tmpQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i ++ {
			tmpNode := queue[i]

			if tmpNode.Left != nil {
				tmpRes = append(tmpRes, tmpNode.Left.Val)
				tmpQueue = append(tmpQueue, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				tmpRes = append(tmpRes, tmpNode.Right.Val)
				tmpQueue = append(tmpQueue, tmpNode.Right)
			}
		}
		queue = tmpQueue
		if len(tmpRes) == 0 {
			continue
		}
		resQ := make([][]int,0)
		resQ = append(resQ, tmpRes)
		for _, key := range res {
			resQ = append(resQ, key)
		}
		res = resQ
	}
	return res
}
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	helper3(root, "",&res)
	return res
}

func helper3(root *TreeNode, str string, res *[]string)  {
	if root == nil {
		return
	}
	str = fmt.Sprintf("%s%d->", str, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, str[:len(str) -2])
	}
	helper3(root.Left, str, res)
	helper3(root.Right, str, res)
}

//修剪二叉搜索树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := make([]int, 0)
	num := 0
	helper4(root, 0, &res)
	for _, val := range res{
		num += val
	}
	return num
}

func helper4(root *TreeNode, num int, res *[]int)  {
	if root == nil {
		return
	}
	num = num * 2 + root.Val
	if root.Left == nil && root.Right == nil {
		*res = append(*res, num)
	}
	helper4(root.Left, num, res)
	helper4(root.Right, num, res)
}

//二叉树的层平均值

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := make([]float64, 0)
	resTmp := make([][]int, 0)
	resTmp = append(resTmp, []int{root.Val})

	for len(queue) >0 {
		var tmpRes []int
		var tmpQue []*TreeNode
		for i := 0; i < len(queue); i ++ {
			tmpNode := queue[i]
			if tmpNode.Left != nil {
				tmpRes = append(tmpRes, tmpNode.Left.Val)
				tmpQue = append(tmpQue, tmpNode.Left)
			}
			if tmpNode.Right != nil{
				tmpRes = append(tmpRes, tmpNode.Right.Val)
				tmpQue =  append(tmpQue, tmpNode.Right)
			}
		}
		queue = tmpQue
		resTmp = append(resTmp, tmpRes)
	}
	if resTmp == nil {
		return nil
	}
	for _, val := range resTmp {
		sum := 0
		num := 0
		if val == nil {
			continue
		}
		for _ , item := range val {
			sum += item
			num ++
		}
		res = append(res, float64(sum)/float64(num))
	}
	return res
}


func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var res1, res2 []int
	if root1 == nil && root2 != nil || root1 !=  nil && root2 == nil {
		return false
	}

	helper5(root1, &res1)
	helper5(root2, &res2)


	if len(res1) != len(res2) {
		return false
	}
	for key, _ := range res1{
		if res1[key] != res2[key]{
			return false
		}
	}
	return true
}

func helper5(root *TreeNode, path *[]int )  {
	if root == nil {
		return
	}

	if root.Right == nil && root.Left == nil {
		*path = append(*path, root.Val)
	}
	helper5 (root.Left, path)
	helper5 (root.Right, path)
}


func convertBST(root *TreeNode) *TreeNode {
	help6(root, 0)
	return root

}

//反向中序遍历迭代, 返回累加值
func help6(root *TreeNode, preVal int ) int {
	if root == nil {
		return preVal
	}
	rightSum := help6(root.Right, preVal)
	root.Val = rightSum + root.Val
	leftSum := help6(root.Left, root.Val)
	//左子树的累加值即整树的累加值
	return leftSum
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, max := helper7(root)
	return max
}

// 第一个返回值表示该树最大深度，第二个返回值表示该树最长路径。深度是节点的个数，直径是节点之间的横线
func helper7(root *TreeNode) (maxDeep int , maxPath int){
	if root == nil {
		return 0,0
	}
	maxDeepL, maxPathL := helper7(root.Left)
	maxDeepR, maxPathR := helper7(root.Right)

	return max(maxDeepL, maxDeepR) + 1, max(maxDeepL + maxDeepR + 1, max(maxPathL, maxPathR))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//验证搜索二叉树

func isValidBST1(root *TreeNode) bool {
	return helper8(root, math.MinInt64, math.MaxInt64)
}

func helper8(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper8(root.Left, lower, root.Val) && helper8(root.Right, root.Val, upper)
}

//中序遍历
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	var res []int
	var node *TreeNode
	node = root
	if root == nil {
		return true
	}

	for len(stack) >0  || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if node ==  nil {
			tmpNode := stack[len(stack) - 1]
			stack = stack[:len(stack) -1]
			if len(res) > 0 && tmpNode.Val <= res[len(res)-1] {
				return false
			}
			res = append(res, tmpNode.Val)
			node = tmpNode.Right
		}

	}
	return true
}

func numTrees(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2 ; i <= n; i ++ {
		for j := 1 ; j <= i ; j ++ {
			res[i] += res[j - 1] * res[i -j]
		}
	}
	return res[n]
}

func rightSideView(root *TreeNode) []int {
	var queue []*TreeNode
	var res [][]int
	var resAns []int
 	if root == nil {
		return nil
	}

	queue = append(queue, root)
	res = append(res, []int{root.Val})

	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		tmpRes := make([]int, 0)
		for i := 0 ; i < len(queue); i ++ {
			tmpNode := queue[i]
			if tmpNode.Left != nil {
				tmpQueue = append(tmpQueue, tmpNode.Left)
				tmpRes = append(tmpRes, tmpNode.Left.Val)

			}
			if tmpNode.Right != nil {
				tmpQueue = append(tmpQueue, tmpNode.Right)
				tmpRes = append(tmpRes, tmpNode.Right.Val)
			}
		}
		queue = tmpQueue
		res = append(res, tmpRes)
	}

	if res == nil {
		return nil
	}

	for _, val := range res {
		if len(val) == 0 {
			continue
		}
		resAns = append(resAns, val[len(val) - 1])

	}
	return resAns
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var res [][]int
	flag := 1
	if root == nil {
		return nil
	}

	queue = append(queue, root)
	res = append(res, []int{root.Val})

	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		tmpRes := make([]int, 0)
		flag ++
		for i := 0 ; i < len(queue); i ++ {
			tmpNode := queue[i]
			if tmpNode.Left != nil {
				tmpQueue = append(tmpQueue, tmpNode.Left)
				if flag % 2 == 0 {
					tmpRes = append([]int{tmpNode.Left.Val}, tmpRes ...)
				}else {
					tmpRes = append(tmpRes, tmpNode.Left.Val)
				}
			}
			if tmpNode.Right != nil {
				tmpQueue = append(tmpQueue, tmpNode.Right)
				if flag % 2 == 0 {
					tmpRes = append([]int{tmpNode.Right.Val}, tmpRes ...)
				}else {
					tmpRes = append(tmpRes, tmpNode.Right.Val)
				}
			}
		}
		if len(tmpQueue) == 0 {
			break
		}
		queue = tmpQueue
		res = append(res, tmpRes)
	}
	return res
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || preorder == nil{
		return nil
	}
	root := &TreeNode{preorder[0],nil,nil}
	i := 0
	for ; i < len(inorder); i ++ {
		if inorder[i] == preorder[0]{
			break
		}
	}
	root.Left = buildTree(preorder[1 : len(inorder[:i]) + 1], inorder[ : i])
	root.Right = buildTree(preorder[len(inorder[:i]) + 1 : ], inorder [i + 1 : ])
	return root
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	tmp := root.Right

	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}

var p, d int

func isCousins(root *TreeNode, x int, y int) bool {
	dfsHelp2(root, 0, 0, x)
	p1, d1 := p, d
	dfsHelp2(root, 0,0, y)
	return p1 != p && d1 == d
}

func dfsHelp2(root *TreeNode, parVal, depth, x int) {
	if root.Val == x {
		p, d = parVal, depth
	}
	if root.Left != nil {
		dfsHelp2(root.Left, root.Val, depth +1 , x)
	}
	if root.Right != nil {
		dfsHelp2(root.Right, root.Val, depth + 1, x)
	}
}
type Node struct {
	Val 	int
	Left 	*Node
	Right	*Node
	Next 	*Node
}

func connect(root *Node) *Node {
	if root == nil{
		return nil
	}

	node := root
	var nextLevel *Node
	for node.Left != nil {
		nextLevel = node.Left
		for node != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
			node = node.Next
		}
		node = nextLevel
	}
	return root
}

func findSecondMinimumValue(root *TreeNode) int {
	return helper10(root, root.Val)
}

func helper10(root *TreeNode, s int) int  {
	if root == nil {
		return -1
	}

	if root.Val > s {
		return root.Val
	}

	sLeft := helper10(root.Left, s)
	sRight := helper10(root.Right, s)

	if sLeft == -1 {
		return sRight
	}

	if sRight == -1 {
		return sLeft
	}
	return min1(sLeft, sRight)
}
func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

