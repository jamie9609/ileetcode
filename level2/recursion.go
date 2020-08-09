package level2

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	result := make([]int,n+1 )
	result[0] = 0
	result[1] = 1
	for i := 2; i <= n ; i ++{
		result[i] = (result[i - 1] + result[i - 2]) % 1000000007
	}
	return result[n]

}