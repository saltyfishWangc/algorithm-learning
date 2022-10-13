package fenzhisuanfa

// BuildHeap 将数组创建成堆
/**
创建堆的想法是：比较A[i],A[i]的左孩子、右孩子A[r]的大小，取最大者与A[i]交换位置。这样原问题就转换成了左子树或右子树的树根
的堆性质维护问题，递归地解决子问题。其中A[i]的左孩子在数组中的下标为2*i+1，右孩子在数组中的下标为2*i+2，当然不能超过数组最
后一个元素的下标，超过则表示没有对应的孩子。

问题的形式化表示为：
输入：数组A[1...n]
输出：重排后的数组A[1...n]，元素间构成一个堆

eg:
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fenzhisuanfa.BuildHeap(arr)
*/
func BuildHeap(arr []int) {
	heapSize := len(arr)
	for i := heapSize / 2; i >= 0; i-- {
		Heapify(arr, i, heapSize)
	}
}

func Heapify(arr []int, i, heapSize int) {
	l := left(i)
	r := right(i)
	var most int

	if l < heapSize && arr[l] > arr[i] {
		most = l
	} else {
		most = i
	}

	if r < heapSize && arr[r] > arr[most] {
		most = r
	}

	if most != i {
		arr[i], arr[most] = arr[most], arr[i]
		Heapify(arr, most, heapSize)
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	if i%2 == 1 {
		return i / 2
	} else {
		return i/2 - 1
	}
}
