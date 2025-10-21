package main

import "fmt"

func main() {
	ch := make(chan int)
	var a int
	go func() {
		a = <-ch
	}()
	go func() {
		ch <- 10
	}()
	fmt.Println("发送成功", a)
}
