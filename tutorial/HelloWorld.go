package main

import (
	"fmt"
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
	for _, 	num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums{
		if num == 2 {
			fmt.Println("2 index:", i)
		}

	}

	kvs := map[string]string{"a": "a", "b": "bb"}
	for k, v := range kvs{
		fmt.Println("k->v", k, v)
	}

}



