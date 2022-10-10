package jianzengxingsuanfa

import "fmt"

// MergeSortedSeq 合并有序序列
/**
问题的形式化表示：
输入：序列A[p...r]。其中，自序列A[p...q]和A[q+1...r]是有序的
输出：A[p...r]所有元素的重排，使之有序。
合并有序序列的逻辑：
	用一个渐增型算法解决此问题。把A[p...q]和A[q+1...r]分别复制到序列L[1...n1]和R[1...n2]，其中n1=q-p+1，n2=r-q。
然后维护3个变量i、j、k，i、j初始化为1，k初始化为p。比较L[i]和R[j]，将较小者复制到A[k]，调整i、j、k使得它们各自指向L、R
和A的合适的位置；若L[i]<R[j]，则i增加1，否则j增加1。无论如何k都要增加1，循环往复，直至L和R之一被扫描完。然后，将另一个序
列中尚存的元素复制到A[k...r]。在此过程中，A[p...k-1]中的元素是A[p...r]中最小的k-p个元素，并且已排好序。随着k的增长，
A[p...k]渐增为一个有序序列.
时间复杂度为o(n)，因为两个子序列中的每个元素都只会被遍历一次，不存在第二个序列里面随着第一个序列的轮询都遍历一次

eg:
	子数组A[9...16]含有序列<2,4,5,7,1,2,3,6>时，A[9...12]和A[13...16]为两个有序的子序列
*/
func MergeSortedSeq(arr []int64, p int64, q int64, r int64) {

	if q-r >= 0 {
		panic("q >= r 会造成arr[q]下组越界")
	}
	if p == q {
		// p 不能等于 q，否则会都多出一个arr[p]的元素
		// L[p]和R[q]起始是同一个元素，最后都写到结果序列了
		return
	}
	// 注意：这样创建切片L、R，它们和arr还是共用的同一块内存，改变arr的同时,对应位置的L、R也变了
	//L := arr[p:q]
	//R := arr[q:r]
	replic := make([]int64, len(arr))

	copy(replic, arr)
	L := replic[p:q]
	R := replic[q:r]

	i := 0
	j := 0

	k := p
	for ; i < len(L); i++ {
		for ; ; j++ {
			if j < len(R) {
				if L[i] < R[j] {
					arr[k] = L[i]
					k++
					fmt.Println(arr[0:k])
					break
				} else {
					arr[k] = R[j]
					k++
					fmt.Println(arr[0:k])
				}
			} else {
				// 表示R被遍历完了，数据都合并到arr了，如果这里是是使用break，那么在for循环里i就会加1，跳过了这次
				goto CopyFinish
			}
		}

	}

CopyFinish:
	for ; k < r; k++ {

		for ; j < len(R); j++ {
			arr[k] = R[j]
		}

		for ; i < len(L); i++ {
			arr[k] = L[i]
		}
	}

}
