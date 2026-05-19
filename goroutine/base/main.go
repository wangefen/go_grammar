package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("王一凡")
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go test()
	for i := 1; i <= 10; i++ {
		fmt.Println("王2凡")
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
	fmt.Println("结束啦")
	
}
