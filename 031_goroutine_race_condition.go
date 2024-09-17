package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"
)

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func rance_condition(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	// 模拟计算延时
	var do_some_thing_clac = func() {
		<-time.After(time.Duration(TrueRandom(5)) * time.Millisecond)
	}
	var data int
	mu.Lock()
	go func() {
		data++
		mu.Unlock()
	}()
	if 0 == data {
		do_some_thing_clac()
		mu.Lock()
		fmt.Println("data:", data)
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 999; i++ {
		wg.Add(1)
		go rance_condition(&wg, &mu)
	}

	wg.Wait()
}
