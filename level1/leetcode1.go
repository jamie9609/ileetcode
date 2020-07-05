package level1
/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]*/

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++{
		for j := i + 1; j < len(nums); j ++ {
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}

		}
	}
	return nil
}


func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make ([]bool, len(candies))
	for i :=0; i< len(candies); i ++{
		sum := candies[i] + extraCandies
		for j:=0; j< len(candies); j++{
			if candies[j] > sum {
				result[i] = false
				break
			} else {result[i] = true}
		}
	}
	return result
}

func XorOperation(n int, start int) int {
	nums := make([]int, n)
	result := 0
	for i := 0; i < n; i++ {
		nums[i] = start + 2 * i
		result = result ^ nums[i]
	}
	return result
}


func Shuffle(nums []int, n int) []int {
	result := make ([]int, 2*n)
	for i := 0 ; i < n ; i ++ {
		result[i * 2] = nums[i]
	}
	for i := n;  i >= n && i < 2*n; i ++ {
		result[ (i-n) *2 + 1] = nums[i]
	}
	return result
}

func ReverseLeftWords(s string, n int) string {
	var s1 string
	var s2 string
	for i := 0; i < len(s); i ++ {
		if i < n {
			s1 += string(s[i])
		} else if i >= n {
			s2 += string(s[i])
		}
	}
	return s2 + s1
}

