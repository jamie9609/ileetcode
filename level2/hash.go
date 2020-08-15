package level2

func findRepeatNumber(nums []int) int {
	target := make(map[int]int)
	for i := 0; i < len(nums); i++{
		if _, ok := target[nums[i]]; ok {
			return nums[i]
		}
		target[nums[i]] = i
	}
	return -1
}

func firstUniqChar(s string) byte {
	var result [26]int

	for i :=0 ; i <len(s); i++{
		result[ int(s[i] - 'a') ]  += 1
	}
	for i :=0 ; i <len(s); i++ {
		if result[ int(s[i] - 'a') ]  == 1{
			return s[i]
		}
	}
	return ' '
}


//回文排列
// 给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

func canPermutePalindrome(s string) bool {
	res := make(map[int32]int)

	for _, val := range s {
		res[val] ++
	}
	sum := 0
	for _, val := range res{
		if val % 2 != 0{
			sum ++
		}
	}
	if sum == 0|| sum == 1{
		return true
	}
	return false

}
