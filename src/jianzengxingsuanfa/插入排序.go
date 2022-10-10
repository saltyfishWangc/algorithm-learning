package jianzengxingsuanfa

// InsertionSort 插入排序
/**
形式化表示为：
输入：一组数<a1,a2,,,an>
输出：输入的一个排列(重排)<a'1,a'2,,,a'n>，满足有序

插入排序的逻辑：
从第二个数开始遍历数组，依次和它前一个数比较，如果比前一个数小，则在这两个位置上进行数据交换，继续和前一个数比较，直到大于前一个数或者不存在前一个数
时间复杂度为 O(n^2)

eg:
	a := []int64{34, 23, 46, 10, 10}
	InsertionSort(a)
*/
func InsertionSort(arr []int64) {
	for j := 1; j < len(arr); j++ {
		for i := j - 1; i >= 0; i-- {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
	}
}
