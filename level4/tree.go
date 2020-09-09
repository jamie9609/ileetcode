package level4

type TreeNode struct {
	Val  	int
	Left    *TreeNode
	Right   *TreeNode
}

//层序遍历，使用队列。
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var tmpQueue []*TreeNode
		var tmpRes []int
		for j := 0; j < len(queue); j ++{
			node := queue[j]
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = tmpQueue
		res = append(res, tmpRes)
	}
	return res
}


//N叉树的前序遍历
/*
type Node struct {
	Val 		int
	Children 	[]*Node
}
var res []int*/

func preorder(root *Node) []int{
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *Node){
	if root != nil {
		res = append(res, root.Val)
		for _, index := range root.Children {
			dfs(index)
		}
	}
}

/*func preorder(root *Node) []int {
	var res []int
	stack := []*Node{root}

	for len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			if len(root.Children) == 0 {
				break
			}
			for i := len(root.Children) - 1; i >0 ; i++{
				stack = append(stack, root.Children[i])
			}
		}
		//出栈
		root = stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]

	}
	return res
}
*/


//迭代实现前序遍历。
func preorder(root *Node) []int {
	var res []int
	stack := []*Node{root}

	if root == nil {
		return nil
	}

	for len(stack) >0	{
		cur := stack[len(stack) - 1]
		stack = stack[ :len(stack) -1]
		res = append(res, cur.Val)
		for i := len(cur.Children) - 1; i >=0; i-- {
			stack = append(stack, cur.Children[i])
		}
	}
	return res
}

//非迭代实现中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) != 0 || root != nil {
		if root == nil {
			tmpNum := stack[len(stack) - 1]
			res = append(res, tmpNum.Val)
			stack = stack[ : len(stack) - 1]
			root = tmpNum.Right
		}else {
			stack = append(stack, root)
			root = root.Left
		}
	}
	return res
}

//二叉树的后序遍历

func postorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)
	var stack []*TreeNode
	var visitRoot *TreeNode
	var node *TreeNode
	node = root
	if root == nil {
		return res
	}
	for len(stack) > 0 || node != nil {
		for  node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		tmpNode := stack[len(stack) - 1]
		if tmpNode.Right == nil || tmpNode.Right == visitRoot{
			res = append(res, tmpNode.Val)
			stack = stack[ :len(stack) -1]
			visitRoot, tmpNode = tmpNode, nil
		} else {
			node = tmpNode.Right
		}
	}
	return res
}

//二叉树前序遍历

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var node *TreeNode
	node = root

	if root == nil {
		return res
	}
	stack = append(stack, node)
	for len(stack) > 0 || node != nil {

		res = append(res, node.Val)
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if len(stack) == 0 {
			return res
		}else {
			node = stack[len(stack)-1]
		}
	}
	return res
}

//n叉树的层序遍历

type Node struct {
	Val   			int
	Children  		[]*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		var tmpQueue []*Node
		var tmpRes []int
		for j := 0; j < len(queue); j ++{
			node := queue[j]
			tmpRes = append(tmpRes, node.Val)
			if len(node.Children) != 0 {
				for i := 0; i < len(node.Children); i ++ {
					tmpQueue = append(tmpQueue, node.Children[i])
				}
			}
		}
		queue = tmpQueue
		res = append(res, tmpRes)
	}
	return res
}