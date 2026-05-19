package main

import "fmt"

//当不确定接受的数据类型时，就可以用空接口，空接口可以接受任意类型的数据
//这个特性还可以用来判断前端传来的数据是什么类型

func test(a interface{}) {
	fmt.Printf("类型：%T, 值：%v\n", a, a)
}

func main() {
	var a interface{}
	a = 123
	fmt.Printf("类型：%T, 值：%v\n", a, a)
	test(a)
	a = "wyf"
	fmt.Printf("类型：%T, 值：%v\n", a, a)
	test(a)
	a = true
	fmt.Printf("类型：%T, 值：%v\n", a, a)
	test(a)

	//map,slice中也能用
	mp := map[string]interface{}{}
	mp["ff"] = "sfa"
	fmt.Println(mp["ff"])

	sl := []interface{}{}
	sl = append(sl, "wer")
	sl = append(sl, 123)
	fmt.Println(sl)
}
