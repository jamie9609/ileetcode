package level7

type TreeNode struct {
	Val 	int
	Left    *TreeNode
	Right   *TreeNode
}
type ListNode struct {
	Val		int
	Next	*ListNode
}


func ListOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	var queue []*TreeNode
	var res []*ListNode
	node := tree
	queue = append(queue, node)
	res = append(res, &ListNode{node.Val,nil})

	for len(queue) > 0{
		var tmpQueue []*TreeNode
		for i := 0; i < len(queue); i ++ {
			tmpNode := queue[i]
			if tmpNode.Left != nil {
				tmpQueue = append(tmpQueue, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				tmpQueue = append(tmpQueue, tmpNode.Right)
			}
		}
		queue = tmpQueue
		var tmpNode *ListNode
		if len(tmpQueue) > 0{
			node := tmpNode
			for i := 0; i < len(tmpQueue); i ++ {
				tmpNode = &ListNode{tmpQueue[i].Val,tmpNode}
			}
			res = append(res, node)
		}

	}
	return res
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	m := []*TreeNode{tree}
	var list []*ListNode

	for len(m) > 0 {

		var temp []*TreeNode
		node:= new(ListNode)
		tempnode :=node
		for i := 0; i < len(m); i++ {
			tempnode.Next = &ListNode{
				Val: m[i].Val,
			}
			tempnode = tempnode.Next
			if m[i].Left != nil {
				temp = append(temp, m[i].Left)
			}
			if m[i].Right != nil {
				temp = append(temp, m[i].Right)
			}
		}
		m = temp
		list = append(list, node.Next)
	}
	return list
}

