package datastruct

/**
树状数组，方便计算前1项到i项的和。
数组0保留位。
*/
type BIT struct {
	arra []int
}

//返回x最小为1的比特位和后面的0  比如8(1000) 返回的就是8
func lowbit(x int) int {
	return x & (-x)
}

func NewBIT(n int) *BIT {
	return &BIT{make([]int, n+1)}
}

//求1~i位的和
func (b *BIT) Query(i int) int {
	ret := 0
	for i > 0 {
		ret += b.arra[i]
		i -= lowbit(i)
	}
	return ret
}

//修改第i位的值
func (b *BIT) Update(i, v int) {
	for i < len(b.arra) {
		b.arra[i] += v
		i += lowbit(i)
	}
}
