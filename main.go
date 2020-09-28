package main

import (
	"ileetcode/level5"

	"fmt"
)

func main(){

	/*var num  = []int{1,2,3,4,5,10,0,3,1}
	results := level1.TwoSum(num,5)*/

	/*var num = []int{2,5,1,3,4,7}
	var target =3
	results:= level1.Shuffle(num, target)*/

	/*str := "abcdefg"
	n := 2
	results := level1.ReverseLeftWords(str, n)*/


	//var testArr = []int{0, 2, 5, 3, 7, 4, 5, 8, 1, 19, 1, 11, -2, 20, 21, 4, -1, 3, 11, 7}
	//level1.QuickSort(testArr)
	//level1.QuickSortDesc(testArr)

	//matrix := [][]int{{1},{5},{9}}

	/*a := map[string]string{"a": "1", "b": "2", "c": "3"}
	var wg sync.WaitGroup
	for _, v := range a{
		wg.Add(1)
		go func(v string) {
			time.Sleep(5 * time.Second)
			defer wg.Done()
			fmt.Println(v)
		}(v)

		wg.Wait()
		fmt.Println("done")
	}*/


	// Node5 := level3.ListNode{5,nil}
	/*Node4 := level3. ListNode{4,nil}
	Node3 := level3.ListNode{3,&Node4}
	Node2 := level3.ListNode{2,&Node3}
	Node1 := level3.ListNode{1,&Node2}
	head := Node1
*/
	/*var s string
	s = "au"*/

	/*nums := []int{3,1}
	target := 1

	res := level5.Search(nums, target)*/

	/*res := level5.GetKthMagicNumber(20)*/

	nums := []int{1,1,2,3,4,4,4,5,5}

	stringL :="我们"
	fmt.Printf("结果是：%d" ,len(stringL))
	res := level5.RemoveDuplicates(nums)

	res2 := level5.RemoveDuplicates1(nums)

	fmt.Printf("结果是：%+v, %+v" ,res, res2)
}