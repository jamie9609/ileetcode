package main

import (
	"fmt"
	level1 "ileetcode/level1"
)

func main(){

	/*var num  = []int{1,2,3,4,5,10,0,3,1}
	results := level1.TwoSum(num,5)*/

	/*var num = []int{2,5,1,3,4,7}
	var target =3
	results:= level1.Shuffle(num, target)*/

	str := "abcdefg"
	n := 2
	results := level1.ReverseLeftWords(str, n)

	fmt.Printf("结果是：%v" , results)
}