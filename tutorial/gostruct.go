package main

import (
	"fmt"
)

type Book struct {
	title  string
	author string
	id     int
}

func printSlice(x []int) {
	fmt.Println("len, cap, slice", len(x), cap(x), x)
}

func factoria(n uint64) uint64 {
	if n > 1 {
		return n * factoria(n-1)
	}
	return 1
}

func finbonacci(n uint64) uint64 {
	if n >= 2 {
		return finbonacci(n-1) + finbonacci(n-2)
	}
	return n
}

type Phone interface {
	call(p string) string
}

type HUaWei struct {
}

func (HUaWei) call(p string) string {
	return p + "hw"
}

type Xiaomi struct {
	
}

func (Xiaomi) call(p string) string {
	return p + "xm"
}

type DivError struct {
	dividee int
	divider int
}

func (de *DivError) Error() string {
	str := `
   can not procced,
   dividee : %d
   divider : 0
	`
	return fmt.Sprintf(str, de.dividee)
}

func Divide(x, y int) (result int, err string){
	if y == 0 {
		divError := DivError{x, y}
		err = divError.Error()
		return
	}
	return x / y, ""
}


func main() {
	book := Book{"java", "Gaosinlin", 1}
	book2 := Book{title: "c++", author: "none", id: 2}
	fmt.Println(book)
	fmt.Println(book2)

	var b *Book
	b = &book2
	fmt.Println(b == &book2)
	fmt.Println(*b == book2)
	fmt.Println(b.author, b.title, b.id)

	numbers := make([]int, 3, 5)
	printSlice(numbers)

	//slice apend and copy
	var nums []int
	printSlice(nums)

	nums = append(nums, 1)
	printSlice(nums)
	nums = append(nums, 2, 3, 4)
	printSlice(nums)
	num1s := make([]int, len(nums), cap(nums)*2)
	printSlice(num1s)

	copy(num1s, nums)

	printSlice(num1s)

	//range iterate
	for _, nn := range num1s {
		fmt.Print(nn, "\t")
	}
	smap := map[string]string{"a": "aa", "b": "bb"}
	fmt.Println()
	for k, v := range smap {
		fmt.Println("key:", k, "value:", v)
	}

	//map
	maps := make(map[string]string)
	maps["Italy"] = "Rome"
	maps["UK"] = "London"
	maps["America"] = "WC"
	for e := range maps {
		fmt.Println(e)
	}

	s := maps["UK"]
	if s != "" {
		fmt.Println(s)
	} else {
		fmt.Println("not exist")
	}

	delete(maps, "UK")
	fmt.Println(maps)

	u := factoria(15)
	fmt.Println(u)

	for i := 0; i < 10; i++ {
		fmt.Print(finbonacci(uint64(i)), "\t")
	}
	fmt.Println()

	var phone Phone
	phone = new(HUaWei)
	fmt.Println(phone.call("phone"))
	phone = new(Xiaomi)
	fmt.Println(phone.call("xphone"))

	result, err := Divide(10, 0)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
