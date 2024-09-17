package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
		err := r.ParseForm()
		if nil == err {
			if 0 == len(r.Form["username"][0]) {
				msg := "未输入用户名"
				fmt.Fprintf(w, msg)
				fmt.Println(msg)
			}

			strage := r.Form.Get("age")
			if "" == strage {
				msg := "未输入年龄"
				fmt.Fprintf(w, msg)
				fmt.Println(msg)
			} else {
				get_age, err := strconv.Atoi(strage)
				if err != nil {
					fmt.Println(err)
					fmt.Fprintln(w, err)
				} else {
					msg := fmt.Sprintf("输入的年龄是:%d", get_age)
					fmt.Println(msg)
					fmt.Fprintln(w, msg)
				}
			}

			if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m {
				msg := fmt.Sprintf("输入的姓名不合法")
				fmt.Println(msg)
				fmt.Fprintln(w, msg)
			}

			//todo

			// 多选
			var mul_slect_check = func(select_name string,
				select_title_text string,
				lega_list map[string]struct{}) {
				_, ok := lega_list[r.Form.Get(select_name)]
				if ok {
					msg := fmt.Sprintf("%s选择了 %v\n", select_title_text, r.Form[select_name])
					fmt.Println(msg)
					fmt.Fprintln(w, msg)
				} else {
					msg := fmt.Sprintf("%v 不在选项\n", r.Form.Get(select_name))
					fmt.Println(msg)
					fmt.Fprintln(w, msg)
				}
			}

			// 下拉菜单
			drop_down_list := map[string]struct{}{
				"apple":  struct{}{},
				"pear":   struct{}{},
				"banana": struct{}{},
			}
			mul_slect_check("fruit", "水果选项", drop_down_list)

			// 单选
			radio := map[string]struct{}{
				"M": struct{}{},
				"F": struct{}{},
			}
			mul_slect_check("gender", "性别", radio)

			// 日期和时间

			// 身份证号
			fmt.Println("username", r.Form["username"])
			fmt.Println("password", r.Form["password"])
		} else {
			fmt.Print(err)
		}
	}
}

func main() {
	http.HandleFunc("/", say_hello_name)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":12138", nil)
	if err != nil {
		log.Fatal("listen_and_server: ", err)
	}

}
