package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

// 打印cpu核心数量
func print_num_goroutine() {
	fmt.Println("num_groutine:", runtime.NumCPU())
}

func calc_pi_by_goroutine(num_of_trem float64) {
	fmt.Println("runtime.NumCPU: ", runtime.NumCPU())
	var sqrt_pi_divie_by6 float64
	part_sum_stream := make(chan float64, runtime.NumCPU())
	//获取cpu核心数量，给到part_sum_stream
	wg := sync.WaitGroup{}

	// 创建匿名函数，计算部分
	var get_Part_sum = func(start float64, step float64, upper float64, sum_stream chan<- float64, wg *sync.WaitGroup) {
		var the_part_sum float64
		defer wg.Done()

		fmt.Println(start, step, upper)
		for i := start; i < upper; i += step {
			the_part_sum += 1.0 / (i * i)
		}
		sum_stream <- the_part_sum
	}

	// 创建cpu核心数量个goroutine
	for core_num := 1; core_num <= runtime.NumCPU(); core_num++ {
		wg.Add(1)
		go get_Part_sum(float64(core_num), num_of_trem/float64(runtime.NumCPU()), num_of_trem, part_sum_stream, &wg)
	}

	print_num_goroutine()
	wg.Wait()
	close(part_sum_stream)

	// 数据汇总
	sqrt_pi_divie_by6 = func() (sum float64) {
		for part := range part_sum_stream {
			fmt.Println("part:", part)
			sum += part
		}
		return sum
	}()

	var pi = math.Pow(sqrt_pi_divie_by6*6, 0.5)
	fmt.Println("pi:", pi)
}

func calc_pi(num_of_trem float64) {
	var sqrt_pi_divie_by6 float64

	for i := 1.0; i <= num_of_trem; i++ {
		sqrt_pi_divie_by6 += 1.0 / (i * i)
	}
	print_num_goroutine()

	var pi = math.Pow(sqrt_pi_divie_by6*6, 0.5)
	fmt.Println("pi:", pi)
}

// 计算pi
func main() {
	var trems_num float64 = 1e11
	var start_time, end_time time.Time

	// 多协程计算pi
	start_time = time.Now()
	calc_pi_by_goroutine(trems_num) //执行函数
	end_time = time.Now()
	fmt.Println("case1:", end_time.Sub(start_time))
	fmt.Println("=======================================")
	// 单协程计算pi
	start_time = time.Now()
	calc_pi(trems_num) //执行函数
	end_time = time.Now()
	fmt.Println("case2:", end_time.Sub(start_time))
}
