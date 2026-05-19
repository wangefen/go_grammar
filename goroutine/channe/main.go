package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}

	close(ch1) //关闭通道，否则for range会认为只要管道没关闭，就一直等下一个数据，从而导致死锁
	// for循环不会这样
	for val := range ch1 {
		fmt.Println(val)
	}
}
