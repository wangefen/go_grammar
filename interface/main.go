package main

import "fmt"

type shuaige interface { //封装在此类接口下的结构体必须有face() string和height()这俩个函数
	face() string
	height()
}

type person struct {
	Name string
}

type people struct{}

func (p people) work(s shuaige) {
	fmt.Println("可以去做模子哥了！")
}

func (p person) face() string {
	if p.Name == "wyf" {
		return "帅逼"
	} else {
		return "丑逼"
	}
}

func (p person) height() {
	fmt.Println("182黄金身高")
}

func (p person) ex() {
	fmt.Println("只属于结构体的方法，接口调用该函数会报错")
}

func main() {
	p := person{Name: "wyf"}
	p3 := person{Name: "678"}
	fmt.Println(p.face())
	var p1 shuaige = p
	p1.face()
	p1.height()
	var p2 shuaige = p3
	fmt.Println(p2.face())
	p.ex()
	pl := people{}
	pl.work(p)
}
