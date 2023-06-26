package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		for i := 0; i <= 3; i++ {
			fmt.Println(i)
		}
		time.Sleep(time.Second * 3)
		ch <- 100
	}()

	fmt.Println(<-ch)
}
