package main

import "fmt"

func main() {
	ch1 := make(chan string, 10)
	ch2 := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- "wyf" + fmt.Sprintf("%d", i)
		ch2 <- i
	}
Exitt:
	for {
		select {
		case v := <-ch1:
			fmt.Println("from ch1:" + v)
		case v := <-ch2:
			fmt.Println("from ch2", v)
		default:
			fmt.Println("over")
			break Exitt
		}
	}
}
