package main

import (
	"fmt"
	"unsafe"
)

//iota const的行索引
const (
	a  = iota // 0
	bb        // 1
	cc        // 2
)

type Human interface {
	show()
}

type Man struct {
}

type Woman struct {
}

func (w Woman) show() {
	fmt.Println("women")
}

func (m Man) show() {
	fmt.Println("man")
}

func main() {
	fmt.Println("hello world")
	var man Human = nil
	man = new(Woman)
	man.show()
	man = new(Man)

	nums := []int{1, 2, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 2 {
			fmt.Println("2 index:", i)
		}

	}

	kvs := map[string]string{"a": "a", "b": "bb"}
	for k, v := range kvs {
		fmt.Println("k->v", k, v)
	}

	//变量声明与赋值
	var aaa = "string"
	var b bool
	c := 10
	fmt.Println(aaa, b, c)

	//_只写变量的 被赋值的变量会被丢弃（原因是因为go强制使用所有被声明的变量）
	_, c = 5, 8
	fmt.Println(c)

	const LENGTH = 10
	const WIDTH = 20
	fmt.Println("area:", LENGTH*WIDTH)
	//iota
	fmt.Println(a, bb, cc)

	//unsafe string
	dd := "hello world"
	//字符串结构包含指向底层数组的指针长度，两个部分都是8字节所以size = 16
	fmt.Println(unsafe.Sizeof(dd) == 16)

	//指针变量
	var ptr *int
	ptr = &c
	fmt.Println(*ptr)

	//if  else
	var e int
	fmt.Println("input")
	fmt.Scan(&e)
	if e == 1 {
		fmt.Println("correct")
	} else {
		fmt.Println("false")
	}

	//switch case  don't need break
	var grade string = "b"
	var marks int = 0
	switch grade {
	case "a":
		grade = "A"
	case "b":
		grade = "B"
	case "c":
		grade = "C"
	}

	switch marks {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70:
		fmt.Println("C")
	default:
		fmt.Println("bad")
	}
	fmt.Println("your grade is :", grade)

	//type case
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("x's type is:", i)
	case float64:
		fmt.Println("x's type is:", i)
	case int:
		fmt.Println("x's type is:", i)
	default:
		fmt.Println("type is unknown")

	}

	//select 每个case必须是一个通信
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received from c1", i1)
	case c2 <- i2:
		fmt.Println("send i2 to c2")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("ok", i3)
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")

	}
	//loop
	var aaaa = 15
	var bbb int
	numbers := []int{1, 3, 54, 5, 6}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for bbb < aaaa {
		bbb++
		fmt.Println("bbb", bbb)
	}

	for i, x := range numbers {
		fmt.Println(i, x)
	}

	//print prime
	var m, n int
	for m = 2; m <= 100; m++ {
		for n = 2; n <= (m / n); n++ {
			if m%n == 0 {
				break
			}
		}
		if n > (m / n) {
			fmt.Println(m)
		}
	}
}
