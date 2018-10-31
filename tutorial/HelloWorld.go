package main

import "fmt"

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

}


