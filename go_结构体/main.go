package main

import "fmt"

type person struct {
	name string
	age  int
	face string
}

func (p person) showyouoff() {
	fmt.Println(p.name, p.age, "是个", p.face)
}

func main() {
	var wyf person
	wyf.name = "王小凡"
	wyf.age = 18
	wyf.face = "大帅逼"
	wyf.showyouoff()
	//	a := make([]int, 2, 3)
	a := []int{1, 2, 3}
	fmt.Println(a)
}
