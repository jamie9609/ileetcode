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
//合并两个有序链表，递归方法

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val{
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}