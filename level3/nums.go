package level3

import (
	"sort"
	"strconv"
	"strings"
)

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。

var result [][]int

func permute(nums []int) [][]int {
	var path []int
	var used = make([]bool, len(nums))

	result = [][]int{}
	backtrack1(nums, path, used)
	return result
}

func backtrack1(nums, path []int, used []bool) {
	if len(nums) == len(path){
		tmp := make([]int, len(nums))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := 0; i< len(nums); i ++{
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			backtrack(nums, path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}


//给定一个 没有重复 数字的序列，返回其所有不重复的全排列。需要剪枝

var result [][]int

func permuteUnique(nums []int) [][]int {
	var path []int
	var used = make([]bool, len(nums))

	result = [][]int{}

	//此处需要给排列一个顺序
	sort.Ints(nums)
	backtrack(nums, path, used)
	return result
}

func backtrack(nums, path []int, used []bool) {
	if len(nums) == len(path){
		tmp := make([]int, len(nums))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := 0; i< len(nums); i ++{
		if !used[i] {
			//此处剪枝
			if i > 0 && nums[i] == nums[i -1]&& used[i - 1] == false{
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(nums, path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}


func minNumber(nums []int) string {
	n  := len(nums)
	if n == 0 {
		return ""
	}
	var arr []string
	for _,v := range nums {
		arr = append(arr, strconv.Itoa(v))
	}
	sort.Slice(arr, func(i,j int) bool {
		return arr[i] + arr[j] < arr[j] + arr[i]
	})
	res := strings.Join(arr, "")
	return res
}


//全排列，N皇后问题

var res [][]string

// 回溯核心
// board： 棋盘
// path：选择列表

func backtrack(board[][]bool, path [][]byte)  {
	if len(path) == len(board){
		v := make([]string, len(path))
		for k, bs := range path{
			v[k] = string(bs)
		}
		res = append(res, v)
	}

	for i := 0; i< len(board); i ++{
		//需要剪枝
		if !isValid(board, len(path), i){
			continue
		}
		bs := printLine(len(board))
		bs[i] = 'Q'
		board[len(path)][i] = true

		path = append(path, bs)
		backtrack(board, path)

		path = path[:len(path)-1]
		board[len(path)][i] = false
	}
}

// 打印每行默认情况,都是'.'
func printLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}
//剪枝条件：不合法
func isValid(board [][]bool, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == true {
			return false
		}
	}
	for j := 0; j < len(board); j++ {
		if board[row][j] == true {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	// 清空变量
	res = [][]string{}
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	backtrack(board, [][]byte{})
	return res
}

/*func backtrack(board [][]bool, path [][]byte) {
	// 结束条件
	if len(path) == len(board) {
		t := make([]string, len(path))
		for k, bs := range path {
			t[k] = string(bs)
		}
		result = append(result, t)
	}

	for i := 0; i < len(board); i++ {
		// 不合法情况,剔除(剪枝)
		if !isValid(board, len(path), i) {
			continue
		}
		// 打印默认位置
		bs := printLine(len(board))
		// 放置皇后
		bs[i] = 'Q'
		// 放入棋盘
		board[len(path)][i] = true
		// 加入选择路径
		path = append(path, bs)
		// 进行下一次决策
		backtrack(board, path)
		// 撤销选择
		path = path[:len(path)-1]
		board[len(path)][i] = false
	}
}*/


func lastRemaining(n int, m int) int {
	res := 0
	if n == 1 {
		return 0
	}
	for i := 2; i <= n ; i ++{
		res = (res + m ) % i
	}
	return res
}


func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sum := 0
	for i := 0; i < len(prices) -1 ; i ++{
		if prices[i +1 ] > prices[i]{
			sum += prices[i +1 ] - prices[i]
		}else {
			continue
		}
	}
	return sum
}

func canJump(nums []int) bool {
	l := len(nums) - 1
	for i := l-1; i >=0 ; i-- {
		if nums[i]+i >= l {
			l = i
		}
	}
	return l <= 0
}
func canJump(nums []int) bool {
	max_index := 0  //能达到的最大索引
	for i,v := range  nums{
		if  max_index>=i&& i+v>max_index{
			max_index = i+v
		}
	}
	return max_index >= len(nums)-1
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) -1
	for right >= left{
		mid := (left + right) >>1
		if nums[mid] <= target {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	rightSum := left

	left = 0
	right = len(nums) -1

	for right >= left{
		mid := (left + right) >>1
		if nums[mid] < target{
			left = mid +1
		} else {
			right = mid -1
		}
	}
	leftSum := right
	return rightSum - leftSum -1
}


func smallestK(arr []int, k int) []int {
	if len (arr) == 0 {
		return nil
	}
	if len (arr) == 1 {
		if k == 0{
			return nil
		}else if k == 1{
			return arr
		}
	}
	Quick(arr)
	return arr[:k]
}

func Quick(arr []int) {
	if len(arr) <=1 {
		return
	}
	left := 0
	right := len(arr) - 1
	key := arr[left]
	for left < right {
		if left < right && key <= arr[right] {
			right --
		}
		arr[left] = arr[right]
		if left < right &&  key > arr[left]{
			left ++
		}
		arr[right] = arr[left]
	}
	arr[left] = key
	Quick(arr[:left])
	Quick(arr[left + 1 :])
}


