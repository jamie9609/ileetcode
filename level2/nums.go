package level2

import "strings"

func reverse(x int) int {

	result := 0
	if x == 0 {
		return 0
	}
	for x != 0 {

		tmp := x % 10
		if result >214748364 || (result == 214748364 && tmp >7) {
			return 0
		}
		if result < -214748364 || (result == -214748364 && tmp < -8) {
			return 0
		}
		result = tmp  + result * 10
		x = int(x /10)
	}
	return result
}


func isPalindrome(x int) bool {

	flag := x
	result := 0
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	for x > 0 {
		tmp := x % 10
		result = tmp  + result * 10
		x = int(x /10)
	}
	if result == flag {
		return true
	}
	return false

}


//数值的整数次方（快速幂，清晰图解）

func myPow(x float64, n int) float64 {
	if x == 0 || x == 1{
		return x
	}
	if  n < 0 {
		x = 1/x
		n = -n
	}
	var res = 1.0
	for n > 0 {
		// 位运算符

		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1

	}
	return res
}

//二维数组中的查找。直接查找，分支情况太多，还不如直接遍历矩阵，放弃。选择左下角标志数法，清晰图解

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0{
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	m := row - 1
	n := 0

	for n < col && m >= 0{
		if matrix[m][n] < target{
			n++
		} else if matrix[m][n] > target{
			m --
		} else {
			return true
		}

	}
	return false

}

//字符串的全排列。输入一个字符串，打印出该字符串中字符的所有排列。
//利用golang的map
func permutation(s string) []string {
	var res []string
	helper([]byte(s), 0, &res)
	return res
}

func helper(s []byte, start int, res *[]string) {
	if start == len(s) {
		*res = append(*res, string(s))
	}
	m := make(map[byte]int)
	for i := start ; i<len(s) ; i++{
		if _, ok := m[s[i]];ok{
			continue
		}
		s[i],s[start] = s[start], s[i]
		helper(s, start+1, res)
		s[i],s[start] = s[start], s[i]
		m[s[i]] = 1
	}
}


//和为s的两个数字。两种方法，方法1：hash
func twoSum(nums []int, target int) []int {
	res := map[int]int{}
	for key , val := range nums{
		res[val] = key
	}
	for i := 0; i <len(nums); i ++{
		if _, ok := res[target - nums[i]]; ok{
			return []int{target - nums[i], nums[i]}
		}
	}
	return []int{}
}
//和为s的两个数字。两种方法，方法2: 双指针碰撞法

func twoSum2(nums []int, target int) []int {
	left := 0
	right := len(nums)-1
	if nums==nil || len(nums)<2 {
		return nil
	}
	for left < right {
		if nums[left] + nums[right] == target{
			return []int{nums[left], nums[right]}
		}

		if nums[left] + nums[right] > target{
			right -= 1
		}
		if nums[left] + nums[right] < target{
			left += 1
		}
	}
	return nil
}

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0{
		return nil
	}
	top, bottom, left, right := 0, len(matrix) - 1, 0, len(matrix[0]) -1
	var res []int
	for top <= bottom {
		for i := left ; i <= right; i ++{
			res = append(res, matrix[top][i])
		}
		top += 1
		if top > bottom {
			break
		}
		for j := top; j <= bottom; j ++{
			res = append(res, matrix[j][right])
		}
		right -= 1
		if right < left {
			break
		}
		for i := right; i >= left; i --{
			res = append(res, matrix[bottom][i])
		}
		bottom -= 1
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i --{
			res = append(res, matrix[i][left])
		}
		left += 1
		if right < left {
			break
		}
	}
	return res
}

//计算  1～n整数中1出现的次数
func countDigitOne(n int) int {
	cur, high, low, digit, res := n%10, n/10, 0, 1, 0

	for cur != 0 || high != 0{
		if cur == 0 {
			res += high * digit
		} else if cur == 1{
			res += high * digit + low + 1
		} else {
			res += high * digit + digit
		}
		low += cur * digit
		cur = high % 10
		high  /=  10
		digit *= 10
	}
	return res
}


func NevWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	var res []int
	res = append(res, 1)
	res = append(res, 1)
	for i := 0; i < n ; i ++{

		res = append(res, (res[i]+res[i+1])%1000000007)
	}
	return res[n]
}


//动态规划，对于数组里求其中的最大子序和。如果前面的子序列>0为增益，需要加上。如果前面的子序列<0为损益，重新开始
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res, preSum, sum := nums[0], nums[0], nums[0]

	for i:= 0; i< len(nums); i ++{
		if preSum >=0 {
			sum = preSum + nums[i]
		} else {
			sum = nums[i]
		}
		preSum = sum
		if sum >= res {
			res = sum
		}
	}
	return res

}

func merge(A []int, m int, B []int, n int)  {
	pointA := m -1
	pointB := n -1
	cur := m + n - 1
	for pointA >= 0 && pointB >= 0 && cur > 0 {
		if A[pointA] > B[pointB] {
			A[cur] = A[pointA]
			cur --
			pointA --
		} else if A[pointA] <= B[pointB]{
			A[cur] = B[pointB]
			cur --
			pointB --
		}
	}
	for pointA  >= 0 && pointB < 0 && cur >=0{
		A[cur] = A[pointA]
		pointA --
		cur --
	}
	for pointB  >= 0 && pointA < 0 && cur >=0{
		A[cur] = B[pointB]
		pointB --
		cur --
	}
}

func MinArray(numbers []int) int {
	n := len(numbers) - 1
	if n == 0 {
		return numbers[n]
	}
	if n == 1 {
		if numbers[0] > numbers[1]{
			return numbers[1]
		}
		return numbers[0]
	}

	for n >= 1 {
		if numbers[n] >  numbers[n - 1] {
			n --
			continue
		} else {
			break
		}
	}
	return numbers[n]
}

func isPalindromeString(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	num := n-1

	for i := 0; i < n; i++{
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9'){
			if (s[num] >= 'a' && s[num] <= 'z') || (s[num] >= '0' && s[num] <= '9'){
				if s[i] != s[num]{
					return false
				}
			}else{
				num --
			}
			num--
		}
		if i >= num{
			return true
		}
	}
	return true
}



func isPalindrome(head *ListNode) bool {
	var res []int
	for head != nil {
		res  = append(res, head.Val)
		head = head.Next
	}
	return isPalindromeNum(res)
}

func isPalindromeNum(res []int) bool {
	if len(res) == 1 {
		return true
	}
	num := len(res) -1
	for i := 0; i <= num; i ++ {
		if res[i] != res[num] {
			return false
		} else {
			num --
		}
	}
	return true
}


