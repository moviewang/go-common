package main

import (
	"fmt"
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

func swap(x, y string) (string, string) {
	return y, x
}

func change(x *int, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}
func getSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Circle struct {
	radius float64
}

//Circle's method
func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func add(x, y int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x + y
	}

}

func main() {
	//go函数默认值传递 将实际参数复制一份到函数中原来参数的值不变
	fmt.Println(max(10, 20))
	a := "world"
	b := "hello"
	fmt.Println(swap(a, b))
	fmt.Println("old a, b", a, b)

	//引用传值修改原来的参数
	x, y := 10, 20
	change(&x, &y)
	fmt.Println("old x, y", x, y)

	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSqrt(9))
	fmt.Println("--------")
	//闭包
	seq := getSeq()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

	getAdd := add(1, 23)
	fmt.Println(getAdd())

	circle := new(Circle)
	circle.radius = 10
	fmt.Println(circle.getArea())

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d \t", j, i, i * j)
		}
		fmt.Println()
	}

}
