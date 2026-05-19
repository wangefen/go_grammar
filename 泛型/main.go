package main

import "fmt"

//泛型申明
func getdata[T any](value T) T {
	return value
}

type container[T any] struct {
	value T
}

func (c *container[T]) Set(value T) {
	c.value = value
}

func (c *container[T]) Get() T {
	return c.value
}

func main() {
	a := "wyf"
	b := 123
	intcon := container[int]{}
	strcon := container[string]{}
	fmt.Println(getdata(a))
	fmt.Println(getdata(b))
	intcon.Set(13)
	strcon.Set("wyf")
	fmt.Println(intcon.Get())
	fmt.Println(strcon.Get())
}
