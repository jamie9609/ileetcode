package level4

func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxNum(nums[0], nums[1])

	for i := 2; i < len(nums); i ++ {
		dp[i] = maxNum(dp[i-2] + nums[i], dp[i - 1])
	}
	return dp[len(nums) -1]

}



func robHelp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxNum(nums[0], nums[1])

	for i := 2 ; i < len(nums); i ++{
		dp[i] = maxNum(dp[i-2] + nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return maxNum(robHelp(nums[:len(nums) -1]), robHelp(nums[1:]))
}

func rob(root *TreeNode) int {
	val := dfsHelp(root)
	return max(val[0], val[1])
}

func dfsHelp(node *TreeNode) []int{
	if node == nil {
		return []int{0,0}
	}
	l := dfsHelp(node.Left)
	r := dfsHelp(node.Right)

	selected  	:= node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max (r[0] , r[1])
	return []int{selected, notSelected}
}


func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func levelOrder9(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([]int, 0)
	res = append(res, root.Val)
	for len(queue) >0 {
		tmpQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i ++ {
			tmpNode := queue[i]
			if tmpNode.Left != nil {
				tmpQueue = append(tmpQueue, tmpNode.Left)
				res = append(res, tmpNode.Left.Val)
			}
			if tmpNode.Right != nil {
				tmpQueue = append(tmpQueue, tmpNode.Right)
				res = append(res, tmpNode.Right.Val)
			}
		}
		if len(tmpQueue) == 0 {
			break
		}
		queue = tmpQueue
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	res = append(res, []int{root.Val})
	flag := 1
	for len(queue) >0 {
		tmpQueue := make([]*TreeNode, 0)
		tmpRes := make([]int, 0)
		flag ++
		for i := 0; i < len(queue); i ++ {
			tmpNode := queue[i]
			if tmpNode.Left != nil {
				tmpQueue = append(tmpQueue, tmpNode.Left)
				if flag % 2 == 0 {
					tmpRes = append([]int{tmpNode.Left.Val}, tmpRes...)
				} else {
					tmpRes = append(tmpRes, tmpNode.Left.Val)
				}
			}
			if tmpNode.Right != nil {
				tmpQueue = append(tmpQueue, tmpNode.Right)
				if flag % 2 == 0 {
					tmpRes = append([]int{tmpNode.Right.Val}, tmpRes...)
				} else {
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


func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return A != nil && B != nil && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))

}
//先判断两个树是否相等。再判断是否是子树
func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val{
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}


//路径总和，双递归
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return allPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

//从头到尾，路径和为sum的条数
func allPath(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	cnt := 0

	if root.Val == sum {
		cnt ++
	}
	cnt += allPath(root.Left, sum - root.Val)
	cnt += allPath(root.Right, sum - root.Val)
	return cnt
}