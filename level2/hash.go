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
