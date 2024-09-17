package statistic

import (
	"fmt"
	"testing"
)

func TestStatistic(t *testing.T) {
	var nums = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var n_avg = Avg(nums)
	fmt.Println(n_avg)
	if n_avg == 5.0 {
		t.Log("test statistic success")
	} else {
		t.Error("test statistic failed")
	}
}

func TestStatistic2(t *testing.T) {
	var nums = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var n_res = STD(nums)

	if n_res == 3.0276503540974917 {
		t.Log("test statistic success")
	} else {
		t.Error("test statistic failed")
	}
}
