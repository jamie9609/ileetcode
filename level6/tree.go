package level6

import "math"

type TreeNode struct {
	Val  		int
	Left 		*TreeNode
	Right 		*TreeNode
}
func findMode(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	var max int
	var res []int
	m := make(map[int]int)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]
		m[node.Val] ++
		if max < m[node.Val] {
			max = m[node.Val]
			res = nil
			res = append(res, node.Val)
		}else if max == m[node.Val] {
			res = append(res, node.Val)
		}
		root = node.Right
	}
	return res
}


func getMinimumDifference(root *TreeNode) int {

	stack := make([]*TreeNode, 0)
	minValue := math.MaxInt64
	preValue := minValue

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]
		abs := getAbs(preValue, node.Val)
		if abs < minValue{
			minValue = abs
		}
		preValue = node.Val
		root = node.Right
	}
	return minValue
}

func getAbs(a int,b int) int {
	if a-b>0{
		return a-b
	}else{
		return b-a
	}
}


func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	node := root
	queue = append(queue, node)
	res = append(res, node.Val)
	for len(queue) != 0 {
		tmpQueue := make([]*TreeNode, 0)
		maxNum := math.MinInt64
		for i := 0; i < len(queue); i ++ {
			if queue[i].Left != nil {
				tmpQueue = append(tmpQueue, queue[i].Left)
				maxNum = maxCompare(maxNum, queue[i].Left.Val)
			}
			if queue[i].Right != nil {
				tmpQueue = append(tmpQueue, queue[i].Right)
				maxNum = maxCompare(maxNum, queue[i].Right.Val)
			}
		}
		if len(tmpQueue) != 0 {
			queue = tmpQueue
			res = append(res, maxNum)
		} else {
			queue = tmpQueue
		}
	}
	return res
}

func maxCompare(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func isValidBST1(root *TreeNode) bool {
	return helper8(root, math.MinInt64, math.MaxInt64)
}

func helper8(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper8(root.Left, lower, root.Val) && helper8(root.Right, root.Val, upper)
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	node := root

	res := make([]int, 0)

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		tmpNode := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, tmpNode.Val)
		node = tmpNode.Right
	}
	return res
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	node := root

	minNum := math.MinInt64

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		tmpNode := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		node = tmpNode.Right
		if tmpNode.Val > minNum {
			minNum = tmpNode.Val
		} else {
			return false
		}
	}
	return true
}


func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := make([]*TreeNode, 0)
	node := root
	stack = append(stack, root)
	var res, preVal int
	preVal, res = math.MaxInt64, math.MaxInt64
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		tmpNode := stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]
		if tmpNode.Val > preVal {
			res = minVal(res, tmpNode.Val - preVal)
		}else {
			res = minVal(res,  preVal - tmpNode.Val)
		}
		preVal = tmpNode.Val
		node = tmpNode.Right
	}
	return res
}

func minVal(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDiffInBST(root *TreeNode) int {

	stack := make([]*TreeNode, 0)
	minValue := math.MaxInt64
	preValue := minValue

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]
		abs := getAbs(preValue, node.Val)
		if abs < minValue{
			minValue = abs
		}
		preValue = node.Val
		root = node.Right
	}
	return minValue
}

func getAbs(a int,b int) int {
	if a-b>0 {
		return a-b
	}else{
		return b-a
	}
}
