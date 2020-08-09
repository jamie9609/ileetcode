package level2

/*type ListNode struct {
	Val int
	Next *ListNode
}*/

/**
判断是否有环
 */

//时间复杂度O(N)，空间复杂度O(1)  快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != nil  && fast != nil && fast.Next != nil {
		if fast == slow{
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}

// 哈希表判定 时间复杂度O(n), 空间复杂度O(n)
func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	tmp := head
	hashMap := make(map[*ListNode]int)
	for tmp != nil {
		if _, ok := hashMap[tmp]; ok {
			return true
		}
		hashMap[tmp] = tmp.Val
		tmp = tmp.Next
	}
	return false
}


