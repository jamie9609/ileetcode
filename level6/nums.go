package level6

import (
	"math"
	"sort"
	"strings"
)

func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B)== 0 {
		return 0
	}
	dp := make([][]int, len(A) + 1)
	for i := 0; i <len(dp); i ++ {
		dp[i] = make([]int, len(B)+1)
	}

	maxL := math.MinInt64
	for i := 1; i <= len(A); i ++ {
		for j := 1; j <=len(B); j ++ {
			if A[i - 1] == B[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxL {
				maxL = dp[i][j]
			}
		}
	}
	return maxL
}

func isValid(s string) bool {
	var stack = make([]byte, 0)
	var bracketsMap = map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	for i := 0; i < len(s); i ++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{'{
			stack = append(stack, s[i])
		} else {
			if val, ok := bracketsMap[s[i]]; ok {
				if val == stack[len(stack) - 1] && len(stack) >= 1 {
					stack = stack[:len(stack) - 1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}


func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs == nil {
		return ""
	}
	flag := 0
	ct := 0
	res := ""
	for flag >= 0 {
		for i := 0; i < len(strs) ; i ++ {
			if flag > len(strs[i]) -1 {
				ct = 1
				break
			}
			tmp := strs[0][flag]

			if strs[i][flag] == tmp {
				ct = 2
			} else {
				ct = 1
				break
			}
		}

		if ct == 1 {
			break
		} else if ct == 2 {
			res = res + string(strs[0][flag])
		}

		flag ++
	}
	return res
}


type KthLargest struct {
	size int	//堆的容量，不包括数组第一个元素
	data []int
	count int 	//当前元素数量

}


/*func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{}
	kth.data = []int{0}
	kth.size = k
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}
*/

func (this *KthLargest) Add(val int) int {
	if this.count < this.size -1 {
		this.data = append(this.data, val)
		this.count += 1
	} else if this.count == this.size-1 {
		this.data = append(this.data, val)
		this.count += 1

		//第一次填满k个数，使堆有序，后续add只需在data[1]覆盖放入新值。
		n:=len(this.data)-1
		for i := n/2; i >= 1; i-- {
			heapify(this.data, i)
		}
	}else {
		if val > this.data[1] {
			this.data[1]=val
			heapify(this.data, 1)
		}
	}
	return this.data[1]
}

//heapify 从给定的i向下堆化为小顶堆
func heapify(a []int, i int) {
	n := len(a) - 1
	for {
		minPos := i
		if i*2 <=n && a[i*2] < a[minPos] {
			minPos=i*2
		}
		if i*2+1 <= n && a[i*2+1] < a[minPos] {
			minPos=i*2+1
		}
		if minPos == i {
			break
		}
		a[minPos], a[i] = a[i], a[minPos]
		i = minPos
	}
}

func getLeastNumbers(arr []int, k int) []int {
	if k==0{
		return []int{}
	}

	heap := arr[:k:k]
	for i:=(k-1)/ 2; i >=0; i--{
		maxHeap1(heap ,i,  k-1)
	}

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			maxHeap1(heap, 0, k-1)
		}
	}

	return heap
}

//构建大顶堆的第一步，构建交换大顶堆的基础
func maxHeap1(arr []int, start int, end int)  {
	root := start
	for true {
		child := root * 2 + 1

		if child > end {
			break
		}

		for child + 1 <= end && arr[child + 1] > arr[child] {
			child ++
		}
		if arr[child] > arr[root] {
			arr[child], arr[root] = arr[root], arr[child]
			root = child
		} else {
			break
		}
	}
	return
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || matrix == nil {
		return 0
	}
	res := make([]int, 0)
	for i := 0 ; i < len(matrix) ; i ++ {
		for j := 0; j < len(matrix[i]); j ++ {
			res = append(res, matrix[i][j])
		}
	}

	heap := res[:k]

	for i := (len(heap) - 1) / 2; i >= 0; i -- {
		maxHeap1(heap, i, k - 1)
	}

	for i := k ; i < len(res) ; i ++ {
		if res[i] < heap[0] {
			heap[0] = res[i]
			maxHeap1(heap, 0, k - 1)
		}
	}
	return  heap[0]
}


//https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
// 输入: [7,1,5,3,6,4]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
// 输入: [7,6,4,3,1]
// 输出: 0
// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

func maxProfit(prices []int) int {
	if len(prices) == 0 || prices == nil || len(prices) == 1{
		return 0
	}
	dp := make([]int, len(prices) + 1)
	min := prices[0]

	for i := 1; i < len(prices) ; i ++ {
		min = minNum(min , prices[i])
		dp[i] = maxNum(dp[i - 1], prices[i] - min)
	}
	return dp[len(prices) -1]
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func minNum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func romanToInt(s string) int {
	dictMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 1 {
		return dictMap[s[0]]
	}
	res := dictMap[s[0]]
	for i := 0; i < len(s) - 1; i ++ {
		if dictMap[s[i + 1]] <= dictMap[s[i]] {
			res = res + dictMap[s[i + 1]]
		} else {
			res = res - dictMap[s[i]] + dictMap[s[i + 1]] - dictMap[s[i]]
		}
	}
	return res
}



func threeSum(nums []int) [][]int {
	if len(nums) < 3 || nums == nil {
		return nil
	}
	res := make([][]int,0)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums) - 1; i ++ {
		if i != 0 && nums[i] == nums[i - 1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left] + nums[right] + nums[i] == 0 {
				res = append(res,[]int{nums[i], nums[left], nums[right]})
				//为了去重，需要把左指针和右指针都移动
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left ++
				right --
			} else if nums[left] + nums[right] + nums[i] > 0 {
				right --
			} else {
				left ++
			}
		}
	}
	return res
}


func plusOne(digits []int) []int {
	for i := len(digits) -1; i >= 0 ;i -- {
		if digits[i] != 9{
			digits[i] ++
			break
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}


func myAtoi(str string) int {
	//去掉收尾空格
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v -'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		// 数值最大检测
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign * result
}



type Twitter struct {
	Tweets 		[]int
	UserTweets 	map[int][]int
	Follows 	map[int][]int
	IsFollowMy 	map[int]bool
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	var Tweets []int
	var UserTweets = make(map[int][]int)
	var Follows = make(map[int][]int)
	var IsFollowMy = make(map[int]bool)

	t := Twitter{
		Tweets: Tweets,
		UserTweets: UserTweets,
		Follows: Follows,
		IsFollowMy: IsFollowMy,
	}
	return t
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.Tweets = append(this.Tweets, tweetId)
	this.UserTweets[userId] = append(this.UserTweets[userId], tweetId)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	fs := this.Follows[userId]
	var allTweets []int

	for _, val := range fs {
		allTweets = append(allTweets, this.UserTweets[val]...)
	}

	if ! this.IsFollowMy[userId] {
		allTweets = append(allTweets, this.UserTweets[userId]...)
	}

	var sortTweets []int

	aTLen := len(this.Tweets)
	s := 0
	// 按照发的推特顺序进行倒序排序
	for i:=aTLen-1; i>=0; i-- {
		if s >= 10 {
			break
		}
		for _,n := range allTweets {

			if this.Tweets[i] == n && s < 10{
				s++
				sortTweets = append(sortTweets,n)
			}
		}
	}
	return sortTweets
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {

	if followerId == followeeId {
		this.IsFollowMy[followerId] = true
	}

	// 下面是判断这人是否关注了，如果已经关注了，那么就不再关注了
	var isFed bool

	for _,v := range this.Follows[followerId] {
		if v == followeeId {
			isFed = true
		}
	}
	if !isFed {
		this.Follows[followerId] = append(this.Follows[followerId],followeeId)
	}

}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followeeId == followerId {
		this.IsFollowMy[followerId] = false
	}
	var temp []int
	for _,v := range this.Follows[followerId] {
		if v != followeeId {
			temp = append(temp,v)
		}
	}
	this.Follows[followerId] = temp
}


// 最接近原点的 K 个点
func kClosest(points [][]int, K int) [][]int {
	if len(points) < K || points == nil {
		return nil
	}

	res := make([]int, 0)
	resAll := make([][]int, 0)
	pointMap := make(map[int][]int)
	for key, val := range points {
		tmpKey := points[key][0] * points[key][0] + points[key][1] * points[key][1]
		pointMap[tmpKey] = val
		res = append(res, tmpKey)
	}
	for i := (K - 1) /2 ; i >= 0; i -- {
		heapSort(res, i, K - 1)
	}

	for i := K ; i < len(res); i ++ {
		if res[i] < res[0] {
			res[0] = res[i]
			heapSort(res, 0, K - 1)
		}
	}

	for i := 0; i <= K -1 ; i ++ {
		tmpPoint := pointMap[res[i]]
		resAll = append(resAll, tmpPoint)
	}
	return resAll
}

//大顶堆
func heapSort(res []int, left int, right int)  {

	root := left
	for true {
		child := root * 2 + 1

		if child > right{
			break
		}

		if child + 1 < right && res[child + 1] > res[child] {
			child ++
		}

		if res[root] < res[child]{
			res[root], res[child] = res[child], res[root]
			root = child
		} else {
			break
		}
	}
	return
}




func tribonacci(n int) int {
	if n ==  0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	res := make([]int, n + 1)
	res[0] = 0
	res[1] = 1
	res[2] = 1
	for i := 3; i <= n; i ++ {
		res[i] = res[i - 1] + res[i - 2] + res[i - 3]
	}
	return res[n]
}


func pathSum(root *TreeNode, sum int) int {
	var m = 0
	solve(root, sum, &m)
	return m
}

func solve(root *TreeNode, sum int, m *int)  {
	if root == nil {
		return
	}
	helpSum(root, sum, m)
	solve(root.Left, sum, m)
	solve(root.Right, sum, m)
}

func helpSum(root *TreeNode, target int, m *int)  {
	if root == nil {
		return
	}
	target -= root.Val
	if target == 0 {
		*m ++
	}
	helpSum(root.Left, target , m)
	helpSum(root.Right, target , m)
}


func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	res := make([]int, 0)
	for i := 0; i <= k ; i ++ {
		res = append(res, longer * i + shorter * (k - i) )
	}
	return res
}

//汉诺塔，经典问题

func hanota(A []int, B []int, C []int) []int {

	if A == nil {
		return nil
	}
	move(len(A), &A, &B, &C)
	return C
}

func move(n int, A *[]int, B *[]int, C *[]int)  {
	if n == 0 {
		return
	}
	if n == 1 {
		*C = append(*C, (*A)[len(*A) - 1])
		*A = (*A)[ : len(*A) - 1]
	}
	if n > 1 {
		move(n - 1, A, C, B)
		move(1, A, B, C)
		move(n - 1, B, A, C)
	}
}


//最长同值路径
func longestUnivaluePath(root *TreeNode) int {
	num := 0
	help(root, &num)
	return num
}
func help(root *TreeNode, num *int) int {
	if root == nil {
		return 0
	}
	l := help(root.Left, num)
	r := help(root.Right, num)
	var al, ar int
	if root.Left != nil && root.Left.Val == root.Val {
		al = l + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		ar += r + 1
	}
	*num = int(math.Max(float64(*num), float64(ar+al)))
	return int(math.Max(float64(al), float64(ar)))
}


func countNegatives(grid [][]int) int {
	var res int
	for i :=0; i < len(grid) ; i ++ {
		c := len(grid[i])
		if grid[i][c-1] >= 0 {
			continue
		}
		if grid[i][0] < 0 {
			res += c
			continue
		}
		r := binary(grid[i])
		res += c - r
	}
	return res
}

func binary(nums []int) int {

	left, right := 0, len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] > 0{
			left = mid + 1
		}else if nums[mid] < 0 {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return left
}

//最小栈


type MinStack struct {
	stack []int
	MinSk []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	if len(this.MinSk) == 0 || this.MinSk[len(this.MinSk) - 1] >= x {
		this.MinSk = append(this.MinSk, x)
	}
	this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
	if this.stack[len(this.stack) - 1] == this.MinSk[len(this.MinSk) - 1]{
		this.MinSk = this.MinSk[: len(this.MinSk) - 1]
	}
	this.stack = this.stack[ : len(this.stack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.MinSk[len(this.MinSk) - 1]
}



func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	//用栈来保存数组下标
	stack := make([]int, 0)
	for i := 0; i < len(T); i ++ {
		temperature := T[i]
		for len(stack) > 0 && temperature > T[stack[len(stack) - 1]]{
			res[stack[len(stack) - 1]] = i - (stack[len(stack) - 1])
			stack = stack[: len(stack) - 1]
		}
		stack = append(stack, i)
	}
	return res
}



type MyQueue struct {
	stackPush []int
	stackPop  []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stackPush = append(this.stackPush, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Pop()
	res := this.stackPop[len(this.stackPop) - 1]
	this.stackPop = this.stackPop[ : len(this.stackPop) - 1]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackPop) == 0 {
		for len(this.stackPush) >0 {
			this.stackPop = append(this.stackPop, this.stackPush[len(this.stackPush) - 1])
			this.stackPush = this.stackPush[ : len(this.stackPush) - 1]
		}
	}
	return this.stackPop[len(this.stackPop) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stackPop) == 0 && len(this.stackPush) == 0 {
		return true
	}
	return  false
}

type MaxQueue struct {
	q1 []int
	max []int

}


func Constructor() MaxQueue {
	return MaxQueue{}
}


func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0{
		return -1
	}
	return this.max[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.q1 = append(this.q1,value)
	for len(this.max) != 0 && value > this.max[len(this.max)-1]{
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max,value)
}


func (this *MaxQueue) Pop_front() int {
	var n int
	if len(this.q1) != 0{
		n = this.q1[0]
		this.q1 = this.q1[1:]
		if this.max[0] == n{
			this.max = this.max[1:]
		}
	}
	return n
}


func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)  + 1)
	for i :=0; i < len(dp) ; i ++ {
		dp[i] = make([]int, len(text2) + 1)
	}

	max := func(x,y int)int{
		if x<y{
			return y
		}
		return x
	}

	for i := 1; i <=len(text1); i ++ {
		for  j := 1 ; j <= len(text2); j ++ {
			if text1[i - 1] == text2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}else {
				dp[i][j] = max (dp[i][j - 1], dp[i -1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func frequencySort(s string) string {
	if len(s) == 0 {
		return ""
	}

	strsMap := make(map[string]string)
	for i := 0; i < len(s); i++ {
		strsMap[string(s[i])] += string(s[i])
	}

	strList := make([]string, len(strsMap))

	for _, val := range strsMap {
		strList = append(strList, val)
	}

	sort.Slice(strList, func(i, j int) bool {
		return len(strList[i]) >= len(strList[j])
	})
	return strings.Join(strList,"")
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0)

	for i := 0; i < len(people); i ++ {
		p := people[i][1]
		if p >= len(res) {
			res = append(res, people[i])
		} else {
			res = append(res, nil)
			copy(res[p + 1: ], res[p : ])
			res[p] = people[i]
		}
	}
	return res
}
