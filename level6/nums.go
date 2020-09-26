package level6

import (
	"math"
)

func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B)== 0 {
		return 0
	}
	dp := make([][]int, len(A) + 1)
	for i := 0; i <len(dp); i ++ {
		dp[i] = make([]int, len(B)+1)
	}

	maxL := math.MinInt64
	for i := 1; i <= len(A); i ++ {
		for j := 1; j <=len(B); j ++ {
			if A[i - 1] == B[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxL {
				maxL = dp[i][j]
			}
		}
	}
	return maxL
}

func isValid(s string) bool {
	var stack = make([]byte, 0)
	var bracketsMap = map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	for i := 0; i < len(s); i ++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{'{
			stack = append(stack, s[i])
		} else {
			if val, ok := bracketsMap[s[i]]; ok {
				if val == stack[len(stack) - 1] && len(stack) >= 1 {
					stack = stack[:len(stack) - 1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}


func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs == nil {
		return ""
	}
	flag := 0
	ct := 0
	res := ""
	for flag >= 0 {
		for i := 0; i < len(strs) ; i ++ {
			if flag > len(strs[i]) -1 {
				ct = 1
				break
			}
			tmp := strs[0][flag]

			if strs[i][flag] == tmp {
				ct = 2
			} else {
				ct = 1
				break
			}
		}

		if ct == 1 {
			break
		} else if ct == 2 {
			res = res + string(strs[0][flag])
		}

		flag ++
	}
	return res
}


type KthLargest struct {
	size int	//堆的容量，不包括数组第一个元素
	data []int
	count int 	//当前元素数量

}


func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{}
	kth.data = []int{0}
	kth.size = k
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}


func (this *KthLargest) Add(val int) int {
	if this.count < this.size -1 {
		this.data = append(this.data, val)
		this.count += 1
	} else if this.count == this.size-1 {
		this.data = append(this.data, val)
		this.count += 1

		//第一次填满k个数，使堆有序，后续add只需在data[1]覆盖放入新值。
		n:=len(this.data)-1
		for i := n/2; i >= 1; i-- {
			heapify(this.data, i)
		}
	}else {
		if val > this.data[1] {
			this.data[1]=val
			heapify(this.data, 1)
		}
	}
	return this.data[1]
}

//heapify 从给定的i向下堆化为小顶堆
func heapify(a []int, i int) {
	n := len(a) - 1
	for {
		minPos := i
		if i*2 <=n && a[i*2] < a[minPos] {
			minPos=i*2
		}
		if i*2+1 <= n && a[i*2+1] < a[minPos] {
			minPos=i*2+1
		}
		if minPos == i {
			break
		}
		a[minPos], a[i] = a[i], a[minPos]
		i = minPos
	}
}

func getLeastNumbers(arr []int, k int) []int {
	if k==0{
		return []int{}
	}

	heap := arr[:k:k]
	for i:=(k-1)/ 2; i >=0; i--{
		maxHeap1(heap ,i,  k-1)
	}

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			maxHeap1(heap, 0, k-1)
		}
	}

	return heap
}

//构建大顶堆的第一步，构建交换大顶堆的基础
func maxHeap1(arr []int, start int, end int)  {
	root := start
	for true {
		child := root * 2 + 1

		if child > end {
			break
		}

		for child + 1 <= end && arr[child + 1] > arr[child] {
			child ++
		}
		if arr[child] > arr[root] {
			arr[child], arr[root] = arr[root], arr[child]
			root = child
		} else {
			break
		}
	}
	return
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || matrix == nil {
		return 0
	}
	res := make([]int, 0)
	for i := 0 ; i < len(matrix) ; i ++ {
		for j := 0; j < len(matrix[i]); j ++ {
			res = append(res, matrix[i][j])
		}
	}

	heap := res[:k]

	for i := (len(heap) - 1) / 2; i >= 0; i -- {
		maxHeap1(heap, i, k - 1)
	}

	for i := k ; i < len(res) ; i ++ {
		if res[i] < heap[0] {
			heap[0] = res[i]
			maxHeap1(heap, 0, k - 1)
		}
	}
	return  heap[0]
}