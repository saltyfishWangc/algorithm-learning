package main

import "fmt"

// GetTheInvertion 统计逆序对数
/**
一般地，设A[1...n]是一个具有n个不同元素的数组。若i<j且A[i]>A[j]，则数偶(i,j)称为A的一个逆序。例如，序列A=<2,3,8,6,1>中
<1,5>,<2,5>,<3,5>,<4,5>,<3,4>是5个逆序，这是因为A[1]=2>1=A[5],A[2]=3>1=A[5],A[3]=8>1=A[5], A[4]=6>1=A[5],A[3]=8>6=A[4]

题目形式化表示为：
输入：数组A=<A[1],A[2],,,A[n]>
输出：A的逆序个数

判断是否是逆序对的逻辑：
从第2个数开始遍历数组，依次和前面的数比较，如果满足前面的数大于当前数，则逆序对总数加1，直到前面的数比较完，再开始下一轮
时间复杂度为 O(n^2)

eg:
	a := []int64{2, 3, 8, 6, 1}
	GetTheInvertion(a)
*/
func GetTheInvertion(arr []int64) {

	invertionCount := 0

	invertionStrFormat := "arr[%d]=%d>%d=arr[%d]\n"
	var invertionStrByte []byte
	for j := 1; j < len(arr); j++ {
		for i := j - 1; i >= 0; i-- {
			if arr[i] > arr[j] {
				invertionCount += 1
				invertionStrByte = fmt.Appendf(invertionStrByte, invertionStrFormat, i, arr[i], arr[j], j)
			}
		}
	}
	fmt.Printf("一共有: %d个逆序对，分别是：\n%s", invertionCount, string(invertionStrByte))

}
