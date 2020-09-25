package level5

func myPow(x float64, n int) float64 {
	if n <0 {
		return mixPow(1.0/x, -n)
	}else if n == 0{
		return 1
	}else {
		return mixPow(x, n)
	}
}

func mixPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}else if n & 1 ==0 {
		return mixPow(x * x, n/2)
	}else {
		return x * mixPow(x, n-1)
	}
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums == nil {
		return []int{-1,-1}
	}
	leftNum := findFirstNums(nums, target)
	if leftNum == -1 {
		return []int{-1,-1}
	}
	rightNum := findSecondNums(nums, target)
	return []int{leftNum, rightNum}
}

func findFirstNums(nums []int, target int) int{
	left := 0
	right := len(nums) -1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			right = mid - 1
		}else if nums[mid] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	if left != len(nums) && nums[left] == target{
		return left
	}
	return -1
}

func findSecondNums(nums []int, target int) int{
	left := 0
	right := len(nums) -1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return right
}

func Search(nums []int, target int) int {
	res := -1
	if nums == nil || len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		if nums[0] == target{
			return 0
		} else {
			return res
		}
	}
	point := -1
	for i :=0; i < len(nums) - 1; i ++ {
		if nums[i + 1] < nums[i]{
			point = i
			break
		}
	}
	if point != -1 {
		if target >= nums[0] {
			res = findNums(nums[:point + 1], target)
		}else if target <= nums[len(nums) -1] {
			res = findNums(nums[point + 1:], target)
			if res != -1 {
				res = res + point + 1
			}
		} else {
			return -1
		}
	}else {
		res = findNums(nums, target)
	}
	return res
}

func findNums(nums []int, target int) int{
	left := 0
	right := len(nums) -1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	if left != len(nums) && nums[left] == target{
		return left
	}
	return -1
}

func findMin1(nums []int) int {
	var res int
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		res = nums[0]
		return res
	} else {
		res = nums[0]
	}

	for i :=0; i < len(nums) - 1; i ++ {
		if nums[i + 1] < nums[i] {
			res = nums[i + 1]
		}
	}
	return res
}

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[left] <= nums[right] {
			break
		}
		if nums[mid] > nums[right] {
			left = mid + 1
		}else {
			right = mid
		}
	}
	return nums[left]
}


func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || triangle == nil {
		return 0
	}
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i ++ {
		for j := 0; j <= i; j ++ {
			dp[i] = append(dp[i], 0)
		}
	}

	for i := len(triangle)-1; i >= 0; i -- {
		if i == len(triangle) - 1 {
			for j := 0; j < len(triangle[i]); j ++ {
				dp[i][j] = triangle[i][j]
			}
		} else {
			for j := 0; j < len(triangle[i]); j ++ {
				dp[i][j] = triangle[i][j] + minNum (dp[i + 1][j], dp[i + 1][j+1])
			}
		}
	}
	return dp[0][0]
}




func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	//先预设定，然后判断是否修改了预设定，如果没有修改，就说明无法取到。
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i ++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i ++ {
		for j := 0; j < len(coins); j ++ {
			if i >= coins[j] {
				dp[i] = minNum( dp[i - coins[j]] + 1, dp[i])
			}
		}
	}
	//说明无法通过最小的coins去实现
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func minNum(x, y int ) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	//l表示i、j的距离
	for l := 0; l < n; l++ {
		for i := 0; i + l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			//找到最大的子串
			if dp[i][j] > 0 && l + 1 > len(ans) {
				ans = s[i:i+l+1]
			}
		}
	}
	return ans
}

func maxSubArray(nums []int) int {
	dp := nums[0]
	maxSum := dp
	for i:=1; i<len(nums); i++{
		dp = maxNum(dp + nums[i],nums[i])
		if dp > maxSum{
			maxSum = dp
		}
	}
	return maxSum
}

func maxNum(a, b int) int {
	if  a > b {
		return a
	}else {
		return b
	}
}


func maxProduct(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxDP := make([]int, len(nums) + 1)
	minDP := make([]int, len(nums) + 1)

	maxDP[1] = nums[0]
	minDP[1] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i ++ {
		maxDP[i + 1] = maxNums(minDP[i] * nums[i], nums[i], maxDP[i] * nums[i])
		minDP[i + 1] = minNums(minDP[i] * nums[i], nums[i], maxDP[i] * nums[i])
		res = maxNum(res, maxDP[i + 1])
	}
	return  res
}

func maxNums(a, b, c int) int {
	res := a
	if  b > a {
		res = b
	}
	if res > c {
		return res
	} else {
		return c
	}
}

func minNums(a, b, c int) int {
	res := a
	if  b < a {
		res = b
	}
	if res > c {
		return c
	} else {
		return res
	}
}











