package level2

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}



/**
从尾到头打印链表。
 */
//很容易联想到栈，或者是利用golang的切片来实现。

func reversePrint(head *ListNode) []int {
	var s []int
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}

	for i ,j := 0, len(s) - 1; i < j ; {
		s[i], s[j] = s[j], s[i]
		i ++
		j --
	}
	return s
}

func main()  {

	var head ListNode
	var node ListNode
	head.Val = 5
	node.Val = 8
	head.Next = &node

	s := reversePrint(&head)
	fmt.Println(s)

}

func reversePrint2(head *ListNode) []int  {
	var s []int
	stack := list.New()
	for head != nil {
		stack.PushBack(head.Val)
		head = head.Next
	}

	for stack.Len() > 0{
		e := stack.Back()
		stack.Remove(e)
		s = append(s, e.Value.(int))
	}
	return s
}



//双指针比一般方法要快。
func getKthFromEnd(head *ListNode, k int) *ListNode {
	former := head
	latter := head
	for i := 0 ; i < k ; i ++ {
		former = former.Next
	}

	for former != nil {
		former = former.Next
		latter = latter.Next
	}
	return latter
}

//删除链表节点，双指针法
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre := head
	cur := head.Next
	for cur.Val != val{
		cur = cur.Next
		pre = pre.Next
	}
	pre.Next = cur.Next
	return head
}
//两个链表的第一个公共节点

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	solder1 := headA
	solder2 := headB
	i := 0
	j := 0

	if solder2 == nil && solder1 == nil {
		return nil
	}
	if solder2 == nil || solder1 == nil {
		return nil
	}
	//需要考虑没有公共节点的情况。
	for solder1 != solder2 {
		solder1 = solder1.Next
		solder2 = solder2.Next

		if solder1 == nil {
			solder1 = headB
			i ++
		}
		if solder2 ==  nil {
			solder2 = headA
			j ++
		}
		if i > 1 || j > 1{
			return nil
		}
	}
	return solder1
}