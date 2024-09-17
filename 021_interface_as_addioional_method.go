package main

import "fmt"

type Cat struct {
	name   string
	weight int
}
type People struct {
	name   string
	weight int
}

type Body interface {
	Eat(food string) error
	Excretion() error
}

func DoEat(someone Body, food string) error {
	return someone.Eat(food)
}
func DoExcretion(someone Body, food string) error {
	return someone.Excretion()
}

func (c *Cat) Eat(food string) error {
	if 10 <= c.weight {
		return fmt.Errorf("%s 太胖了,不能再吃,体重是 %d\n", c.name, c.weight)
	}
	c.weight++
	fmt.Printf("%s 吃完了,体重是 %d \n", c.name, c.weight)
	return nil
}
func (c *Cat) Excretion() error {
	if 2 > c.weight {
		return fmt.Errorf("%s 已经太瘦,无法拉,体重是 %d\n", c.name, c.weight)
	}
	c.weight--
	fmt.Printf("%s 拉在了猫砂上,体重更新为 %d \n", c.name, c.weight)
	return nil
}
func (p *People) Eat(food string) error {
	if 100 <= p.weight {
		return fmt.Errorf("%s 已经太重,不可以再吃 %d\n", p.name, p.weight)
	}
	p.weight++
	fmt.Printf("%s 吃了一顿饭,体重更新为 %d \n", p.name, p.weight)
	return nil
}
func (p *People) Excretion() error {
	if 95 > p.weight {
		return fmt.Errorf("%s 太瘦了,不可以上厕所,体重是 %d\n", p.name, p.weight)
	}
	p.weight--
	fmt.Printf("%s 上完了厕所,体重是 %d \n", p.name, p.weight)
	return nil
}

func main() {
	fmt.Println("===============init info===================")
	var fir_cat Cat = Cat{name: "cat1", weight: 7}
	var fir_peo *People = &People{name: "peo1", weight: 99}
	fmt.Println("fir_cat`s name is ", fir_cat.name, " and weight is ", fir_cat.weight)
	fmt.Println("fir_peo`s name is ", fir_peo.name, " and weight is ", fir_peo.weight)

	fmt.Println("===============add weight===================")
	DoEat(fir_peo, "肉")
	DoEat(&fir_cat, "肉")
	fmt.Printf("i'm %s,my body weight is %d\n", fir_cat.name, fir_cat.weight)
	fmt.Printf("i'm %s,my body weight is %d\n", fir_peo.name, fir_peo.weight)

	fmt.Println("===============try weight===================")
	DoExcretion(fir_peo, "肉")
	DoExcretion(fir_peo, "肉")
	DoExcretion(fir_peo, "肉")
	DoExcretion(fir_peo, "肉")
	DoExcretion(fir_peo, "肉")
	DoExcretion(fir_peo, "肉")
	DoExcretion(fir_peo, "肉")
	DoExcretion(fir_peo, "肉")
	err := DoExcretion(fir_peo, "肉")
	if err != nil {
		fmt.Println(err)
	}

	err = DoExcretion(fir_peo, "肉")
	DoExcretion(&fir_cat, "肉")
	DoExcretion(&fir_cat, "肉")
	DoExcretion(&fir_cat, "肉")
	DoExcretion(&fir_cat, "肉")
	DoExcretion(&fir_cat, "肉")
	DoExcretion(&fir_cat, "肉")
	DoExcretion(&fir_cat, "肉")
	DoExcretion(&fir_cat, "肉")
	err = DoExcretion(&fir_cat, "肉")
	if err != nil {
		fmt.Println(err)
	}
	err = DoExcretion(&fir_cat, "肉")
	if err != nil {
		fmt.Println(err)
	}
}
