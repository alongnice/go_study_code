package main

import (
	"fmt"
	"reflect"
)

// 定义了一个名为 tag_type 的结构体，其中包含三个字段，每个字段都有一个字符串标签。
type tag_type struct {
	field1 bool   "==this is bool field==“"
	field2 string "==this is string field==“"
	field3 int    "==this is int field==“"
}

// 实现了一个 get_field_by_mem_index 函数，该函数使用反射来获取结构体中指定索引的字段信息。
func get_field_by_mem_index(tt tag_type, index int) reflect.StructField {
	ttType := reflect.TypeOf(tt)
	ix_field := ttType.Field(index)
	return ix_field
}

func main() {
	tt := tag_type{}
	for i := 0; i < 3; i++ {
		// 对每个字段，调用 get_field_by_mem_index 函数获取字段信息，并打印出该字段的标签。
		struct_field := get_field_by_mem_index(tt, i)
		fmt.Printf("field %d`s tag is %v \n", i+1, struct_field.Tag)
	}
}
