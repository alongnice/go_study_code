package main

import "fmt"

func demo_select_is_random2eash_case() {
	var close_channels = func(channels ...chan interface{}) {
		for _, ch := range channels {
			close(ch)
		}
	}
	// 构建一个闭包函数,关闭多个信道

	c1 := make(chan interface{})
	c2 := make(chan interface{})
	close_channels(c2)
	// 创建两个无缓冲的信道

	var c1_count, c2_count, def_count int
	// 定义三个计数器
	defer fmt.Printf("c1Count:%d, c2Count:%d, defaultCount:%d\n", c1_count, c2_count, def_count)

	fmt.Println("into for loop")
	for i := 1; i < 1000; i++ { //循环999次
		select {
		case <-c1:
			c1_count++
			// if 600 == i {
			// 	close_channels(c1, c2)
			// }
		case <-c2:
			c2_count++
			// if 600 == i {
			// 	close_channels(c1, c2)
			// }
		default:
			def_count++
			if 600 == i {
				close_channels(c1, c2)
				fmt.Println("c1_count:", c1_count, "c2_count:", c2_count, "def_count:", def_count)
				fmt.Println("c1 and c2 closed!")
			}
		}
	}
	// select 语句会随机选择一个 case 执行,如果多个 case 都满足条件,也会随机选择一个
	// 权重则是 close 信道 > default > open 信道
	// runtime 无法解析select的意图
	// close_channels(c1, c2)
	fmt.Println("c1_count:", c1_count, "c2_count:", c2_count, "def_count:", def_count)

}

func main() {
	demo_select_is_random2eash_case()
}
