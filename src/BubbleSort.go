package main

import "fmt"

// BubbleSort 冒泡排序
/**
形式化表示为：
输入：一组数<a1,a2,,,an>
输出：输入的一个排列(重排)<a'1,a'2,,,a'n>，满足有序

冒泡排序的逻辑：
从第1个数开始遍历数组，依次和数组中的数比较，取两者小的数据，继续和下一个数比较，直到将本轮的数据都比较完
比如有n个数，那么会循环n次，每次取数组的第一个数依次和后面的数开始比较，每次比较n - i个数，最后得到的就是每轮最小的数
例如：数组有5个数，那么
第一轮：数组的第一个数和第二个数开始比较，一共比较4次，第一轮结束，数组的最后一个数就是整个数组最小的数
第二轮：数组的第一个数和第二个数开始比较，一共比较3次，第二轮结束，数组的倒数第二个数就是这次比较中最小的数，因为第一轮中得到的最后一个数就是整个数组的最小数，所以不需要再和上一轮的最后一个数进行比较了
时间复杂度为 O(n^2)

eg:
	a := []int64{5, 34, 23, 100, 46, 10, 10, 78}
	BubbleSort(a)
*/
func BubbleSort(arr []int64) {
	for i := 1; i < len(arr)+1; i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
		fmt.Printf("第：%d 次：结果：%v \n", i, arr)
	}
}
