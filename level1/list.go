package level1

type ListNode struct {
	Val int
	Next *ListNode
}

func DeleteNode(node *ListNode) {
	if node == nil {
		return
	}
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}


type CQueue struct {
	stack1 []int
	stack2 []int
}


/*func Constructor() CQueue {
	return CQueue{}
}*/


func (this *CQueue) AppendTail(value int)  {
	this.stack1 = append(this.stack1, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0{
			return -1
		}
		for len(this.stack1) > 0 {
			index := len(this.stack1) - 1
			value := this.stack1[index]
			this.stack2 = append(this.stack2, value)
			//用切片方式裁剪掉尾部
			this.stack1 = this.stack1[:index]
		}
	}
	index2 := len(this.stack2) - 1
	value2 := this.stack2[index2]
	this.stack2 = this.stack2[:index2]
	return value2
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

/*两个栈实现一个队列*/
type MyQueue struct {
	stack1 []int
 	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}
 
func (this *MyQueue) Push(x int)  {
	this.stack1 = append(this.stack1, x)
}
func (this *MyQueue) Pop() int {
	if len(this.stack1) == 0 {
		return -1
	}
	this.stack2 = make([]int, len(this.stack1))
	for i, j := len(this.stack1)-1, 0 ; i >= 0 ; i , j = i-1, j +1  {
		this.stack2[j] = this.stack1[i]
	}
	for i, j := len(this.stack2) -2, 0; i >= 0; i, j = i - 1, j + 1{
		this.stack1[j] = this.stack2[i]
	}
	index := len (this.stack1)
	this.stack1 = this.stack1[:index - 1]
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Peek() int {
	if len(this.stack1) == 0 {
		return -1
	}
	return this.stack1[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}


//Go解决二进制中1的个数(移位操作)

func hammingWeight(num uint32) int {
	sum := 0
	//将二进制数字num无符号位右移一位(golang中>>)
	for ; num > 0; num >>=1 {
		if 1 & num == 1 {
			sum ++
		}
	}
	return sum
}