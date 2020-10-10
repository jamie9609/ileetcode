package level7

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func singleNumber(nums []int) int {
	flag := 0
	for _, num := range nums {
		flag = flag ^ num
	}
	return flag
}

func containsDuplicate(nums []int) bool {
	res := make(map[int]int, 0)
	for _, val := range nums {
		if _, ok := res[val]; ok {
			return true
		}
		res[val] ++
	}
	return false
}

func combine(n int, k int) [][]int {
	var res [][]int
	if k == 0 {
		return res
	}
	var nums = make([]int, n)
	for i, _:= range nums {
		nums[i] = i + 1
	}
	backtracking(&res, nums, 0, k)
	return res
}

func backtracking(res *[][]int, nums []int, index, k int)  {
	if index == k {
		* res = append(*res, append([]int{}, nums[0:k]...))
		return
	}

	for start := index; start < len(nums); start ++ {
		if index == 0 || nums[start] > nums[index - 1] {
			nums[start], nums[index] = nums[index], nums[start]
			backtracking(res, nums, index + 1, k)
			nums[index], nums[start] = nums[start], nums[index]
		}
	}
	return
}

func addStrings(num1 string, num2 string) string {
	add := 0
	res := ""

	for i,j := len(num1) - 1,len(num2) - 1; i >= 0 || j >= 0 || add != 0; i ,j = i -1, j - 1{
		var x, y int
		if i >= 0{
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		tmp := x + y + add
		add = tmp / 10
		res = strconv.Itoa(tmp % 10) + res
	}
	return res
}

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	g := gcd(l1, l2)
	subStr := str1[:g]

	if check(str1, subStr) && check(str2, subStr) {
		return subStr
	}
	return ""
}

func check(s, sub string) bool {
	l := len(s) / len(sub)
	r := ""
	for i := 0; i < l ; i ++ {
		r += sub
	}
	return r == r
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd (b, a%b)
}


func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return checkHelp(rec1, rec2) || checkHelp(rec2, rec1)
}

func checkHelp(rec1 []int, rec2 []int) bool {
	if ( rec2[0] > rec1[0] && rec2[0] < rec1[2] ) && (rec2[1] > rec1[1] && rec2[1] < rec1[3]) {
		return true
	}

	if ( rec2[0] > rec1[0] && rec2[0] < rec1[2] ) && (rec2[3] > rec1[1] && rec2[3] < rec1[3]) {
		return true
	}
	if ( rec2[2] > rec1[0] && rec2[2] < rec1[2] ) && (rec2[3] > rec1[1] && rec2[3] < rec1[3]) {
		return true
	}
	if ( rec2[2] > rec1[0] && rec2[2] < rec1[2] ) && (rec2[1] > rec1[1] && rec2[1] < rec1[3]) {
		return true
	}
	return false
}


func singleNumbers(nums []int) []int {
	if nums == nil {
		return nil
	}

	xor := 0

	for _, val := range nums {
		xor = xor ^ val
	}
	position := 0

	for i := xor; i & 1 == 0; i >>= 1 {
		position ++
	}
	tempXor := xor
	for _, v := range nums {
		if (uint(v) >> uint(position)) & 1 == 1{
			xor ^= v
		}
	}

	return []int{xor, xor ^ tempXor}
}


func getRow(rowIndex int) []int {
	nums := []int{1}
	for i := 1; i <= rowIndex; i ++ {
		nums = append(nums, 1)
		for j := i - 1 ; j > 0; j -- {
			nums[j] += nums[j  - 1]

		}
	}
	return nums
}

func hammingDistance(x int, y int) int {
	c := x ^ y
	count := 0
	for c > 0 {
		if c & 1 > 0 {
			count ++
		}
		c >>= 1
	}
	return count
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1{
		return 1
	}
	res := make([]int, N+1 )
	res[1] = 1
	for i := 2; i <= N; i ++ {
		res[i] = res[i - 2] + res[i -1]
	}
	return res[N]
}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num & 1 > 0 {
			res ++
		}
		num >>= 1
	}
	return res
}

func majorityElement(nums []int) int {
	resMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i ++ {
		resMap[nums[i]] ++
	}
	for key, val := range resMap {
		if val > len(nums) / 2 {
			return key
		}
	}
	return -1
}
func readBinaryWatch(num int) []string {
	var res []string
	backtracking2(0, 10, num, []int{}, &res)
	return res
}
func backtracking2(start, cap, target int,path []int,res *[]string)  {
	if len(path) == target {
		min, hour := 0, 0
		for _, v := range path {
			if v >= 4 {
				min += 1 << uint(v - 4)
			}else {
				hour += 1 << uint(v)
			}
		}
		if min < 60 && hour < 12 {
			*res = append(*res, fmt.Sprintf("%d:%02d", hour,min))
		}
		path = []int{}
		return
	}

	for i := start ; i < cap ; i ++ {
		path = append(path, i)
		backtracking2(i + 1, cap, target - 1, path, res)
		path = path[:len(path) - 1]
	}
}


func lemonadeChange(bills []int) bool {
	extra := make(map[int]int, 0)
	for i := 0 ; i < len(bills); i ++ {
		if bills[i] == 5 {
			extra[5] ++
		}
		if bills[i] == 10 {
			if extra[5] >= 0 {
				extra[5] --
				extra[10] ++
			} else {
				return false
			}
		}
		if bills[i] == 20 {
			if extra[10] > 0 && extra[5] > 0 {
				extra[10] --
				extra[5] --
			} else {
				if extra[5] < 3 {
					return false
				}
				extra[5] -= 3
			}
		}
	}
	return true
}


func isPowerOfFour(num int) bool {
	if num == 1 {
		return true
	}
	index := 0
	for num >= 1 {
		index = num % 4
		if index != 0 {
			return false
		}
		num = num / 4
		if index == 0 && num == 1{
			return true
		}
	}
	return false
}

func peakIndexInMountainArray(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	left := 0
	right := len(arr) -1

	for left <= right {
		mid := left + (right - left) / 2
		if arr[mid + 1] > arr[mid] {
			left = mid + 1
		}else if arr[mid - 1] > arr[mid] {
			right = mid - 1
		}else {
			return arr[mid]
		}
	}
	return -1
}

func toHex(num int) string {
	if num < 0 {    // 如果是负数，转换成补码形式，比如：-1 + 4294967296 = 4294967295 = 0xffffffff
		num += 4294967296
	}
	var ans []rune
	hash := [16]rune{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'}
	for {           // 迭代提取每个位上的十六进制字符
		t := num%16
		num /= 16
		ans = append([]rune{hash[t]}, ans...)
		if num == 0 {
			break
		}
	}
	return string(ans)
}

//不用加法符号的加法运算。
func getSum(a int, b int) int {
	and := (a & b) << 1
	xor := a ^ b
	for and != 0 {
		b = and
		a = xor
		and = (a & b) << 1
		xor = a ^ b
	}
	return xor
}

func removeDuplicates(S string) string {
	if len(S) == 0 || len(S) == 1 {
		return ""
	}
	left := 0
	right := 1
	for right < len(S) {
		if S[right] == S[left] && left == 0{
			if len(S) == 2 {
				return ""
			}else {
				S = S[2:]
			}
		}else if S[right] == S[left] {
			tmp := S[right + 1:]
			S = S[:left]
			S = S + tmp
			left --
			right --
		}else {
			left ++
			right ++
		}
	}
	return S
}

func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i ++ {
		for j := 0; j < len(grid[0]); j ++ {
			if i == 0 && j == 0{
				continue
			}
			if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			}
			if i != 0 && j == 0 {
				grid[i][j] += grid[i - 1][j]
			}
			if i != 0 && j != 0 {
				grid[i][j] += maxNum(grid[i - 1][j], grid[i][j-1])
			}
		}
	}
	return grid[len(grid) - 1][ len(grid[0]) - 1]
}
func maxNum1(x int,y int) int{
	if x>y{
		return x
	}else{
		return y
	}
}

func nextGreatestLetter1(letters []byte, target byte) byte {
	for i := 0 ; i < len(letters); i ++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	if letters[right] <= target {
		return letters[left]
	}
	for left <= right {
		mid := (left + right) / 2
		if letters[mid] > target {
			right = mid - 1
		} else if letters[mid] == target {
			left = mid + 1
		} else {
			left = mid + 1
		}
	}
	return letters[left]
}


func pruneTree(root *TreeNode) *TreeNode {
	return deal(root)
}

func deal(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	node.Left = deal(node.Left)
	node.Right = deal(node.Right)

	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return nil
	}
 	return node
}

//K个鸡蛋，N层建筑
func superEggDrop1(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}
	if K == 1 {
		return N
	}
	if N == 1{
		return 1
	}

	dp := make([][]int, K + 1)
	for i := 0; i < K + 1; i ++ {
		dp[i] = make([]int, N + 1)
	}

	for i := 0; i < N + 1 ; i ++ {
		dp[0][i] = 0
	}
	for i := 0; i < K + 1 ; i++ {
		dp[i][1] = 1
	}
	for i := 0; i < K + 1 ; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < N + 1 ; i++ {
		dp[1][i] = i
	}

	for i := 2 ; i <= K ; i ++ {
		for j := 2; j <= N; j ++ {
			for x := 1; x <= j ; x ++ {
				if dp[i][j] == 0 {
					dp[i][j] = math.MaxInt32
				}
				dp[i][j] = minNum(dp[i][j], maxNum(dp[i -1][x - 1], dp[i][j - x])+1)
			}
		}
	}
	return dp[K][N]

}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func minNum(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//K个鸡蛋，N层建筑
func superEggDrop(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i ++ {
		dp[i] = make([]int, K + 1)
	}

	for i := 0; i < N+1; i ++ {
		dp[i][0] = 0
		dp[i][1] = i
	}
	for i := 0; i < K+1; i ++ {
		dp[0][i] = 0
		dp[1][i] = 1
	}

	for i := 2 ; i < N + 1; i ++ {
		for j := 2; j < K + 1; j ++ {
			left := 1
			right := i
			for left <= right{
				mid := left + (right - left) / 2
				breakCount := dp[mid - 1][j - 1]
				notBreakCount := dp[i - mid][j]
				if breakCount > notBreakCount {
					right = mid - 1
				}else {
					left = mid
				}
			}
			dp[i][j] = maxNum(dp[left - 1][j - 1], dp[i - left][j]) + 1
		}
	}
	return dp[N][K]
}



var nums [1690]int
func init() {
	nums[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < 1690; i++ {
		ugly := min(nums[i2]*2, nums[i3]*3, nums[i5]*5)
		nums[i] = ugly
		if nums[i2] * 2 == ugly {
			i2++
		}
		if nums[i3] * 3 == ugly {
			i3++
		}
		if nums[i5] * 5 == ugly {
			i5++
		}
	}

}
func min2(a, b int)int {
	if a < b {
		return a
	}
	return b
}
func min(a, b, c int) int {
	return min2(a,min2(b,c))
}
func nthUglyNumber(n int) int {
	return nums[n-1]
}






func cuttingRope(n int) int {
	dp := make([]int, n + 1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j * (i - j), j * dp[i - j]))
		}
		dp[i] = curMax
	}
	return dp[n] % 1000000007
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}



func kWeakestRows(mat [][]int, k int) []int {
	ans := make([]int, 0, k)
	resMap := make(map[int][]int, len(mat))

	for key, val := range mat {
		tmp := 0
		for i := 0; i < len(val); i ++ {
			if val[i] == 1 {
				tmp ++
			}
		}
		resMap[tmp] = append(resMap[tmp], key)
	}

	for i := 0; i <= len(mat[0]); i ++ {
		ans = append(ans, resMap[i]...)
		if len(ans) >= k {
			ans = ans[:k]
			break
		}
	}
	return ans
}




func Subsets(nums []int) [][]int {
	res := make([][]int, 0 )
	res = append(res, []int{})

	for i := 0; i < len(nums); i ++ {
		tmpRes := make([][]int, 0)
		oldRes := res
		for _,val := range res {
			val = append(val, nums[i])
			tmpRes = append(tmpRes, val)
		}
		oldRes = append(oldRes, tmpRes...)
		res = oldRes
	}
	return res
}







func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	ab := make(map[int]int)
	cd := make(map[int]int)

	for i := 0; i < len(A); i ++ {
		for j := 0; j < len(B); j ++ {
			num := A[i] + B[j]
			if _, ok := ab[num]; ok {
				ab[num] ++
			}else {
				ab[num] = 1
			}
		}
	}

	for k := 0;k < len(C);k ++{
		for l:=0;l<len(D);l++{
			num := C[k]+D[l]
			if _,ok:=cd[num];ok{
				cd[num]++
			}else {
				cd[num] = 1
			}
		}
	}

	for abkey, abval := range ab {
		if _, ok := cd[-abkey]; ok {
			res += abval * cd[-abkey]
		}
	}
	return res

}



func largestPerimeter(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	for i := len(A) - 1; i >= 2; i -- {
		if A[i - 1] + A[i - 2] > A[i] {
			return A[i - 1] + A[i - 2] + A[i]
		}
	}
	return 0

}


func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	left, right := 0, len(nums) - 1
	mid := left + (right - left) / 2
	return nums[mid]
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right{
		mid := left + (right - left) / 2
		if isBadVersion(mid) == true {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return left
}


type NumArray struct {
	Value []int
}


func Constructor(nums []int) NumArray {
	return NumArray{nums}
}


func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for ;i<=j; i++ {
		sum += this.Value[i]
	}
	return sum
}

func majorityElement(nums []int) int {
	maps := make(map[int]int)
	for _,value := range nums{
		maps[value] ++
	}

	for keyMap,valueMap := range maps{
		if valueMap > len(nums) {
			return keyMap
		}
	}
	return -1
}

func waysToStep(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1{
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n + 1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n ; i ++ {
		dp[i] = (dp[i - 1]  + dp[i - 2]  + dp[i - 3] ) % 1000000007

	}
	return dp[n]
}


func minSubsequence(nums []int) []int {
	quickSort(nums, 0, len(nums) -1)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	sub := 0
	for i := 0; i < len(nums); i++ {
		sub += nums[i]
		if sub > (sum - sub) {
			return nums[0 : i+1]
		}
	}

	return nil
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l, r
	sepVal := nums[r]

	for i < j {
		if i < j && nums[i] >=sepVal {
			i ++
		}
		nums[j] = nums[i]

		if i < j && nums[j] < sepVal {
			j --
		}
		nums[i] = nums[j]
	}

	nums[i] = sepVal
	quickSort(nums, l , i - 1)
	quickSort(nums, i , r)
}

func sortString(s string) string {
	list := make([]int, 26)
	for i := 0; i < len(s); i++ {
		list[int(s[i]) - int('a')]++
	}

	ret := make([]byte, 0, len(s))
	n := 0
	for n <= len(s) {
		if n & 1 == 0 {
			for i := 0; i < len(list); i ++ {
				if list[i] != 0 {
					ret = append(ret, byte(i + 'a'))
					list[i] --
				}
			}
		} else {
			for i := len(list)-1; i >= 0; i-- {
				if list[i] != 0 {
					ret = append(ret, byte(i + int('a')))
					list[i]--
				}
			}
		}
		n ++
	}
	return string(ret)
}





