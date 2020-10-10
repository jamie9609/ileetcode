package level6

import (
	"math"
	"sort"
)

func maxPathSum(root *TreeNode) int {
	maxNum := math.MinInt64
	maxPath(root, &maxNum)
	return maxNum
}

func maxPath(root *TreeNode, maxNum *int) int {

	if root == nil {
		return 0
	}
	//返回的是以当前节点为根节点的最大贡献值，不包括子树组成的回路
	left := maxPath(root.Left, maxNum)
	right := maxPath(root.Right, maxNum)
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	if left + right + root.Val > *maxNum {
		*maxNum = left + right + root.Val
	}
	//当前节点的路径最大值，也即是节点贡献值。
	if left > right {
		return root.Val + left
	}
	return root.Val + right
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int][]int)
	res, others := make([]int, 0), make([]int, 0)
	for key, item := range arr2 {
		arr2Map[key] = append(arr2Map[key], item)
		arr2Map[key] = append(arr2Map[key], 0)
	}

	for i := 0; i < len(arr1); i ++ {
		for key, item := range arr2 {
			if item == arr1[i] {
				arr2Map[key][1] ++
			}  else {
				others = append(others, arr1[i])
				break
			}
		}
	}

	for i := 0; i < len(arr2Map); i ++ {
		for j := 0; j < arr2Map[i][1]; j ++ {
			res = append(res, arr2Map[i][0])
		}
	}

	for i := len(others) - 1 ; i >=0; i -- {
		res = append(res, others[i])
	}
	return res
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	indexNum := make([]int, len(arr2))
	otherNum := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(arr1) ; i ++ {
		flag := 0
		for j := 0; j < len(arr2) ; j ++ {
			if arr1[i] == arr2[j] {
				indexNum[j] ++
				flag = 1
				break
			}
		}
		if flag == 0 {
			otherNum = append(otherNum, arr1[i])
		}
	}

	for j := 0; j < len(arr2) ; j ++ {
		if indexNum[j] == 0 {
			continue
		}
		for i := 1; i <= indexNum[j] ; i ++ {
			res = append(res, arr2[j])
		}
	}
	sort.Slice(otherNum, func(i, j int) bool {
		return otherNum[i] < otherNum[j]
	})
	res = append(res, otherNum...)
	return res
}




func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	mid := make(map[int]int)
	for i := 0; i < len(nums2); i ++ {
		mid[nums2[i]] = i
	}

	for i := 0; i < len(nums1) ; i ++ {
		tag := true
		for j := mid[nums1[i]] + 1; j < len(nums2); j ++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				tag = false
				break
			}
		}
		if tag == true {
			res = append(res, -1)
		}
	}
	return res
}


type MinStack struct {
	stack []int
	min []int

}

func Constructor() MinStack {
	return MinStack{
		stack:   []int{},
		min:	 []int{},
	}

}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] { // 没有元素或者x小于最小值的时候，append到min最后
		this.min = append(this.min, x)
	}
}


func (this *MinStack) Pop()  {
	if this.stack[len(this.stack) - 1] == this.min[len(this.min) - 1] {
		this.min = this.min[:len(this.min) - 1]
	}
	this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) Min() int {
	return this.min[len(this.min) - 1]
}



func maxArea(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	left := 0
	right := len(height) - 1

	maxNum := math.MinInt64

	for left <= right {
		maxNum = maxx(maxNum, (right - left) * minx(height[left], height[right]))
		if height[left] > height[right] {
			right --
		} else {
			left ++
		}
	}
	return maxNum
}

func minx(a , b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxx(a , b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}



func letterCombinations(digits string) []string {
	//如果是空，则直接返回
	if len(digits) == 0 {
		return nil
	}

	//创建存储结果的切片
	res := make([]string, 0)
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	//将第一个元素添加到切片中，方便下面的操作
	for _, v := range m[string(digits[0])] {
		res = append(res, string(v))
	}
	//如果只有一个元素，则返回结果
	if len(digits) == 1 {
		return res
	}

	//循环，依次添加
	for i := 1; i < len(digits); i++ {
		res = combination(res, m[string(digits[i])])
	}

	return res

}

func combination(arr []string, s string) (res []string) {
	for _, v := range s {
		for i := 0; i < len(arr); i++ {
			res = append(res, arr[i]+string(v))
		}
	}
	return
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if len(nums2) == 0 {
		return
	}
	if m == 0 {
		for i := 0 ; i < len(nums2) ; i ++ {
			nums1[i] = nums2[i]
		}
		return
	}

	flag1 := m - 1
	flag2 := n - 1
	item := len(nums1) - 1

	for flag2 >= 0 &&  flag1 >= 0{
		if nums2[flag2] >= nums1[flag1] {
			nums1[item] = nums2[flag2]
			item --
			flag2 --
		}else {
			nums1[item] = nums1[flag1]
			item --
			flag1 --
		}
	}
	if flag1 < 0 {
		for i := 0; i <= item ; i ++ {
			nums1[i] = nums2[i]
		}
	}
	return
}

func moveZeroes(nums []int)  {
	if nums == nil {
		return
	}

	var j int

	for i := 0; i < len(nums) ; i ++ {
		if nums [i] == 0 {
			continue
		}
		nums [i], nums [j] = nums [j], nums [i]
		j ++
	}
	return
}

func findContinuousSequence(target int) [][]int {
	if target <= 0 || target == 2 {
		return nil
	}
	if target == 1 {
		return [][]int{{1}}
	}

	res := make([][]int, 0)

	left, right := 1, 2
	for left < right {
		sum := (right + left) * (right - left + 1) / 2
		if sum == target {
			tmpRes := make([]int, 0)
			for i := left; i <= right; i ++ {
				tmpRes = append(tmpRes, i)
			}
			res = append(res, tmpRes)
			right ++
			left ++
		}else if sum < target {
			right ++
		}else {
			left ++
		}
	}
	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalL := len(nums1) + len(nums2)
	if totalL % 2 == 1 {
		midIndex := totalL / 2
		return float64(getKthElement(nums1, nums2, midIndex + 1))
	}else {
		midIndex1, midIndex2 := totalL/2 - 1, totalL/2
		return float64(getKthElement(nums1, nums2, midIndex1 + 1) + getKthElement(nums1, nums2, midIndex2 + 1)) / 2.0
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return  nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1 + half, len(nums1)) - 1
		newIndex2 := min(index2 + half, len(nums2)) - 1

		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]

		if pivot1 <= pivot2 {
			k = k - (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k = k - (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


func sortArrayByParityII(A []int) []int {
	odd := 0
	even := 1

	for odd < len(A) && even < len(A) {
		if odd & 1 == A[odd] & 1 {
			odd += 2
			continue
		}
		if even & 1 == A[even] & 1 {
			even += 2
			continue
		}
		A[odd], A[even] = A[even], A[odd]
	}
	return A
}

func isAnagram(s string, t string) bool {
	/*a := [26]int{}
	n := [26]int{}*/
	var a [26]int
	var n [26]int
	for _, v := range s {
		a[v - 'a'] += 1
	}
	for _, v := range t {
		n[v - 'a'] += 1
	}
	return a == n
}

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	cur, pre := 1, 1
	for i := 1; i < len(s); i ++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			} else {
				cur = pre
			}
		} else if s[i - 1] == '1' ||  s[i-1] == '2' && s[i] <= '6'{
			cur += pre
		}
		pre = tmp
	}
	return cur
}

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)

	if n < 4 {
		return nil
	}

	sort.Ints(nums)

	for a := 0; a <= n-4; a ++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b <= n-3; b ++ {
			if b > a+1 && nums[b] == nums[b-1] { // 剪枝
				continue
			}
			c, d := b+1, n-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c] == nums[c+1] {
						c ++
					}
					for c < d && nums[d] == nums[d-1] {
						d --
					}
					c ++
					d --
				} else if sum < target {
					c ++
				} else {
					d --
				}
			}
		}
	}
	return ans
}




func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i ++ {
		dp[i] = cost[i] + minNum(dp[i - 1], dp[i - 2])
	}
	return dp[len(cost) - 1]
}

func minNum(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {

}






