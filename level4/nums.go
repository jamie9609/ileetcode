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

//股票买卖，贪心算法
func maxProfit(prices []int) int {
	maxVal := 0
	if len(prices) == 0 {
		return maxVal
	}

	minPrice := prices[0]
	for i := 1; i < len(prices); i ++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else {
			v := prices[i] - minPrice
			maxVal = max(maxVal, v)
		}
	}
	return maxVal
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, 0)

	for i := 0; i < m; i ++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := 0; i < n; i ++ {
		dp[0][i] = 1
	}

	for i := 1; i < m ; i ++ {
		for j := 1; j < n ; j ++ {
			dp[i][j] = dp[i-1][j] + dp[i][j -1]
		}
	}
	return dp[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0{
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < m ; i ++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m ; i ++ {
		dp[i][0] = dp[i -1][0] + grid[i][0]
	}
	for j := 1; j < n; j ++ {
		dp[0][j] = dp[0][j - 1] + grid[0][j]
	}
	for i := 1; i < m ; i ++ {
		for j := 1; j < n ; j ++ {
			dp[i][j] = min(dp[i][j-1],dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(x,y int) int {
	if x<y {
		return x
	}
	return y
}


func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0{
		return -1
	}

	left := 0
	right := len(nums) -1

	for left <= right {
		mid := (left + right ) / 2 + 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target{
			right = mid - 1
		}
		if nums[mid] < target{
			left = mid + 1
		}
	}
	return left
}

func mySqrt(x int) int {
	if x <= 1 {
		return 1
	}
	left := 0
	right := x
	for right - left > 1 {
		mid := (left + right) / 2
		if mid * mid == x {
			return mid
		}else if mid * mid > x {
			right = mid
		}else {
			left = mid
		}
	}
	return left
}


func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) -1

	for left <= right {
		if numbers[left] + numbers[right] == target {
			return []int{left+1, right+1}
		}else if numbers[left] + numbers[right] > target {
			right --
		}else {
			left ++
		}
	}
	return nil
}

func missingNumber(nums []int) int {
	if nums == nil {
		return -1
	}
	left1 := 0
	right1 := len(nums) - 1

	for left1 <= right1 {
		mid := (left1 + right1) /2
		if nums[mid] == mid {
			left1 = mid + 1
		}else {
			right1 = mid -1
		}
	}
	return left1
}