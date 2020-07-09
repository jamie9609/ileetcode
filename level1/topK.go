package level1

//快速排序法 从小到大排

func QuickSort(arr []int){
	head := 0
	tail := len(arr) - 1
	flag := 0
	if tail <= head {
		return
	}
	for head < tail {
		for head < tail{
			if arr[tail] >=  arr[flag] {
				tail --
				continue
			}
			arr[tail], arr[flag] = arr[flag], arr[tail]
			flag = tail
			break

		}
		for head < tail {
			if arr[head] <=  arr[flag]{
				head ++
				continue
			}
			arr[head], arr[flag] = arr[flag], arr[head]
			flag = head
			break
		}
	}
	QuickSort(arr[:flag])
	QuickSort(arr[flag+1:])
}

//快速排序法 从大到小排

func QuickSortDesc(arr []int){
	head := 0
	tail := len(arr) - 1
	flag := 0
	if tail <= head {
		return
	}
	for head < tail {
		for head < tail{
			if arr[tail] <=  arr[flag] {
				tail --
				continue
			}
			arr[tail], arr[flag] = arr[flag], arr[tail]
			flag = tail
			break

		}
		for head < tail {
			if arr[head] >=  arr[flag]{
				head ++
				continue
			}
			arr[head], arr[flag] = arr[flag], arr[head]
			flag = head
			break
		}
	}
	QuickSortDesc(arr[:flag])
	QuickSortDesc(arr[flag+1:])
}


func FindKthLargest(nums []int, k int) int {
	QuickSortDesc(nums)
	return nums[k-1]
}