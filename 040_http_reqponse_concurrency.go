package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func get_gid() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func main() {
	var accmulate_res = 0
	var print_num_goroutines = func(w http.ResponseWriter, r *http.Request) {
		go func() {
			accmulate_res++
			msg := fmt.Sprintf("ID:%v Number of goroutines: %d\tnum of goroutine: %d\n", get_gid(), accmulate_res, runtime.NumGoroutine())
			<-time.After(time.Second)
			fmt.Fprintf(w, msg)
			fmt.Println(msg)
		}()
	}

	http.HandleFunc("/concurrency_exam", print_num_goroutines)
	err := http.ListenAndServe(":12138", nil)
	if nil != err {
		log.Fatal("ListenAndServe: ", err)
	}

}
