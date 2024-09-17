package main

import "fmt"

type cat struct {
	name    string
	age     uint8
	weightt uint8
	color   string
}

func (c cat) bark() string       { return "Meow" }
func (c cat) get_age() uint8     { return c.age }
func (c *cat) get_weight() uint8 { return c.weightt }

func (c cat) set_weight(weight uint8) { //不可改
	c.weightt = weight
}
func (c *cat) set_weight_r(weight uint8) { //可改
	c.weightt = weight
}

func main() {
	var fir_cat cat = cat{
		name:    "cat_1",
		age:     2,
		weightt: 10,
		color:   "black",
	}
	fmt.Println(fir_cat)
	fmt.Println(fir_cat.bark())
	fmt.Println(fir_cat.get_age())
	fmt.Println("set before", fir_cat.get_weight())
	fir_cat.set_weight(100)
	fmt.Println("set after", fir_cat.get_weight())

	fmt.Println("set before", fir_cat.get_weight())
	fir_cat.set_weight_r(66)
	fmt.Println("set after", fir_cat.get_weight())

	fmt.Println("====================================")

	cpy_cat := cat(fir_cat)
	fmt.Println(fir_cat, " => ", cpy_cat)
	fmt.Println(fir_cat.get_weight(), " => ", cpy_cat.get_weight())
	fmt.Println("is copy", cpy_cat == fir_cat)

	any_mous_str := struct {
		Name string
		age  int
	}{Name: "aa", age: -1}
	fmt.Println(any_mous_str)

}
