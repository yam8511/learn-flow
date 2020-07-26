package demo

// 宣告函式: public 外部可以引用
func Sum2Num(a, b int) (sum int) {
	sum = sumNum(a, b)
	return
}

// 宣告函式: private 外部無法引用
func sumNum(a, b int) (sum int) {
	sum = a + b
	return
}

// 宣告函式: 可選參數
func SumNum(a int, b ...int) (sum int) {
	sum += a
	for _, n := range b {
		sum += n
	}
	return
}
