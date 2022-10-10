package fenzhisuanfa

import "fmt"

var count int64 = 0

// HanNuoTa 汉诺塔问题
/**
汉诺塔是个典型的递归问题
n ：表示阶层数，也就是有多少个盘子
a : 起始柱
b : 中转柱
c : 目标柱

算法逻辑：
	如果n=1，那可以直接从起始柱转移到目标柱，用不上中转柱
	在n > 1的情况，在每次轮回中的步骤一定是如下3步：
		1.把上面的n-1个盘转移到中转柱上
		2.将最后一个盘(也就是第n个盘),从起始柱转移到目标柱
		3.开启下一轮，但是起始柱和中转柱互换了
	因为总共就只有3个柱子，每次轮回中的目标都是将最下面的盘子从起始柱转移到目标柱，
那其余的柱子必须得让道，它们也没别的柱子可以去了，所以只能去中转柱。注意；这里说的n-1个盘转到中转柱上去也并不是一步到位的，也是通过递归让n-1-1继续如此
	这就是递归的思想，在递归函数中，它做的就是每个核心步骤
*/
func HanNuoTa(n int64, a, b, c byte) int64 {
	if n == 1 {
		move(a, c)
	} else {
		HanNuoTa(n-1, a, c, b)
		move(a, c)
		HanNuoTa(n-1, b, a, c)
	}
	return count
}

func move(a, c byte) {
	fmt.Println(string(a) + "------>" + string(c))
	count++
}
