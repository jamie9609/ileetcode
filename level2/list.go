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

//删除环。
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

func reverseList1(head *ListNode) *ListNode {
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


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := &ListNode{0, nil}
	list.Next = head
	pre, cur  := list, head
	num := 1
	for cur != nil {
		if num == n {
			cur = cur.Next
			if cur == nil {
				pre.Next = pre.Next.Next
				break
			}
			pre = pre.Next
			continue
		}
		cur = cur.Next
		num ++
	}
	return list.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	soldier := &ListNode{0, nil}
	soldier.Next = head
	res := head.Next

	for head != nil && head.Next != nil {

		cur := head
		pre := head.Next

		soldier.Next = pre
		cur.Next = pre.Next
		pre.Next = cur

		soldier = cur
		head = cur.Next
	}

	return res
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	soldier := &ListNode{0, nil}
	pre := soldier
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val{
			pre.Next = l2
			if l2 == nil {
				continue
			}else {l2 = l2.Next}
		} else {
			pre.Next = l1
			if l1 == nil {
				continue
			} else {
				l1 = l1.Next
			}
		}
		pre = pre.Next
	}
	if l1== nil {
		pre.Next = l2
	}
	if l2== nil {
		pre.Next = l1
	}
	return soldier.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0{
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	return mergeTwoLists2(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))

}

//k个一组反转链表

func reverseKGroup1(head *ListNode, k int) *ListNode {
	soldier := &ListNode{0 ,head}
	pre := soldier

	for head != nil{
		tail := pre
		for i := 0; i < k; i ++{
			tail = tail.Next
			//此处是如果为空，直接返回
			if tail == nil {
				return soldier.Next
			}
		}
		tmp := tail.Next
		head, tail = reverseList3(head, tail)
		pre.Next = head
		tail.Next = tmp
		pre = tail
		head = tail.Next
	}
	return  soldier.Next
}

func reverseList3(head *ListNode, tail *ListNode) (heads *ListNode, tails *ListNode) {

	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	soldier := &ListNode{Next: head}
	pre := soldier

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return soldier.Next
			}
		}
		tmp := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = tmp
		pre = tail
		head = tail.Next
	}
	return soldier.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}


func rotateRight(head *ListNode, k int) *ListNode {
	hair := &ListNode{0, head}
	pre, soldier := head, head
	num := 1
	for pre != nil && pre.Next != nil{
		pre = pre.Next
		num ++
	}
	k = k % num
	if head == nil || k == 0{
		return head
	}
	for i := 1; i < num - k ; i ++ {
		head = head.Next
	}
	hair.Next = head.Next
	head.Next = nil
	pre.Next = soldier
	return hair.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	pre := head
	if head == nil || head.Next == nil {
		return head
	} else {
		for head != nil {
			if head.Next != nil {
				for head.Val == head.Next.Val {
					head.Next = head.Next.Next
					if head.Next == nil {
						break
					}
				}
			}
			head = head.Next
		}
		return pre
	}
}



func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hair := &ListNode{0, head}
	pre := hair
	cur := hair.Next

	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val{
			cur = cur.Next
		}
		if pre.Next != cur{
			cur = cur.Next
			pre.Next = cur
			continue
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}
	return hair.Next
}


func partition(head *ListNode, x int) *ListNode {
	listA := &ListNode{}
	listB := &ListNode{}
	small, big := listA, listB
	for head != nil {
		if head.Val >= x {
			big.Next = head
			big = big.Next
		}
		if head.Val < x {
			small.Next = head
			small = small.Next
		}
		head = head.Next
	}
	small.Next = listB.Next
	big.Next = nil

	return listA.Next
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{0, head}
	prem := dummy

	for i := 1; i <= m-1; i++ {
		prem  = prem.Next
	}
	current := prem.Next
	var pre1 *ListNode
	for i := m ; i <=n ; i ++ {
		temp := current.Next
		current.Next = pre1
		pre1 = current
		current = temp
	}
	prem.Next.Next = current
	prem.Next = pre1
	return dummy.Next
}


type Node struct {
	Val     int
	Next    *Node
	Random	*Node
}


func copyRandomList(head *Node) *Node {
	res := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		tmp := &Node{cur.Val, nil ,nil}
		res[cur] = tmp
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		res[cur].Next = res[cur.Next]
		res[cur].Random = res[cur.Random]
		cur = cur.Next
	}
	return res[head]
}


func reorderList(head *ListNode)  {
	num := 0
	i := 0

	curA, curB, flag := head, head, head
	for head != nil {
		head = head.Next
		num ++
	}
	num = num >> 1
	//cur := &ListNode{0,nil}

	for flag != nil {
		flag = flag.Next
		i ++
		if i >= num {
			curB = reverseList4(flag)
			break
		}
	}
	for curA != nil{
		tmpA := curA.Next
		tmpB := curB.Next
		curA.Next = curB
		curB.Next = tmpA
		curA = tmpA
		curB = tmpB

	}
}

func reverseList4(head *ListNode) *ListNode {
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



func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	pre := dummy
	cur := head
	for cur != nil {
		curNext := cur.Next
		if pre.Next == nil {
			pre.Next = cur
			cur = curNext
		}
		for pre!= nil || pre.Next != nil{
			if pre.Next.Val >= cur.Val {
				cur.Next = pre.Next
				pre.Next = cur
				cur = curNext
				pre = dummy
				break
			} else if pre.Next == nil {
				pre.Next = cur
				cur = curNext
				pre = dummy
				break
			} else {
				pre = pre.Next
				cur = curNext
				continue
			}
		}
	}
	return dummy.Next
}


func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	if pre.Next == nil {
		return nil
	}
	preNext := dummy.Next
	for preNext != nil {
		if preNext.Val == val{
			if preNext.Next == nil {
				pre.Next = nil
				break
			} else {
				preNext = preNext.Next
				pre.Next = preNext
				continue
			}
		}
		pre = preNext
		preNext = preNext.Next
	}
	return dummy.Next

}