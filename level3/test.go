package level3

type NodeList struct {
	Val 	int64
	Next 	*NodeList
}

func Marge(firstHead *NodeList,  secondHead *NodeList) *NodeList {
	h1 := &NodeList{0,firstHead}
	h2 := &NodeList{0,secondHead}
	flag1 := h1.Next
	flag2 := h2.Next
	for flag1 != nil && flag2 != nil {
		//如果flag1.Val <  flag2.Val，跳到下一个
		if flag1.Val >= flag2.Val {
			h1.Next = h2.Next
			flag2 = flag2.Next
			if flag2.Val >= flag1.Val {
				h1.Next.Next = flag1
				flag1 = flag1.Next
			}
			h1.Next.Next = flag2
			flag2 = flag2.Next
		}
		flag1 = flag1.Next
	}
	if flag1 == nil && flag2 != nil{
		flag1.Next = flag2
	}
	if flag2 == nil && flag1 != nil {
		flag2.Next = flag1
	}
	return h1
}


type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

//从小到大
func searchNums (nums []int, target int) int{
	left, mid := 0, 0
	right := len(nums)
	if len(nums) == 0 {
		return -1
	}
	for mid != right && mid != left {
		if nums[mid] > nums[target] {
			right = mid
			mid = (left + right) /2
			continue
		}
		if nums[mid] < nums[target] {
			left = mid
			mid = (left + right) /2
			continue
		}
		return mid
	}
	// 返回没有找到的错误码
	return -1
}
