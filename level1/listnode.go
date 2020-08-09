package level1

func lengthOfLongestSubstring(s string) int {
	head := 0
	tail := 0
	res := 1
	if len(s) <= 2 {
	return len(s)
	}

	for tail = 1; tail < len(s); tail ++ {
		for {
			if exitStr(string(s[tail]), s[head:tail]) == false {
				break
			}
			head += 1
		}

		if res >= tail - head + 1 {
			continue
		}
		if res < tail - head + 1 {
			res = tail - head + 1
		}
	}
	return res

}

func exitStr (strKey string, str2 string) bool{
	for _, val := range str2 {
		if string(val)== strKey {
			return true
		}
		continue
	}
	return false
}