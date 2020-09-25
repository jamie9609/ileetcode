package level5

import (
	"math"
	"sort"
)

//长度最小的子数组

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	left, right := 0, 0
	length := len(nums) - 1
	minLen := math.MaxInt64
	if sumNums (nums,left,length) < s {
		return 0
	}
	for right <= length && left <=  length && left <= right {
		if sumNums(nums,left,right) >= s {
			minLen = min(minLen, right-left +1)
			left ++
			continue
		}
		right ++

	}
	return minLen
}
func sumNums(slice []int, i int, j int) int {
	sum := 0
	for key, val := range slice {
		if key >= i && key <= j {
			sum += val
		}
	}
	return sum
}
func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}


func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool)
	res := make([]int, 0)

	for _, val := range nums1 {
		numMap[val] = true
	}

	for _, val := range nums2 {
		if numMap[val] == true {
			res = append(res, val)
			numMap[val] = false
		}
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]int)
	res := make([]int, 0)

	for _, val := range nums1 {
		numsMap[val] += 1
	}

	for _, val :=range nums2 {
		if numsMap[val] >= 1 {
			res = append(res, val)
			numsMap[val] --
		}
	}
	return res
}

//寻找重复数
func findDuplicate(nums []int) int {
	var numsmap = make(map[int]bool)
	for _,num := range nums {
		if !numsmap[num] {
			numsmap[num] = true
		} else {
			return num
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if nums[len(nums) - 1] < target {
		return len(nums)
	}
	left := 0
	right := len(nums) -1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}


func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || matrix == nil {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := n * m -1

	for left <= right {
		mid := left + (right - left) / 2
		xm := mid / n
		xn := mid % n
		if matrix[xm][xn] == target {
			return true
		}else if matrix[xm][xn] > target {
			right = mid - 1
		}else {
			left = mid +1
		}
	}
	return false
}


func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || nums  == nil {
		return 0
	}
	dp := make([]int, len(nums))
	for key, _ := range dp {
		dp[key] = 1
	}
	ans := 1
	for i := 1; i < len(nums); i ++ {
		for j := 0; j < i; j ++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
				ans = max(ans, dp[i])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

//单调栈
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	start, end, ans := 0, 0, 0
	m := make(map[int32]int, 0)

	for i, v := range s{

		if val, ok := m[v]; ok {
			start = max(val, start)
		}
		m[v] = i + 1
		ans = max(end - start + 1, ans)
		end ++
	}

	return ans
}
//无重叠区间

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	//按照end升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1

	xEnd := intervals[0][1]

	for _, val := range intervals {
		start := val[0]
		if start >= xEnd {
			count ++
			xEnd = val[1]
		}
	}
	return len(intervals) - count
}


func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	dp := make([]int, len(pairs))
	for key, _ := range dp {
		dp[key] = 1
	}
	res := 0
	for i := 0; i < len(pairs) ; i ++ {
		for j := 0; j < i; j ++ {
			if pairs[i][0] > pairs[j][1]{
				dp[i] = max(dp[i], dp[j] + 1)
				res = max(res, dp[i])
			}
		}
	}
	return res
}













