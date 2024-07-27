package main

type MyFloat float64

// funcとメソッドの間にあるのがレシーバ引数
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}