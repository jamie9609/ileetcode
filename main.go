package main

import (
	"fmt"
	"ileetcode/level2"
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

	matrix := [][]int{{1},{5},{9}}
	result := level2.SpiralOrder(matrix)

	fmt.Printf("结果是：%v" , result)
}