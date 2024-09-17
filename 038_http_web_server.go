package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func say_hello_name(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "===============Hello along!==============")
}

func main() {
	http.HandleFunc("/", say_hello_name)
	err := http.ListenAndServe(":12318", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 启动服务后 浏览器访问  http://localhost:12318/?url_long=value1&url_long=value2&key1=val1&key2=val2&key3=val3
