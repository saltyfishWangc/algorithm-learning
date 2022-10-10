package fenzhisuanfa

type Result struct {
	Res int64
}

// BinarySearch 二分查找
/**
	二分查找算法思想：
	每次检测序列的中间点与v的大小关系，并在下一步考虑中消除序列的一半。重复这一过程，每次重复将保留的序列缩小一半。这样的算法称为二分查找算法
	起始下标和结束下标之和的一半，比较和目标值的大小，如果小于目标值，说明要从大于该下标的部分继续找；大于目标值，说明要从小于该下标的部分继续找；
因为是要得到第一个等于v的下标，所以当等于目标值时，应该先记录该下标，但是还不能返回，应该将该下标作为结束下标，继续从小于该下标的部分继续找。
	时间复杂度：因为每次通过比较，只会选择一半去继续下一次递归，所以永远都只会选择一半，时间复杂度为O(lgn)

	注意：递归函数一定要return，否则，会一直递归下去

	问题的形式化表示：
	输入：排好序的序列A，数值v
	输出：若A中存在元素其值等于v，则输出从A[0]起第一个这样的元素的下标，否则输出-1表示无此元素的特殊值

	eg:
	result := &fenzhisuanfa.Result{
		Res: -1,
	}
	a := []int64{1, 2, 2, 3, 4, 5, 6, 7, 9, 9}
	fmt.Println(fenzhisuanfa.BinarySearch(a, result, 9, 0, int64(len(a)-1)))
*/
func BinarySearch(arr []int64, res *Result, v, beginIndex, endIndex int64) int64 {

	if beginIndex == endIndex {
		if arr[beginIndex] == v {
			res.Res = beginIndex
			return res.Res
		} else {
			return res.Res
		}
	} else if beginIndex > endIndex {
		return res.Res
	}
	cur := (beginIndex + endIndex) / 2

	if v == arr[cur] {
		res.Res = cur
		endIndex = cur - 1
	} else if v > arr[cur] {
		beginIndex = cur + 1
	} else {
		endIndex = cur - 1
	}
	return BinarySearch(arr, res, v, beginIndex, endIndex)
}
