package fenzhisuanfa

import "algorithm-learning/src/jianzengxingsuanfa"

// MergeSort 归并排序
/**
算法思想：
将序列arr[p:r]是一段无序的序列，将这段序列分成两个子序列，各自做排序，然后在合并两个子序列，最终就是一个有序的序列
归并排序是一个典型的分治算法：将序列A[p...r]分成两半A[p...q]和A[q+1...r]，分别递归地对这两个子序列排序，将得到的排好序的子序列
A[p...q]和A[q+1...r]调用合并有序序列jianzengxingsuanfa.MergeSortedSeq(arr []int64, p int64, q int64, r int64)合并成
整个有序序列A[p...r]
时间复杂度为O(nlgn)，其中jianzengxingsuanfa.MergeSortedSeq的时间复杂度为o(n)，最终也比插入排序(O(n^2))要有效

注意：递归思想里面一定要想清楚满足申请条件时就要跳出递归

eg:
	数组A=<5,2,4,7,1,3,2,6,9>
	输入参数：
		arr：待排序的序列
		p：需要被排序子序列的第一个元素的下标
		r：需要被排序子序列的最后一个元素的下个下标(因为不包含该下标的元素，所以需要指向下一个)
	例如：需要对A={5,2,4,7,1,3,2,6,9}中的{2,4,7,1,3,2}子序列进行归并排序，最后A={5,1,2,2,3,4,7,6,9}，所以p=1,r=7
	a := []int64{5, 2, 4, 7, 1, 3, 2, 6, 9}
	fenzhisuanfa.MergeSort(a, 0, 7)
*/
func MergeSort(arr []int64, p int64, r int64) {
	if p < r {
		q := (p + r) / 2
		/**
		不能无限的递归拆分序列，最终会栈溢出，经过分析，
		当arr[p:r]序列只包含两个元素时就不能再继续拆分递归了，
		而是应该进行合并有序序列得到一个有序的两个元素的序列

		分析思维图可以参考：https://processon.com/diagraming/632888f2637689341d661bda
		*/
		if p+2 < r {
			MergeSort(arr, p, q)
			MergeSort(arr, q, r)
		}
		jianzengxingsuanfa.MergeSortedSeq(arr, p, q, r)
	}
}
