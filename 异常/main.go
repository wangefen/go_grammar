package main

import "fmt"

func test(a, b int) int {
	defer func() {
		err := recover() //recover catched the error
		if err != nil {
			fmt.Println("除数不能为0", err)
		}
	}()
	return a / b //return不是原子语句，Go 会先把 a/b 的计算结果存到一个临时的变量中，在return
	//因为这里有了defer，所以顺序是：1.存a/b 2.defer 3.return
}

func main() {
	a := test(5, 0)
	fmt.Println(a)
}
