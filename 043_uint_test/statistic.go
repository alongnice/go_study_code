package statistic

import "math"

func Avg(nums []float64) (ans float64) {
	ans = 0.0
	for _, num := range nums {
		ans += num
	}
	return ans / float64(len(nums))
}

func STD(nums []float64) (ans float64) {
	// 自由度？
	degree_of_freedom := len(nums) - 1
	// 计算均值
	sum_of_dev_square := 0.0
	for _, n := range nums {
		sum_of_dev_square += math.Pow(n-Avg(nums), 2)
	}
	return math.Sqrt(sum_of_dev_square / float64(degree_of_freedom))
}
