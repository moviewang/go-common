package main

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {

	go Say("world")
	Say("hello")
	arr := []int{14, 3, 1, 5, 67}
	ints := make(chan int)
	go Sum(arr[:len(arr)/2], ints)
	go Sum(arr[len(arr)/2:], ints)

	x, y := <-ints, <-ints
	fmt.Printf("x: %v, y：%v\n", x, y)
	fmt.Println(x + y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("------------")
	c := make(chan int, 10)
	go MyFibonacci(cap(c), c)
	for e := range c {
		fmt.Println(e)
	}
	fmt.Println("------------")
	cc := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cc)
		}
		quit <- 0
	}()
	FibonacciSelect(cc, quit)
	fmt.Println("///////////////")
	tick := time.Tick(100 * time.Millisecond)
	after := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-after:
			fmt.Println("boom")
			return
		default:
			fmt.Println("....")
			time.Sleep(50 * time.Millisecond)
		}
	}


}

func MyFibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func FibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("wait")
		}
	}
}
