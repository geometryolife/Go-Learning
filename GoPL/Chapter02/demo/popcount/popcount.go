// 定义一个 PopCount 函数，它返回一个数字中被置位的个数，即在一个
// uint64 的值中，值为1的位的个数，这称为种群统计。它使用 init 函
// 数来针对每一个可能的8位值预计算一个结果表 pc，这样 PopCount 只
// 需要将8个快查表的结果相加而不用进行64步计算。(这个不是最快的统
// 计位算法，只是方便用来说明 init 函数，用来展示如何预计算一个数
// 值表，它是一种很有用的编程技术。)
package popcount

// pc[i] 是 i 的种群统计
var pc [256]byte

// init 中的 range 循环只使用索引；值不是必须的，所以没必要包含进
// 来。循环可以重写为下面的形式:
// for i, _ := range pc {
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 返回 x 的种群统计(置位的个数)
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
