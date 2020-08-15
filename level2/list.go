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


func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	tmp := head
	for tmp != slow {
		tmp = tmp.Next
		slow = slow.Next
	}
	return slow
}


//爬楼梯，斐波拉契数列

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	var res []int
	res = append(res, 1)
	res = append(res, 1)
	for i := 0; i < n ; i ++{
		res = append(res, res[i]+res[i+1])
	}
	return res[n]
}


type MyStack struct {
	slice []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}

}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.slice = append(this.slice, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.slice) == 0{
		return -1
	}
	top := this.slice[len(this.slice) -1]
	this.slice = this.slice[ : len(this.slice) -1 ]
	return top

}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.slice[len(this.slice) - 1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.slice) == 0{
		return true
	} else {
		return false
	}
}

func reverseList2(head *ListNode) *ListNode {
	tmp, cur:= head, head
	var pre *ListNode = nil
	for cur != nil{
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}


func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}


func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	result := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}
	return result.Next
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	res := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}
	return res.Next
}


