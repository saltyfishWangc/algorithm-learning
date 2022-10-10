package jianzengxingsuanfa

// Partition 根据要求，将序列的下标p到r的子序列按照arr[r]进行排序，使得小于arr[r]的在左侧，大于arr[r]的在右侧，并得到最终arr[r]元素所在的下标
/**
	题目的形式化表示为：
	输入：序列A[p...r]
	输出：下标q(p<q<r)，原序列A[p...r]的一个重排：使得A[p...q]中的元素值不超过A[q](=原A[r])的值，而A[q+1...r]中的元素值均大于A[q]。
	解决此问题的算法对序列A[p...r]进行原地重组。算法选择元素x=A[r]作为主元素，以它为分界点对序列A[p...r]进行划分，目标是将其分成前后两段
A[p...q]和A[q+1...r]，使得A[p...q]中的元素值不超过x，而A[q+1...r]中的元素值大于x。算法维护两个下标值i和j，初始时分别为p-1和p。让j在[p...r]
中扫描，若A[j]<A[r]，则将A[j]和A[i+1]交换，然后i增加1。这样，在此过程中，A[p...i]中的元素均不超过A[r]，而A[i+1,j]中的元素大于A[r]。随着j的增加，
A[p...i]和A[i+1...j]也随之增长，最终j达到r-1，并将A[i+1]与A[r]交换，返回i+1即为所求的q。
	时间复杂度：假定序列A[p...r]中含有n个元素，在此过程中，针对j的for循环重复了n次，所以该算法的最坏情形运行时间为o(n)，即时间复杂度为o(n)
	eg:
		[1, 3, 99, 50, 25, 45, 33, 12, 2, 4, 5, 7, 1, 2, 3, 6, 10, 9, 5], 下标p,r分别是5，9，则是对45, 33, 12, 2, 4进行分序
*/

func Partition(arr []int64, p int64, r int64) (q int64) {
	i := p - 1
	j := p
	for ; j < r; j++ {
		if arr[j] < arr[r] {
			// 此时
			arr[i+1], arr[j] = arr[j], arr[i+1]
			i++
			// i记录的是在p-r这段序列中值小于arr[r]的最后一个元素的下标
		}
	}
	// 最后一定要将下标为i+1的元素与arr[r]互换
	/**
	因为i记录的是每次对比后当前小于arr[r]的最后一个元素的下标，所以如果当j=r-1时，此时遍历完了，
	那么arr[i+1]>arr[i]且arr[i+1]>arr[r]一定是成立的，所以将arr[i+1]与arr[r]互换，
	就满足了以arr[r]为分界线，小于它的在左侧，大于它的在右侧
	*/
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
