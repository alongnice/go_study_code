package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if nil != err {
		fmt.Errorf("%s", err)
		panic(err)
	}
}

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}
func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	checkErr(err)

	//如果不存在就建立資料表users
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, username VARCHAR(30) NOT NULL, age INT(3) NOT NULL, sex VARCHAR(1) NOT NULL, timestamp TIMESTAMP NOT NULL)")
	checkErr(err)
	/*
	   CREATE TABLE IF NOT EXISTS users
	   	(id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	   	username VARCHAR(30) NOT NULL,
	   	age INT(3) NOT NULL,
	   	sex VARCHAR(1) NOT NULL,
	   	timestamp TIMESTAMP NOT NULL)")
	   创建一个表，如果不存在的话。叫做users
	   	字段1: id 设置为6位无符号整数，自动递增，作为主键
	   	字段2: username 设置为30位字符串，不能为空
	   	字段3: age 设置为3位整数，不能为空
	   	字段4: sex 设置为1位字符串，不能为空
	*/

	//插入資料
	stmt, err := db.Prepare("INSERT users SET username=?,age=?,sex=?,timestamp=?")
	checkErr(err)
	/*
		创建一个处理语句；给users表插入数据
		这里是还没有执行的
	*/

	// 執行 query 或 prepared
	res, err := stmt.Exec("along", TrueRandom(100), "F", time.Now())
	// 给定参数
	checkErr(err)

	id, err := res.LastInsertId()
	//获取最后插入的id
	fmt.Println("LastInsertId():", id)
	// 打印执行id
	checkErr(err)

	// =========================================================
	// 更新数据
	stmt, err = db.Prepare("update users set username=? where id=?")
	checkErr(err)
	res, err = stmt.Exec("ALONG", 7)
	checkErr(err)

	affect, err := res.RowsAffected() //操作影响的数据库行数
	checkErr(err)
	fmt.Println("更新了:", affect, "行")

	// =========================================================
	// 查询数据
	rows, err := db.Query("SELECT * FROM test.users")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		var age int
		var sex string
		var timestamp interface{}
		err = rows.Scan(&id, &username, &age, &sex, &timestamp)
		checkErr(err)

		if nil != timestamp {
			switch v := timestamp.(type) {
			case []uint8:
				t, _ := time.Parse("2006-01-02 15:04:05", string(v))
				fmt.Printf("%v %v %v %v %v %v\n", id, username, age, sex, t, reflect.TypeOf(v))
			case time.Time:
				fmt.Printf("%d %s %d %s %s %v\n", id, username, age, sex, v, reflect.TypeOf(v))
			case string:
				//t, err := time.Parse("2006-01-02 15:04:05", v)
				t, err := time.Parse(time.Now().String(), v)
				if nil != err {
					fmt.Printf("%v, %v, %v, %v, %v, %v\n", id, username, age, sex, t, reflect.TypeOf(t))
				} else {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Printf("%v, %v, %v, %v, %v\n", id, username, age, sex, "timestamp is null")
		}
	}
	// =========================================================
	// 删除数据
	fmt.Println("====================================")
	stmt, err = db.Prepare("delete from users where id=13")
	fmt.Println(stmt)
	checkErr(err)
	fmt.Println("=NULL:", err)

	res, err = stmt.Exec()
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("影响行数：", affect)

	db.Close()
}
