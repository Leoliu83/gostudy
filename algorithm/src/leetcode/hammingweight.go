package leetcode

/*
原题：https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
二进制中1的个数
请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。
例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
*/
func HammingWeight(num uint32) int {
	var rs uint32 = 0
	for i := num; i > 0; i >>= 1 {
		rs += (i & 1)
	}
	//log.Println(rs)
	return int(rs)
}
