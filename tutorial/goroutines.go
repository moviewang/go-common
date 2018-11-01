package main

import (
	"fmt"
	"time"
)

func say(s string)  {
	for i := 0; i < 5; i++  {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int)  {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
}

//range and close
func fibonacci(n int,  c chan int)  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x +y
	}
	close(c)
}

func fibonacci2(c, quit chan int)  {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x +y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func main()  {
	go say("world")
	say("hello")

	//two go routines sum arrasy
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s) / 2], c)
	go sum(s[len(s) / 2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+ y)

	//buffered channel  buffer length is 2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	fmt.Println("/////////////////")

	chl := make(chan int, 10)
	go fibonacci(cap(chl), chl)
	for i := range chl{
		fmt.Println(i)
	}
	fmt.Println("//////////////////////")
	//select
	cc := make(chan int)
	quit := make(chan int)
	go func () {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cc)
		}
		quit <- 0
	}()
	fibonacci2(cc, quit)
	fmt.Println("************")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick")
		case <- boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("        .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

