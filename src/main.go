package main

import (
	"algorithm-learning/src/fenzhisuanfa"
	"fmt"
)

func main() {
	fmt.Println("hello")

	//a := []int64{2, 3, 8, 6, 1}
	//GetTheInvertion(a)

	//子数组A[9...16]含有序列<2,4,5,7,1,2,3,6>   45, 33, 12, 2, 4
	//a := []int64{1, 3, 99, 50, 25, 45, 33, 12, 2, 4, 5, 7, 1, 2, 3, 6, 10, 9, 5}
	//idx := jianzengxingsuanfa.Partition(a, 5, 9)
	//fmt.Printf("下标：%d\n", idx)
	//fmt.Println(a)

	//var a byte = 'A'
	//var b byte = 'B'
	//var c byte = 'C'
	//var n int64 = 2
	//fmt.Printf("%d阶的汉诺塔总共需要移动:%d次", n, fenzhisuanfa.HanNuoTa(n, a, b, c))

	//arr := []int64{1, 3, 99, 50, 25, 45, 33, 12, 2, 4, 5, 7, 1, 2, 3, 6, 10, 9, 5}
	//fenzhisuanfa.QuickSort(arr, 0, 18)
	//fmt.Println(arr)

	arr := []int64{1, 3, 99, 50, 25, 45, 33, 12, 2, 4, 5, 7, 1, 2, 3, 6, 10, 9, 5}
	fmt.Println(fenzhisuanfa.FindTargetSortNumber(arr, 0, int64(len(arr)-1), 5))

}
