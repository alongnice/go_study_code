package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
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

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 解读协议
	if r.Method == "GET" {           //如果是get
		t, _ := template.ParseFiles("login.gtpl") //解析模板文件
		t.Execute(w, nil)                         //执行模板
	} else {
		err := r.ParseForm() //解析参数，默认是不会解析的
		if err != nil {      //解析失败
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/", say_hello_name) //静态路由
	http.HandleFunc("/login", login)     //身份认证逻辑
	err := http.ListenAndServe(":12138", nil)
	if err != nil {
		log.Fatal("listen_and_server: ", err)
	}
}
