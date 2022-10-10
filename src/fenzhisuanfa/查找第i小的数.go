package fenzhisuanfa

import (
	"algorithm-learning/src/jianzengxingsuanfa"
	"fmt"
)

// FindTargetSortNumber 查找第i小的数
/**
arr：表示目标序列
p：表示查找序列范围的第一个元素下标
r: 表示查找序列范围的最后一个元素下标
i: 表示需要查找的在序列arr中第i小的数，在序列arr中对应的下标是i-1

eg:
	arr := []int64{1, 3, 99, 50, 25, 45, 33, 12, 2, 4, 5, 7, 1, 2, 3, 6, 10, 9, 5}
	fmt.Println(fenzhisuanfa.FindTargetSortNumber(arr, 0, int64(len(arr)-1), 5))
*/
func FindTargetSortNumber(arr []int64, p, r, i int64) int64 {
	// 先对arr进行分区，得到分区点
	partitionIndex := jianzengxingsuanfa.Partition(arr, p, r)
	// 将分区点和i-1比较,，如果相等，则表示partitionIndex下标对应的数就是第i小的数
	if partitionIndex == (i - 1) {
		fmt.Printf("下标为%d的数就是第%d小的数\n", partitionIndex, i)
	} else if partitionIndex > i-1 {
		// 将0，partitionIndex - 1作为序列区间进行分区
		partitionIndex = FindTargetSortNumber(arr, p, partitionIndex-1, i)
	} else {
		// 将partitionIndex + 1，len(arr) - 1作为序列区间进行分区
		partitionIndex = FindTargetSortNumber(arr, partitionIndex+1, r, i)
	}
	return arr[partitionIndex]
}
