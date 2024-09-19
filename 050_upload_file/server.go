package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request) {
	// 打印客户端的请求类别
	fmt.Println("method", r.Method)
	// 如果是GET请求，则显示上传的页面
	if r.Method == "GET" {
		crutime := time.Now().Unix()                      //获取当前时间
		h := md5.New()                                    //创建一个md5算法
		io.WriteString(h, strconv.FormatInt(crutime, 10)) //将时间戳写入h
		token := fmt.Sprintf("%x", h.Sum(nil))            //构建一个token

		t, _ := template.ParseFiles("upload.gtpl") //解析模板
		t.Execute(w, token)                        //将token写入模板
	} else {
		// 如果是POST请求，则处理上传的文件
		r.ParseMultipartForm(32 << 20)                 //设置内存大小,32MB(这里是上限)
		file, handler, err := r.FormFile("uploadfile") //获取上传的文件
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close() //文件句柄关闭
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(""+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

	}
}
func main() {
	// 增加一个路径映射
	http.HandleFunc("/upload", upload)
	// 设置监听端口
	err := http.ListenAndServe(":12138", nil)
	if err != nil {
		panic(err)
	}
}
