package main

import (
	"fmt"
	"sync"
	"time"
)

var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func sync_wait_groups_style() {
	wg := sync.WaitGroup{}
	for _, num := range numbers {
		//并发匿名函数
		wg.Add(1)
		go func() {
			fmt.Println("A", num)
			wg.Done()
		}()
	}
}

func sync_wait_groups_style_trap() {
	wg := sync.WaitGroup{}
	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			fmt.Println("B", num)
			wg.Done()
		}(num)
	}
}

func manul_counter_sytle() {
	var goroutine_count int

	for _, num := range numbers {
		goroutine_count++
		go func(num int) {
			fmt.Println("C", num)
			goroutine_count--
		}(num)
	}
	for 0 != goroutine_count {
	}
}

func main() {
	// sync_wait_groups_style()
	// sync_wait_groups_style_trap()
	manul_counter_sytle()
	time.Sleep((5 * time.Second))
}
