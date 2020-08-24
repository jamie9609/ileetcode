package level3

import "github.com/pingcap/parser/ast"

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast.Next !=  nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//递归，分割链表
	r := sortList(slow.Next)
	slow.Next = nil
	l := sortList(head)
	//合并，将两个list合并为一个

	tmpHead := &ListNode{0,nil}

	return mergeList(tmpHead, l, r)
}


func mergeList(tmpHead, l, r *ListNode) *ListNode {
	p := tmpHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l == nil {
		p.Next = r
	}
	if r == nil {
		p.Next = l
	}
	return tmpHead.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	preA, preB := headA, headB
	num := 0
	if headA == headB {
		return headA
	}
	for preA != nil && preB != nil {
		if preA == preB{
			return preA
		}
		preA = preA.Next
		preB = preB.Next
		num ++
		if preA == nil {
			preA = headB
			continue
		}
		if preB == nil {
			preB = headA
			continue
		}
	}
	return nil

}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func oddEvenList1(head *ListNode) *ListNode {
	if head ==nil || head.Next == nil || head.Next.Next == nil  {
		return head
	}
	slow, fast := head, head.Next
	tmp := head.Next
	for fast != nil  && fast.Next != nil {
		slow.Next = slow.Next.Next
		fast.Next = fast.Next.Next
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = tmp
	return head
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	p := slow
	q := fast
	for q != nil && q.Next != nil {
		p.Next = p.Next.Next
		q.Next = q.Next.Next

		p = p.Next
		q = q.Next
	}
	p.Next = fast
	return slow
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	} else if l2 == nil {
		return l1
	}
	var num1, num2, num3  []int
	flag := 0
	l3 := &ListNode{0,nil}
	head := l3
	for l1 != nil {
		num1 = append(num1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		num2 = append(num2, l2.Val)
		l2 = l2.Next
	}
	for i, j := 0, 0; i < len(num1) || j < len(num2); i, j = i+1, j+1 {
		tmp := 0
		if i >= len(num1) {
			tmp = num2[len(num2) - j - 1] + flag
			if tmp >= 10{
				tmp /= 10; flag = 1
			} else {
				flag = 0
			}
			num3 = append(num3, tmp)
			continue
		}
		if j >= len(num2) {
			tmp = num1[len(num1) - i - 1] + flag
			if tmp >= 10{
				tmp /= 10; flag = 1
			} else {
				flag = 0
			}
			num3 = append(num3, tmp)
			continue
		}
		tmp = num1[len(num1) - i - 1] + num2[len(num2) - j - 1] + flag

		if tmp >= 10{
			tmp /= 10; flag = 1
		} else {
			flag = 0
		}
		num3 = append(num3, tmp)
	}
	for i := 0; i < len(num3); i ++{
		l3.Val = num3[len(num3) -i -1]
		l3 = l3.Next
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a,b []int
	for p:=l1;p!=nil;p=p.Next  {
		a = append(a,p.Val)

	}
	for p:=l2;p!=nil;p=p.Next {
		b = append(b,p.Val)
	}
	var res *ListNode
	carry := 0
	aLen,bLen := len(a),len(b)
	for i,j:=0,0; i<aLen || j<bLen; i,j=i+1,j+1{
		sum := carry
		if i< aLen{
			sum += a[aLen-i-1]
		}
		if j< bLen{
			sum += b[bLen-j-1]
		}
		node := ListNode{Val:sum%10,Next:nil}
		if res == nil{
			res = &node
		}else{
			node.Next = res
			res = &node
		}
		carry = sum / 10
	}
	//还需要考虑最高位的进制情况
	if carry != 0{
		res = &ListNode{Val:carry,Next:res}
	}
	return res
}


type ListNode1 struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	List *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var a MyLinkedList
	return a
}

func (this *MyLinkedList) AddAtHead(val int) {
	p := new(ListNode)
	p.Val = val
	p.Next = this.List
	this.List = p
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	p := this.List
	if p == nil || index < 0 {
		return -1
	}
	num := 0
	for p != nil {
		if num == index {
			return p.Val
		}
		num++
		p = p.Next
	}
	return -1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	t := this.List
	for t.Next != nil {
		t = t.Next
	}
	p := MyLinkedList{new(ListNode)}
	p.List.Val = val
	p.List.Next = nil
	t.Next = p.List
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if this.List != nil {
		pre := MyLinkedList{new(ListNode)}.List
		num := 0
		p := this.List

		for p != nil {
			if num == index {
				pre.Next = MyLinkedList{new(ListNode)}.List
				pre.Next.Val = val
				pre.Next.Next = p
				return
			}
			num++
			pre = p
			p = p.Next
		}

		if num == index {
			this.AddAtTail(val)
		}

	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.List = this.List.Next
	} else if index > 0 {
		num := 0
		p := this.List
		pre := MyLinkedList{new(ListNode)}.List
		for p != nil {
			if index == num {
				if p.Next == nil {
					pre.Next = nil
				} else {
					pre.Next = p.Next
					pre.Next.Val = p.Next.Val
				}

			}
			num++
			pre = p
			p = p.Next
		}
	}
}

func SplitListToParts(root *ListNode, k int) []*ListNode {
	var res []*ListNode
	sum := 0
	head := root
	if root == nil {
		for  i  := 0; i < k; i ++ {
			res = append(res, nil)
			continue
		}
		return res
	}
	for root != nil {
		sum++
		root = root.Next
	}
	remainder := sum % k
	cur, pre := head, head
	for i  := 0; i < k; i ++{
		base := sum/k
		if remainder != 0{
			//循环，知道base == -1, 放入res
			for pre != nil{
				base --
				if base == -1 {
					tmp := pre.Next
					pre.Next = nil
					res = append(res, cur)
					pre = tmp
					cur = tmp
					remainder --
					break
				}
				pre = pre.Next
			}
			continue
		}
		if remainder == 0 {
			if pre == nil && base == 0  {
				res = append(res, nil)
				continue
			}
			for pre != nil {
				base --
				if base == 0 {
					tmp := pre.Next
					pre.Next = nil
					res = append(res, cur)
					pre = tmp
					cur = tmp
					break
				}
				pre = pre.Next
			}
			continue
		}
		if pre.Next == nil {
			res = append(res, nil)
			continue
		}
	}
	return res
}

func numComponents(head *ListNode, G []int) int {
	mp := make(map[int]int)
	for _,v:=range G{
		mp[v]++
	}
	res:=0
	flag:=0
	for head!=nil{
		if mp[head.Val]>0{
			if flag==0{
				res++
				flag=1
			}
		}else{
			flag=0
		}
		head=head.Next
	}
	return res
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{0,head}
	slow, fast := pre, pre

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast == nil {
		return slow
	}
	if fast.Next == nil {
		return slow.Next
	}
	return nil
}

func nextLargerNodes(head *ListNode) []int {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	stack, ans := make([]int, len(data)), make([]int, len(data))

	for index := 0; index < len(data); index++ {
		if len(stack) == 0 {
			stack = append(stack, index)
		} else {
			for len(stack) > 0 && data[index] > data[stack[len(stack)-1]] {
				pop := len(stack) - 1
				ans[stack[pop]] = data[index]
				stack = stack[:pop]
			}
			stack = append(stack, index)
		}
	}
	return ans
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil{
		return nil
	}
	pointA := headA
	pointB := headB
	for pointA != pointB {
		if pointA == nil {
			pointA = headB
		} else {
			pointA = pointA.Next
		}

		if pointB == nil {
			pointB = headA
		} else {
			pointB = pointB.Next
		}
		if pointA ==nil && pointB == nil {
			return nil
		}
	}
	return pointA
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{0,head}
	slow, fast := pre.Next, pre.Next
	for fast != nil && slow != nil{
		if fast.Next == nil {
			if fast.Val >= x{
				break
			} else {
				changeValue(fast, slow)
				slow = slow.Next
				for slow.Val >= fast.Val{
					changeValue(fast, slow)
					slow = slow.Next
				}
				break
			}
		}
		fast = fast.Next
		if fast.Val >= x{
			fast = fast.Next
		} else {
			changeValue(fast, slow)
			slow = slow.Next
			for slow.Val >= fast.Val{
				changeValue(fast, slow)
				slow = slow.Next
			}
		}
	}
	return pre.Next
}

func changeValue(bigger *ListNode, smaller *ListNode)  {
	tmp := bigger.Val
	bigger.Val = smaller.Val
	smaller.Val = tmp
}

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return -1
	}
	fast, slow := head, head
	for i := 0; i < k ; i ++{
		fast = fast.Next
	}
	for fast!= nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

type Node struct {
	Val int
	Next *Node
	Random *Node
}
var (
	visit = make(map[*Node]*Node)
)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	if newNode, ok := visit[head]; ok{
		return newNode
	}
	copy := &Node{
		Val:head.Val,
	}
	visit[head] = copy
	copy.Next = copyRandomList(head.Next)
	copy.Random = copyRandomList(head.Random)
	return copy
}



func dfs(hd *Node) *Node {
	if hd == nil {
		return nil
	}
	if newNode,ok := visit[hd];ok {
		return newNode
	}
	copy := &Node{
		Val: hd.Val,
	}
	visit[hd] = copy
	copy.Next = dfs(hd.Next)
	copy.Random = dfs(hd.Random)
	return copy
}


func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow, point := head, head, head
	for {
		if slow == nil || fast == nil || slow.Next == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			break
		}
	}
	for point != nil {
		if slow == point {
			return slow
		}
		point = point.Next
		slow = slow.Next
	}
	return nil
}
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head

	for slow != nil {
		fast := slow
		for fast.Next != nil {
			if slow.Val == fast.Next.Val {
				fast.Next = fast.Next.Next
			} else {
				fast = fast.Next
			}
		}
		slow = slow.Next
	}
	return head
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	hair := &ListNode{0,head}
	slow, fast := hair, hair.Next

	for slow != nil && fast != nil {
		sum := 0
		for fast != nil {
			sum += fast.Val
			if sum == 0{
				break
			}
			fast = fast.Next
		}

		if sum ==0 {
			fast = fast.Next
			slow.Next = fast
		} else{
			slow = slow.Next
			fast = slow.Next
		}
	}
	return hair.Next

}
























