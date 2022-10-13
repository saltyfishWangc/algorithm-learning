package fenzhisuanfa

// HeapSort 堆排序
/**
	利用最大堆的特性(A[1]为最大值)可以对数组A进行快速排序：首先使用BuildHeap在输入数组A[1...n]上建立一个最大堆，
其中n=len(A)。由于数组中的最大元素存于根A[1]，可以通过将其与A[n]交换而是其至于数组最终的位置。然后，从堆中"放弃"
节点n(通过减少heap-size(A))，此时，根A[1]的孩子们仍然是最大堆，但根元素可能违背了最大堆性质。调用Heapify(A, 1),
它使得A[1...(n-1)]为一个最大堆。接下来，算法对此规模为n-1到2的堆重复此过程。

	算法的运行时间：在HeapSort过程中，第1行调用BuildHeap耗时O(nlgn)，第2～5行的for循环重复n-1次，每次重复调用Heapify(A,1)，
耗时O(lgn)。所以，该循环耗时O(nlgn)，并且过程的总耗时为O(nlgn)。

	eg:
		arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
		fenzhisuanfa.HeapSort(arr)
*/
func HeapSort(arr []int) {
	heapSize := len(arr)
	BuildHeap(arr)
	for i := heapSize - 1; i > 0; i-- {
		// 将第一个和最后一个替换
		arr[0], arr[i] = arr[i], arr[0]
		// 因为每次循环后当前范围内最大的元素都是最后一个，是排好序的，所以需要将最后一个剔除，继续对剩下的进行循环
		heapSize--
		Heapify(arr, 0, heapSize)
	}
}
