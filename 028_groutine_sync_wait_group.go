package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const most_delay_millisecond int = 99

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func Print0(repeat int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < repeat; i++ {
		time.Sleep(time.Duration(TrueRandom(most_delay_millisecond)) * time.Millisecond)
		fmt.Printf("0")
	}
	fmt.Println("end of print0")
}
func Print1(repeat int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < repeat; i++ {
		time.Sleep(time.Duration(TrueRandom(most_delay_millisecond)) * time.Millisecond)
		fmt.Printf("1")
	}
	fmt.Println("end of print1")
}

func main() {

	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(2)

	go Print0(99, wg)
	go Print1(99, wg)
	wg.Wait() // 等待所有 goroutine 完成
}
