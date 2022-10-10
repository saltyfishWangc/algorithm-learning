package fenzhisuanfa

import "algorithm-learning/src/jianzengxingsuanfa"

// QuickSort 快速排序
/**
	快速排序采用的是分区算法，每次逻辑是对序列进行分区，得到区间点的下标，这样就可以将序列以区间点分成两个子序列，下标小于该区间点的一定都是小于等于
改值的，下标大于该区间点的一定都是大于等于该值的。所以每次递归的逻辑是 划分序列得到区间点，以区间点得到两个子序列的起始下标和结束下标，继续递归调用。
每次划分区间后该区间点的位置就是最终在整个序列中的位置，相当于每次划分区间就确定了一个元素。
	时间复杂度：O(nlgn)

	eg:
	arr := []int64{1, 3, 99, 50, 25, 45, 33, 12, 2, 4, 5, 7, 1, 2, 3, 6, 10, 9, 5}
	fenzhisuanfa.QuickSort(arr, 0, 18)
*/
func QuickSort(arr []int64, p, r int64) {
	/**
	在对一个序列进行划分时，可能第一个元素就是区间点元素，那么该下标就和p相等，在进行下一次递归时，
	q - 1就小于了p，对应着就是r < q，所以这种情况就是不再递归的判断条件
	*/
	if p < r {
		q := jianzengxingsuanfa.Partition(arr, p, r)
		/**
		在序列arr中，下标小于q的元素都是小于等于arr[q]，
		下标大于q的元素都是大于等于q，所以递归时q就不再参与了，它在序列中的位置就是q，不用改变了
		*/
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}
